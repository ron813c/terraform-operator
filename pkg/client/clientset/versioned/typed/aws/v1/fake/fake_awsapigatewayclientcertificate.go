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

// FakeAwsApiGatewayClientCertificates implements AwsApiGatewayClientCertificateInterface
type FakeAwsApiGatewayClientCertificates struct {
	Fake *FakeTrussleV1
	ns   string
}

var awsapigatewayclientcertificatesResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awsapigatewayclientcertificates"}

var awsapigatewayclientcertificatesKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsApiGatewayClientCertificate"}

// Get takes name of the awsApiGatewayClientCertificate, and returns the corresponding awsApiGatewayClientCertificate object, and an error if there is any.
func (c *FakeAwsApiGatewayClientCertificates) Get(name string, options v1.GetOptions) (result *aws_v1.AwsApiGatewayClientCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsapigatewayclientcertificatesResource, c.ns, name), &aws_v1.AwsApiGatewayClientCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsApiGatewayClientCertificate), err
}

// List takes label and field selectors, and returns the list of AwsApiGatewayClientCertificates that match those selectors.
func (c *FakeAwsApiGatewayClientCertificates) List(opts v1.ListOptions) (result *aws_v1.AwsApiGatewayClientCertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsapigatewayclientcertificatesResource, awsapigatewayclientcertificatesKind, c.ns, opts), &aws_v1.AwsApiGatewayClientCertificateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsApiGatewayClientCertificateList{}
	for _, item := range obj.(*aws_v1.AwsApiGatewayClientCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsApiGatewayClientCertificates.
func (c *FakeAwsApiGatewayClientCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsapigatewayclientcertificatesResource, c.ns, opts))

}

// Create takes the representation of a awsApiGatewayClientCertificate and creates it.  Returns the server's representation of the awsApiGatewayClientCertificate, and an error, if there is any.
func (c *FakeAwsApiGatewayClientCertificates) Create(awsApiGatewayClientCertificate *aws_v1.AwsApiGatewayClientCertificate) (result *aws_v1.AwsApiGatewayClientCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsapigatewayclientcertificatesResource, c.ns, awsApiGatewayClientCertificate), &aws_v1.AwsApiGatewayClientCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsApiGatewayClientCertificate), err
}

// Update takes the representation of a awsApiGatewayClientCertificate and updates it. Returns the server's representation of the awsApiGatewayClientCertificate, and an error, if there is any.
func (c *FakeAwsApiGatewayClientCertificates) Update(awsApiGatewayClientCertificate *aws_v1.AwsApiGatewayClientCertificate) (result *aws_v1.AwsApiGatewayClientCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsapigatewayclientcertificatesResource, c.ns, awsApiGatewayClientCertificate), &aws_v1.AwsApiGatewayClientCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsApiGatewayClientCertificate), err
}

// Delete takes name of the awsApiGatewayClientCertificate and deletes it. Returns an error if one occurs.
func (c *FakeAwsApiGatewayClientCertificates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsapigatewayclientcertificatesResource, c.ns, name), &aws_v1.AwsApiGatewayClientCertificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsApiGatewayClientCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsapigatewayclientcertificatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsApiGatewayClientCertificateList{})
	return err
}

// Patch applies the patch and returns the patched awsApiGatewayClientCertificate.
func (c *FakeAwsApiGatewayClientCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsApiGatewayClientCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsapigatewayclientcertificatesResource, c.ns, name, data, subresources...), &aws_v1.AwsApiGatewayClientCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsApiGatewayClientCertificate), err
}