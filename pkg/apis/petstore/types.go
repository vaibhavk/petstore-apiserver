package petstore

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PetStore specifies an offered pizza with toppings.
type PetStore struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   PetStoreSpec
	Status PetStoreStatus
}

type PetStoreSpec struct {
	// toppings is a list of Topping names. They don't have to be unique. Order does not matter.
	PetList []Pet
}

type Pet struct {
	// name is the name of a Topping object .
	Name string
	// quantity is the number of how often the topping is put onto the pizza.
	Quantity int
}

type PetStoreStatus struct {
	Population int64
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PetStoreList is a list of PetStore objects.
type PetStoreList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []PetStore
}
