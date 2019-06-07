package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NodeinfoSpec defines the desired state of Nodeinfo
// +k8s:openapi-gen=true
type NodeinfoSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	Nodename   string `json:"nodename"`
	Macaddress string `json:"macaddress"`
}

// NodeinfoStatus defines the observed state of Nodeinfo
// +k8s:openapi-gen=true
type NodeinfoStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Nodeinfo is the Schema for the nodeinfos API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type Nodeinfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeinfoSpec   `json:"spec,omitempty"`
	Status NodeinfoStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeinfoList contains a list of Nodeinfo
type NodeinfoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Nodeinfo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Nodeinfo{}, &NodeinfoList{})
}
