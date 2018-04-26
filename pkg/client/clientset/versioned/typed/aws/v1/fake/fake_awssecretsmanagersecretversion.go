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

// FakeAwsSecretsmanagerSecretVersions implements AwsSecretsmanagerSecretVersionInterface
type FakeAwsSecretsmanagerSecretVersions struct {
	Fake *FakeTrussleV1
	ns   string
}

var awssecretsmanagersecretversionsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awssecretsmanagersecretversions"}

var awssecretsmanagersecretversionsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsSecretsmanagerSecretVersion"}

// Get takes name of the awsSecretsmanagerSecretVersion, and returns the corresponding awsSecretsmanagerSecretVersion object, and an error if there is any.
func (c *FakeAwsSecretsmanagerSecretVersions) Get(name string, options v1.GetOptions) (result *aws_v1.AwsSecretsmanagerSecretVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awssecretsmanagersecretversionsResource, c.ns, name), &aws_v1.AwsSecretsmanagerSecretVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSecretsmanagerSecretVersion), err
}

// List takes label and field selectors, and returns the list of AwsSecretsmanagerSecretVersions that match those selectors.
func (c *FakeAwsSecretsmanagerSecretVersions) List(opts v1.ListOptions) (result *aws_v1.AwsSecretsmanagerSecretVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awssecretsmanagersecretversionsResource, awssecretsmanagersecretversionsKind, c.ns, opts), &aws_v1.AwsSecretsmanagerSecretVersionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsSecretsmanagerSecretVersionList{}
	for _, item := range obj.(*aws_v1.AwsSecretsmanagerSecretVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSecretsmanagerSecretVersions.
func (c *FakeAwsSecretsmanagerSecretVersions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awssecretsmanagersecretversionsResource, c.ns, opts))

}

// Create takes the representation of a awsSecretsmanagerSecretVersion and creates it.  Returns the server's representation of the awsSecretsmanagerSecretVersion, and an error, if there is any.
func (c *FakeAwsSecretsmanagerSecretVersions) Create(awsSecretsmanagerSecretVersion *aws_v1.AwsSecretsmanagerSecretVersion) (result *aws_v1.AwsSecretsmanagerSecretVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awssecretsmanagersecretversionsResource, c.ns, awsSecretsmanagerSecretVersion), &aws_v1.AwsSecretsmanagerSecretVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSecretsmanagerSecretVersion), err
}

// Update takes the representation of a awsSecretsmanagerSecretVersion and updates it. Returns the server's representation of the awsSecretsmanagerSecretVersion, and an error, if there is any.
func (c *FakeAwsSecretsmanagerSecretVersions) Update(awsSecretsmanagerSecretVersion *aws_v1.AwsSecretsmanagerSecretVersion) (result *aws_v1.AwsSecretsmanagerSecretVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awssecretsmanagersecretversionsResource, c.ns, awsSecretsmanagerSecretVersion), &aws_v1.AwsSecretsmanagerSecretVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSecretsmanagerSecretVersion), err
}

// Delete takes name of the awsSecretsmanagerSecretVersion and deletes it. Returns an error if one occurs.
func (c *FakeAwsSecretsmanagerSecretVersions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awssecretsmanagersecretversionsResource, c.ns, name), &aws_v1.AwsSecretsmanagerSecretVersion{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSecretsmanagerSecretVersions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awssecretsmanagersecretversionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsSecretsmanagerSecretVersionList{})
	return err
}

// Patch applies the patch and returns the patched awsSecretsmanagerSecretVersion.
func (c *FakeAwsSecretsmanagerSecretVersions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsSecretsmanagerSecretVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awssecretsmanagersecretversionsResource, c.ns, name, data, subresources...), &aws_v1.AwsSecretsmanagerSecretVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsSecretsmanagerSecretVersion), err
}
