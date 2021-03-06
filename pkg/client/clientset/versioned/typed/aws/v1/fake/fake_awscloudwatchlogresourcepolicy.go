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

// FakeAwsCloudwatchLogResourcePolicies implements AwsCloudwatchLogResourcePolicyInterface
type FakeAwsCloudwatchLogResourcePolicies struct {
	Fake *FakeTrussleV1
	ns   string
}

var awscloudwatchlogresourcepoliciesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awscloudwatchlogresourcepolicies"}

var awscloudwatchlogresourcepoliciesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsCloudwatchLogResourcePolicy"}

// Get takes name of the awsCloudwatchLogResourcePolicy, and returns the corresponding awsCloudwatchLogResourcePolicy object, and an error if there is any.
func (c *FakeAwsCloudwatchLogResourcePolicies) Get(name string, options v1.GetOptions) (result *aws_v1.AwsCloudwatchLogResourcePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awscloudwatchlogresourcepoliciesResource, c.ns, name), &aws_v1.AwsCloudwatchLogResourcePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCloudwatchLogResourcePolicy), err
}

// List takes label and field selectors, and returns the list of AwsCloudwatchLogResourcePolicies that match those selectors.
func (c *FakeAwsCloudwatchLogResourcePolicies) List(opts v1.ListOptions) (result *aws_v1.AwsCloudwatchLogResourcePolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awscloudwatchlogresourcepoliciesResource, awscloudwatchlogresourcepoliciesKind, c.ns, opts), &aws_v1.AwsCloudwatchLogResourcePolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsCloudwatchLogResourcePolicyList{}
	for _, item := range obj.(*aws_v1.AwsCloudwatchLogResourcePolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsCloudwatchLogResourcePolicies.
func (c *FakeAwsCloudwatchLogResourcePolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awscloudwatchlogresourcepoliciesResource, c.ns, opts))

}

// Create takes the representation of a awsCloudwatchLogResourcePolicy and creates it.  Returns the server's representation of the awsCloudwatchLogResourcePolicy, and an error, if there is any.
func (c *FakeAwsCloudwatchLogResourcePolicies) Create(awsCloudwatchLogResourcePolicy *aws_v1.AwsCloudwatchLogResourcePolicy) (result *aws_v1.AwsCloudwatchLogResourcePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awscloudwatchlogresourcepoliciesResource, c.ns, awsCloudwatchLogResourcePolicy), &aws_v1.AwsCloudwatchLogResourcePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCloudwatchLogResourcePolicy), err
}

// Update takes the representation of a awsCloudwatchLogResourcePolicy and updates it. Returns the server's representation of the awsCloudwatchLogResourcePolicy, and an error, if there is any.
func (c *FakeAwsCloudwatchLogResourcePolicies) Update(awsCloudwatchLogResourcePolicy *aws_v1.AwsCloudwatchLogResourcePolicy) (result *aws_v1.AwsCloudwatchLogResourcePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awscloudwatchlogresourcepoliciesResource, c.ns, awsCloudwatchLogResourcePolicy), &aws_v1.AwsCloudwatchLogResourcePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCloudwatchLogResourcePolicy), err
}

// Delete takes name of the awsCloudwatchLogResourcePolicy and deletes it. Returns an error if one occurs.
func (c *FakeAwsCloudwatchLogResourcePolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awscloudwatchlogresourcepoliciesResource, c.ns, name), &aws_v1.AwsCloudwatchLogResourcePolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsCloudwatchLogResourcePolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awscloudwatchlogresourcepoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsCloudwatchLogResourcePolicyList{})
	return err
}

// Patch applies the patch and returns the patched awsCloudwatchLogResourcePolicy.
func (c *FakeAwsCloudwatchLogResourcePolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsCloudwatchLogResourcePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awscloudwatchlogresourcepoliciesResource, c.ns, name, data, subresources...), &aws_v1.AwsCloudwatchLogResourcePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCloudwatchLogResourcePolicy), err
}
