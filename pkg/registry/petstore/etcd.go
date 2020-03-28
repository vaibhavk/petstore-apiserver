package petstore

import (
	"github.com/vaibhavk/petstore-apiserver/pkg/apis/petstore"
	"github.com/vaibhavk/petstore-apiserver/pkg/registry"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
)

// returns RESTStorage object that will work against API services
func NewREST(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (*registry.REST, error) {
	strategy := NewStrategy(scheme)
	store := &genericregistry.Store{
		NewFunc:                  func() runtime.Object { return &petstore.PetStore{} },
		NewListFunc:              func() runtime.Object { return &petstore.PetStoreList{} },
		PredicateFunc:            MatchPetStore,
		DefaultQualifiedResource: petstore.Resource("petstore"),
		CreateStrategy:           strategy,
		UpdateStrategy:           strategy,
		DeleteStrategy:           strategy,
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &registry.REST{store}, nil
}
