
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbSubnetGroup describes a AwsDbSubnetGroup resource
type AwsDbSubnetGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDbSubnetGroupSpec	`json:"spec"`
}


// AwsDbSubnetGroupSpec is the spec for a AwsDbSubnetGroup Resource
type AwsDbSubnetGroupSpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	Description	string	`json:"description"`
	SubnetIds	string	`json:"subnet_ids"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbSubnetGroupList is a list of AwsDbSubnetGroup resources
type AwsDbSubnetGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDbSubnetGroup	`json:"items"`
}

