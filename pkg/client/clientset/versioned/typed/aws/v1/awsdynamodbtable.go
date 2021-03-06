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

// AwsDynamodbTablesGetter has a method to return a AwsDynamodbTableInterface.
// A group's client should implement this interface.
type AwsDynamodbTablesGetter interface {
	AwsDynamodbTables(namespace string) AwsDynamodbTableInterface
}

// AwsDynamodbTableInterface has methods to work with AwsDynamodbTable resources.
type AwsDynamodbTableInterface interface {
	Create(*v1.AwsDynamodbTable) (*v1.AwsDynamodbTable, error)
	Update(*v1.AwsDynamodbTable) (*v1.AwsDynamodbTable, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsDynamodbTable, error)
	List(opts meta_v1.ListOptions) (*v1.AwsDynamodbTableList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsDynamodbTable, err error)
	AwsDynamodbTableExpansion
}

// awsDynamodbTables implements AwsDynamodbTableInterface
type awsDynamodbTables struct {
	client rest.Interface
	ns     string
}

// newAwsDynamodbTables returns a AwsDynamodbTables
func newAwsDynamodbTables(c *TrussleV1Client, namespace string) *awsDynamodbTables {
	return &awsDynamodbTables{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsDynamodbTable, and returns the corresponding awsDynamodbTable object, and an error if there is any.
func (c *awsDynamodbTables) Get(name string, options meta_v1.GetOptions) (result *v1.AwsDynamodbTable, err error) {
	result = &v1.AwsDynamodbTable{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdynamodbtables").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDynamodbTables that match those selectors.
func (c *awsDynamodbTables) List(opts meta_v1.ListOptions) (result *v1.AwsDynamodbTableList, err error) {
	result = &v1.AwsDynamodbTableList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdynamodbtables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDynamodbTables.
func (c *awsDynamodbTables) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsdynamodbtables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsDynamodbTable and creates it.  Returns the server's representation of the awsDynamodbTable, and an error, if there is any.
func (c *awsDynamodbTables) Create(awsDynamodbTable *v1.AwsDynamodbTable) (result *v1.AwsDynamodbTable, err error) {
	result = &v1.AwsDynamodbTable{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsdynamodbtables").
		Body(awsDynamodbTable).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDynamodbTable and updates it. Returns the server's representation of the awsDynamodbTable, and an error, if there is any.
func (c *awsDynamodbTables) Update(awsDynamodbTable *v1.AwsDynamodbTable) (result *v1.AwsDynamodbTable, err error) {
	result = &v1.AwsDynamodbTable{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsdynamodbtables").
		Name(awsDynamodbTable.Name).
		Body(awsDynamodbTable).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDynamodbTable and deletes it. Returns an error if one occurs.
func (c *awsDynamodbTables) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdynamodbtables").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDynamodbTables) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdynamodbtables").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDynamodbTable.
func (c *awsDynamodbTables) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsDynamodbTable, err error) {
	result = &v1.AwsDynamodbTable{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsdynamodbtables").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
