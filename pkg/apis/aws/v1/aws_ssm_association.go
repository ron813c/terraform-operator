
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmAssociation describes a AwsSsmAssociation resource
type AwsSsmAssociation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSsmAssociationSpec	`json:"spec"`
}


// AwsSsmAssociationSpec is the spec for a AwsSsmAssociation Resource
type AwsSsmAssociationSpec struct {
	Name	string	`json:"name"`
	ScheduleExpression	string	`json:"schedule_expression"`
	OutputLocation	[]vATdjZyE	`json:"output_location"`
	AssociationName	string	`json:"association_name"`
	InstanceId	string	`json:"instance_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmAssociationList is a list of AwsSsmAssociation resources
type AwsSsmAssociationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmAssociation	`json:"items"`
}


// iJiWmAPK is a iJiWmAPK
type iJiWmAPK struct {
	Key	string	`json:"key"`
	Values	[]string	`json:"values"`
}

// vATdjZyE is a vATdjZyE
type vATdjZyE struct {
	S3BucketName	string	`json:"s3_bucket_name"`
	S3KeyPrefix	string	`json:"s3_key_prefix"`
}
