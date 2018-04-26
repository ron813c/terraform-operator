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

// AwsWafXssMatchSetLister helps list AwsWafXssMatchSets.
type AwsWafXssMatchSetLister interface {
	// List lists all AwsWafXssMatchSets in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsWafXssMatchSet, err error)
	// AwsWafXssMatchSets returns an object that can list and get AwsWafXssMatchSets.
	AwsWafXssMatchSets(namespace string) AwsWafXssMatchSetNamespaceLister
	AwsWafXssMatchSetListerExpansion
}

// awsWafXssMatchSetLister implements the AwsWafXssMatchSetLister interface.
type awsWafXssMatchSetLister struct {
	indexer cache.Indexer
}

// NewAwsWafXssMatchSetLister returns a new AwsWafXssMatchSetLister.
func NewAwsWafXssMatchSetLister(indexer cache.Indexer) AwsWafXssMatchSetLister {
	return &awsWafXssMatchSetLister{indexer: indexer}
}

// List lists all AwsWafXssMatchSets in the indexer.
func (s *awsWafXssMatchSetLister) List(selector labels.Selector) (ret []*v1.AwsWafXssMatchSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsWafXssMatchSet))
	})
	return ret, err
}

// AwsWafXssMatchSets returns an object that can list and get AwsWafXssMatchSets.
func (s *awsWafXssMatchSetLister) AwsWafXssMatchSets(namespace string) AwsWafXssMatchSetNamespaceLister {
	return awsWafXssMatchSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsWafXssMatchSetNamespaceLister helps list and get AwsWafXssMatchSets.
type AwsWafXssMatchSetNamespaceLister interface {
	// List lists all AwsWafXssMatchSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsWafXssMatchSet, err error)
	// Get retrieves the AwsWafXssMatchSet from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsWafXssMatchSet, error)
	AwsWafXssMatchSetNamespaceListerExpansion
}

// awsWafXssMatchSetNamespaceLister implements the AwsWafXssMatchSetNamespaceLister
// interface.
type awsWafXssMatchSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsWafXssMatchSets in the indexer for a given namespace.
func (s awsWafXssMatchSetNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsWafXssMatchSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsWafXssMatchSet))
	})
	return ret, err
}

// Get retrieves the AwsWafXssMatchSet from the indexer for a given namespace and name.
func (s awsWafXssMatchSetNamespaceLister) Get(name string) (*v1.AwsWafXssMatchSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awswafxssmatchset"), name)
	}
	return obj.(*v1.AwsWafXssMatchSet), nil
}