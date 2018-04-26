/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/trussle/terraform-operator/pkg/apis/aws/v1"
	scheme "github.com/trussle/terraform-operator/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AwsApiGatewayStagesGetter has a method to return a AwsApiGatewayStageInterface.
// A group's client should implement this interface.
type AwsApiGatewayStagesGetter interface {
	AwsApiGatewayStages(namespace string) AwsApiGatewayStageInterface
}

// AwsApiGatewayStageInterface has methods to work with AwsApiGatewayStage resources.
type AwsApiGatewayStageInterface interface {
	Create(*v1.AwsApiGatewayStage) (*v1.AwsApiGatewayStage, error)
	Update(*v1.AwsApiGatewayStage) (*v1.AwsApiGatewayStage, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsApiGatewayStage, error)
	List(opts meta_v1.ListOptions) (*v1.AwsApiGatewayStageList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsApiGatewayStage, err error)
	AwsApiGatewayStageExpansion
}

// awsApiGatewayStages implements AwsApiGatewayStageInterface
type awsApiGatewayStages struct {
	client rest.Interface
	ns     string
}

// newAwsApiGatewayStages returns a AwsApiGatewayStages
func newAwsApiGatewayStages(c *TrussleV1Client, namespace string) *awsApiGatewayStages {
	return &awsApiGatewayStages{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsApiGatewayStage, and returns the corresponding awsApiGatewayStage object, and an error if there is any.
func (c *awsApiGatewayStages) Get(name string, options meta_v1.GetOptions) (result *v1.AwsApiGatewayStage, err error) {
	result = &v1.AwsApiGatewayStage{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewaystages").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsApiGatewayStages that match those selectors.
func (c *awsApiGatewayStages) List(opts meta_v1.ListOptions) (result *v1.AwsApiGatewayStageList, err error) {
	result = &v1.AwsApiGatewayStageList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewaystages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayStages.
func (c *awsApiGatewayStages) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsapigatewaystages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsApiGatewayStage and creates it.  Returns the server's representation of the awsApiGatewayStage, and an error, if there is any.
func (c *awsApiGatewayStages) Create(awsApiGatewayStage *v1.AwsApiGatewayStage) (result *v1.AwsApiGatewayStage, err error) {
	result = &v1.AwsApiGatewayStage{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsapigatewaystages").
		Body(awsApiGatewayStage).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsApiGatewayStage and updates it. Returns the server's representation of the awsApiGatewayStage, and an error, if there is any.
func (c *awsApiGatewayStages) Update(awsApiGatewayStage *v1.AwsApiGatewayStage) (result *v1.AwsApiGatewayStage, err error) {
	result = &v1.AwsApiGatewayStage{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsapigatewaystages").
		Name(awsApiGatewayStage.Name).
		Body(awsApiGatewayStage).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsApiGatewayStage and deletes it. Returns an error if one occurs.
func (c *awsApiGatewayStages) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewaystages").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsApiGatewayStages) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsapigatewaystages").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsApiGatewayStage.
func (c *awsApiGatewayStages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsApiGatewayStage, err error) {
	result = &v1.AwsApiGatewayStage{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsapigatewaystages").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}