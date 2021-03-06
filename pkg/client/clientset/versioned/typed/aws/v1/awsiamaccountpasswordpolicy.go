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

// AwsIamAccountPasswordPoliciesGetter has a method to return a AwsIamAccountPasswordPolicyInterface.
// A group's client should implement this interface.
type AwsIamAccountPasswordPoliciesGetter interface {
	AwsIamAccountPasswordPolicies(namespace string) AwsIamAccountPasswordPolicyInterface
}

// AwsIamAccountPasswordPolicyInterface has methods to work with AwsIamAccountPasswordPolicy resources.
type AwsIamAccountPasswordPolicyInterface interface {
	Create(*v1.AwsIamAccountPasswordPolicy) (*v1.AwsIamAccountPasswordPolicy, error)
	Update(*v1.AwsIamAccountPasswordPolicy) (*v1.AwsIamAccountPasswordPolicy, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsIamAccountPasswordPolicy, error)
	List(opts meta_v1.ListOptions) (*v1.AwsIamAccountPasswordPolicyList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsIamAccountPasswordPolicy, err error)
	AwsIamAccountPasswordPolicyExpansion
}

// awsIamAccountPasswordPolicies implements AwsIamAccountPasswordPolicyInterface
type awsIamAccountPasswordPolicies struct {
	client rest.Interface
	ns     string
}

// newAwsIamAccountPasswordPolicies returns a AwsIamAccountPasswordPolicies
func newAwsIamAccountPasswordPolicies(c *TrussleV1Client, namespace string) *awsIamAccountPasswordPolicies {
	return &awsIamAccountPasswordPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsIamAccountPasswordPolicy, and returns the corresponding awsIamAccountPasswordPolicy object, and an error if there is any.
func (c *awsIamAccountPasswordPolicies) Get(name string, options meta_v1.GetOptions) (result *v1.AwsIamAccountPasswordPolicy, err error) {
	result = &v1.AwsIamAccountPasswordPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiamaccountpasswordpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsIamAccountPasswordPolicies that match those selectors.
func (c *awsIamAccountPasswordPolicies) List(opts meta_v1.ListOptions) (result *v1.AwsIamAccountPasswordPolicyList, err error) {
	result = &v1.AwsIamAccountPasswordPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsiamaccountpasswordpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsIamAccountPasswordPolicies.
func (c *awsIamAccountPasswordPolicies) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsiamaccountpasswordpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsIamAccountPasswordPolicy and creates it.  Returns the server's representation of the awsIamAccountPasswordPolicy, and an error, if there is any.
func (c *awsIamAccountPasswordPolicies) Create(awsIamAccountPasswordPolicy *v1.AwsIamAccountPasswordPolicy) (result *v1.AwsIamAccountPasswordPolicy, err error) {
	result = &v1.AwsIamAccountPasswordPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsiamaccountpasswordpolicies").
		Body(awsIamAccountPasswordPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsIamAccountPasswordPolicy and updates it. Returns the server's representation of the awsIamAccountPasswordPolicy, and an error, if there is any.
func (c *awsIamAccountPasswordPolicies) Update(awsIamAccountPasswordPolicy *v1.AwsIamAccountPasswordPolicy) (result *v1.AwsIamAccountPasswordPolicy, err error) {
	result = &v1.AwsIamAccountPasswordPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsiamaccountpasswordpolicies").
		Name(awsIamAccountPasswordPolicy.Name).
		Body(awsIamAccountPasswordPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsIamAccountPasswordPolicy and deletes it. Returns an error if one occurs.
func (c *awsIamAccountPasswordPolicies) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiamaccountpasswordpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsIamAccountPasswordPolicies) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsiamaccountpasswordpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsIamAccountPasswordPolicy.
func (c *awsIamAccountPasswordPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsIamAccountPasswordPolicy, err error) {
	result = &v1.AwsIamAccountPasswordPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsiamaccountpasswordpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
