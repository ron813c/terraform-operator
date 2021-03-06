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

// FakeAwsWafregionalIpsets implements AwsWafregionalIpsetInterface
type FakeAwsWafregionalIpsets struct {
	Fake *FakeTrussleV1
	ns   string
}

var awswafregionalipsetsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awswafregionalipsets"}

var awswafregionalipsetsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsWafregionalIpset"}

// Get takes name of the awsWafregionalIpset, and returns the corresponding awsWafregionalIpset object, and an error if there is any.
func (c *FakeAwsWafregionalIpsets) Get(name string, options v1.GetOptions) (result *aws_v1.AwsWafregionalIpset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awswafregionalipsetsResource, c.ns, name), &aws_v1.AwsWafregionalIpset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsWafregionalIpset), err
}

// List takes label and field selectors, and returns the list of AwsWafregionalIpsets that match those selectors.
func (c *FakeAwsWafregionalIpsets) List(opts v1.ListOptions) (result *aws_v1.AwsWafregionalIpsetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awswafregionalipsetsResource, awswafregionalipsetsKind, c.ns, opts), &aws_v1.AwsWafregionalIpsetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsWafregionalIpsetList{}
	for _, item := range obj.(*aws_v1.AwsWafregionalIpsetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsWafregionalIpsets.
func (c *FakeAwsWafregionalIpsets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awswafregionalipsetsResource, c.ns, opts))

}

// Create takes the representation of a awsWafregionalIpset and creates it.  Returns the server's representation of the awsWafregionalIpset, and an error, if there is any.
func (c *FakeAwsWafregionalIpsets) Create(awsWafregionalIpset *aws_v1.AwsWafregionalIpset) (result *aws_v1.AwsWafregionalIpset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awswafregionalipsetsResource, c.ns, awsWafregionalIpset), &aws_v1.AwsWafregionalIpset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsWafregionalIpset), err
}

// Update takes the representation of a awsWafregionalIpset and updates it. Returns the server's representation of the awsWafregionalIpset, and an error, if there is any.
func (c *FakeAwsWafregionalIpsets) Update(awsWafregionalIpset *aws_v1.AwsWafregionalIpset) (result *aws_v1.AwsWafregionalIpset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awswafregionalipsetsResource, c.ns, awsWafregionalIpset), &aws_v1.AwsWafregionalIpset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsWafregionalIpset), err
}

// Delete takes name of the awsWafregionalIpset and deletes it. Returns an error if one occurs.
func (c *FakeAwsWafregionalIpsets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awswafregionalipsetsResource, c.ns, name), &aws_v1.AwsWafregionalIpset{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsWafregionalIpsets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awswafregionalipsetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsWafregionalIpsetList{})
	return err
}

// Patch applies the patch and returns the patched awsWafregionalIpset.
func (c *FakeAwsWafregionalIpsets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsWafregionalIpset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awswafregionalipsetsResource, c.ns, name, data, subresources...), &aws_v1.AwsWafregionalIpset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsWafregionalIpset), err
}
