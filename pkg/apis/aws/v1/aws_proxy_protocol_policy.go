
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsProxyProtocolPolicy describes a AwsProxyProtocolPolicy resource
type AwsProxyProtocolPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsProxyProtocolPolicySpec	`json:"spec"`
}


// AwsProxyProtocolPolicySpec is the spec for a AwsProxyProtocolPolicy Resource
type AwsProxyProtocolPolicySpec struct {
	InstancePorts	string	`json:"instance_ports"`
	LoadBalancer	string	`json:"load_balancer"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsProxyProtocolPolicyList is a list of AwsProxyProtocolPolicy resources
type AwsProxyProtocolPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsProxyProtocolPolicy	`json:"items"`
}

