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

// AwsRoute53HealthChecksGetter has a method to return a AwsRoute53HealthCheckInterface.
// A group's client should implement this interface.
type AwsRoute53HealthChecksGetter interface {
	AwsRoute53HealthChecks(namespace string) AwsRoute53HealthCheckInterface
}

// AwsRoute53HealthCheckInterface has methods to work with AwsRoute53HealthCheck resources.
type AwsRoute53HealthCheckInterface interface {
	Create(*v1.AwsRoute53HealthCheck) (*v1.AwsRoute53HealthCheck, error)
	Update(*v1.AwsRoute53HealthCheck) (*v1.AwsRoute53HealthCheck, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsRoute53HealthCheck, error)
	List(opts meta_v1.ListOptions) (*v1.AwsRoute53HealthCheckList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsRoute53HealthCheck, err error)
	AwsRoute53HealthCheckExpansion
}

// awsRoute53HealthChecks implements AwsRoute53HealthCheckInterface
type awsRoute53HealthChecks struct {
	client rest.Interface
	ns     string
}

// newAwsRoute53HealthChecks returns a AwsRoute53HealthChecks
func newAwsRoute53HealthChecks(c *TrussleV1Client, namespace string) *awsRoute53HealthChecks {
	return &awsRoute53HealthChecks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsRoute53HealthCheck, and returns the corresponding awsRoute53HealthCheck object, and an error if there is any.
func (c *awsRoute53HealthChecks) Get(name string, options meta_v1.GetOptions) (result *v1.AwsRoute53HealthCheck, err error) {
	result = &v1.AwsRoute53HealthCheck{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsroute53healthchecks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsRoute53HealthChecks that match those selectors.
func (c *awsRoute53HealthChecks) List(opts meta_v1.ListOptions) (result *v1.AwsRoute53HealthCheckList, err error) {
	result = &v1.AwsRoute53HealthCheckList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsroute53healthchecks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsRoute53HealthChecks.
func (c *awsRoute53HealthChecks) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsroute53healthchecks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsRoute53HealthCheck and creates it.  Returns the server's representation of the awsRoute53HealthCheck, and an error, if there is any.
func (c *awsRoute53HealthChecks) Create(awsRoute53HealthCheck *v1.AwsRoute53HealthCheck) (result *v1.AwsRoute53HealthCheck, err error) {
	result = &v1.AwsRoute53HealthCheck{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsroute53healthchecks").
		Body(awsRoute53HealthCheck).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsRoute53HealthCheck and updates it. Returns the server's representation of the awsRoute53HealthCheck, and an error, if there is any.
func (c *awsRoute53HealthChecks) Update(awsRoute53HealthCheck *v1.AwsRoute53HealthCheck) (result *v1.AwsRoute53HealthCheck, err error) {
	result = &v1.AwsRoute53HealthCheck{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsroute53healthchecks").
		Name(awsRoute53HealthCheck.Name).
		Body(awsRoute53HealthCheck).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsRoute53HealthCheck and deletes it. Returns an error if one occurs.
func (c *awsRoute53HealthChecks) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsroute53healthchecks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsRoute53HealthChecks) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsroute53healthchecks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsRoute53HealthCheck.
func (c *awsRoute53HealthChecks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsRoute53HealthCheck, err error) {
	result = &v1.AwsRoute53HealthCheck{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsroute53healthchecks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
