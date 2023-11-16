// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20201201 "github.com/Azure/azure-service-operator/v2/api/cache/v1api20201201"
	v20201201s "github.com/Azure/azure-service-operator/v2/api/cache/v1api20201201/storage"
	v20230401 "github.com/Azure/azure-service-operator/v2/api/cache/v1api20230401"
	v20230401s "github.com/Azure/azure-service-operator/v2/api/cache/v1api20230401/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type RedisExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *RedisExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20201201.Redis{},
		&v20201201s.Redis{},
		&v20230401.Redis{},
		&v20230401s.Redis{}}
}
