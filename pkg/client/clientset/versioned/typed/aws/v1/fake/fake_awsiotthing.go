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

package fake

import (
	aws_v1 "github.com/trussle/terraform-operator/pkg/apis/aws/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAwsIotThings implements AwsIotThingInterface
type FakeAwsIotThings struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsiotthingsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsiotthings"}

var awsiotthingsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsIotThing"}

// Get takes name of the awsIotThing, and returns the corresponding awsIotThing object, and an error if there is any.
func (c *FakeAwsIotThings) Get(name string, options v1.GetOptions) (result *aws_v1.AwsIotThing, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsiotthingsResource, c.ns, name), &aws_v1.AwsIotThing{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIotThing), err
}

// List takes label and field selectors, and returns the list of AwsIotThings that match those selectors.
func (c *FakeAwsIotThings) List(opts v1.ListOptions) (result *aws_v1.AwsIotThingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsiotthingsResource, awsiotthingsKind, c.ns, opts), &aws_v1.AwsIotThingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsIotThingList{}
	for _, item := range obj.(*aws_v1.AwsIotThingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsIotThings.
func (c *FakeAwsIotThings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsiotthingsResource, c.ns, opts))

}

// Create takes the representation of a awsIotThing and creates it.  Returns the server's representation of the awsIotThing, and an error, if there is any.
func (c *FakeAwsIotThings) Create(awsIotThing *aws_v1.AwsIotThing) (result *aws_v1.AwsIotThing, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsiotthingsResource, c.ns, awsIotThing), &aws_v1.AwsIotThing{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIotThing), err
}

// Update takes the representation of a awsIotThing and updates it. Returns the server's representation of the awsIotThing, and an error, if there is any.
func (c *FakeAwsIotThings) Update(awsIotThing *aws_v1.AwsIotThing) (result *aws_v1.AwsIotThing, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsiotthingsResource, c.ns, awsIotThing), &aws_v1.AwsIotThing{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIotThing), err
}

// Delete takes name of the awsIotThing and deletes it. Returns an error if one occurs.
func (c *FakeAwsIotThings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsiotthingsResource, c.ns, name), &aws_v1.AwsIotThing{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsIotThings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsiotthingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsIotThingList{})
	return err
}

// Patch applies the patch and returns the patched awsIotThing.
func (c *FakeAwsIotThings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsIotThing, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsiotthingsResource, c.ns, name, data, subresources...), &aws_v1.AwsIotThing{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsIotThing), err
}
