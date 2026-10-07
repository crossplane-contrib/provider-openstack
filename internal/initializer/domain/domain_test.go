/*
Copyright 2026 The Crossplane Authors.
*/

package domain

import (
	"context"
	"strings"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"

	identityv1alpha1 "github.com/crossplane-contrib/provider-openstack/apis/cluster/identity/v1alpha1"
	nsidentityv1alpha1 "github.com/crossplane-contrib/provider-openstack/apis/namespaced/identity/v1alpha1"
)

func ptr[T any](v T) *T { return &v }

func newScheme(t *testing.T) *runtime.Scheme {
	t.Helper()
	s := runtime.NewScheme()
	if err := identityv1alpha1.SchemeBuilder.AddToScheme(s); err != nil {
		t.Fatal(err)
	}
	return s
}

func newDomain(name string) *identityv1alpha1.ProjectV3 {
	return &identityv1alpha1.ProjectV3{
		ObjectMeta: metav1.ObjectMeta{Name: name},
		Spec: identityv1alpha1.ProjectV3Spec{
			ForProvider: identityv1alpha1.ProjectV3Parameters{
				IsDomain: ptr(true),
				Name:     ptr(name),
			},
		},
	}
}

func newProject(name string) *identityv1alpha1.ProjectV3 {
	return &identityv1alpha1.ProjectV3{
		ObjectMeta: metav1.ObjectMeta{Name: name},
		Spec: identityv1alpha1.ProjectV3Spec{
			ForProvider: identityv1alpha1.ProjectV3Parameters{
				Name: ptr(name),
			},
		},
	}
}

func newProjectWithLabels(name string, labels map[string]string) *identityv1alpha1.ProjectV3 {
	p := newProject(name)
	p.Labels = labels
	return p
}

func newDomainWithLabels(name string, labels map[string]string) *identityv1alpha1.ProjectV3 {
	d := newDomain(name)
	d.Labels = labels
	return d
}

