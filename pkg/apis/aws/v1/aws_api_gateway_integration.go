
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayIntegration describes a AwsApiGatewayIntegration resource
type AwsApiGatewayIntegration struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayIntegrationSpec	`json:"spec"`
}


// AwsApiGatewayIntegrationSpec is the spec for a AwsApiGatewayIntegration Resource
type AwsApiGatewayIntegrationSpec struct {
	RestApiId	string	`json:"rest_api_id"`
	RequestParameters	map[string]string	`json:"request_parameters"`
	ResourceId	string	`json:"resource_id"`
	HttpMethod	string	`json:"http_method"`
	ConnectionType	string	`json:"connection_type"`
	IntegrationHttpMethod	string	`json:"integration_http_method"`
	RequestParametersInJson	string	`json:"request_parameters_in_json"`
	CacheKeyParameters	string	`json:"cache_key_parameters"`
	Type	string	`json:"type"`
	ConnectionId	string	`json:"connection_id"`
	Uri	string	`json:"uri"`
	Credentials	string	`json:"credentials"`
	ContentHandling	string	`json:"content_handling"`
	RequestTemplates	map[string]string	`json:"request_templates"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayIntegrationList is a list of AwsApiGatewayIntegration resources
type AwsApiGatewayIntegrationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayIntegration	`json:"items"`
}

