
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchLogMetricFilter describes a AwsCloudwatchLogMetricFilter resource
type AwsCloudwatchLogMetricFilter struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudwatchLogMetricFilterSpec	`json:"spec"`
}


// AwsCloudwatchLogMetricFilterSpec is the spec for a AwsCloudwatchLogMetricFilter Resource
type AwsCloudwatchLogMetricFilterSpec struct {
	Name	string	`json:"name"`
	Pattern	string	`json:"pattern"`
	LogGroupName	string	`json:"log_group_name"`
	MetricTransformation	[]AwsCloudwatchLogMetricFilterMetricTransformation	`json:"metric_transformation"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchLogMetricFilterList is a list of AwsCloudwatchLogMetricFilter resources
type AwsCloudwatchLogMetricFilterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchLogMetricFilter	`json:"items"`
}


// AwsCloudwatchLogMetricFilterMetricTransformation is a AwsCloudwatchLogMetricFilterMetricTransformation
type AwsCloudwatchLogMetricFilterMetricTransformation struct {
	Name	string	`json:"name"`
	Namespace	string	`json:"namespace"`
	Value	string	`json:"value"`
	DefaultValue	float64	`json:"default_value"`
}
