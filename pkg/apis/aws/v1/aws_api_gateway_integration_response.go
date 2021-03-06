
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayIntegrationResponse describes a AwsApiGatewayIntegrationResponse resource
type AwsApiGatewayIntegrationResponse struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayIntegrationResponseSpec	`json:"spec"`
}


// AwsApiGatewayIntegrationResponseSpec is the spec for a AwsApiGatewayIntegrationResponse Resource
type AwsApiGatewayIntegrationResponseSpec struct {
	ResourceId	string	`json:"resource_id"`
	SelectionPattern	string	`json:"selection_pattern"`
	ResponseTemplates	map[string]string	`json:"response_templates"`
	ResponseParameters	map[string]string	`json:"response_parameters"`
	ContentHandling	string	`json:"content_handling"`
	RestApiId	string	`json:"rest_api_id"`
	HttpMethod	string	`json:"http_method"`
	StatusCode	string	`json:"status_code"`
	ResponseParametersInJson	string	`json:"response_parameters_in_json"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayIntegrationResponseList is a list of AwsApiGatewayIntegrationResponse resources
type AwsApiGatewayIntegrationResponseList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayIntegrationResponse	`json:"items"`
}

