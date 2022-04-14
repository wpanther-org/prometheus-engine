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

// ClusterRulesLister helps list ClusterRules.
// All objects returned here must be treated as read-only.
type ClusterRulesLister interface {
	// List lists all ClusterRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterRules, err error)
	// Get retrieves the ClusterRules from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ClusterRules, error)
	ClusterRulesListerExpansion
}

// clusterRulesLister implements the ClusterRulesLister interface.
type clusterRulesLister struct {
	indexer cache.Indexer
}

// NewClusterRulesLister returns a new ClusterRulesLister.
func NewClusterRulesLister(indexer cache.Indexer) ClusterRulesLister {
	return &clusterRulesLister{indexer: indexer}
}

// List lists all ClusterRules in the indexer.
func (s *clusterRulesLister) List(selector labels.Selector) (ret []*v1.ClusterRules, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterRules))
	})
	return ret, err
}

// Get retrieves the ClusterRules from the index for a given name.
func (s *clusterRulesLister) Get(name string) (*v1.ClusterRules, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clusterrules"), name)
	}
	return obj.(*v1.ClusterRules), nil
}
