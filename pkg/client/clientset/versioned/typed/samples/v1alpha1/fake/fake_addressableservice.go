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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "knative.dev/super-controller/pkg/apis/samples/v1alpha1"
)

// FakeAddressableServices implements AddressableServiceInterface
type FakeAddressableServices struct {
	Fake *FakeSamplesV1alpha1
	ns   string
}

var addressableservicesResource = schema.GroupVersionResource{Group: "samples.knative.dev", Version: "v1alpha1", Resource: "addressableservices"}

var addressableservicesKind = schema.GroupVersionKind{Group: "samples.knative.dev", Version: "v1alpha1", Kind: "AddressableService"}

// Get takes name of the addressableService, and returns the corresponding addressableService object, and an error if there is any.
func (c *FakeAddressableServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AddressableService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(addressableservicesResource, c.ns, name), &v1alpha1.AddressableService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AddressableService), err
}

// List takes label and field selectors, and returns the list of AddressableServices that match those selectors.
func (c *FakeAddressableServices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AddressableServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(addressableservicesResource, addressableservicesKind, c.ns, opts), &v1alpha1.AddressableServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AddressableServiceList{ListMeta: obj.(*v1alpha1.AddressableServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AddressableServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested addressableServices.
func (c *FakeAddressableServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(addressableservicesResource, c.ns, opts))

}

// Create takes the representation of a addressableService and creates it.  Returns the server's representation of the addressableService, and an error, if there is any.
func (c *FakeAddressableServices) Create(ctx context.Context, addressableService *v1alpha1.AddressableService, opts v1.CreateOptions) (result *v1alpha1.AddressableService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(addressableservicesResource, c.ns, addressableService), &v1alpha1.AddressableService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AddressableService), err
}

// Update takes the representation of a addressableService and updates it. Returns the server's representation of the addressableService, and an error, if there is any.
func (c *FakeAddressableServices) Update(ctx context.Context, addressableService *v1alpha1.AddressableService, opts v1.UpdateOptions) (result *v1alpha1.AddressableService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(addressableservicesResource, c.ns, addressableService), &v1alpha1.AddressableService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AddressableService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAddressableServices) UpdateStatus(ctx context.Context, addressableService *v1alpha1.AddressableService, opts v1.UpdateOptions) (*v1alpha1.AddressableService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(addressableservicesResource, "status", c.ns, addressableService), &v1alpha1.AddressableService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AddressableService), err
}

// Delete takes name of the addressableService and deletes it. Returns an error if one occurs.
func (c *FakeAddressableServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(addressableservicesResource, c.ns, name), &v1alpha1.AddressableService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAddressableServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(addressableservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AddressableServiceList{})
	return err
}

// Patch applies the patch and returns the patched addressableService.
func (c *FakeAddressableServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AddressableService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(addressableservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AddressableService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AddressableService), err
}
