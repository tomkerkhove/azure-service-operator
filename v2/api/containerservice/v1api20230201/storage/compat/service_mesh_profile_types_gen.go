// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package compat

import (
	v20231001s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/pkg/errors"
)

// Storage version of v1api20230202preview.ServiceMeshProfile
// Service mesh profile for a managed cluster.
type ServiceMeshProfile struct {
	Istio       *IstioServiceMesh      `json:"istio,omitempty"`
	Mode        *string                `json:"mode,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_ServiceMeshProfile populates our ServiceMeshProfile from the provided source ServiceMeshProfile
func (profile *ServiceMeshProfile) AssignProperties_From_ServiceMeshProfile(source *v20231001s.ServiceMeshProfile) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Istio
	if source.Istio != nil {
		var istio IstioServiceMesh
		err := istio.AssignProperties_From_IstioServiceMesh(source.Istio)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_IstioServiceMesh() to populate field Istio")
		}
		profile.Istio = &istio
	} else {
		profile.Istio = nil
	}

	// Mode
	profile.Mode = genruntime.ClonePointerToString(source.Mode)

	// Update the property bag
	if len(propertyBag) > 0 {
		profile.PropertyBag = propertyBag
	} else {
		profile.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignProperties_To_ServiceMeshProfile populates the provided destination ServiceMeshProfile from our ServiceMeshProfile
func (profile *ServiceMeshProfile) AssignProperties_To_ServiceMeshProfile(destination *v20231001s.ServiceMeshProfile) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(profile.PropertyBag)

	// Istio
	if profile.Istio != nil {
		var istio v20231001s.IstioServiceMesh
		err := profile.Istio.AssignProperties_To_IstioServiceMesh(&istio)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_IstioServiceMesh() to populate field Istio")
		}
		destination.Istio = &istio
	} else {
		destination.Istio = nil
	}

	// Mode
	destination.Mode = genruntime.ClonePointerToString(profile.Mode)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}
