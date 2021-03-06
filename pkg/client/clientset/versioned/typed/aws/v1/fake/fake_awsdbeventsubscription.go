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

// FakeAwsDbEventSubscriptions implements AwsDbEventSubscriptionInterface
type FakeAwsDbEventSubscriptions struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsdbeventsubscriptionsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsdbeventsubscriptions"}

var awsdbeventsubscriptionsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsDbEventSubscription"}

// Get takes name of the awsDbEventSubscription, and returns the corresponding awsDbEventSubscription object, and an error if there is any.
func (c *FakeAwsDbEventSubscriptions) Get(name string, options v1.GetOptions) (result *aws_v1.AwsDbEventSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsdbeventsubscriptionsResource, c.ns, name), &aws_v1.AwsDbEventSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsDbEventSubscription), err
}

// List takes label and field selectors, and returns the list of AwsDbEventSubscriptions that match those selectors.
func (c *FakeAwsDbEventSubscriptions) List(opts v1.ListOptions) (result *aws_v1.AwsDbEventSubscriptionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsdbeventsubscriptionsResource, awsdbeventsubscriptionsKind, c.ns, opts), &aws_v1.AwsDbEventSubscriptionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsDbEventSubscriptionList{}
	for _, item := range obj.(*aws_v1.AwsDbEventSubscriptionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDbEventSubscriptions.
func (c *FakeAwsDbEventSubscriptions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsdbeventsubscriptionsResource, c.ns, opts))

}

// Create takes the representation of a awsDbEventSubscription and creates it.  Returns the server's representation of the awsDbEventSubscription, and an error, if there is any.
func (c *FakeAwsDbEventSubscriptions) Create(awsDbEventSubscription *aws_v1.AwsDbEventSubscription) (result *aws_v1.AwsDbEventSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsdbeventsubscriptionsResource, c.ns, awsDbEventSubscription), &aws_v1.AwsDbEventSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsDbEventSubscription), err
}

// Update takes the representation of a awsDbEventSubscription and updates it. Returns the server's representation of the awsDbEventSubscription, and an error, if there is any.
func (c *FakeAwsDbEventSubscriptions) Update(awsDbEventSubscription *aws_v1.AwsDbEventSubscription) (result *aws_v1.AwsDbEventSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsdbeventsubscriptionsResource, c.ns, awsDbEventSubscription), &aws_v1.AwsDbEventSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsDbEventSubscription), err
}

// Delete takes name of the awsDbEventSubscription and deletes it. Returns an error if one occurs.
func (c *FakeAwsDbEventSubscriptions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsdbeventsubscriptionsResource, c.ns, name), &aws_v1.AwsDbEventSubscription{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDbEventSubscriptions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsdbeventsubscriptionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsDbEventSubscriptionList{})
	return err
}

// Patch applies the patch and returns the patched awsDbEventSubscription.
func (c *FakeAwsDbEventSubscriptions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsDbEventSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsdbeventsubscriptionsResource, c.ns, name, data, subresources...), &aws_v1.AwsDbEventSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsDbEventSubscription), err
}
