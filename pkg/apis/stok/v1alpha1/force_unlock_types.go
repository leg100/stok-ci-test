// Code generated by go generate; DO NOT EDIT.
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ForceUnlock is the Schema for the forceunlocks API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=forceunlocks,scope=Namespaced
// +genclient
type ForceUnlock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	CommandSpec   `json:"spec,omitempty"`
	CommandStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ForceUnlockList contains a list of ForceUnlock
type ForceUnlockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ForceUnlock `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ForceUnlock{}, &ForceUnlockList{})
}
