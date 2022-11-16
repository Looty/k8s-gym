/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type Check struct {
	// +kubebuilder:validation:MaxLength=36
	// +kubebuilder:validation:MinLength=3
	Name string `json:"name"`

	// +kubebuilder:validation:MinLength=3
	Cmd string `json:"cmd"`

	ExpectedResult string `json:"expectedResult"`
}

// QuestionSpec defines the desired state of Question
type QuestionSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:MaxLength=256
	// +kubebuilder:validation:MinLength=3
	Description string `json:"description"`

	// +kubebuilder:validation:MaxLength=256
	// +kubebuilder:validation:MinLength=3
	Tip string `json:"tip"`

	// +kubebuilder:validation:MaxLength=512
	// +kubebuilder:validation:MinLength=1
	Answer           string  `json:"answer"`
	ValidationChecks []Check `json:"validationChecks"`
}

// QuestionStatus defines the observed state of Question
type QuestionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:printcolumn:name="Description",type=string,JSONPath=`.spec.description`

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// Question is the Schema for the questions API
type Question struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   QuestionSpec   `json:"spec,omitempty"`
	Status QuestionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// QuestionList contains a list of Question
type QuestionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Question `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Question{}, &QuestionList{})
}
