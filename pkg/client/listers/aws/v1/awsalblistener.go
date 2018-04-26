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

// AwsAlbListenerLister helps list AwsAlbListeners.
type AwsAlbListenerLister interface {
	// List lists all AwsAlbListeners in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsAlbListener, err error)
	// AwsAlbListeners returns an object that can list and get AwsAlbListeners.
	AwsAlbListeners(namespace string) AwsAlbListenerNamespaceLister
	AwsAlbListenerListerExpansion
}

// awsAlbListenerLister implements the AwsAlbListenerLister interface.
type awsAlbListenerLister struct {
	indexer cache.Indexer
}

// NewAwsAlbListenerLister returns a new AwsAlbListenerLister.
func NewAwsAlbListenerLister(indexer cache.Indexer) AwsAlbListenerLister {
	return &awsAlbListenerLister{indexer: indexer}
}

// List lists all AwsAlbListeners in the indexer.
func (s *awsAlbListenerLister) List(selector labels.Selector) (ret []*v1.AwsAlbListener, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsAlbListener))
	})
	return ret, err
}

// AwsAlbListeners returns an object that can list and get AwsAlbListeners.
func (s *awsAlbListenerLister) AwsAlbListeners(namespace string) AwsAlbListenerNamespaceLister {
	return awsAlbListenerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsAlbListenerNamespaceLister helps list and get AwsAlbListeners.
type AwsAlbListenerNamespaceLister interface {
	// List lists all AwsAlbListeners in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsAlbListener, err error)
	// Get retrieves the AwsAlbListener from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsAlbListener, error)
	AwsAlbListenerNamespaceListerExpansion
}

// awsAlbListenerNamespaceLister implements the AwsAlbListenerNamespaceLister
// interface.
type awsAlbListenerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsAlbListeners in the indexer for a given namespace.
func (s awsAlbListenerNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsAlbListener, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsAlbListener))
	})
	return ret, err
}

// Get retrieves the AwsAlbListener from the indexer for a given namespace and name.
func (s awsAlbListenerNamespaceLister) Get(name string) (*v1.AwsAlbListener, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsalblistener"), name)
	}
	return obj.(*v1.AwsAlbListener), nil
}