
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesDomainMailFrom struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSesDomainMailFromSpec	`json:"spec"`
}

type AwsSesDomainMailFromList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesDomainMailFrom	`json:"items"`
}

type AwsSesDomainMailFromSpec struct {
	Domain	string	`json:"domain"`
	MailFromDomain	string	`json:"mail_from_domain"`
	BehaviorOnMxFailure	string	`json:"behavior_on_mx_failure"`
}