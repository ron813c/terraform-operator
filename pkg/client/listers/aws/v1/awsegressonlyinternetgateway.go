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

// AwsEgressOnlyInternetGatewayLister helps list AwsEgressOnlyInternetGateways.
type AwsEgressOnlyInternetGatewayLister interface {
	// List lists all AwsEgressOnlyInternetGateways in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsEgressOnlyInternetGateway, err error)
	// AwsEgressOnlyInternetGateways returns an object that can list and get AwsEgressOnlyInternetGateways.
	AwsEgressOnlyInternetGateways(namespace string) AwsEgressOnlyInternetGatewayNamespaceLister
	AwsEgressOnlyInternetGatewayListerExpansion
}

// awsEgressOnlyInternetGatewayLister implements the AwsEgressOnlyInternetGatewayLister interface.
type awsEgressOnlyInternetGatewayLister struct {
	indexer cache.Indexer
}

// NewAwsEgressOnlyInternetGatewayLister returns a new AwsEgressOnlyInternetGatewayLister.
func NewAwsEgressOnlyInternetGatewayLister(indexer cache.Indexer) AwsEgressOnlyInternetGatewayLister {
	return &awsEgressOnlyInternetGatewayLister{indexer: indexer}
}

// List lists all AwsEgressOnlyInternetGateways in the indexer.
func (s *awsEgressOnlyInternetGatewayLister) List(selector labels.Selector) (ret []*v1.AwsEgressOnlyInternetGateway, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsEgressOnlyInternetGateway))
	})
	return ret, err
}

// AwsEgressOnlyInternetGateways returns an object that can list and get AwsEgressOnlyInternetGateways.
func (s *awsEgressOnlyInternetGatewayLister) AwsEgressOnlyInternetGateways(namespace string) AwsEgressOnlyInternetGatewayNamespaceLister {
	return awsEgressOnlyInternetGatewayNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsEgressOnlyInternetGatewayNamespaceLister helps list and get AwsEgressOnlyInternetGateways.
type AwsEgressOnlyInternetGatewayNamespaceLister interface {
	// List lists all AwsEgressOnlyInternetGateways in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsEgressOnlyInternetGateway, err error)
	// Get retrieves the AwsEgressOnlyInternetGateway from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsEgressOnlyInternetGateway, error)
	AwsEgressOnlyInternetGatewayNamespaceListerExpansion
}

// awsEgressOnlyInternetGatewayNamespaceLister implements the AwsEgressOnlyInternetGatewayNamespaceLister
// interface.
type awsEgressOnlyInternetGatewayNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsEgressOnlyInternetGateways in the indexer for a given namespace.
func (s awsEgressOnlyInternetGatewayNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsEgressOnlyInternetGateway, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsEgressOnlyInternetGateway))
	})
	return ret, err
}

// Get retrieves the AwsEgressOnlyInternetGateway from the indexer for a given namespace and name.
func (s awsEgressOnlyInternetGatewayNamespaceLister) Get(name string) (*v1.AwsEgressOnlyInternetGateway, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsegressonlyinternetgateway"), name)
	}
	return obj.(*v1.AwsEgressOnlyInternetGateway), nil
}
