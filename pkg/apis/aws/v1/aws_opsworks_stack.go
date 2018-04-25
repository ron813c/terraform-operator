
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksStack describes a AwsOpsworksStack resource
type AwsOpsworksStack struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsOpsworksStackSpec	`json:"spec"`
}


// AwsOpsworksStackSpec is the spec for a AwsOpsworksStack Resource
type AwsOpsworksStackSpec struct {
	ServiceRoleArn	string	`json:"service_role_arn"`
	DefaultInstanceProfileArn	string	`json:"default_instance_profile_arn"`
	Color	string	`json:"color"`
	ConfigurationManagerVersion	string	`json:"configuration_manager_version"`
	ManageBerkshelf	bool	`json:"manage_berkshelf"`
	DefaultRootDeviceType	string	`json:"default_root_device_type"`
	Name	string	`json:"name"`
	Tags	map[string]interface{}	`json:"tags"`
	DefaultSshKeyName	string	`json:"default_ssh_key_name"`
	UseCustomCookbooks	bool	`json:"use_custom_cookbooks"`
	UseOpsworksSecurityGroups	bool	`json:"use_opsworks_security_groups"`
	ConfigurationManagerName	string	`json:"configuration_manager_name"`
	DefaultOs	string	`json:"default_os"`
	HostnameTheme	string	`json:"hostname_theme"`
	Region	string	`json:"region"`
	BerkshelfVersion	string	`json:"berkshelf_version"`
	CustomJson	string	`json:"custom_json"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsOpsworksStackList is a list of AwsOpsworksStack resources
type AwsOpsworksStackList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsOpsworksStack	`json:"items"`
}

