// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/GoogleCloudPlatform/prometheus-engine/pkg/operator/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PodMonitoringLister helps list PodMonitorings.
// All objects returned here must be treated as read-only.
type PodMonitoringLister interface {
	// List lists all PodMonitorings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PodMonitoring, err error)
	// PodMonitorings returns an object that can list and get PodMonitorings.
	PodMonitorings(namespace string) PodMonitoringNamespaceLister
	PodMonitoringListerExpansion
}

// podMonitoringLister implements the PodMonitoringLister interface.
type podMonitoringLister struct {
	indexer cache.Indexer
}

// NewPodMonitoringLister returns a new PodMonitoringLister.
func NewPodMonitoringLister(indexer cache.Indexer) PodMonitoringLister {
	return &podMonitoringLister{indexer: indexer}
}

// List lists all PodMonitorings in the indexer.
func (s *podMonitoringLister) List(selector labels.Selector) (ret []*v1.PodMonitoring, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PodMonitoring))
	})
	return ret, err
}

// PodMonitorings returns an object that can list and get PodMonitorings.
func (s *podMonitoringLister) PodMonitorings(namespace string) PodMonitoringNamespaceLister {
	return podMonitoringNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PodMonitoringNamespaceLister helps list and get PodMonitorings.
// All objects returned here must be treated as read-only.
type PodMonitoringNamespaceLister interface {
	// List lists all PodMonitorings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PodMonitoring, err error)
	// Get retrieves the PodMonitoring from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.PodMonitoring, error)
	PodMonitoringNamespaceListerExpansion
}

// podMonitoringNamespaceLister implements the PodMonitoringNamespaceLister
// interface.
type podMonitoringNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PodMonitorings in the indexer for a given namespace.
func (s podMonitoringNamespaceLister) List(selector labels.Selector) (ret []*v1.PodMonitoring, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PodMonitoring))
	})
	return ret, err
}

// Get retrieves the PodMonitoring from the indexer for a given namespace and name.
func (s podMonitoringNamespaceLister) Get(name string) (*v1.PodMonitoring, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("podmonitoring"), name)
	}
	return obj.(*v1.PodMonitoring), nil
}
