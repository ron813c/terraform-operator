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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/trussle/terraform-operator/pkg/apis/aws/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AwsLambdaEventSourceMappingLister helps list AwsLambdaEventSourceMappings.
type AwsLambdaEventSourceMappingLister interface {
	// List lists all AwsLambdaEventSourceMappings in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsLambdaEventSourceMapping, err error)
	// AwsLambdaEventSourceMappings returns an object that can list and get AwsLambdaEventSourceMappings.
	AwsLambdaEventSourceMappings(namespace string) AwsLambdaEventSourceMappingNamespaceLister
	AwsLambdaEventSourceMappingListerExpansion
}

// awsLambdaEventSourceMappingLister implements the AwsLambdaEventSourceMappingLister interface.
type awsLambdaEventSourceMappingLister struct {
	indexer cache.Indexer
}

// NewAwsLambdaEventSourceMappingLister returns a new AwsLambdaEventSourceMappingLister.
func NewAwsLambdaEventSourceMappingLister(indexer cache.Indexer) AwsLambdaEventSourceMappingLister {
	return &awsLambdaEventSourceMappingLister{indexer: indexer}
}

// List lists all AwsLambdaEventSourceMappings in the indexer.
func (s *awsLambdaEventSourceMappingLister) List(selector labels.Selector) (ret []*v1.AwsLambdaEventSourceMapping, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsLambdaEventSourceMapping))
	})
	return ret, err
}

// AwsLambdaEventSourceMappings returns an object that can list and get AwsLambdaEventSourceMappings.
func (s *awsLambdaEventSourceMappingLister) AwsLambdaEventSourceMappings(namespace string) AwsLambdaEventSourceMappingNamespaceLister {
	return awsLambdaEventSourceMappingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsLambdaEventSourceMappingNamespaceLister helps list and get AwsLambdaEventSourceMappings.
type AwsLambdaEventSourceMappingNamespaceLister interface {
	// List lists all AwsLambdaEventSourceMappings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsLambdaEventSourceMapping, err error)
	// Get retrieves the AwsLambdaEventSourceMapping from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsLambdaEventSourceMapping, error)
	AwsLambdaEventSourceMappingNamespaceListerExpansion
}

// awsLambdaEventSourceMappingNamespaceLister implements the AwsLambdaEventSourceMappingNamespaceLister
// interface.
type awsLambdaEventSourceMappingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsLambdaEventSourceMappings in the indexer for a given namespace.
func (s awsLambdaEventSourceMappingNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsLambdaEventSourceMapping, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsLambdaEventSourceMapping))
	})
	return ret, err
}

// Get retrieves the AwsLambdaEventSourceMapping from the indexer for a given namespace and name.
func (s awsLambdaEventSourceMappingNamespaceLister) Get(name string) (*v1.AwsLambdaEventSourceMapping, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awslambdaeventsourcemapping"), name)
	}
	return obj.(*v1.AwsLambdaEventSourceMapping), nil
}