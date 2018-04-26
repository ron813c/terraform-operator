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

// FakeAwsS3BucketMetrics implements AwsS3BucketMetricInterface
type FakeAwsS3BucketMetrics struct {
	Fake *FakeTrussleV1
	ns   string
}

var awss3bucketmetricsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awss3bucketmetrics"}

var awss3bucketmetricsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsS3BucketMetric"}

// Get takes name of the awsS3BucketMetric, and returns the corresponding awsS3BucketMetric object, and an error if there is any.
func (c *FakeAwsS3BucketMetrics) Get(name string, options v1.GetOptions) (result *aws_v1.AwsS3BucketMetric, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awss3bucketmetricsResource, c.ns, name), &aws_v1.AwsS3BucketMetric{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsS3BucketMetric), err
}

// List takes label and field selectors, and returns the list of AwsS3BucketMetrics that match those selectors.
func (c *FakeAwsS3BucketMetrics) List(opts v1.ListOptions) (result *aws_v1.AwsS3BucketMetricList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awss3bucketmetricsResource, awss3bucketmetricsKind, c.ns, opts), &aws_v1.AwsS3BucketMetricList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsS3BucketMetricList{}
	for _, item := range obj.(*aws_v1.AwsS3BucketMetricList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsS3BucketMetrics.
func (c *FakeAwsS3BucketMetrics) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awss3bucketmetricsResource, c.ns, opts))

}

// Create takes the representation of a awsS3BucketMetric and creates it.  Returns the server's representation of the awsS3BucketMetric, and an error, if there is any.
func (c *FakeAwsS3BucketMetrics) Create(awsS3BucketMetric *aws_v1.AwsS3BucketMetric) (result *aws_v1.AwsS3BucketMetric, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awss3bucketmetricsResource, c.ns, awsS3BucketMetric), &aws_v1.AwsS3BucketMetric{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsS3BucketMetric), err
}

// Update takes the representation of a awsS3BucketMetric and updates it. Returns the server's representation of the awsS3BucketMetric, and an error, if there is any.
func (c *FakeAwsS3BucketMetrics) Update(awsS3BucketMetric *aws_v1.AwsS3BucketMetric) (result *aws_v1.AwsS3BucketMetric, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awss3bucketmetricsResource, c.ns, awsS3BucketMetric), &aws_v1.AwsS3BucketMetric{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsS3BucketMetric), err
}

// Delete takes name of the awsS3BucketMetric and deletes it. Returns an error if one occurs.
func (c *FakeAwsS3BucketMetrics) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awss3bucketmetricsResource, c.ns, name), &aws_v1.AwsS3BucketMetric{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsS3BucketMetrics) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awss3bucketmetricsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsS3BucketMetricList{})
	return err
}

// Patch applies the patch and returns the patched awsS3BucketMetric.
func (c *FakeAwsS3BucketMetrics) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsS3BucketMetric, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awss3bucketmetricsResource, c.ns, name, data, subresources...), &aws_v1.AwsS3BucketMetric{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsS3BucketMetric), err
}