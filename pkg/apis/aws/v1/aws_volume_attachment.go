
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVolumeAttachment describes a AwsVolumeAttachment resource
type AwsVolumeAttachment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVolumeAttachmentSpec	`json:"spec"`
}


// AwsVolumeAttachmentSpec is the spec for a AwsVolumeAttachment Resource
type AwsVolumeAttachmentSpec struct {
	DeviceName	string	`json:"device_name"`
	InstanceId	string	`json:"instance_id"`
	VolumeId	string	`json:"volume_id"`
	ForceDetach	bool	`json:"force_detach"`
	SkipDestroy	bool	`json:"skip_destroy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVolumeAttachmentList is a list of AwsVolumeAttachment resources
type AwsVolumeAttachmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVolumeAttachment	`json:"items"`
}

