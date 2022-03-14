/*
Copyright 2022 The Kubernetes Authors.

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
package cloud

import (
	"errors"

	"github.com/IBM-Cloud/power-go-client/power/models"
	"k8s.io/utils/pointer"
)

const (
	PowerVSInstanceStateSHUTOFF = "SHUTOFF"
	PowerVSInstanceStateACTIVE  = "ACTIVE"
	PowerVSInstanceHealthOK     = "OK"
	StoragePoolAffinity         = false
	ProviderIDValidLength       = 6
)

type NodeUpdateScopeParams struct {
	ServiceInstanceId string
	InstanceId        string
}

type NodeUpdateScope struct {
	Cloud
	ServiceInstanceId string
	InstanceId        string
}

func NewNodeUpdateScope(params NodeUpdateScopeParams) (scope *NodeUpdateScope, err error) {
	scope = &NodeUpdateScope{}

	if params.ServiceInstanceId == "" {
		err = errors.New("ServiceInstanceId is required when creating a NodeUpdateScope")
		return
	}
	scope.ServiceInstanceId = params.ServiceInstanceId

	if params.InstanceId == "" {
		err = errors.New("InstanceId is required when creating a NodeUpdateScope")
		return
	}
	scope.InstanceId = params.InstanceId

	c, err := NewPowerVSCloud(scope.ServiceInstanceId, false)
	if err != nil {
		panic(err)
	}
	scope.Cloud = c

	return scope, nil
}

func (p *powerVSCloud) GetPVMInstanceDetails(instanceID string) (*models.PVMInstance, error) {
	insDetails, err := p.pvmInstancesClient.Get(instanceID)
	if err != nil {
		return nil, err
	}
	return insDetails, nil
}

func (p *powerVSCloud) UpdateStoragePoolAffinity(instanceID string) error {

	body := &models.PVMInstanceUpdate{
		StoragePoolAffinity: pointer.Bool(StoragePoolAffinity),
	}

	_, err := p.pvmInstancesClient.Update(instanceID, body)
	if err != nil {
		return err
	}
	return nil
}