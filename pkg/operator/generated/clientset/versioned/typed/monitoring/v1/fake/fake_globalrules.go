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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	monitoringv1 "github.com/GoogleCloudPlatform/prometheus-engine/pkg/operator/apis/monitoring/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGlobalRules implements GlobalRulesInterface
type FakeGlobalRules struct {
	Fake *FakeMonitoringV1
}

var globalrulesResource = schema.GroupVersionResource{Group: "monitoring.googleapis.com", Version: "v1", Resource: "globalrules"}

var globalrulesKind = schema.GroupVersionKind{Group: "monitoring.googleapis.com", Version: "v1", Kind: "GlobalRules"}

// Get takes name of the globalRules, and returns the corresponding globalRules object, and an error if there is any.
func (c *FakeGlobalRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *monitoringv1.GlobalRules, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(globalrulesResource, name), &monitoringv1.GlobalRules{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.GlobalRules), err
}

// List takes label and field selectors, and returns the list of GlobalRules that match those selectors.
func (c *FakeGlobalRules) List(ctx context.Context, opts v1.ListOptions) (result *monitoringv1.GlobalRulesList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(globalrulesResource, globalrulesKind, opts), &monitoringv1.GlobalRulesList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &monitoringv1.GlobalRulesList{ListMeta: obj.(*monitoringv1.GlobalRulesList).ListMeta}
	for _, item := range obj.(*monitoringv1.GlobalRulesList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested globalRules.
func (c *FakeGlobalRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(globalrulesResource, opts))
}

// Create takes the representation of a globalRules and creates it.  Returns the server's representation of the globalRules, and an error, if there is any.
func (c *FakeGlobalRules) Create(ctx context.Context, globalRules *monitoringv1.GlobalRules, opts v1.CreateOptions) (result *monitoringv1.GlobalRules, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(globalrulesResource, globalRules), &monitoringv1.GlobalRules{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.GlobalRules), err
}

// Update takes the representation of a globalRules and updates it. Returns the server's representation of the globalRules, and an error, if there is any.
func (c *FakeGlobalRules) Update(ctx context.Context, globalRules *monitoringv1.GlobalRules, opts v1.UpdateOptions) (result *monitoringv1.GlobalRules, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(globalrulesResource, globalRules), &monitoringv1.GlobalRules{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.GlobalRules), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGlobalRules) UpdateStatus(ctx context.Context, globalRules *monitoringv1.GlobalRules, opts v1.UpdateOptions) (*monitoringv1.GlobalRules, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(globalrulesResource, "status", globalRules), &monitoringv1.GlobalRules{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.GlobalRules), err
}

// Delete takes name of the globalRules and deletes it. Returns an error if one occurs.
func (c *FakeGlobalRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(globalrulesResource, name), &monitoringv1.GlobalRules{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGlobalRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(globalrulesResource, listOpts)

	_, err := c.Fake.Invokes(action, &monitoringv1.GlobalRulesList{})
	return err
}

// Patch applies the patch and returns the patched globalRules.
func (c *FakeGlobalRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *monitoringv1.GlobalRules, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(globalrulesResource, name, pt, data, subresources...), &monitoringv1.GlobalRules{})
	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.GlobalRules), err
}
