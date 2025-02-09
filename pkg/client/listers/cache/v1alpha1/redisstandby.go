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

package v1alpha1

import (
	v1alpha1 "redis-priv-operator/pkg/apis/cache/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RedisStandbyLister helps list RedisStandbies.
type RedisStandbyLister interface {
	// List lists all RedisStandbies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RedisStandby, err error)
	// RedisStandbies returns an object that can list and get RedisStandbies.
	RedisStandbies(namespace string) RedisStandbyNamespaceLister
	RedisStandbyListerExpansion
}

// redisStandbyLister implements the RedisStandbyLister interface.
type redisStandbyLister struct {
	indexer cache.Indexer
}

// NewRedisStandbyLister returns a new RedisStandbyLister.
func NewRedisStandbyLister(indexer cache.Indexer) RedisStandbyLister {
	return &redisStandbyLister{indexer: indexer}
}

// List lists all RedisStandbies in the indexer.
func (s *redisStandbyLister) List(selector labels.Selector) (ret []*v1alpha1.RedisStandby, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RedisStandby))
	})
	return ret, err
}

// RedisStandbies returns an object that can list and get RedisStandbies.
func (s *redisStandbyLister) RedisStandbies(namespace string) RedisStandbyNamespaceLister {
	return redisStandbyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RedisStandbyNamespaceLister helps list and get RedisStandbies.
type RedisStandbyNamespaceLister interface {
	// List lists all RedisStandbies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RedisStandby, err error)
	// Get retrieves the RedisStandby from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RedisStandby, error)
	RedisStandbyNamespaceListerExpansion
}

// redisStandbyNamespaceLister implements the RedisStandbyNamespaceLister
// interface.
type redisStandbyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RedisStandbies in the indexer for a given namespace.
func (s redisStandbyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RedisStandby, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RedisStandby))
	})
	return ret, err
}

// Get retrieves the RedisStandby from the indexer for a given namespace and name.
func (s redisStandbyNamespaceLister) Get(name string) (*v1alpha1.RedisStandby, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("redisstandby"), name)
	}
	return obj.(*v1alpha1.RedisStandby), nil
}
