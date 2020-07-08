// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
	"reflect"
)

// Volumes are the basic storage units from which the total capacity is apportioned. The terms volume and LUN are used interchangeably.The number of volumes per array depends on
// storage allocation.
const (
	volumePath = "volumes"
)

// VolumeObjectSet
type VolumeObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new Volume object
func (objectSet *VolumeObjectSet) CreateObject(payload *model.Volume) (*model.Volume, error) {
	newPayload, err := model.EncodeVolume(payload)
	volumeObjectSetResp, err := objectSet.Client.Post(volumePath, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if volumeObjectSetResp == nil {
		return nil, nil
	}

	return model.DecodeVolume(volumeObjectSetResp)
}

// UpdateObject Modify existing Volume object
func (objectSet *VolumeObjectSet) UpdateObject(id string, payload *model.Volume) (*model.Volume, error) {
	newPayload, err := model.EncodeVolume(payload)
	volumeObjectSetResp, err := objectSet.Client.Put(volumePath, id, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if volumeObjectSetResp == nil {
		return nil, nil
	}
	return model.DecodeVolume(volumeObjectSetResp)
}

// DeleteObject deletes the Volume object with the specified ID
func (objectSet *VolumeObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(volumePath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a Volume object with the given ID
func (objectSet *VolumeObjectSet) GetObject(id string) (*model.Volume, error) {
	volumeObjectSetResp, err := objectSet.Client.Get(volumePath, id, model.Volume{})
	if err != nil {
		return nil, err
	}

	// null check
	if volumeObjectSetResp == nil {
		return nil, nil
	}
	return volumeObjectSetResp.(*model.Volume), err
}

// GetObjectList returns the list of Volume objects
func (objectSet *VolumeObjectSet) GetObjectList() ([]*model.Volume, error) {
	volumeObjectSetResp, err := objectSet.Client.List(volumePath)
	if err != nil {
		return nil, err
	}
	return buildVolumeObjectSet(volumeObjectSetResp), err
}

// GetObjectListFromParams returns the list of Volume objects using the given params query info
func (objectSet *VolumeObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.Volume, error) {
	volumeObjectSetResp, err := objectSet.Client.ListFromParams(volumePath, params)
	if err != nil {
		return nil, err
	}
	return buildVolumeObjectSet(volumeObjectSetResp), err
}

// generated function to build the appropriate response types
func buildVolumeObjectSet(response interface{}) []*model.Volume {
	values := reflect.ValueOf(response)
	results := make([]*model.Volume, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.Volume{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
