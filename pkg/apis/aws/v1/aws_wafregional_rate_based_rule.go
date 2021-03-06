
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalRateBasedRule describes a AwsWafregionalRateBasedRule resource
type AwsWafregionalRateBasedRule struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsWafregionalRateBasedRuleSpec	`json:"spec"`
}


// AwsWafregionalRateBasedRuleSpec is the spec for a AwsWafregionalRateBasedRule Resource
type AwsWafregionalRateBasedRuleSpec struct {
	Name	string	`json:"name"`
	MetricName	string	`json:"metric_name"`
	Predicate	AwsWafregionalRateBasedRulePredicate	`json:"predicate"`
	RateKey	string	`json:"rate_key"`
	RateLimit	int	`json:"rate_limit"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafregionalRateBasedRuleList is a list of AwsWafregionalRateBasedRule resources
type AwsWafregionalRateBasedRuleList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsWafregionalRateBasedRule	`json:"items"`
}


// AwsWafregionalRateBasedRulePredicate is a AwsWafregionalRateBasedRulePredicate
type AwsWafregionalRateBasedRulePredicate struct {
	Negated	bool	`json:"negated"`
	DataId	string	`json:"data_id"`
	Type	string	`json:"type"`
}
