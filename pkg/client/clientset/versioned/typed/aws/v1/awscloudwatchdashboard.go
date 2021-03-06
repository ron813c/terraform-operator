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

// AwsCloudwatchDashboardsGetter has a method to return a AwsCloudwatchDashboardInterface.
// A group's client should implement this interface.
type AwsCloudwatchDashboardsGetter interface {
	AwsCloudwatchDashboards(namespace string) AwsCloudwatchDashboardInterface
}

// AwsCloudwatchDashboardInterface has methods to work with AwsCloudwatchDashboard resources.
type AwsCloudwatchDashboardInterface interface {
	Create(*v1.AwsCloudwatchDashboard) (*v1.AwsCloudwatchDashboard, error)
	Update(*v1.AwsCloudwatchDashboard) (*v1.AwsCloudwatchDashboard, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsCloudwatchDashboard, error)
	List(opts meta_v1.ListOptions) (*v1.AwsCloudwatchDashboardList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsCloudwatchDashboard, err error)
	AwsCloudwatchDashboardExpansion
}

// awsCloudwatchDashboards implements AwsCloudwatchDashboardInterface
type awsCloudwatchDashboards struct {
	client rest.Interface
	ns     string
}

// newAwsCloudwatchDashboards returns a AwsCloudwatchDashboards
func newAwsCloudwatchDashboards(c *TrussleV1Client, namespace string) *awsCloudwatchDashboards {
	return &awsCloudwatchDashboards{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsCloudwatchDashboard, and returns the corresponding awsCloudwatchDashboard object, and an error if there is any.
func (c *awsCloudwatchDashboards) Get(name string, options meta_v1.GetOptions) (result *v1.AwsCloudwatchDashboard, err error) {
	result = &v1.AwsCloudwatchDashboard{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscloudwatchdashboards").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsCloudwatchDashboards that match those selectors.
func (c *awsCloudwatchDashboards) List(opts meta_v1.ListOptions) (result *v1.AwsCloudwatchDashboardList, err error) {
	result = &v1.AwsCloudwatchDashboardList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awscloudwatchdashboards").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsCloudwatchDashboards.
func (c *awsCloudwatchDashboards) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awscloudwatchdashboards").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsCloudwatchDashboard and creates it.  Returns the server's representation of the awsCloudwatchDashboard, and an error, if there is any.
func (c *awsCloudwatchDashboards) Create(awsCloudwatchDashboard *v1.AwsCloudwatchDashboard) (result *v1.AwsCloudwatchDashboard, err error) {
	result = &v1.AwsCloudwatchDashboard{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awscloudwatchdashboards").
		Body(awsCloudwatchDashboard).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsCloudwatchDashboard and updates it. Returns the server's representation of the awsCloudwatchDashboard, and an error, if there is any.
func (c *awsCloudwatchDashboards) Update(awsCloudwatchDashboard *v1.AwsCloudwatchDashboard) (result *v1.AwsCloudwatchDashboard, err error) {
	result = &v1.AwsCloudwatchDashboard{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awscloudwatchdashboards").
		Name(awsCloudwatchDashboard.Name).
		Body(awsCloudwatchDashboard).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsCloudwatchDashboard and deletes it. Returns an error if one occurs.
func (c *awsCloudwatchDashboards) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscloudwatchdashboards").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsCloudwatchDashboards) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awscloudwatchdashboards").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsCloudwatchDashboard.
func (c *awsCloudwatchDashboards) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsCloudwatchDashboard, err error) {
	result = &v1.AwsCloudwatchDashboard{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awscloudwatchdashboards").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
