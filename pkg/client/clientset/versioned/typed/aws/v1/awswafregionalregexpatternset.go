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

// AwsWafregionalRegexPatternSetsGetter has a method to return a AwsWafregionalRegexPatternSetInterface.
// A group's client should implement this interface.
type AwsWafregionalRegexPatternSetsGetter interface {
	AwsWafregionalRegexPatternSets(namespace string) AwsWafregionalRegexPatternSetInterface
}

// AwsWafregionalRegexPatternSetInterface has methods to work with AwsWafregionalRegexPatternSet resources.
type AwsWafregionalRegexPatternSetInterface interface {
	Create(*v1.AwsWafregionalRegexPatternSet) (*v1.AwsWafregionalRegexPatternSet, error)
	Update(*v1.AwsWafregionalRegexPatternSet) (*v1.AwsWafregionalRegexPatternSet, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsWafregionalRegexPatternSet, error)
	List(opts meta_v1.ListOptions) (*v1.AwsWafregionalRegexPatternSetList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsWafregionalRegexPatternSet, err error)
	AwsWafregionalRegexPatternSetExpansion
}

// awsWafregionalRegexPatternSets implements AwsWafregionalRegexPatternSetInterface
type awsWafregionalRegexPatternSets struct {
	client rest.Interface
	ns     string
}

// newAwsWafregionalRegexPatternSets returns a AwsWafregionalRegexPatternSets
func newAwsWafregionalRegexPatternSets(c *TrussleV1Client, namespace string) *awsWafregionalRegexPatternSets {
	return &awsWafregionalRegexPatternSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsWafregionalRegexPatternSet, and returns the corresponding awsWafregionalRegexPatternSet object, and an error if there is any.
func (c *awsWafregionalRegexPatternSets) Get(name string, options meta_v1.GetOptions) (result *v1.AwsWafregionalRegexPatternSet, err error) {
	result = &v1.AwsWafregionalRegexPatternSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awswafregionalregexpatternsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsWafregionalRegexPatternSets that match those selectors.
func (c *awsWafregionalRegexPatternSets) List(opts meta_v1.ListOptions) (result *v1.AwsWafregionalRegexPatternSetList, err error) {
	result = &v1.AwsWafregionalRegexPatternSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awswafregionalregexpatternsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsWafregionalRegexPatternSets.
func (c *awsWafregionalRegexPatternSets) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awswafregionalregexpatternsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsWafregionalRegexPatternSet and creates it.  Returns the server's representation of the awsWafregionalRegexPatternSet, and an error, if there is any.
func (c *awsWafregionalRegexPatternSets) Create(awsWafregionalRegexPatternSet *v1.AwsWafregionalRegexPatternSet) (result *v1.AwsWafregionalRegexPatternSet, err error) {
	result = &v1.AwsWafregionalRegexPatternSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awswafregionalregexpatternsets").
		Body(awsWafregionalRegexPatternSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsWafregionalRegexPatternSet and updates it. Returns the server's representation of the awsWafregionalRegexPatternSet, and an error, if there is any.
func (c *awsWafregionalRegexPatternSets) Update(awsWafregionalRegexPatternSet *v1.AwsWafregionalRegexPatternSet) (result *v1.AwsWafregionalRegexPatternSet, err error) {
	result = &v1.AwsWafregionalRegexPatternSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awswafregionalregexpatternsets").
		Name(awsWafregionalRegexPatternSet.Name).
		Body(awsWafregionalRegexPatternSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsWafregionalRegexPatternSet and deletes it. Returns an error if one occurs.
func (c *awsWafregionalRegexPatternSets) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awswafregionalregexpatternsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsWafregionalRegexPatternSets) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awswafregionalregexpatternsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsWafregionalRegexPatternSet.
func (c *awsWafregionalRegexPatternSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsWafregionalRegexPatternSet, err error) {
	result = &v1.AwsWafregionalRegexPatternSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awswafregionalregexpatternsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
