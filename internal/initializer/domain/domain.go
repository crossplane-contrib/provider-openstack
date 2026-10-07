/*
Copyright 2026 The Crossplane Authors.
*/

// Package domain provides a managed resource initializer that reports a clear
// error when a domainId reference points to a ProjectV3 that is not a domain.
package domain

import (
	"context"
	"fmt"

	apimeta "k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/v2/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/v2/pkg/meta"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
)

const projectKind = "ProjectV3"

// isDomainPaths are the fields consulted to decide whether a ProjectV3 is a
// domain. They must stay in sync with the ExtractDomainID reference extractors
// in apis/{cluster,namespaced}/identity/v1alpha1, otherwise this initializer
// would accept a reference that the resolver later rejects.
var isDomainPaths = []string{"spec.forProvider.isDomain", "status.atProvider.isDomain"}

// sites are the parameter blocks that can carry a domainId reference.
var sites = []string{"spec.forProvider", "spec.initProvider"}

// NewValidator returns an Initializer that verifies domainId references point
// to a ProjectV3 that is a domain. It satisfies upjet's config.NewInitializerFn.
func NewValidator(c client.Client) managed.Initializer {
	return &validator{client: c}
}

type validator struct {
	client client.Client
}

// Initialize returns an error describing the problem when a domainId reference
// resolves to a ProjectV3 that is not a domain. Every other situation, in
// particular a reference to an object that does not exist yet, is left to the
// regular reference resolution that runs after initialization.
func (v *validator) Initialize(ctx context.Context, mg resource.Managed) error {
	gvk, err := v.client.GroupVersionKindFor(mg)
	if err != nil {
		return nil //nolint:nilerr // Not our concern; the reconciler reports it.
	}

	paved, err := fieldpath.PaveObject(mg)
	if err != nil {
		return nil //nolint:nilerr // Not our concern; the reconciler reports it.
	}

	for _, site := range sites {
		// A resolved value is cached in the spec, so there is nothing left to
		// check and we avoid re-reading the target on every reconcile.
		if id, err := paved.GetString(site + ".domainId"); err == nil && id != "" {
			continue
		}
		if err := v.checkSite(ctx, mg, paved, gvk.GroupVersion(), site); err != nil {
			return err
		}
	}

	return nil
}

func (v *validator) checkSite(ctx context.Context, mg resource.Managed, paved *fieldpath.Paved, gv schema.GroupVersion, site string) error {
	if err := v.checkRef(ctx, mg, paved, gv, site); err != nil {
		return err
	}
	return v.checkSelector(ctx, mg, paved, gv, site)
}

func (v *validator) checkRef(ctx context.Context, mg resource.Managed, paved *fieldpath.Paved, gv schema.GroupVersion, site string) error {
	name, err := paved.GetString(site + ".domainIdRef.name")
	if err != nil {
		return nil //nolint:nilerr // Field not present; nothing to validate.
	}
	if name == "" {
		return nil
	}
	ns := refNamespace(paved, site+".domainIdRef.namespace", mg)
	target, err := v.getProject(ctx, gv, types.NamespacedName{Name: name, Namespace: ns})
	if err != nil || target == nil {
		return nil //nolint:nilerr // Unresolvable for now; the resolver reports it.
	}
	if isDomain(target) {
		return nil
	}
	return fmt.Errorf("%s.domainIdRef: referenced %s %q is a project, not a domain: reference a %s created with isDomain: true", site, projectKind, name, projectKind)
}

func (v *validator) checkSelector(ctx context.Context, mg resource.Managed, paved *fieldpath.Paved, gv schema.GroupVersion, site string) error {
	labels, err := paved.GetStringObject(site + ".domainIdSelector.matchLabels")
	if err != nil || len(labels) == 0 {
		return nil //nolint:nilerr // No reference configured for this site.
	}

	name, target, err := v.selectProject(ctx, mg, paved, gv, site, labels)
	if err != nil || target == nil {
		return nil //nolint:nilerr // Unresolvable for now; the resolver reports it.
	}
	if isDomain(target) {
		return nil
	}

	return fmt.Errorf("%s.domainIdSelector: selected %s %q is a project, not a domain: adjust the selector to match a %s created with isDomain: true", site, projectKind, name, projectKind)
}

// selectProject mirrors how reference.APIResolver picks a target for a
// selector, so the reported name is the one the resolver would use.
func (v *validator) selectProject(ctx context.Context, mg resource.Managed, paved *fieldpath.Paved, gv schema.GroupVersion, site string, labels map[string]string) (string, *fieldpath.Paved, error) {
	lo, err := v.client.Scheme().New(gv.WithKind(projectKind + "List"))
	if err != nil {
		return "", nil, err
	}
	list, ok := lo.(client.ObjectList)
	if !ok {
		return "", nil, fmt.Errorf("%T is not a client.ObjectList", lo)
	}
	if err := v.client.List(ctx, list, client.MatchingLabels(labels), client.InNamespace(refNamespace(paved, site+".domainIdSelector.namespace", mg))); err != nil {
		return "", nil, err
	}

	items, err := apimeta.ExtractList(list)
	if err != nil {
		return "", nil, err
	}

	mustMatchController, _ := paved.GetBool(site + ".domainIdSelector.matchControllerRef")
	for _, item := range items {
		obj, ok := item.(client.Object)
		if !ok {
			continue
		}
		if mustMatchController && !meta.HaveSameController(mg, obj) {
			continue
		}
		target, err := fieldpath.PaveObject(obj)
		if err != nil {
			return "", nil, err
		}
		return obj.GetName(), target, nil
	}

	return "", nil, nil
}

func (v *validator) getProject(ctx context.Context, gv schema.GroupVersion, nn types.NamespacedName) (*fieldpath.Paved, error) {
	o, err := v.client.Scheme().New(gv.WithKind(projectKind))
	if err != nil {
		return nil, err
	}
	obj, ok := o.(client.Object)
	if !ok {
		return nil, fmt.Errorf("%T is not a client.Object", o)
	}
	if err := v.client.Get(ctx, nn, obj); err != nil {
		return nil, err
	}

	return fieldpath.PaveObject(obj)
}

func isDomain(p *fieldpath.Paved) bool {
	for _, path := range isDomainPaths {
		if v, err := p.GetBool(path); err == nil && v {
			return true
		}
	}

	return false
}

// refNamespace reads an explicit namespace from the paved object at the given
// path, falling back to the managed resource's own namespace. This mirrors
// how APINamespacedResolver resolves NamespacedReference.Namespace in
// crossplane-runtime.
func refNamespace(paved *fieldpath.Paved, path string, mg resource.Managed) string {
	if ns, err := paved.GetString(path); err == nil && ns != "" {
		return ns
	}
	return mg.GetNamespace()
}
