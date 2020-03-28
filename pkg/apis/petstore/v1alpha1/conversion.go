package v1alpha1

import (
	"github.com/vaibhavk/petstore-apiserver/pkg/apis/petstore"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
)

func addConversionFuncs(scheme *runtime.Scheme) error {
	err := scheme.AddConversionFuncs(
		Convert_v1alpha1_PetStoreSpec_To_petstore_PetStoreSpec,
		Convert_petstore_PetStoreSpec_To_v1alpha1_PetStoreSpec,
	)
	if err != nil {
		return err
	}

	return nil
}

func Convert_v1alpha1_PetStoreSpec_To_petstore_PetStoreSpec(in *PetStoreSpec, out *petstore.PetStoreSpec, s conversion.Scope) error {
	for _, top := range in.PetList {
		out.PetList = append(out.PetList, petstore.Pet{
			Name:     top.Name,
			Quantity: top.Count,
		})
	}
	return nil
}

func Convert_petstore_PetStoreSpec_To_v1alpha1_PetStoreSpec(in *petstore.PetStoreSpec, out *PetStoreSpec, s conversion.Scope) error {
	for _, top := range in.PetList {
		out.PetList = append(out.PetList, Pet{
			Name:  top.Name,
			Count: top.Quantity,
		})
	}
	return nil
}
