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

// FakeAwsSpotInstanceRequests implements AwsSpotInstanceRequestInterface
type FakeAwsSpotInstanceRequests struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsspotinstancerequestsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsspotinstancerequests"}

var awsspotinstancerequestsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsSpotInstanceRequest"}

// Get takes name of the awsSpotInstanceRequest, and returns the corresponding awsSpotInstanceRequest object, and an error if there is any.
func (c *FakeAwsSpotInstanceRequests) Get(name string, options v1.GetOptions) (result *aws_v1.AwsSpotInstanceRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsspotinstancerequestsResource, c.ns, name), &aws_v1.AwsSpotInstanceRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSpotInstanceRequest), err
}

// List takes label and field selectors, and returns the list of AwsSpotInstanceRequests that match those selectors.
func (c *FakeAwsSpotInstanceRequests) List(opts v1.ListOptions) (result *aws_v1.AwsSpotInstanceRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsspotinstancerequestsResource, awsspotinstancerequestsKind, c.ns, opts), &aws_v1.AwsSpotInstanceRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsSpotInstanceRequestList{}
	for _, item := range obj.(*aws_v1.AwsSpotInstanceRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSpotInstanceRequests.
func (c *FakeAwsSpotInstanceRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsspotinstancerequestsResource, c.ns, opts))

}

// Create takes the representation of a awsSpotInstanceRequest and creates it.  Returns the server's representation of the awsSpotInstanceRequest, and an error, if there is any.
func (c *FakeAwsSpotInstanceRequests) Create(awsSpotInstanceRequest *aws_v1.AwsSpotInstanceRequest) (result *aws_v1.AwsSpotInstanceRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsspotinstancerequestsResource, c.ns, awsSpotInstanceRequest), &aws_v1.AwsSpotInstanceRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSpotInstanceRequest), err
}

// Update takes the representation of a awsSpotInstanceRequest and updates it. Returns the server's representation of the awsSpotInstanceRequest, and an error, if there is any.
func (c *FakeAwsSpotInstanceRequests) Update(awsSpotInstanceRequest *aws_v1.AwsSpotInstanceRequest) (result *aws_v1.AwsSpotInstanceRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsspotinstancerequestsResource, c.ns, awsSpotInstanceRequest), &aws_v1.AwsSpotInstanceRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSpotInstanceRequest), err
}

// Delete takes name of the awsSpotInstanceRequest and deletes it. Returns an error if one occurs.
func (c *FakeAwsSpotInstanceRequests) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsspotinstancerequestsResource, c.ns, name), &aws_v1.AwsSpotInstanceRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSpotInstanceRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsspotinstancerequestsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsSpotInstanceRequestList{})
	return err
}

// Patch applies the patch and returns the patched awsSpotInstanceRequest.
func (c *FakeAwsSpotInstanceRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsSpotInstanceRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsspotinstancerequestsResource, c.ns, name, data, subresources...), &aws_v1.AwsSpotInstanceRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSpotInstanceRequest), err
}
