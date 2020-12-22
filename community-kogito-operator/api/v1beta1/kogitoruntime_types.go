/*


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

package v1beta1

import (
	"github.com/vaibhavjainwiz/kogito-operator/community-kogito-operator/core/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KogitoRuntimeSpec defines the desired state of KogitoRuntime
type KogitoRuntimeSpec struct {
	api.KogitoServiceSpec `json:",inline"`

	// Annotates the pods managed by the operator with the required metadata for Istio to setup its sidecars, enabling the mesh. Defaults to false.
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors.displayName="Enable Istio"
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors.x-descriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	EnableIstio bool `json:"enableIstio,omitempty"`

	// The name of the runtime used, either Quarkus or SpringBoot.
	// Default value: quarkus
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors.displayName="runtime"
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors.x-descriptors="urn:alm:descriptor:com.tectonic.ui:label"
	// +kubebuilder:validation:Enum=quarkus;springboot
	Runtime api.RuntimeType `json:"runtime,omitempty"`
}

// KogitoRuntimeStatus defines the observed state of KogitoRuntime
type KogitoRuntimeStatus struct {
	api.KogitoServiceStatus `json:",inline"`
}

// +k8s:openapi-gen=true
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KogitoRuntime is the Schema for the kogitoruntimes API
type KogitoRuntime struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KogitoRuntimeSpec   `json:"spec,omitempty"`
	Status KogitoRuntimeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KogitoRuntimeList contains a list of KogitoRuntime
type KogitoRuntimeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KogitoRuntime `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KogitoRuntime{}, &KogitoRuntimeList{})
}
