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

// FakeAwsLaunchConfigurations implements AwsLaunchConfigurationInterface
type FakeAwsLaunchConfigurations struct {
	Fake *FakeTrussleV1
	ns   string
}

var awslaunchconfigurationsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awslaunchconfigurations"}

var awslaunchconfigurationsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsLaunchConfiguration"}

// Get takes name of the awsLaunchConfiguration, and returns the corresponding awsLaunchConfiguration object, and an error if there is any.
func (c *FakeAwsLaunchConfigurations) Get(name string, options v1.GetOptions) (result *aws_v1.AwsLaunchConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awslaunchconfigurationsResource, c.ns, name), &aws_v1.AwsLaunchConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLaunchConfiguration), err
}

// List takes label and field selectors, and returns the list of AwsLaunchConfigurations that match those selectors.
func (c *FakeAwsLaunchConfigurations) List(opts v1.ListOptions) (result *aws_v1.AwsLaunchConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awslaunchconfigurationsResource, awslaunchconfigurationsKind, c.ns, opts), &aws_v1.AwsLaunchConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsLaunchConfigurationList{}
	for _, item := range obj.(*aws_v1.AwsLaunchConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsLaunchConfigurations.
func (c *FakeAwsLaunchConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awslaunchconfigurationsResource, c.ns, opts))

}

// Create takes the representation of a awsLaunchConfiguration and creates it.  Returns the server's representation of the awsLaunchConfiguration, and an error, if there is any.
func (c *FakeAwsLaunchConfigurations) Create(awsLaunchConfiguration *aws_v1.AwsLaunchConfiguration) (result *aws_v1.AwsLaunchConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awslaunchconfigurationsResource, c.ns, awsLaunchConfiguration), &aws_v1.AwsLaunchConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLaunchConfiguration), err
}

// Update takes the representation of a awsLaunchConfiguration and updates it. Returns the server's representation of the awsLaunchConfiguration, and an error, if there is any.
func (c *FakeAwsLaunchConfigurations) Update(awsLaunchConfiguration *aws_v1.AwsLaunchConfiguration) (result *aws_v1.AwsLaunchConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awslaunchconfigurationsResource, c.ns, awsLaunchConfiguration), &aws_v1.AwsLaunchConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLaunchConfiguration), err
}

// Delete takes name of the awsLaunchConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeAwsLaunchConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awslaunchconfigurationsResource, c.ns, name), &aws_v1.AwsLaunchConfiguration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsLaunchConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awslaunchconfigurationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsLaunchConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched awsLaunchConfiguration.
func (c *FakeAwsLaunchConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsLaunchConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awslaunchconfigurationsResource, c.ns, name, data, subresources...), &aws_v1.AwsLaunchConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLaunchConfiguration), err
}
