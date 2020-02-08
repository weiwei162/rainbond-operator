// RAINBOND, Application Management Platform
// Copyright (C) 2014-2020 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/goodrain/rainbond-operator/pkg/apis/rainbond/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RainbondClusterLister helps list RainbondClusters.
type RainbondClusterLister interface {
	// List lists all RainbondClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RainbondCluster, err error)
	// RainbondClusters returns an object that can list and get RainbondClusters.
	RainbondClusters(namespace string) RainbondClusterNamespaceLister
	RainbondClusterListerExpansion
}

// rainbondClusterLister implements the RainbondClusterLister interface.
type rainbondClusterLister struct {
	indexer cache.Indexer
}

// NewRainbondClusterLister returns a new RainbondClusterLister.
func NewRainbondClusterLister(indexer cache.Indexer) RainbondClusterLister {
	return &rainbondClusterLister{indexer: indexer}
}

// List lists all RainbondClusters in the indexer.
func (s *rainbondClusterLister) List(selector labels.Selector) (ret []*v1alpha1.RainbondCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RainbondCluster))
	})
	return ret, err
}

// RainbondClusters returns an object that can list and get RainbondClusters.
func (s *rainbondClusterLister) RainbondClusters(namespace string) RainbondClusterNamespaceLister {
	return rainbondClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RainbondClusterNamespaceLister helps list and get RainbondClusters.
type RainbondClusterNamespaceLister interface {
	// List lists all RainbondClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RainbondCluster, err error)
	// Get retrieves the RainbondCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RainbondCluster, error)
	RainbondClusterNamespaceListerExpansion
}

// rainbondClusterNamespaceLister implements the RainbondClusterNamespaceLister
// interface.
type rainbondClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RainbondClusters in the indexer for a given namespace.
func (s rainbondClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RainbondCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RainbondCluster))
	})
	return ret, err
}

// Get retrieves the RainbondCluster from the indexer for a given namespace and name.
func (s rainbondClusterNamespaceLister) Get(name string) (*v1alpha1.RainbondCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("rainbondcluster"), name)
	}
	return obj.(*v1alpha1.RainbondCluster), nil
}
