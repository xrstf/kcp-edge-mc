/*
Copyright The KubeStellar Authors.

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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "github.com/kcp-dev/edge-mc/pkg/apis/edge/v1alpha1"
	scheme "github.com/kcp-dev/edge-mc/pkg/client/clientset/versioned/scheme"
)

// EdgeSyncConfigsGetter has a method to return a EdgeSyncConfigInterface.
// A group's client should implement this interface.
type EdgeSyncConfigsGetter interface {
	EdgeSyncConfigs() EdgeSyncConfigInterface
}

// EdgeSyncConfigInterface has methods to work with EdgeSyncConfig resources.
type EdgeSyncConfigInterface interface {
	Create(ctx context.Context, edgeSyncConfig *v1alpha1.EdgeSyncConfig, opts v1.CreateOptions) (*v1alpha1.EdgeSyncConfig, error)
	Update(ctx context.Context, edgeSyncConfig *v1alpha1.EdgeSyncConfig, opts v1.UpdateOptions) (*v1alpha1.EdgeSyncConfig, error)
	UpdateStatus(ctx context.Context, edgeSyncConfig *v1alpha1.EdgeSyncConfig, opts v1.UpdateOptions) (*v1alpha1.EdgeSyncConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.EdgeSyncConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.EdgeSyncConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EdgeSyncConfig, err error)
	EdgeSyncConfigExpansion
}

// edgeSyncConfigs implements EdgeSyncConfigInterface
type edgeSyncConfigs struct {
	client rest.Interface
}

// newEdgeSyncConfigs returns a EdgeSyncConfigs
func newEdgeSyncConfigs(c *EdgeV1alpha1Client) *edgeSyncConfigs {
	return &edgeSyncConfigs{
		client: c.RESTClient(),
	}
}

// Get takes name of the edgeSyncConfig, and returns the corresponding edgeSyncConfig object, and an error if there is any.
func (c *edgeSyncConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EdgeSyncConfig, err error) {
	result = &v1alpha1.EdgeSyncConfig{}
	err = c.client.Get().
		Resource("edgesyncconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EdgeSyncConfigs that match those selectors.
func (c *edgeSyncConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EdgeSyncConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EdgeSyncConfigList{}
	err = c.client.Get().
		Resource("edgesyncconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested edgeSyncConfigs.
func (c *edgeSyncConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("edgesyncconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a edgeSyncConfig and creates it.  Returns the server's representation of the edgeSyncConfig, and an error, if there is any.
func (c *edgeSyncConfigs) Create(ctx context.Context, edgeSyncConfig *v1alpha1.EdgeSyncConfig, opts v1.CreateOptions) (result *v1alpha1.EdgeSyncConfig, err error) {
	result = &v1alpha1.EdgeSyncConfig{}
	err = c.client.Post().
		Resource("edgesyncconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(edgeSyncConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a edgeSyncConfig and updates it. Returns the server's representation of the edgeSyncConfig, and an error, if there is any.
func (c *edgeSyncConfigs) Update(ctx context.Context, edgeSyncConfig *v1alpha1.EdgeSyncConfig, opts v1.UpdateOptions) (result *v1alpha1.EdgeSyncConfig, err error) {
	result = &v1alpha1.EdgeSyncConfig{}
	err = c.client.Put().
		Resource("edgesyncconfigs").
		Name(edgeSyncConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(edgeSyncConfig).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *edgeSyncConfigs) UpdateStatus(ctx context.Context, edgeSyncConfig *v1alpha1.EdgeSyncConfig, opts v1.UpdateOptions) (result *v1alpha1.EdgeSyncConfig, err error) {
	result = &v1alpha1.EdgeSyncConfig{}
	err = c.client.Put().
		Resource("edgesyncconfigs").
		Name(edgeSyncConfig.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(edgeSyncConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the edgeSyncConfig and deletes it. Returns an error if one occurs.
func (c *edgeSyncConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("edgesyncconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *edgeSyncConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("edgesyncconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched edgeSyncConfig.
func (c *edgeSyncConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EdgeSyncConfig, err error) {
	result = &v1alpha1.EdgeSyncConfig{}
	err = c.client.Patch(pt).
		Resource("edgesyncconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
