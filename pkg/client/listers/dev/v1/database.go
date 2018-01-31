/*
Copyright 2018 The Openshift Evangelists

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

// This file was automatically generated by lister-gen

package v1

import (
	v1 "crds/pkg/apis/dev.kubernetes.pavanvelidandla.com/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DatabaseLister helps list Databases.
type DatabaseLister interface {
	// List lists all Databases in the indexer.
	List(selector labels.Selector) (ret []*v1.Database, err error)
	// Databases returns an object that can list and get Databases.
	Databases(namespace string) DatabaseNamespaceLister
	DatabaseListerExpansion
}

// databaseLister implements the DatabaseLister interface.
type databaseLister struct {
	indexer cache.Indexer
}

// NewDatabaseLister returns a new DatabaseLister.
func NewDatabaseLister(indexer cache.Indexer) DatabaseLister {
	return &databaseLister{indexer: indexer}
}

// List lists all Databases in the indexer.
func (s *databaseLister) List(selector labels.Selector) (ret []*v1.Database, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Database))
	})
	return ret, err
}

// Databases returns an object that can list and get Databases.
func (s *databaseLister) Databases(namespace string) DatabaseNamespaceLister {
	return databaseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DatabaseNamespaceLister helps list and get Databases.
type DatabaseNamespaceLister interface {
	// List lists all Databases in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Database, err error)
	// Get retrieves the Database from the indexer for a given namespace and name.
	Get(name string) (*v1.Database, error)
	DatabaseNamespaceListerExpansion
}

// databaseNamespaceLister implements the DatabaseNamespaceLister
// interface.
type databaseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Databases in the indexer for a given namespace.
func (s databaseNamespaceLister) List(selector labels.Selector) (ret []*v1.Database, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Database))
	})
	return ret, err
}

// Get retrieves the Database from the indexer for a given namespace and name.
func (s databaseNamespaceLister) Get(name string) (*v1.Database, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("database"), name)
	}
	return obj.(*v1.Database), nil
}
