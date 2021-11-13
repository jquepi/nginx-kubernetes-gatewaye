// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/nginxinc/nginx-gateway-kubernetes/pkg/apis/v1alpha1"
	scheme "github.com/nginxinc/nginx-gateway-kubernetes/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GatewayConfigsGetter has a method to return a GatewayConfigInterface.
// A group's client should implement this interface.
type GatewayConfigsGetter interface {
	GatewayConfigs() GatewayConfigInterface
}

// GatewayConfigInterface has methods to work with GatewayConfig resources.
type GatewayConfigInterface interface {
	Create(ctx context.Context, gatewayConfig *v1alpha1.GatewayConfig, opts v1.CreateOptions) (*v1alpha1.GatewayConfig, error)
	Update(ctx context.Context, gatewayConfig *v1alpha1.GatewayConfig, opts v1.UpdateOptions) (*v1alpha1.GatewayConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.GatewayConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.GatewayConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GatewayConfig, err error)
	GatewayConfigExpansion
}

// gatewayConfigs implements GatewayConfigInterface
type gatewayConfigs struct {
	client rest.Interface
}

// newGatewayConfigs returns a GatewayConfigs
func newGatewayConfigs(c *GatewayV1alpha1Client) *gatewayConfigs {
	return &gatewayConfigs{
		client: c.RESTClient(),
	}
}

// Get takes name of the gatewayConfig, and returns the corresponding gatewayConfig object, and an error if there is any.
func (c *gatewayConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GatewayConfig, err error) {
	result = &v1alpha1.GatewayConfig{}
	err = c.client.Get().
		Resource("gatewayconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GatewayConfigs that match those selectors.
func (c *gatewayConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GatewayConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GatewayConfigList{}
	err = c.client.Get().
		Resource("gatewayconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gatewayConfigs.
func (c *gatewayConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("gatewayconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a gatewayConfig and creates it.  Returns the server's representation of the gatewayConfig, and an error, if there is any.
func (c *gatewayConfigs) Create(ctx context.Context, gatewayConfig *v1alpha1.GatewayConfig, opts v1.CreateOptions) (result *v1alpha1.GatewayConfig, err error) {
	result = &v1alpha1.GatewayConfig{}
	err = c.client.Post().
		Resource("gatewayconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gatewayConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a gatewayConfig and updates it. Returns the server's representation of the gatewayConfig, and an error, if there is any.
func (c *gatewayConfigs) Update(ctx context.Context, gatewayConfig *v1alpha1.GatewayConfig, opts v1.UpdateOptions) (result *v1alpha1.GatewayConfig, err error) {
	result = &v1alpha1.GatewayConfig{}
	err = c.client.Put().
		Resource("gatewayconfigs").
		Name(gatewayConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gatewayConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the gatewayConfig and deletes it. Returns an error if one occurs.
func (c *gatewayConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("gatewayconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gatewayConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("gatewayconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched gatewayConfig.
func (c *gatewayConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GatewayConfig, err error) {
	result = &v1alpha1.GatewayConfig{}
	err = c.client.Patch(pt).
		Resource("gatewayconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
