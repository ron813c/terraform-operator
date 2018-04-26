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

// FakeAwsOrganizationsPolicies implements AwsOrganizationsPolicyInterface
type FakeAwsOrganizationsPolicies struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsorganizationspoliciesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsorganizationspolicies"}

var awsorganizationspoliciesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsOrganizationsPolicy"}

// Get takes name of the awsOrganizationsPolicy, and returns the corresponding awsOrganizationsPolicy object, and an error if there is any.
func (c *FakeAwsOrganizationsPolicies) Get(name string, options v1.GetOptions) (result *aws_v1.AwsOrganizationsPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsorganizationspoliciesResource, c.ns, name), &aws_v1.AwsOrganizationsPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsOrganizationsPolicy), err
}

// List takes label and field selectors, and returns the list of AwsOrganizationsPolicies that match those selectors.
func (c *FakeAwsOrganizationsPolicies) List(opts v1.ListOptions) (result *aws_v1.AwsOrganizationsPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsorganizationspoliciesResource, awsorganizationspoliciesKind, c.ns, opts), &aws_v1.AwsOrganizationsPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsOrganizationsPolicyList{}
	for _, item := range obj.(*aws_v1.AwsOrganizationsPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsOrganizationsPolicies.
func (c *FakeAwsOrganizationsPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsorganizationspoliciesResource, c.ns, opts))

}

// Create takes the representation of a awsOrganizationsPolicy and creates it.  Returns the server's representation of the awsOrganizationsPolicy, and an error, if there is any.
func (c *FakeAwsOrganizationsPolicies) Create(awsOrganizationsPolicy *aws_v1.AwsOrganizationsPolicy) (result *aws_v1.AwsOrganizationsPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsorganizationspoliciesResource, c.ns, awsOrganizationsPolicy), &aws_v1.AwsOrganizationsPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsOrganizationsPolicy), err
}

// Update takes the representation of a awsOrganizationsPolicy and updates it. Returns the server's representation of the awsOrganizationsPolicy, and an error, if there is any.
func (c *FakeAwsOrganizationsPolicies) Update(awsOrganizationsPolicy *aws_v1.AwsOrganizationsPolicy) (result *aws_v1.AwsOrganizationsPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsorganizationspoliciesResource, c.ns, awsOrganizationsPolicy), &aws_v1.AwsOrganizationsPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsOrganizationsPolicy), err
}

// Delete takes name of the awsOrganizationsPolicy and deletes it. Returns an error if one occurs.
func (c *FakeAwsOrganizationsPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsorganizationspoliciesResource, c.ns, name), &aws_v1.AwsOrganizationsPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsOrganizationsPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsorganizationspoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsOrganizationsPolicyList{})
	return err
}

// Patch applies the patch and returns the patched awsOrganizationsPolicy.
func (c *FakeAwsOrganizationsPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsOrganizationsPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsorganizationspoliciesResource, c.ns, name, data, subresources...), &aws_v1.AwsOrganizationsPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsOrganizationsPolicy), err
}