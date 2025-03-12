package apis

import (
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var s = runtime.NewScheme()

// GetManagedResource returns a new object and list object for
// the specified GVK. The module's type registry is used to
// obtain these objects from the GVKs.
func GetManagedResource(group, version, kind, listKind string) (xpresource.Managed, xpresource.ManagedList, error) {
	gv := schema.GroupVersion{
		Group:   group,
		Version: version,
	}
	kingGVK := gv.WithKind(kind)
	m, err := s.New(kingGVK)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to get a new API object of GVK %q from the runtime scheme", kingGVK)
	}

	listGVK := gv.WithKind(listKind)
	l, err := s.New(listGVK)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to get a new API object list of GVK %q from the runtime scheme", listGVK)
	}
	return m.(xpresource.Managed), l.(xpresource.ManagedList), nil
}

// BuildScheme builds the module's type registry using the specified
// runtime.SchemeBuilder.
func BuildScheme(sb runtime.SchemeBuilder) error {
	return errors.Wrap(sb.AddToScheme(s), "failed to register the GVKs with the runtime scheme")
}
