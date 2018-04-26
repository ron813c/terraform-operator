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

// AwsSesReceiptRuleSetsGetter has a method to return a AwsSesReceiptRuleSetInterface.
// A group's client should implement this interface.
type AwsSesReceiptRuleSetsGetter interface {
	AwsSesReceiptRuleSets(namespace string) AwsSesReceiptRuleSetInterface
}

// AwsSesReceiptRuleSetInterface has methods to work with AwsSesReceiptRuleSet resources.
type AwsSesReceiptRuleSetInterface interface {
	Create(*v1.AwsSesReceiptRuleSet) (*v1.AwsSesReceiptRuleSet, error)
	Update(*v1.AwsSesReceiptRuleSet) (*v1.AwsSesReceiptRuleSet, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsSesReceiptRuleSet, error)
	List(opts meta_v1.ListOptions) (*v1.AwsSesReceiptRuleSetList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsSesReceiptRuleSet, err error)
	AwsSesReceiptRuleSetExpansion
}

// awsSesReceiptRuleSets implements AwsSesReceiptRuleSetInterface
type awsSesReceiptRuleSets struct {
	client rest.Interface
	ns     string
}

// newAwsSesReceiptRuleSets returns a AwsSesReceiptRuleSets
func newAwsSesReceiptRuleSets(c *TrussleV1Client, namespace string) *awsSesReceiptRuleSets {
	return &awsSesReceiptRuleSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsSesReceiptRuleSet, and returns the corresponding awsSesReceiptRuleSet object, and an error if there is any.
func (c *awsSesReceiptRuleSets) Get(name string, options meta_v1.GetOptions) (result *v1.AwsSesReceiptRuleSet, err error) {
	result = &v1.AwsSesReceiptRuleSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssesreceiptrulesets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsSesReceiptRuleSets that match those selectors.
func (c *awsSesReceiptRuleSets) List(opts meta_v1.ListOptions) (result *v1.AwsSesReceiptRuleSetList, err error) {
	result = &v1.AwsSesReceiptRuleSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awssesreceiptrulesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsSesReceiptRuleSets.
func (c *awsSesReceiptRuleSets) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awssesreceiptrulesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsSesReceiptRuleSet and creates it.  Returns the server's representation of the awsSesReceiptRuleSet, and an error, if there is any.
func (c *awsSesReceiptRuleSets) Create(awsSesReceiptRuleSet *v1.AwsSesReceiptRuleSet) (result *v1.AwsSesReceiptRuleSet, err error) {
	result = &v1.AwsSesReceiptRuleSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awssesreceiptrulesets").
		Body(awsSesReceiptRuleSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsSesReceiptRuleSet and updates it. Returns the server's representation of the awsSesReceiptRuleSet, and an error, if there is any.
func (c *awsSesReceiptRuleSets) Update(awsSesReceiptRuleSet *v1.AwsSesReceiptRuleSet) (result *v1.AwsSesReceiptRuleSet, err error) {
	result = &v1.AwsSesReceiptRuleSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awssesreceiptrulesets").
		Name(awsSesReceiptRuleSet.Name).
		Body(awsSesReceiptRuleSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsSesReceiptRuleSet and deletes it. Returns an error if one occurs.
func (c *awsSesReceiptRuleSets) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssesreceiptrulesets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsSesReceiptRuleSets) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awssesreceiptrulesets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsSesReceiptRuleSet.
func (c *awsSesReceiptRuleSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsSesReceiptRuleSet, err error) {
	result = &v1.AwsSesReceiptRuleSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awssesreceiptrulesets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}