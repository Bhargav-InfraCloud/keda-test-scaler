/*
Copyright 2024.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TestScalerSpec defines the desired state of TestScaler
type TestScalerSpec struct {
	//+kubebuilder:validation:Minimum=0

	// Replicas count
	Replicas *int64 `json:"replicas,omitempty"`
}

// TestScalerStatus defines the observed state of TestScaler
type TestScalerStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TestScaler is the Schema for the testscalers API
type TestScaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TestScalerSpec   `json:"spec,omitempty"`
	Status TestScalerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TestScalerList contains a list of TestScaler
type TestScalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TestScaler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TestScaler{}, &TestScalerList{})
}
