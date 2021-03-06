
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEip describes a AwsEip resource
type AwsEip struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEipSpec	`json:"spec"`
}


// AwsEipSpec is the spec for a AwsEip Resource
type AwsEipSpec struct {
	Tags	map[string]string	`json:"tags"`
	AssociateWithPrivateIp	string	`json:"associate_with_private_ip"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEipList is a list of AwsEip resources
type AwsEipList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEip	`json:"items"`
}

