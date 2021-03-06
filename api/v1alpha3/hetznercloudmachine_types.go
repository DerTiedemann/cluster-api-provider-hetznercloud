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

package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// MachineFinalizer allows reconciliation loop to clean up resources associated with HetznerCloudMachine before
	// removing it from the apiserver.
	MachineFinalizer = "hetznercloudmachine.infrastructure.cluster.x-k8s.io"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HetznerCloudMachineSpec defines the desired state of HetznerCloudMachine
type HetznerCloudMachineSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of HetznerCloudMachine. Edit HetznerCloudMachine_types.go to remove/update
	Type string `json:"type"`

	// name of the ssh key to access machine
	SSHKey string `json:"sshKey"`

	// +optional
	ProviderId *string `json:"providerID,omitempty"` // WTF?
}

// HetznerCloudMachineStatus defines the observed state of HetznerCloudMachine
type HetznerCloudMachineStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// true if the corresponding cloud vm is ready
	Ready bool `json:"ready"`
	// +optional
	ProviderId      *string `json:"providerID"`
	HetznerServerId *string `json:"hetznerServerID"`
}

// +kubebuilder:subresource:status
// +kubebuilder:object:root=true

// HetznerCloudMachine is the Schema for the hetznercloudmachines API
type HetznerCloudMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HetznerCloudMachineSpec   `json:"spec,omitempty"`
	Status HetznerCloudMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HetznerCloudMachineList contains a list of HetznerCloudMachine
type HetznerCloudMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HetznerCloudMachine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HetznerCloudMachine{}, &HetznerCloudMachineList{})
}
