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

// FakeApps implements AppInterface
type FakeApps struct {
	Fake *FakeAppsV1alpha1
}

var appsResource = schema.GroupVersionResource{Group: "apps.cloudrobotics.com", Version: "v1alpha1", Resource: "apps"}

var appsKind = schema.GroupVersionKind{Group: "apps.cloudrobotics.com", Version: "v1alpha1", Kind: "App"}

// Get takes name of the app, and returns the corresponding app object, and an error if there is any.
func (c *FakeApps) Get(name string, options v1.GetOptions) (result *v1alpha1.App, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(appsResource, name), &v1alpha1.App{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.App), err
}

// List takes label and field selectors, and returns the list of Apps that match those selectors.
func (c *FakeApps) List(opts v1.ListOptions) (result *v1alpha1.AppList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(appsResource, appsKind, opts), &v1alpha1.AppList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AppList{ListMeta: obj.(*v1alpha1.AppList).ListMeta}
	for _, item := range obj.(*v1alpha1.AppList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apps.
func (c *FakeApps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(appsResource, opts))
}

// Create takes the representation of a app and creates it.  Returns the server's representation of the app, and an error, if there is any.
func (c *FakeApps) Create(app *v1alpha1.App) (result *v1alpha1.App, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(appsResource, app), &v1alpha1.App{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.App), err
}

// Update takes the representation of a app and updates it. Returns the server's representation of the app, and an error, if there is any.
func (c *FakeApps) Update(app *v1alpha1.App) (result *v1alpha1.App, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(appsResource, app), &v1alpha1.App{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.App), err
}

// Delete takes name of the app and deletes it. Returns an error if one occurs.
func (c *FakeApps) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(appsResource, name), &v1alpha1.App{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(appsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AppList{})
	return err
}

// Patch applies the patch and returns the patched app.
func (c *FakeApps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.App, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(appsResource, name, pt, data, subresources...), &v1alpha1.App{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.App), err
}
