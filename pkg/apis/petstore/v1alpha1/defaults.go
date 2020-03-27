package v1alpha1

func init() {
	localSchemeBuilder.Register(RegisterDefaults)
}

func SetDefaults_PetStoreSpec(obj *PetStoreSpec) {
	if len(obj.PetList) == 0 {
		obj.PetList = []Pet{{
			Name:  "default-pet",
			Count: 5,
		}}
	}
}
