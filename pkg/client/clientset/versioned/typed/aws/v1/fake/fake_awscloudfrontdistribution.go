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

// FakeAwsCloudfrontDistributions implements AwsCloudfrontDistributionInterface
type FakeAwsCloudfrontDistributions struct {
	Fake *FakeTrussleV1
	ns   string
}

var awscloudfrontdistributionsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awscloudfrontdistributions"}

var awscloudfrontdistributionsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsCloudfrontDistribution"}

// Get takes name of the awsCloudfrontDistribution, and returns the corresponding awsCloudfrontDistribution object, and an error if there is any.
func (c *FakeAwsCloudfrontDistributions) Get(name string, options v1.GetOptions) (result *aws_v1.AwsCloudfrontDistribution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awscloudfrontdistributionsResource, c.ns, name), &aws_v1.AwsCloudfrontDistribution{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCloudfrontDistribution), err
}

// List takes label and field selectors, and returns the list of AwsCloudfrontDistributions that match those selectors.
func (c *FakeAwsCloudfrontDistributions) List(opts v1.ListOptions) (result *aws_v1.AwsCloudfrontDistributionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awscloudfrontdistributionsResource, awscloudfrontdistributionsKind, c.ns, opts), &aws_v1.AwsCloudfrontDistributionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsCloudfrontDistributionList{}
	for _, item := range obj.(*aws_v1.AwsCloudfrontDistributionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsCloudfrontDistributions.
func (c *FakeAwsCloudfrontDistributions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awscloudfrontdistributionsResource, c.ns, opts))

}

// Create takes the representation of a awsCloudfrontDistribution and creates it.  Returns the server's representation of the awsCloudfrontDistribution, and an error, if there is any.
func (c *FakeAwsCloudfrontDistributions) Create(awsCloudfrontDistribution *aws_v1.AwsCloudfrontDistribution) (result *aws_v1.AwsCloudfrontDistribution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awscloudfrontdistributionsResource, c.ns, awsCloudfrontDistribution), &aws_v1.AwsCloudfrontDistribution{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCloudfrontDistribution), err
}

// Update takes the representation of a awsCloudfrontDistribution and updates it. Returns the server's representation of the awsCloudfrontDistribution, and an error, if there is any.
func (c *FakeAwsCloudfrontDistributions) Update(awsCloudfrontDistribution *aws_v1.AwsCloudfrontDistribution) (result *aws_v1.AwsCloudfrontDistribution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awscloudfrontdistributionsResource, c.ns, awsCloudfrontDistribution), &aws_v1.AwsCloudfrontDistribution{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCloudfrontDistribution), err
}

// Delete takes name of the awsCloudfrontDistribution and deletes it. Returns an error if one occurs.
func (c *FakeAwsCloudfrontDistributions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awscloudfrontdistributionsResource, c.ns, name), &aws_v1.AwsCloudfrontDistribution{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsCloudfrontDistributions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awscloudfrontdistributionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsCloudfrontDistributionList{})
	return err
}

// Patch applies the patch and returns the patched awsCloudfrontDistribution.
func (c *FakeAwsCloudfrontDistributions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsCloudfrontDistribution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awscloudfrontdistributionsResource, c.ns, name, data, subresources...), &aws_v1.AwsCloudfrontDistribution{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCloudfrontDistribution), err
}
