package validation

import (
	"github.com/vaibhavk/petstore-apiserver/pkg/apis/petstore"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// ValidatePizza validates a Pizza.
func ValidatePizza(f *petstore.Pizza) field.ErrorList {
	allErrs := field.ErrorList{}

	allErrs = append(allErrs, ValidatePizzaSpec(&f.Spec, field.NewPath("spec"))...)

	return allErrs
}

// ValidatePizzaSpec validates a PizzaSpec.
func ValidatePizzaSpec(s *petstore.PizzaSpec, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if len(s.PetLists) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("petlists"), "cannot be empty"))
	}
	for i := range s.PetLists {
		if s.PetLists[i].Quantity <= 0 {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("petlists").Index(i).Child("quantity"), s.PetLists[i].Quantity, "cannot be negative or zero"))
		}
	}

	return allErrs
}
