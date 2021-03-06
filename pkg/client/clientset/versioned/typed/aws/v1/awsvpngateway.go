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

// AwsVpnGatewaysGetter has a method to return a AwsVpnGatewayInterface.
// A group's client should implement this interface.
type AwsVpnGatewaysGetter interface {
	AwsVpnGateways(namespace string) AwsVpnGatewayInterface
}

// AwsVpnGatewayInterface has methods to work with AwsVpnGateway resources.
type AwsVpnGatewayInterface interface {
	Create(*v1.AwsVpnGateway) (*v1.AwsVpnGateway, error)
	Update(*v1.AwsVpnGateway) (*v1.AwsVpnGateway, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsVpnGateway, error)
	List(opts meta_v1.ListOptions) (*v1.AwsVpnGatewayList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsVpnGateway, err error)
	AwsVpnGatewayExpansion
}

// awsVpnGateways implements AwsVpnGatewayInterface
type awsVpnGateways struct {
	client rest.Interface
	ns     string
}

// newAwsVpnGateways returns a AwsVpnGateways
func newAwsVpnGateways(c *TrussleV1Client, namespace string) *awsVpnGateways {
	return &awsVpnGateways{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsVpnGateway, and returns the corresponding awsVpnGateway object, and an error if there is any.
func (c *awsVpnGateways) Get(name string, options meta_v1.GetOptions) (result *v1.AwsVpnGateway, err error) {
	result = &v1.AwsVpnGateway{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsvpngateways").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsVpnGateways that match those selectors.
func (c *awsVpnGateways) List(opts meta_v1.ListOptions) (result *v1.AwsVpnGatewayList, err error) {
	result = &v1.AwsVpnGatewayList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsvpngateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsVpnGateways.
func (c *awsVpnGateways) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsvpngateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsVpnGateway and creates it.  Returns the server's representation of the awsVpnGateway, and an error, if there is any.
func (c *awsVpnGateways) Create(awsVpnGateway *v1.AwsVpnGateway) (result *v1.AwsVpnGateway, err error) {
	result = &v1.AwsVpnGateway{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsvpngateways").
		Body(awsVpnGateway).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsVpnGateway and updates it. Returns the server's representation of the awsVpnGateway, and an error, if there is any.
func (c *awsVpnGateways) Update(awsVpnGateway *v1.AwsVpnGateway) (result *v1.AwsVpnGateway, err error) {
	result = &v1.AwsVpnGateway{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsvpngateways").
		Name(awsVpnGateway.Name).
		Body(awsVpnGateway).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsVpnGateway and deletes it. Returns an error if one occurs.
func (c *awsVpnGateways) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsvpngateways").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsVpnGateways) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsvpngateways").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsVpnGateway.
func (c *awsVpnGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsVpnGateway, err error) {
	result = &v1.AwsVpnGateway{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsvpngateways").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
