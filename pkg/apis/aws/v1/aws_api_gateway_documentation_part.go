
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayDocumentationPart describes a AwsApiGatewayDocumentationPart resource
type AwsApiGatewayDocumentationPart struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayDocumentationPartSpec	`json:"spec"`
}


// AwsApiGatewayDocumentationPartSpec is the spec for a AwsApiGatewayDocumentationPart Resource
type AwsApiGatewayDocumentationPartSpec struct {
	Location	[]interface{}	`json:"location"`
	Properties	string	`json:"properties"`
	RestApiId	string	`json:"rest_api_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayDocumentationPartList is a list of AwsApiGatewayDocumentationPart resources
type AwsApiGatewayDocumentationPartList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayDocumentationPart	`json:"items"`
}

