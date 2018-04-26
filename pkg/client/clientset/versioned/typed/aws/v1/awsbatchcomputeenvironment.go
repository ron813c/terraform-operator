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

// AwsBatchComputeEnvironmentsGetter has a method to return a AwsBatchComputeEnvironmentInterface.
// A group's client should implement this interface.
type AwsBatchComputeEnvironmentsGetter interface {
	AwsBatchComputeEnvironments(namespace string) AwsBatchComputeEnvironmentInterface
}

// AwsBatchComputeEnvironmentInterface has methods to work with AwsBatchComputeEnvironment resources.
type AwsBatchComputeEnvironmentInterface interface {
	Create(*v1.AwsBatchComputeEnvironment) (*v1.AwsBatchComputeEnvironment, error)
	Update(*v1.AwsBatchComputeEnvironment) (*v1.AwsBatchComputeEnvironment, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsBatchComputeEnvironment, error)
	List(opts meta_v1.ListOptions) (*v1.AwsBatchComputeEnvironmentList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsBatchComputeEnvironment, err error)
	AwsBatchComputeEnvironmentExpansion
}

// awsBatchComputeEnvironments implements AwsBatchComputeEnvironmentInterface
type awsBatchComputeEnvironments struct {
	client rest.Interface
	ns     string
}

// newAwsBatchComputeEnvironments returns a AwsBatchComputeEnvironments
func newAwsBatchComputeEnvironments(c *TrussleV1Client, namespace string) *awsBatchComputeEnvironments {
	return &awsBatchComputeEnvironments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsBatchComputeEnvironment, and returns the corresponding awsBatchComputeEnvironment object, and an error if there is any.
func (c *awsBatchComputeEnvironments) Get(name string, options meta_v1.GetOptions) (result *v1.AwsBatchComputeEnvironment, err error) {
	result = &v1.AwsBatchComputeEnvironment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsbatchcomputeenvironments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsBatchComputeEnvironments that match those selectors.
func (c *awsBatchComputeEnvironments) List(opts meta_v1.ListOptions) (result *v1.AwsBatchComputeEnvironmentList, err error) {
	result = &v1.AwsBatchComputeEnvironmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsbatchcomputeenvironments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsBatchComputeEnvironments.
func (c *awsBatchComputeEnvironments) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsbatchcomputeenvironments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsBatchComputeEnvironment and creates it.  Returns the server's representation of the awsBatchComputeEnvironment, and an error, if there is any.
func (c *awsBatchComputeEnvironments) Create(awsBatchComputeEnvironment *v1.AwsBatchComputeEnvironment) (result *v1.AwsBatchComputeEnvironment, err error) {
	result = &v1.AwsBatchComputeEnvironment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsbatchcomputeenvironments").
		Body(awsBatchComputeEnvironment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsBatchComputeEnvironment and updates it. Returns the server's representation of the awsBatchComputeEnvironment, and an error, if there is any.
func (c *awsBatchComputeEnvironments) Update(awsBatchComputeEnvironment *v1.AwsBatchComputeEnvironment) (result *v1.AwsBatchComputeEnvironment, err error) {
	result = &v1.AwsBatchComputeEnvironment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsbatchcomputeenvironments").
		Name(awsBatchComputeEnvironment.Name).
		Body(awsBatchComputeEnvironment).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsBatchComputeEnvironment and deletes it. Returns an error if one occurs.
func (c *awsBatchComputeEnvironments) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsbatchcomputeenvironments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsBatchComputeEnvironments) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsbatchcomputeenvironments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsBatchComputeEnvironment.
func (c *awsBatchComputeEnvironments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsBatchComputeEnvironment, err error) {
	result = &v1.AwsBatchComputeEnvironment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsbatchcomputeenvironments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}