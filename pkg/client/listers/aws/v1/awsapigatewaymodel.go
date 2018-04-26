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

// AwsApiGatewayModelLister helps list AwsApiGatewayModels.
type AwsApiGatewayModelLister interface {
	// List lists all AwsApiGatewayModels in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsApiGatewayModel, err error)
	// AwsApiGatewayModels returns an object that can list and get AwsApiGatewayModels.
	AwsApiGatewayModels(namespace string) AwsApiGatewayModelNamespaceLister
	AwsApiGatewayModelListerExpansion
}

// awsApiGatewayModelLister implements the AwsApiGatewayModelLister interface.
type awsApiGatewayModelLister struct {
	indexer cache.Indexer
}

// NewAwsApiGatewayModelLister returns a new AwsApiGatewayModelLister.
func NewAwsApiGatewayModelLister(indexer cache.Indexer) AwsApiGatewayModelLister {
	return &awsApiGatewayModelLister{indexer: indexer}
}

// List lists all AwsApiGatewayModels in the indexer.
func (s *awsApiGatewayModelLister) List(selector labels.Selector) (ret []*v1.AwsApiGatewayModel, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsApiGatewayModel))
	})
	return ret, err
}

// AwsApiGatewayModels returns an object that can list and get AwsApiGatewayModels.
func (s *awsApiGatewayModelLister) AwsApiGatewayModels(namespace string) AwsApiGatewayModelNamespaceLister {
	return awsApiGatewayModelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsApiGatewayModelNamespaceLister helps list and get AwsApiGatewayModels.
type AwsApiGatewayModelNamespaceLister interface {
	// List lists all AwsApiGatewayModels in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsApiGatewayModel, err error)
	// Get retrieves the AwsApiGatewayModel from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsApiGatewayModel, error)
	AwsApiGatewayModelNamespaceListerExpansion
}

// awsApiGatewayModelNamespaceLister implements the AwsApiGatewayModelNamespaceLister
// interface.
type awsApiGatewayModelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsApiGatewayModels in the indexer for a given namespace.
func (s awsApiGatewayModelNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsApiGatewayModel, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsApiGatewayModel))
	})
	return ret, err
}

// Get retrieves the AwsApiGatewayModel from the indexer for a given namespace and name.
func (s awsApiGatewayModelNamespaceLister) Get(name string) (*v1.AwsApiGatewayModel, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsapigatewaymodel"), name)
	}
	return obj.(*v1.AwsApiGatewayModel), nil
}