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
package fake

import (
	dev_kubernetes_pavanvelidandla_com_v1 "crds/pkg/apis/dev.kubernetes.pavanvelidandla.com/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConfigFromGits implements ConfigFromGitInterface
type FakeConfigFromGits struct {
	Fake *FakeDevV1
	ns   string
}

var configfromgitsResource = schema.GroupVersionResource{Group: "dev.kubernetes.pavanvelidandla.com", Version: "v1", Resource: "configfromgits"}

var configfromgitsKind = schema.GroupVersionKind{Group: "dev.kubernetes.pavanvelidandla.com", Version: "v1", Kind: "ConfigFromGit"}

// Get takes name of the configFromGit, and returns the corresponding configFromGit object, and an error if there is any.
func (c *FakeConfigFromGits) Get(name string, options v1.GetOptions) (result *dev_kubernetes_pavanvelidandla_com_v1.ConfigFromGit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(configfromgitsResource, c.ns, name), &dev_kubernetes_pavanvelidandla_com_v1.ConfigFromGit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*dev_kubernetes_pavanvelidandla_com_v1.ConfigFromGit), err
}

// List takes label and field selectors, and returns the list of ConfigFromGits that match those selectors.
func (c *FakeConfigFromGits) List(opts v1.ListOptions) (result *dev_kubernetes_pavanvelidandla_com_v1.ConfigFromGitList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(configfromgitsResource, configfromgitsKind, c.ns, opts), &dev_kubernetes_pavanvelidandla_com_v1.ConfigFromGitList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &dev_kubernetes_pavanvelidandla_com_v1.ConfigFromGitList{}
	for _, item := range obj.(*dev_kubernetes_pavanvelidandla_com_v1.ConfigFromGitList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested configFromGits.
func (c *FakeConfigFromGits) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(configfromgitsResource, c.ns, opts))

}

// Create takes the representation of a configFromGit and creates it.  Returns the server's representation of the configFromGit, and an error, if there is any.
func (c *FakeConfigFromGits) Create(configFromGit *dev_kubernetes_pavanvelidandla_com_v1.ConfigFromGit) (result *dev_kubernetes_pavanvelidandla_com_v1.ConfigFromGit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(configfromgitsResource, c.ns, configFromGit), &dev_kubernetes_pavanvelidandla_com_v1.ConfigFromGit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*dev_kubernetes_pavanvelidandla_com_v1.ConfigFromGit), err
}

// Update takes the representation of a configFromGit and updates it. Returns the server's representation of the configFromGit, and an error, if there is any.
func (c *FakeConfigFromGits) Update(configFromGit *dev_kubernetes_pavanvelidandla_com_v1.ConfigFromGit) (result *dev_kubernetes_pavanvelidandla_com_v1.ConfigFromGit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(configfromgitsResource, c.ns, configFromGit), &dev_kubernetes_pavanvelidandla_com_v1.ConfigFromGit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*dev_kubernetes_pavanvelidandla_com_v1.ConfigFromGit), err
}

// Delete takes name of the configFromGit and deletes it. Returns an error if one occurs.
func (c *FakeConfigFromGits) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(configfromgitsResource, c.ns, name), &dev_kubernetes_pavanvelidandla_com_v1.ConfigFromGit{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConfigFromGits) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(configfromgitsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &dev_kubernetes_pavanvelidandla_com_v1.ConfigFromGitList{})
	return err
}

// Patch applies the patch and returns the patched configFromGit.
func (c *FakeConfigFromGits) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *dev_kubernetes_pavanvelidandla_com_v1.ConfigFromGit, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(configfromgitsResource, c.ns, name, data, subresources...), &dev_kubernetes_pavanvelidandla_com_v1.ConfigFromGit{})

	if obj == nil {
		return nil, err
	}
	return obj.(*dev_kubernetes_pavanvelidandla_com_v1.ConfigFromGit), err
}
