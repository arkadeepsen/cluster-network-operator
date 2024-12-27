// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/config/v1"
)

// ProfileCustomizationsApplyConfiguration represents a declarative configuration of the ProfileCustomizations type for use
// with apply.
type ProfileCustomizationsApplyConfiguration struct {
	DynamicResourceAllocation *v1.DRAEnablement `json:"dynamicResourceAllocation,omitempty"`
}

// ProfileCustomizationsApplyConfiguration constructs a declarative configuration of the ProfileCustomizations type for use with
// apply.
func ProfileCustomizations() *ProfileCustomizationsApplyConfiguration {
	return &ProfileCustomizationsApplyConfiguration{}
}

// WithDynamicResourceAllocation sets the DynamicResourceAllocation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DynamicResourceAllocation field is set to the value of the last call.
func (b *ProfileCustomizationsApplyConfiguration) WithDynamicResourceAllocation(value v1.DRAEnablement) *ProfileCustomizationsApplyConfiguration {
	b.DynamicResourceAllocation = &value
	return b
}