func TestInitialize(t *testing.T) {
	tests := []struct {
		name      string
		existing  []runtime.Object
		managed   *identityv1alpha1.ProjectV3
		wantErr   bool
		errSubstr string
	}{
		{
			name: "NoRef_NoError",
			managed: &identityv1alpha1.ProjectV3{
				ObjectMeta: metav1.ObjectMeta{Name: "my-project"},
				Spec: identityv1alpha1.ProjectV3Spec{
					ForProvider: identityv1alpha1.ProjectV3Parameters{
						Name: ptr("my-project"),
					},
				},
			},
		},
		{
			name:     "RefToDomain_NoError",
			existing: []runtime.Object{newDomain("my-domain")},
			managed: &identityv1alpha1.ProjectV3{
				ObjectMeta: metav1.ObjectMeta{Name: "my-project"},
				Spec: identityv1alpha1.ProjectV3Spec{
					ForProvider: identityv1alpha1.ProjectV3Parameters{
						Name:        ptr("my-project"),
						DomainIDRef: &xpv1.Reference{Name: "my-domain"},
					},
				},
			},
		},
		{
			name:     "RefToProject_Error",
			existing: []runtime.Object{newProject("not-a-domain")},
			managed: &identityv1alpha1.ProjectV3{
				ObjectMeta: metav1.ObjectMeta{Name: "child-project"},
				Spec: identityv1alpha1.ProjectV3Spec{
					ForProvider: identityv1alpha1.ProjectV3Parameters{
						Name:        ptr("child-project"),
						DomainIDRef: &xpv1.Reference{Name: "not-a-domain"},
					},
				},
			},
			wantErr:   true,
			errSubstr: "is a project, not a domain",
		},
		{
			name: "RefToNonExistent_NoError",
			managed: &identityv1alpha1.ProjectV3{
				ObjectMeta: metav1.ObjectMeta{Name: "my-project"},
				Spec: identityv1alpha1.ProjectV3Spec{
					ForProvider: identityv1alpha1.ProjectV3Parameters{
						Name:        ptr("my-project"),
						DomainIDRef: &xpv1.Reference{Name: "does-not-exist"},
					},
				},
			},
		},
		{
			name:     "AlreadyResolved_SkipsCheck",
			existing: []runtime.Object{newProject("not-a-domain")},
			managed: &identityv1alpha1.ProjectV3{
				ObjectMeta: metav1.ObjectMeta{Name: "my-project"},
				Spec: identityv1alpha1.ProjectV3Spec{
					ForProvider: identityv1alpha1.ProjectV3Parameters{
						Name:        ptr("my-project"),
						DomainID:    ptr("some-uuid-already-resolved"),
						DomainIDRef: &xpv1.Reference{Name: "not-a-domain"},
					},
				},
			},
		},
		{
			name: "SelectorToDomain_NoError",
			existing: []runtime.Object{newDomainWithLabels("sel-domain", map[string]string{
				"app": "test",
			})},
			managed: &identityv1alpha1.ProjectV3{
				ObjectMeta: metav1.ObjectMeta{Name: "my-project"},
				Spec: identityv1alpha1.ProjectV3Spec{
					ForProvider: identityv1alpha1.ProjectV3Parameters{
						Name: ptr("my-project"),
						DomainIDSelector: &xpv1.Selector{
							MatchLabels: map[string]string{"app": "test"},
						},
					},
				},
			},
		},
		{
			name: "SelectorToProject_Error",
			existing: []runtime.Object{newProjectWithLabels("sel-project", map[string]string{
				"app": "test",
			})},
			managed: &identityv1alpha1.ProjectV3{
				ObjectMeta: metav1.ObjectMeta{Name: "my-project"},
				Spec: identityv1alpha1.ProjectV3Spec{
					ForProvider: identityv1alpha1.ProjectV3Parameters{
						Name: ptr("my-project"),
						DomainIDSelector: &xpv1.Selector{
							MatchLabels: map[string]string{"app": "test"},
						},
					},
				},
			},
			wantErr:   true,
			errSubstr: "is a project, not a domain",
		},
		{
			name: "SelectorNoMatch_NoError",
			managed: &identityv1alpha1.ProjectV3{
				ObjectMeta: metav1.ObjectMeta{Name: "my-project"},
				Spec: identityv1alpha1.ProjectV3Spec{
					ForProvider: identityv1alpha1.ProjectV3Parameters{
						Name: ptr("my-project"),
						DomainIDSelector: &xpv1.Selector{
							MatchLabels: map[string]string{"missing": "label"},
						},
					},
				},
			},
		},
		{
			name:     "InitProviderRefToProject_Error",
			existing: []runtime.Object{newProject("init-project")},
			managed: &identityv1alpha1.ProjectV3{
				ObjectMeta: metav1.ObjectMeta{Name: "my-project"},
				Spec: identityv1alpha1.ProjectV3Spec{
					ForProvider: identityv1alpha1.ProjectV3Parameters{
						Name: ptr("my-project"),
					},
					InitProvider: identityv1alpha1.ProjectV3InitParameters{
						DomainIDRef: &xpv1.Reference{Name: "init-project"},
					},
				},
			},
			wantErr:   true,
			errSubstr: "spec.initProvider.domainIdRef",
		},
		{
			name:     "IsDomainInStatus_NoError",
			existing: []runtime.Object{statusDomain("status-domain")},
			managed: &identityv1alpha1.ProjectV3{
				ObjectMeta: metav1.ObjectMeta{Name: "my-project"},
				Spec: identityv1alpha1.ProjectV3Spec{
					ForProvider: identityv1alpha1.ProjectV3Parameters{
						Name:        ptr("my-project"),
						DomainIDRef: &xpv1.Reference{Name: "status-domain"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scheme := newScheme(t)
			builder := fake.NewClientBuilder().WithScheme(scheme)
			if len(tt.existing) > 0 {
				builder = builder.WithRuntimeObjects(tt.existing...)
			}
			c := builder.Build()

			v := NewValidator(c)
			err := v.Initialize(context.Background(), tt.managed)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if !strings.Contains(err.Error(), tt.errSubstr) {
					t.Errorf("error %q does not contain %q", err.Error(), tt.errSubstr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func statusDomain(name string) *identityv1alpha1.ProjectV3 {
	return &identityv1alpha1.ProjectV3{
		ObjectMeta: metav1.ObjectMeta{Name: name},
		Spec: identityv1alpha1.ProjectV3Spec{
			ForProvider: identityv1alpha1.ProjectV3Parameters{
				Name: ptr(name),
			},
		},
		Status: identityv1alpha1.ProjectV3Status{
			AtProvider: identityv1alpha1.ProjectV3Observation{
				IsDomain: ptr(true),
			},
		},
	}
}

func newNSScheme(t *testing.T) *runtime.Scheme {
	t.Helper()
	s := runtime.NewScheme()
	if err := nsidentityv1alpha1.SchemeBuilder.AddToScheme(s); err != nil {
		t.Fatal(err)
	}
	return s
}

func newNSDomain(name, ns string) *nsidentityv1alpha1.ProjectV3 {
	return &nsidentityv1alpha1.ProjectV3{
		ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: ns},
		Spec: nsidentityv1alpha1.ProjectV3Spec{
			ForProvider: nsidentityv1alpha1.ProjectV3Parameters{
				IsDomain: ptr(true),
				Name:     ptr(name),
			},
		},
	}
}

func newNSProject(name, ns string) *nsidentityv1alpha1.ProjectV3 {
	return &nsidentityv1alpha1.ProjectV3{
		ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: ns},
		Spec: nsidentityv1alpha1.ProjectV3Spec{
			ForProvider: nsidentityv1alpha1.ProjectV3Parameters{
				Name: ptr(name),
			},
		},
	}
}

func TestInitializeCrossNamespace(t *testing.T) {
	tests := []struct {
		name      string
		existing  []runtime.Object
		managed   *nsidentityv1alpha1.ProjectV3
		wantErr   bool
		errSubstr string
	}{
		{
			name: "RefCrossNS_DomainInTargetNS_NoError",
			existing: []runtime.Object{
				newNSDomain("shared-domain", "infrastructure"),
				newNSProject("shared-domain", "application"),
			},
			managed: &nsidentityv1alpha1.ProjectV3{
				ObjectMeta: metav1.ObjectMeta{Name: "my-project", Namespace: "application"},
				Spec: nsidentityv1alpha1.ProjectV3Spec{
					ForProvider: nsidentityv1alpha1.ProjectV3Parameters{
						Name: ptr("my-project"),
						DomainIDRef: &xpv1.NamespacedReference{
							Name:      "shared-domain",
							Namespace: "infrastructure",
						},
					},
				},
			},
		},
		{
			name: "RefCrossNS_ProjectInTargetNS_Error",
			existing: []runtime.Object{
				newNSProject("not-a-domain", "infrastructure"),
			},
			managed: &nsidentityv1alpha1.ProjectV3{
				ObjectMeta: metav1.ObjectMeta{Name: "my-project", Namespace: "application"},
				Spec: nsidentityv1alpha1.ProjectV3Spec{
					ForProvider: nsidentityv1alpha1.ProjectV3Parameters{
						Name: ptr("my-project"),
						DomainIDRef: &xpv1.NamespacedReference{
							Name:      "not-a-domain",
							Namespace: "infrastructure",
						},
					},
				},
			},
			wantErr:   true,
			errSubstr: "is a project, not a domain",
		},
		{
			name: "RefNoExplicitNS_FallsBackToMgNS",
			existing: []runtime.Object{
				newNSDomain("local-domain", "application"),
			},
			managed: &nsidentityv1alpha1.ProjectV3{
				ObjectMeta: metav1.ObjectMeta{Name: "my-project", Namespace: "application"},
				Spec: nsidentityv1alpha1.ProjectV3Spec{
					ForProvider: nsidentityv1alpha1.ProjectV3Parameters{
						Name: ptr("my-project"),
						DomainIDRef: &xpv1.NamespacedReference{
							Name: "local-domain",
						},
					},
				},
			},
		},
		{
			name: "SelectorCrossNS_DomainInTargetNS_NoError",
			existing: []runtime.Object{
				newNSProjectWithLabels("shared-domain", "infrastructure", map[string]string{"type": "domain"}, true),
				newNSProjectWithLabels("same-name", "application", map[string]string{"type": "domain"}, false),
			},
			managed: &nsidentityv1alpha1.ProjectV3{
				ObjectMeta: metav1.ObjectMeta{Name: "my-project", Namespace: "application"},
				Spec: nsidentityv1alpha1.ProjectV3Spec{
					ForProvider: nsidentityv1alpha1.ProjectV3Parameters{
						Name: ptr("my-project"),
						DomainIDSelector: &xpv1.NamespacedSelector{
							MatchLabels: map[string]string{"type": "domain"},
							Namespace:   "infrastructure",
						},
					},
				},
			},
		},
		{
			name: "SelectorCrossNS_ProjectInTargetNS_Error",
			existing: []runtime.Object{
				newNSProjectWithLabels("infra-project", "infrastructure", map[string]string{"type": "domain"}, false),
			},
			managed: &nsidentityv1alpha1.ProjectV3{
				ObjectMeta: metav1.ObjectMeta{Name: "my-project", Namespace: "application"},
				Spec: nsidentityv1alpha1.ProjectV3Spec{
					ForProvider: nsidentityv1alpha1.ProjectV3Parameters{
						Name: ptr("my-project"),
						DomainIDSelector: &xpv1.NamespacedSelector{
							MatchLabels: map[string]string{"type": "domain"},
							Namespace:   "infrastructure",
						},
					},
				},
			},
			wantErr:   true,
			errSubstr: "is a project, not a domain",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scheme := newNSScheme(t)
			builder := fake.NewClientBuilder().WithScheme(scheme)
			if len(tt.existing) > 0 {
				builder = builder.WithRuntimeObjects(tt.existing...)
			}
			c := builder.Build()

			v := NewValidator(c)
			err := v.Initialize(context.Background(), tt.managed)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if !strings.Contains(err.Error(), tt.errSubstr) {
					t.Errorf("error %q does not contain %q", err.Error(), tt.errSubstr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func newNSProjectWithLabels(name, ns string, labels map[string]string, isDomain bool) *nsidentityv1alpha1.ProjectV3 {
	p := &nsidentityv1alpha1.ProjectV3{
		ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: ns, Labels: labels},
		Spec: nsidentityv1alpha1.ProjectV3Spec{
			ForProvider: nsidentityv1alpha1.ProjectV3Parameters{
				Name: ptr(name),
			},
		},
	}
	if isDomain {
		p.Spec.ForProvider.IsDomain = ptr(true)
	}
	return p
}
