// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/openshift/api/operator/v1alpha1"
	operatorv1alpha1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOLMs implements OLMInterface
type FakeOLMs struct {
	Fake *FakeOperatorV1alpha1
}

var olmsResource = v1alpha1.SchemeGroupVersion.WithResource("olms")

var olmsKind = v1alpha1.SchemeGroupVersion.WithKind("OLM")

// Get takes name of the oLM, and returns the corresponding oLM object, and an error if there is any.
func (c *FakeOLMs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OLM, err error) {
	emptyResult := &v1alpha1.OLM{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(olmsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.OLM), err
}

// List takes label and field selectors, and returns the list of OLMs that match those selectors.
func (c *FakeOLMs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OLMList, err error) {
	emptyResult := &v1alpha1.OLMList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(olmsResource, olmsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OLMList{ListMeta: obj.(*v1alpha1.OLMList).ListMeta}
	for _, item := range obj.(*v1alpha1.OLMList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested oLMs.
func (c *FakeOLMs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(olmsResource, opts))
}

// Create takes the representation of a oLM and creates it.  Returns the server's representation of the oLM, and an error, if there is any.
func (c *FakeOLMs) Create(ctx context.Context, oLM *v1alpha1.OLM, opts v1.CreateOptions) (result *v1alpha1.OLM, err error) {
	emptyResult := &v1alpha1.OLM{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(olmsResource, oLM, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.OLM), err
}

// Update takes the representation of a oLM and updates it. Returns the server's representation of the oLM, and an error, if there is any.
func (c *FakeOLMs) Update(ctx context.Context, oLM *v1alpha1.OLM, opts v1.UpdateOptions) (result *v1alpha1.OLM, err error) {
	emptyResult := &v1alpha1.OLM{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(olmsResource, oLM, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.OLM), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOLMs) UpdateStatus(ctx context.Context, oLM *v1alpha1.OLM, opts v1.UpdateOptions) (result *v1alpha1.OLM, err error) {
	emptyResult := &v1alpha1.OLM{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(olmsResource, "status", oLM, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.OLM), err
}

// Delete takes name of the oLM and deletes it. Returns an error if one occurs.
func (c *FakeOLMs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(olmsResource, name, opts), &v1alpha1.OLM{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOLMs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(olmsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OLMList{})
	return err
}

// Patch applies the patch and returns the patched oLM.
func (c *FakeOLMs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OLM, err error) {
	emptyResult := &v1alpha1.OLM{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(olmsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.OLM), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied oLM.
func (c *FakeOLMs) Apply(ctx context.Context, oLM *operatorv1alpha1.OLMApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.OLM, err error) {
	if oLM == nil {
		return nil, fmt.Errorf("oLM provided to Apply must not be nil")
	}
	data, err := json.Marshal(oLM)
	if err != nil {
		return nil, err
	}
	name := oLM.Name
	if name == nil {
		return nil, fmt.Errorf("oLM.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.OLM{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(olmsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.OLM), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeOLMs) ApplyStatus(ctx context.Context, oLM *operatorv1alpha1.OLMApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.OLM, err error) {
	if oLM == nil {
		return nil, fmt.Errorf("oLM provided to Apply must not be nil")
	}
	data, err := json.Marshal(oLM)
	if err != nil {
		return nil, err
	}
	name := oLM.Name
	if name == nil {
		return nil, fmt.Errorf("oLM.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.OLM{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(olmsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.OLM), err
}