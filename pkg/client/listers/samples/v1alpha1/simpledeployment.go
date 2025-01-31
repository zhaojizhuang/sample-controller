/*
Copyright 2020 The Knative Authors

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "knative.dev/super-controller/pkg/apis/samples/v1alpha1"
)

// SimpleDeploymentLister helps list SimpleDeployments.
// All objects returned here must be treated as read-only.
type SimpleDeploymentLister interface {
	// List lists all SimpleDeployments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SimpleDeployment, err error)
	// SimpleDeployments returns an object that can list and get SimpleDeployments.
	SimpleDeployments(namespace string) SimpleDeploymentNamespaceLister
	SimpleDeploymentListerExpansion
}

// simpleDeploymentLister implements the SimpleDeploymentLister interface.
type simpleDeploymentLister struct {
	indexer cache.Indexer
}

// NewSimpleDeploymentLister returns a new SimpleDeploymentLister.
func NewSimpleDeploymentLister(indexer cache.Indexer) SimpleDeploymentLister {
	return &simpleDeploymentLister{indexer: indexer}
}

// List lists all SimpleDeployments in the indexer.
func (s *simpleDeploymentLister) List(selector labels.Selector) (ret []*v1alpha1.SimpleDeployment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SimpleDeployment))
	})
	return ret, err
}

// SimpleDeployments returns an object that can list and get SimpleDeployments.
func (s *simpleDeploymentLister) SimpleDeployments(namespace string) SimpleDeploymentNamespaceLister {
	return simpleDeploymentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SimpleDeploymentNamespaceLister helps list and get SimpleDeployments.
// All objects returned here must be treated as read-only.
type SimpleDeploymentNamespaceLister interface {
	// List lists all SimpleDeployments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SimpleDeployment, err error)
	// Get retrieves the SimpleDeployment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SimpleDeployment, error)
	SimpleDeploymentNamespaceListerExpansion
}

// simpleDeploymentNamespaceLister implements the SimpleDeploymentNamespaceLister
// interface.
type simpleDeploymentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SimpleDeployments in the indexer for a given namespace.
func (s simpleDeploymentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SimpleDeployment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SimpleDeployment))
	})
	return ret, err
}

// Get retrieves the SimpleDeployment from the indexer for a given namespace and name.
func (s simpleDeploymentNamespaceLister) Get(name string) (*v1alpha1.SimpleDeployment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("simpledeployment"), name)
	}
	return obj.(*v1alpha1.SimpleDeployment), nil
}
