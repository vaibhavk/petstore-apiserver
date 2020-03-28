package validation

import (
	"github.com/vaibhavk/petstore-apiserver/pkg/apis/petstore"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// ValidatePetStore validates a PetStore.
func ValidatePetStore(f *petstore.PetStore) field.ErrorList {
	allErrs := field.ErrorList{}

	allErrs = append(allErrs, ValidatePetStoreSpec(&f.Spec, field.NewPath("spec"))...)

	return allErrs
}

// ValidatePetStoreSpec validates a PetStoreSpec.
func ValidatePetStoreSpec(s *petstore.PetStoreSpec, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if len(s.PetList) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("petlists"), "cannot be empty"))
	}
	for i := range s.PetList {
		if s.PetList[i].Count <= 0 {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("petlists").Index(i).Child("count"), s.PetList[i].Count, "cannot be negative or zero"))
		}
	}

	return allErrs
}
