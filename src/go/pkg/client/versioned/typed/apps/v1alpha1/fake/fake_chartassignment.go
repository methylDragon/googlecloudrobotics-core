// Copyright 2019 The Google Cloud Robotics Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/googlecloudrobotics/core/src/go/pkg/apis/apps/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeChartAssignments implements ChartAssignmentInterface
type FakeChartAssignments struct {
	Fake *FakeAppsV1alpha1
}

var chartassignmentsResource = schema.GroupVersionResource{Group: "apps.cloudrobotics.com", Version: "v1alpha1", Resource: "chartassignments"}

var chartassignmentsKind = schema.GroupVersionKind{Group: "apps.cloudrobotics.com", Version: "v1alpha1", Kind: "ChartAssignment"}

// Get takes name of the chartAssignment, and returns the corresponding chartAssignment object, and an error if there is any.
func (c *FakeChartAssignments) Get(name string, options v1.GetOptions) (result *v1alpha1.ChartAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(chartassignmentsResource, name), &v1alpha1.ChartAssignment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChartAssignment), err
}

// List takes label and field selectors, and returns the list of ChartAssignments that match those selectors.
func (c *FakeChartAssignments) List(opts v1.ListOptions) (result *v1alpha1.ChartAssignmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(chartassignmentsResource, chartassignmentsKind, opts), &v1alpha1.ChartAssignmentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ChartAssignmentList{ListMeta: obj.(*v1alpha1.ChartAssignmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.ChartAssignmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested chartAssignments.
func (c *FakeChartAssignments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(chartassignmentsResource, opts))
}

// Create takes the representation of a chartAssignment and creates it.  Returns the server's representation of the chartAssignment, and an error, if there is any.
func (c *FakeChartAssignments) Create(chartAssignment *v1alpha1.ChartAssignment) (result *v1alpha1.ChartAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(chartassignmentsResource, chartAssignment), &v1alpha1.ChartAssignment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChartAssignment), err
}

// Update takes the representation of a chartAssignment and updates it. Returns the server's representation of the chartAssignment, and an error, if there is any.
func (c *FakeChartAssignments) Update(chartAssignment *v1alpha1.ChartAssignment) (result *v1alpha1.ChartAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(chartassignmentsResource, chartAssignment), &v1alpha1.ChartAssignment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChartAssignment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeChartAssignments) UpdateStatus(chartAssignment *v1alpha1.ChartAssignment) (*v1alpha1.ChartAssignment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(chartassignmentsResource, "status", chartAssignment), &v1alpha1.ChartAssignment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChartAssignment), err
}

// Delete takes name of the chartAssignment and deletes it. Returns an error if one occurs.
func (c *FakeChartAssignments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(chartassignmentsResource, name), &v1alpha1.ChartAssignment{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeChartAssignments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(chartassignmentsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ChartAssignmentList{})
	return err
}

// Patch applies the patch and returns the patched chartAssignment.
func (c *FakeChartAssignments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ChartAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(chartassignmentsResource, name, pt, data, subresources...), &v1alpha1.ChartAssignment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChartAssignment), err
}
