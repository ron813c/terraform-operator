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

// AwsDmsEndpointLister helps list AwsDmsEndpoints.
type AwsDmsEndpointLister interface {
	// List lists all AwsDmsEndpoints in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsDmsEndpoint, err error)
	// AwsDmsEndpoints returns an object that can list and get AwsDmsEndpoints.
	AwsDmsEndpoints(namespace string) AwsDmsEndpointNamespaceLister
	AwsDmsEndpointListerExpansion
}

// awsDmsEndpointLister implements the AwsDmsEndpointLister interface.
type awsDmsEndpointLister struct {
	indexer cache.Indexer
}

// NewAwsDmsEndpointLister returns a new AwsDmsEndpointLister.
func NewAwsDmsEndpointLister(indexer cache.Indexer) AwsDmsEndpointLister {
	return &awsDmsEndpointLister{indexer: indexer}
}

// List lists all AwsDmsEndpoints in the indexer.
func (s *awsDmsEndpointLister) List(selector labels.Selector) (ret []*v1.AwsDmsEndpoint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsDmsEndpoint))
	})
	return ret, err
}

// AwsDmsEndpoints returns an object that can list and get AwsDmsEndpoints.
func (s *awsDmsEndpointLister) AwsDmsEndpoints(namespace string) AwsDmsEndpointNamespaceLister {
	return awsDmsEndpointNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsDmsEndpointNamespaceLister helps list and get AwsDmsEndpoints.
type AwsDmsEndpointNamespaceLister interface {
	// List lists all AwsDmsEndpoints in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsDmsEndpoint, err error)
	// Get retrieves the AwsDmsEndpoint from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsDmsEndpoint, error)
	AwsDmsEndpointNamespaceListerExpansion
}

// awsDmsEndpointNamespaceLister implements the AwsDmsEndpointNamespaceLister
// interface.
type awsDmsEndpointNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsDmsEndpoints in the indexer for a given namespace.
func (s awsDmsEndpointNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsDmsEndpoint, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsDmsEndpoint))
	})
	return ret, err
}

// Get retrieves the AwsDmsEndpoint from the indexer for a given namespace and name.
func (s awsDmsEndpointNamespaceLister) Get(name string) (*v1.AwsDmsEndpoint, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsdmsendpoint"), name)
	}
	return obj.(*v1.AwsDmsEndpoint), nil
}