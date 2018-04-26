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

// FakeAwsLambdaPermissions implements AwsLambdaPermissionInterface
type FakeAwsLambdaPermissions struct {
	Fake *FakeTrussleV1
	ns   string
}

var awslambdapermissionsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awslambdapermissions"}

var awslambdapermissionsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsLambdaPermission"}

// Get takes name of the awsLambdaPermission, and returns the corresponding awsLambdaPermission object, and an error if there is any.
func (c *FakeAwsLambdaPermissions) Get(name string, options v1.GetOptions) (result *aws_v1.AwsLambdaPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awslambdapermissionsResource, c.ns, name), &aws_v1.AwsLambdaPermission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLambdaPermission), err
}

// List takes label and field selectors, and returns the list of AwsLambdaPermissions that match those selectors.
func (c *FakeAwsLambdaPermissions) List(opts v1.ListOptions) (result *aws_v1.AwsLambdaPermissionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awslambdapermissionsResource, awslambdapermissionsKind, c.ns, opts), &aws_v1.AwsLambdaPermissionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsLambdaPermissionList{}
	for _, item := range obj.(*aws_v1.AwsLambdaPermissionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsLambdaPermissions.
func (c *FakeAwsLambdaPermissions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awslambdapermissionsResource, c.ns, opts))

}

// Create takes the representation of a awsLambdaPermission and creates it.  Returns the server's representation of the awsLambdaPermission, and an error, if there is any.
func (c *FakeAwsLambdaPermissions) Create(awsLambdaPermission *aws_v1.AwsLambdaPermission) (result *aws_v1.AwsLambdaPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awslambdapermissionsResource, c.ns, awsLambdaPermission), &aws_v1.AwsLambdaPermission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLambdaPermission), err
}

// Update takes the representation of a awsLambdaPermission and updates it. Returns the server's representation of the awsLambdaPermission, and an error, if there is any.
func (c *FakeAwsLambdaPermissions) Update(awsLambdaPermission *aws_v1.AwsLambdaPermission) (result *aws_v1.AwsLambdaPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awslambdapermissionsResource, c.ns, awsLambdaPermission), &aws_v1.AwsLambdaPermission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLambdaPermission), err
}

// Delete takes name of the awsLambdaPermission and deletes it. Returns an error if one occurs.
func (c *FakeAwsLambdaPermissions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awslambdapermissionsResource, c.ns, name), &aws_v1.AwsLambdaPermission{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsLambdaPermissions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awslambdapermissionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsLambdaPermissionList{})
	return err
}

// Patch applies the patch and returns the patched awsLambdaPermission.
func (c *FakeAwsLambdaPermissions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsLambdaPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awslambdapermissionsResource, c.ns, name, data, subresources...), &aws_v1.AwsLambdaPermission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLambdaPermission), err
}