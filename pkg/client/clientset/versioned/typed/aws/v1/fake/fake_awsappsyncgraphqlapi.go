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

// FakeAwsAppsyncGraphqlApis implements AwsAppsyncGraphqlApiInterface
type FakeAwsAppsyncGraphqlApis struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsappsyncgraphqlapisResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsappsyncgraphqlapis"}

var awsappsyncgraphqlapisKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsAppsyncGraphqlApi"}

// Get takes name of the awsAppsyncGraphqlApi, and returns the corresponding awsAppsyncGraphqlApi object, and an error if there is any.
func (c *FakeAwsAppsyncGraphqlApis) Get(name string, options v1.GetOptions) (result *aws_v1.AwsAppsyncGraphqlApi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsappsyncgraphqlapisResource, c.ns, name), &aws_v1.AwsAppsyncGraphqlApi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAppsyncGraphqlApi), err
}

// List takes label and field selectors, and returns the list of AwsAppsyncGraphqlApis that match those selectors.
func (c *FakeAwsAppsyncGraphqlApis) List(opts v1.ListOptions) (result *aws_v1.AwsAppsyncGraphqlApiList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsappsyncgraphqlapisResource, awsappsyncgraphqlapisKind, c.ns, opts), &aws_v1.AwsAppsyncGraphqlApiList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsAppsyncGraphqlApiList{}
	for _, item := range obj.(*aws_v1.AwsAppsyncGraphqlApiList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsAppsyncGraphqlApis.
func (c *FakeAwsAppsyncGraphqlApis) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsappsyncgraphqlapisResource, c.ns, opts))

}

// Create takes the representation of a awsAppsyncGraphqlApi and creates it.  Returns the server's representation of the awsAppsyncGraphqlApi, and an error, if there is any.
func (c *FakeAwsAppsyncGraphqlApis) Create(awsAppsyncGraphqlApi *aws_v1.AwsAppsyncGraphqlApi) (result *aws_v1.AwsAppsyncGraphqlApi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsappsyncgraphqlapisResource, c.ns, awsAppsyncGraphqlApi), &aws_v1.AwsAppsyncGraphqlApi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAppsyncGraphqlApi), err
}

// Update takes the representation of a awsAppsyncGraphqlApi and updates it. Returns the server's representation of the awsAppsyncGraphqlApi, and an error, if there is any.
func (c *FakeAwsAppsyncGraphqlApis) Update(awsAppsyncGraphqlApi *aws_v1.AwsAppsyncGraphqlApi) (result *aws_v1.AwsAppsyncGraphqlApi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsappsyncgraphqlapisResource, c.ns, awsAppsyncGraphqlApi), &aws_v1.AwsAppsyncGraphqlApi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAppsyncGraphqlApi), err
}

// Delete takes name of the awsAppsyncGraphqlApi and deletes it. Returns an error if one occurs.
func (c *FakeAwsAppsyncGraphqlApis) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsappsyncgraphqlapisResource, c.ns, name), &aws_v1.AwsAppsyncGraphqlApi{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsAppsyncGraphqlApis) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsappsyncgraphqlapisResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsAppsyncGraphqlApiList{})
	return err
}

// Patch applies the patch and returns the patched awsAppsyncGraphqlApi.
func (c *FakeAwsAppsyncGraphqlApis) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsAppsyncGraphqlApi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsappsyncgraphqlapisResource, c.ns, name, data, subresources...), &aws_v1.AwsAppsyncGraphqlApi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsAppsyncGraphqlApi), err
}