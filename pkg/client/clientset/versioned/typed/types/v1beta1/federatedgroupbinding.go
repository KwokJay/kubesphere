/*
Copyright 2020 The KubeSphere Authors.

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

package v1beta1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1beta1 "kubesphere.io/kubesphere/pkg/apis/types/v1beta1"
	scheme "kubesphere.io/kubesphere/pkg/client/clientset/versioned/scheme"
)

// FederatedGroupBindingsGetter has a method to return a FederatedGroupBindingInterface.
// A group's client should implement this interface.
type FederatedGroupBindingsGetter interface {
	FederatedGroupBindings() FederatedGroupBindingInterface
}

// FederatedGroupBindingInterface has methods to work with FederatedGroupBinding resources.
type FederatedGroupBindingInterface interface {
	Create(*v1beta1.FederatedGroupBinding) (*v1beta1.FederatedGroupBinding, error)
	Update(*v1beta1.FederatedGroupBinding) (*v1beta1.FederatedGroupBinding, error)
	UpdateStatus(*v1beta1.FederatedGroupBinding) (*v1beta1.FederatedGroupBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.FederatedGroupBinding, error)
	List(opts v1.ListOptions) (*v1beta1.FederatedGroupBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedGroupBinding, err error)
	FederatedGroupBindingExpansion
}

// federatedGroupBindings implements FederatedGroupBindingInterface
type federatedGroupBindings struct {
	client rest.Interface
}

// newFederatedGroupBindings returns a FederatedGroupBindings
func newFederatedGroupBindings(c *TypesV1beta1Client) *federatedGroupBindings {
	return &federatedGroupBindings{
		client: c.RESTClient(),
	}
}

// Get takes name of the federatedGroupBinding, and returns the corresponding federatedGroupBinding object, and an error if there is any.
func (c *federatedGroupBindings) Get(name string, options v1.GetOptions) (result *v1beta1.FederatedGroupBinding, err error) {
	result = &v1beta1.FederatedGroupBinding{}
	err = c.client.Get().
		Resource("federatedgroupbindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedGroupBindings that match those selectors.
func (c *federatedGroupBindings) List(opts v1.ListOptions) (result *v1beta1.FederatedGroupBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.FederatedGroupBindingList{}
	err = c.client.Get().
		Resource("federatedgroupbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedGroupBindings.
func (c *federatedGroupBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("federatedgroupbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a federatedGroupBinding and creates it.  Returns the server's representation of the federatedGroupBinding, and an error, if there is any.
func (c *federatedGroupBindings) Create(federatedGroupBinding *v1beta1.FederatedGroupBinding) (result *v1beta1.FederatedGroupBinding, err error) {
	result = &v1beta1.FederatedGroupBinding{}
	err = c.client.Post().
		Resource("federatedgroupbindings").
		Body(federatedGroupBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a federatedGroupBinding and updates it. Returns the server's representation of the federatedGroupBinding, and an error, if there is any.
func (c *federatedGroupBindings) Update(federatedGroupBinding *v1beta1.FederatedGroupBinding) (result *v1beta1.FederatedGroupBinding, err error) {
	result = &v1beta1.FederatedGroupBinding{}
	err = c.client.Put().
		Resource("federatedgroupbindings").
		Name(federatedGroupBinding.Name).
		Body(federatedGroupBinding).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *federatedGroupBindings) UpdateStatus(federatedGroupBinding *v1beta1.FederatedGroupBinding) (result *v1beta1.FederatedGroupBinding, err error) {
	result = &v1beta1.FederatedGroupBinding{}
	err = c.client.Put().
		Resource("federatedgroupbindings").
		Name(federatedGroupBinding.Name).
		SubResource("status").
		Body(federatedGroupBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the federatedGroupBinding and deletes it. Returns an error if one occurs.
func (c *federatedGroupBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("federatedgroupbindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedGroupBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("federatedgroupbindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched federatedGroupBinding.
func (c *federatedGroupBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedGroupBinding, err error) {
	result = &v1beta1.FederatedGroupBinding{}
	err = c.client.Patch(pt).
		Resource("federatedgroupbindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
