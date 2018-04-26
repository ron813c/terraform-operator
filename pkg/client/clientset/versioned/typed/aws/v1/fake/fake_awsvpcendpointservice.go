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

// FakeAwsVpcEndpointServices implements AwsVpcEndpointServiceInterface
type FakeAwsVpcEndpointServices struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsvpcendpointservicesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsvpcendpointservices"}

var awsvpcendpointservicesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsVpcEndpointService"}

// Get takes name of the awsVpcEndpointService, and returns the corresponding awsVpcEndpointService object, and an error if there is any.
func (c *FakeAwsVpcEndpointServices) Get(name string, options v1.GetOptions) (result *aws_v1.AwsVpcEndpointService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsvpcendpointservicesResource, c.ns, name), &aws_v1.AwsVpcEndpointService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsVpcEndpointService), err
}

// List takes label and field selectors, and returns the list of AwsVpcEndpointServices that match those selectors.
func (c *FakeAwsVpcEndpointServices) List(opts v1.ListOptions) (result *aws_v1.AwsVpcEndpointServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsvpcendpointservicesResource, awsvpcendpointservicesKind, c.ns, opts), &aws_v1.AwsVpcEndpointServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsVpcEndpointServiceList{}
	for _, item := range obj.(*aws_v1.AwsVpcEndpointServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsVpcEndpointServices.
func (c *FakeAwsVpcEndpointServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsvpcendpointservicesResource, c.ns, opts))

}

// Create takes the representation of a awsVpcEndpointService and creates it.  Returns the server's representation of the awsVpcEndpointService, and an error, if there is any.
func (c *FakeAwsVpcEndpointServices) Create(awsVpcEndpointService *aws_v1.AwsVpcEndpointService) (result *aws_v1.AwsVpcEndpointService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsvpcendpointservicesResource, c.ns, awsVpcEndpointService), &aws_v1.AwsVpcEndpointService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsVpcEndpointService), err
}

// Update takes the representation of a awsVpcEndpointService and updates it. Returns the server's representation of the awsVpcEndpointService, and an error, if there is any.
func (c *FakeAwsVpcEndpointServices) Update(awsVpcEndpointService *aws_v1.AwsVpcEndpointService) (result *aws_v1.AwsVpcEndpointService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsvpcendpointservicesResource, c.ns, awsVpcEndpointService), &aws_v1.AwsVpcEndpointService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsVpcEndpointService), err
}

// Delete takes name of the awsVpcEndpointService and deletes it. Returns an error if one occurs.
func (c *FakeAwsVpcEndpointServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsvpcendpointservicesResource, c.ns, name), &aws_v1.AwsVpcEndpointService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsVpcEndpointServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsvpcendpointservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsVpcEndpointServiceList{})
	return err
}

// Patch applies the patch and returns the patched awsVpcEndpointService.
func (c *FakeAwsVpcEndpointServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsVpcEndpointService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsvpcendpointservicesResource, c.ns, name, data, subresources...), &aws_v1.AwsVpcEndpointService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsVpcEndpointService), err
}