
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamRole describes a AwsIamRole resource
type AwsIamRole struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsIamRoleSpec	`json:"spec"`
}


// AwsIamRoleSpec is the spec for a AwsIamRole Resource
type AwsIamRoleSpec struct {
	MaxSessionDuration	int	`json:"max_session_duration"`
	Path	string	`json:"path"`
	AssumeRolePolicy	string	`json:"assume_role_policy"`
	Description	string	`json:"description"`
	ForceDetachPolicies	bool	`json:"force_detach_policies"`
	NamePrefix	string	`json:"name_prefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamRoleList is a list of AwsIamRole resources
type AwsIamRoleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsIamRole	`json:"items"`
}

