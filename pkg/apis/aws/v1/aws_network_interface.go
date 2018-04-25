
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsNetworkInterface describes a AwsNetworkInterface resource
type AwsNetworkInterface struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsNetworkInterfaceSpec	`json:"spec"`
}


// AwsNetworkInterfaceSpec is the spec for a AwsNetworkInterface Resource
type AwsNetworkInterfaceSpec struct {
	SubnetId	string	`json:"subnet_id"`
	SourceDestCheck	bool	`json:"source_dest_check"`
	Description	string	`json:"description"`
	Tags	map[string]interface{}	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsNetworkInterfaceList is a list of AwsNetworkInterface resources
type AwsNetworkInterfaceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsNetworkInterface	`json:"items"`
}

