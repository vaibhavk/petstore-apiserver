package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PetStore specifies an offered pizza with toppings.
type PetStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   PetStoreSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status PetStoreStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type PetStoreSpec struct {
	Petlist []Pet `json:"petlist" protobuf:"bytes,1,rep,name=petlist"`
}

type Pet struct {
	Name  string `json:"name" protobuf:"bytes,1,rep,name=name"`
	Count int64  `json:"count" protobuf:"bytes,2,rep,name=count"`
}

type PetStoreStatus struct {
	// cost is the cost of the whole pizza including all toppings.
	Population int64 `json:"population,omitempty" protobuf:"bytes,1,opt,name=cost"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PetStoreList is a list of PetStore objects.
type PetStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []PetStore `json:"items" protobuf:"bytes,2,rep,name=items"`
}
