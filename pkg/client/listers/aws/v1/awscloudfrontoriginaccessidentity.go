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

// AwsCloudfrontOriginAccessIdentityLister helps list AwsCloudfrontOriginAccessIdentities.
type AwsCloudfrontOriginAccessIdentityLister interface {
	// List lists all AwsCloudfrontOriginAccessIdentities in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsCloudfrontOriginAccessIdentity, err error)
	// AwsCloudfrontOriginAccessIdentities returns an object that can list and get AwsCloudfrontOriginAccessIdentities.
	AwsCloudfrontOriginAccessIdentities(namespace string) AwsCloudfrontOriginAccessIdentityNamespaceLister
	AwsCloudfrontOriginAccessIdentityListerExpansion
}

// awsCloudfrontOriginAccessIdentityLister implements the AwsCloudfrontOriginAccessIdentityLister interface.
type awsCloudfrontOriginAccessIdentityLister struct {
	indexer cache.Indexer
}

// NewAwsCloudfrontOriginAccessIdentityLister returns a new AwsCloudfrontOriginAccessIdentityLister.
func NewAwsCloudfrontOriginAccessIdentityLister(indexer cache.Indexer) AwsCloudfrontOriginAccessIdentityLister {
	return &awsCloudfrontOriginAccessIdentityLister{indexer: indexer}
}

// List lists all AwsCloudfrontOriginAccessIdentities in the indexer.
func (s *awsCloudfrontOriginAccessIdentityLister) List(selector labels.Selector) (ret []*v1.AwsCloudfrontOriginAccessIdentity, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsCloudfrontOriginAccessIdentity))
	})
	return ret, err
}

// AwsCloudfrontOriginAccessIdentities returns an object that can list and get AwsCloudfrontOriginAccessIdentities.
func (s *awsCloudfrontOriginAccessIdentityLister) AwsCloudfrontOriginAccessIdentities(namespace string) AwsCloudfrontOriginAccessIdentityNamespaceLister {
	return awsCloudfrontOriginAccessIdentityNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsCloudfrontOriginAccessIdentityNamespaceLister helps list and get AwsCloudfrontOriginAccessIdentities.
type AwsCloudfrontOriginAccessIdentityNamespaceLister interface {
	// List lists all AwsCloudfrontOriginAccessIdentities in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsCloudfrontOriginAccessIdentity, err error)
	// Get retrieves the AwsCloudfrontOriginAccessIdentity from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsCloudfrontOriginAccessIdentity, error)
	AwsCloudfrontOriginAccessIdentityNamespaceListerExpansion
}

// awsCloudfrontOriginAccessIdentityNamespaceLister implements the AwsCloudfrontOriginAccessIdentityNamespaceLister
// interface.
type awsCloudfrontOriginAccessIdentityNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsCloudfrontOriginAccessIdentities in the indexer for a given namespace.
func (s awsCloudfrontOriginAccessIdentityNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsCloudfrontOriginAccessIdentity, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsCloudfrontOriginAccessIdentity))
	})
	return ret, err
}

// Get retrieves the AwsCloudfrontOriginAccessIdentity from the indexer for a given namespace and name.
func (s awsCloudfrontOriginAccessIdentityNamespaceLister) Get(name string) (*v1.AwsCloudfrontOriginAccessIdentity, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awscloudfrontoriginaccessidentity"), name)
	}
	return obj.(*v1.AwsCloudfrontOriginAccessIdentity), nil
}