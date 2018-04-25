
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsNetworkAclRule describes a AwsNetworkAclRule resource
type AwsNetworkAclRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsNetworkAclRuleSpec	`json:"spec"`
}


// AwsNetworkAclRuleSpec is the spec for a AwsNetworkAclRule Resource
type AwsNetworkAclRuleSpec struct {
	IcmpType	string	`json:"icmp_type"`
	Protocol	string	`json:"protocol"`
	RuleAction	string	`json:"rule_action"`
	ToPort	int	`json:"to_port"`
	CidrBlock	string	`json:"cidr_block"`
	Ipv6CidrBlock	string	`json:"ipv6_cidr_block"`
	FromPort	int	`json:"from_port"`
	IcmpCode	string	`json:"icmp_code"`
	NetworkAclId	string	`json:"network_acl_id"`
	RuleNumber	int	`json:"rule_number"`
	Egress	bool	`json:"egress"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsNetworkAclRuleList is a list of AwsNetworkAclRule resources
type AwsNetworkAclRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsNetworkAclRule	`json:"items"`
}

