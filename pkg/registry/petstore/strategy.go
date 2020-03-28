package petstore

import (
	"context"
	"fmt"

	"github.com/vaibhavk/petstore-apiserver/pkg/apis/petstore/validation"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"

	"github.com/vaibhavk/petstore-apiserver/pkg/apis/petstore"
)

// NewStrategy creates and returns a petstoreStrategy instance
func NewStrategy(typer runtime.ObjectTyper) petstoreStrategy {
	return petstoreStrategy{typer, names.SimpleNameGenerator}
}

// GetAttrs returns labels.Set, fields.Set, the presence of Initializers if any
// and error in case the given runtime.Object is not a PetStore
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	apiserver, ok := obj.(*petstore.PetStore)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a PetStore")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), SelectableFields(apiserver), nil
}

// MatchPetStore is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchPetStore(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// SelectableFields returns a field set that represents the object.
func SelectableFields(obj *petstore.PetStore) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

type petstoreStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (petstoreStrategy) NamespaceScoped() bool {
	return true
}

func (petstoreStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (petstoreStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

func (petstoreStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	petstore := obj.(*petstore.PetStore)
	return validation.ValidatePetStore(petstore)
}

func (petstoreStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (petstoreStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (petstoreStrategy) Canonicalize(obj runtime.Object) {
}

func (petstoreStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
