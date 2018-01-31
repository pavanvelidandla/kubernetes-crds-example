/*
Copyright 2018 The Openshift Evangelists

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
package v1

import (
	v1 "crds/pkg/apis/dev.kubernetes.pavanvelidandla.com/v1"
	scheme "crds/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ConfigFromGitsGetter has a method to return a ConfigFromGitInterface.
// A group's client should implement this interface.
type ConfigFromGitsGetter interface {
	ConfigFromGits(namespace string) ConfigFromGitInterface
}

// ConfigFromGitInterface has methods to work with ConfigFromGit resources.
type ConfigFromGitInterface interface {
	Create(*v1.ConfigFromGit) (*v1.ConfigFromGit, error)
	Update(*v1.ConfigFromGit) (*v1.ConfigFromGit, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.ConfigFromGit, error)
	List(opts meta_v1.ListOptions) (*v1.ConfigFromGitList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ConfigFromGit, err error)
	ConfigFromGitExpansion
}

// configFromGits implements ConfigFromGitInterface
type configFromGits struct {
	client rest.Interface
	ns     string
}

// newConfigFromGits returns a ConfigFromGits
func newConfigFromGits(c *DevV1Client, namespace string) *configFromGits {
	return &configFromGits{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the configFromGit, and returns the corresponding configFromGit object, and an error if there is any.
func (c *configFromGits) Get(name string, options meta_v1.GetOptions) (result *v1.ConfigFromGit, err error) {
	result = &v1.ConfigFromGit{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("configfromgits").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ConfigFromGits that match those selectors.
func (c *configFromGits) List(opts meta_v1.ListOptions) (result *v1.ConfigFromGitList, err error) {
	result = &v1.ConfigFromGitList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("configfromgits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested configFromGits.
func (c *configFromGits) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("configfromgits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a configFromGit and creates it.  Returns the server's representation of the configFromGit, and an error, if there is any.
func (c *configFromGits) Create(configFromGit *v1.ConfigFromGit) (result *v1.ConfigFromGit, err error) {
	result = &v1.ConfigFromGit{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("configfromgits").
		Body(configFromGit).
		Do().
		Into(result)
	return
}

// Update takes the representation of a configFromGit and updates it. Returns the server's representation of the configFromGit, and an error, if there is any.
func (c *configFromGits) Update(configFromGit *v1.ConfigFromGit) (result *v1.ConfigFromGit, err error) {
	result = &v1.ConfigFromGit{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("configfromgits").
		Name(configFromGit.Name).
		Body(configFromGit).
		Do().
		Into(result)
	return
}

// Delete takes name of the configFromGit and deletes it. Returns an error if one occurs.
func (c *configFromGits) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("configfromgits").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *configFromGits) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("configfromgits").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched configFromGit.
func (c *configFromGits) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ConfigFromGit, err error) {
	result = &v1.ConfigFromGit{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("configfromgits").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
