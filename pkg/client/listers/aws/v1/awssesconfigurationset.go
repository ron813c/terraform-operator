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

// AwsSesConfigurationSetLister helps list AwsSesConfigurationSets.
type AwsSesConfigurationSetLister interface {
	// List lists all AwsSesConfigurationSets in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsSesConfigurationSet, err error)
	// AwsSesConfigurationSets returns an object that can list and get AwsSesConfigurationSets.
	AwsSesConfigurationSets(namespace string) AwsSesConfigurationSetNamespaceLister
	AwsSesConfigurationSetListerExpansion
}

// awsSesConfigurationSetLister implements the AwsSesConfigurationSetLister interface.
type awsSesConfigurationSetLister struct {
	indexer cache.Indexer
}

// NewAwsSesConfigurationSetLister returns a new AwsSesConfigurationSetLister.
func NewAwsSesConfigurationSetLister(indexer cache.Indexer) AwsSesConfigurationSetLister {
	return &awsSesConfigurationSetLister{indexer: indexer}
}

// List lists all AwsSesConfigurationSets in the indexer.
func (s *awsSesConfigurationSetLister) List(selector labels.Selector) (ret []*v1.AwsSesConfigurationSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsSesConfigurationSet))
	})
	return ret, err
}

// AwsSesConfigurationSets returns an object that can list and get AwsSesConfigurationSets.
func (s *awsSesConfigurationSetLister) AwsSesConfigurationSets(namespace string) AwsSesConfigurationSetNamespaceLister {
	return awsSesConfigurationSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsSesConfigurationSetNamespaceLister helps list and get AwsSesConfigurationSets.
type AwsSesConfigurationSetNamespaceLister interface {
	// List lists all AwsSesConfigurationSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsSesConfigurationSet, err error)
	// Get retrieves the AwsSesConfigurationSet from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsSesConfigurationSet, error)
	AwsSesConfigurationSetNamespaceListerExpansion
}

// awsSesConfigurationSetNamespaceLister implements the AwsSesConfigurationSetNamespaceLister
// interface.
type awsSesConfigurationSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsSesConfigurationSets in the indexer for a given namespace.
func (s awsSesConfigurationSetNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsSesConfigurationSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsSesConfigurationSet))
	})
	return ret, err
}

// Get retrieves the AwsSesConfigurationSet from the indexer for a given namespace and name.
func (s awsSesConfigurationSetNamespaceLister) Get(name string) (*v1.AwsSesConfigurationSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awssesconfigurationset"), name)
	}
	return obj.(*v1.AwsSesConfigurationSet), nil
}