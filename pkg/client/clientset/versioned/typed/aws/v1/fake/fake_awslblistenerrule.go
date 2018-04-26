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

// FakeAwsLbListenerRules implements AwsLbListenerRuleInterface
type FakeAwsLbListenerRules struct {
	Fake *FakeTrussleV1
	ns   string
}

var awslblistenerrulesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awslblistenerrules"}

var awslblistenerrulesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsLbListenerRule"}

// Get takes name of the awsLbListenerRule, and returns the corresponding awsLbListenerRule object, and an error if there is any.
func (c *FakeAwsLbListenerRules) Get(name string, options v1.GetOptions) (result *aws_v1.AwsLbListenerRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awslblistenerrulesResource, c.ns, name), &aws_v1.AwsLbListenerRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLbListenerRule), err
}

// List takes label and field selectors, and returns the list of AwsLbListenerRules that match those selectors.
func (c *FakeAwsLbListenerRules) List(opts v1.ListOptions) (result *aws_v1.AwsLbListenerRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awslblistenerrulesResource, awslblistenerrulesKind, c.ns, opts), &aws_v1.AwsLbListenerRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsLbListenerRuleList{}
	for _, item := range obj.(*aws_v1.AwsLbListenerRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsLbListenerRules.
func (c *FakeAwsLbListenerRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awslblistenerrulesResource, c.ns, opts))

}

// Create takes the representation of a awsLbListenerRule and creates it.  Returns the server's representation of the awsLbListenerRule, and an error, if there is any.
func (c *FakeAwsLbListenerRules) Create(awsLbListenerRule *aws_v1.AwsLbListenerRule) (result *aws_v1.AwsLbListenerRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awslblistenerrulesResource, c.ns, awsLbListenerRule), &aws_v1.AwsLbListenerRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLbListenerRule), err
}

// Update takes the representation of a awsLbListenerRule and updates it. Returns the server's representation of the awsLbListenerRule, and an error, if there is any.
func (c *FakeAwsLbListenerRules) Update(awsLbListenerRule *aws_v1.AwsLbListenerRule) (result *aws_v1.AwsLbListenerRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awslblistenerrulesResource, c.ns, awsLbListenerRule), &aws_v1.AwsLbListenerRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLbListenerRule), err
}

// Delete takes name of the awsLbListenerRule and deletes it. Returns an error if one occurs.
func (c *FakeAwsLbListenerRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awslblistenerrulesResource, c.ns, name), &aws_v1.AwsLbListenerRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsLbListenerRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awslblistenerrulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsLbListenerRuleList{})
	return err
}

// Patch applies the patch and returns the patched awsLbListenerRule.
func (c *FakeAwsLbListenerRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsLbListenerRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awslblistenerrulesResource, c.ns, name, data, subresources...), &aws_v1.AwsLbListenerRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLbListenerRule), err
}