
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesConfigurationSet describes a AwsSesConfigurationSet resource
type AwsSesConfigurationSet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSesConfigurationSetSpec	`json:"spec"`
}


// AwsSesConfigurationSetSpec is the spec for a AwsSesConfigurationSet Resource
type AwsSesConfigurationSetSpec struct {
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesConfigurationSetList is a list of AwsSesConfigurationSet resources
type AwsSesConfigurationSetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesConfigurationSet	`json:"items"`
}

