package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type crdtraining struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec crdtrainingSpec `json:"spec"`
}

type crdtrainingSpec struct {
	message string `json:"message"`
	count   int    `json:"count"`
}

type crdtrainingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []crdtraining `json:"items"`
}
