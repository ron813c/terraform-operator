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

// AwsCloudwatchLogDestinationPoliciesGetter has a method to return a AwsCloudwatchLogDestinationPolicyInterface.
// A group's client should implement this interface.
type AwsCloudwatchLogDestinationPoliciesGetter interface {
	AwsCloudwatchLogDestinationPolicies(namespace string) AwsCloudwatchLogDestinationPolicyInterface
}

// AwsCloudwatchLogDestinationPolicyInterface has methods to work with AwsCloudwatchLogDestinationPolicy resources.
type AwsCloudwatchLogDestinationPolicyInterface interface {
	Create(*v1.AwsCloudwatchLogDestinationPolicy) (*v1.AwsCloudwatchLogDestinationPolicy, error)
	Update(*v1.AwsCloudwatchLogDestinationPolicy) (*v1.AwsCloudwatchLogDestinationPolicy, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsCloudwatchLogDestinationPolicy, error)
	List(opts meta_v1.ListOptions) (*v1.AwsCloudwatchLogDestinationPolicyList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsCloudwatchLogDestinationPolicy, err error)
	AwsCloudwatchLogDestinationPolicyExpansion
}

// awsCloudwatchLogDestinationPolicies implements AwsCloudwatchLogDestinationPolicyInterface
type awsCloudwatchLogDestinationPolicies struct {
	client rest.Interface
	ns     string
}

// newAwsCloudwatchLogDestinationPolicies returns a AwsCloudwatchLogDestinationPolicies
func newAwsCloudwatchLogDestinationPolicies(c *TrussleV1Client, namespace string) *awsCloudwatchLogDestinationPolicies {
	return &awsCloudwatchLogDestinationPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsCloudwatchLogDestinationPolicy, and returns the corresponding awsCloudwatchLogDestinationPolicy object, and an error if there is any.
func (c *awsCloudwatchLogDestinationPolicies) Get(name string, options meta_v1.GetOptions) (result *v1.AwsCloudwatchLogDestinationPolicy, err error) {
	result = &v1.AwsCloudwatchLogDestinationPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscloudwatchlogdestinationpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsCloudwatchLogDestinationPolicies that match those selectors.
func (c *awsCloudwatchLogDestinationPolicies) List(opts meta_v1.ListOptions) (result *v1.AwsCloudwatchLogDestinationPolicyList, err error) {
	result = &v1.AwsCloudwatchLogDestinationPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscloudwatchlogdestinationpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsCloudwatchLogDestinationPolicies.
func (c *awsCloudwatchLogDestinationPolicies) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awscloudwatchlogdestinationpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsCloudwatchLogDestinationPolicy and creates it.  Returns the server's representation of the awsCloudwatchLogDestinationPolicy, and an error, if there is any.
func (c *awsCloudwatchLogDestinationPolicies) Create(awsCloudwatchLogDestinationPolicy *v1.AwsCloudwatchLogDestinationPolicy) (result *v1.AwsCloudwatchLogDestinationPolicy, err error) {
	result = &v1.AwsCloudwatchLogDestinationPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awscloudwatchlogdestinationpolicies").
		Body(awsCloudwatchLogDestinationPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsCloudwatchLogDestinationPolicy and updates it. Returns the server's representation of the awsCloudwatchLogDestinationPolicy, and an error, if there is any.
func (c *awsCloudwatchLogDestinationPolicies) Update(awsCloudwatchLogDestinationPolicy *v1.AwsCloudwatchLogDestinationPolicy) (result *v1.AwsCloudwatchLogDestinationPolicy, err error) {
	result = &v1.AwsCloudwatchLogDestinationPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awscloudwatchlogdestinationpolicies").
		Name(awsCloudwatchLogDestinationPolicy.Name).
		Body(awsCloudwatchLogDestinationPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsCloudwatchLogDestinationPolicy and deletes it. Returns an error if one occurs.
func (c *awsCloudwatchLogDestinationPolicies) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscloudwatchlogdestinationpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsCloudwatchLogDestinationPolicies) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscloudwatchlogdestinationpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsCloudwatchLogDestinationPolicy.
func (c *awsCloudwatchLogDestinationPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsCloudwatchLogDestinationPolicy, err error) {
	result = &v1.AwsCloudwatchLogDestinationPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awscloudwatchlogdestinationpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
