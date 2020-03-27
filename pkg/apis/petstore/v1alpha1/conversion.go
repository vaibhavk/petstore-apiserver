package v1alpha1

import (
	"github.com/vaibhavk/petstore-apiserver/pkg/apis/petstore"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
)

func addConversionFuncs(scheme *runtime.Scheme) error {
	err := scheme.AddConversionFuncs(
		Convert_v1alpha1_PizzaSpec_To_petstore_PizzaSpec,
		Convert_petstore_PizzaSpec_To_v1alpha1_PizzaSpec,
	)
	if err != nil {
		return err
	}

	return nil
}

func Convert_v1alpha1_PizzaSpec_To_petstore_PizzaSpec(in *PizzaSpec, out *petstore.PizzaSpec, s conversion.Scope) error {
	for _, top := range in.PetList {
		out.PetList = append(out.PetList, petstore.PetList{
			Name:     top.Name,
			Quantity: top.Count,
		})
	}
	return nil
}

func Convert_petstore_PizzaSpec_To_v1alpha1_PizzaSpec(in *petstore.PizzaSpec, out *PizzaSpec, s conversion.Scope) error {
	for _, top := range in.PetList {
		out.PetList = append(out.PetList, PetList{
			Name:  top.Name,
			Count: top.Quantity,
		})
	}
	return nil
}
