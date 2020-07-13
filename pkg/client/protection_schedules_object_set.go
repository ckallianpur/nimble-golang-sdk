// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"reflect"
)

// Manage protection schedules used in protection templates.
const (
	protectionSchedulePath = "protection_schedules"
)

// ProtectionScheduleObjectSet
type ProtectionScheduleObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new ProtectionSchedule object
func (objectSet *ProtectionScheduleObjectSet) CreateObject(payload *nimbleos.ProtectionSchedule) (*nimbleos.ProtectionSchedule, error) {
	newPayload, err := nimbleos.EncodeProtectionSchedule(payload)
	resp, err := objectSet.Client.Post(protectionSchedulePath, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}

	return nimbleos.DecodeProtectionSchedule(resp)
}

// UpdateObject Modify existing ProtectionSchedule object
func (objectSet *ProtectionScheduleObjectSet) UpdateObject(id string, payload *nimbleos.ProtectionSchedule) (*nimbleos.ProtectionSchedule, error) {
	newPayload, err := nimbleos.EncodeProtectionSchedule(payload)
	resp, err := objectSet.Client.Put(protectionSchedulePath, id, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if resp == nil {
		return nil, nil
	}
	return nimbleos.DecodeProtectionSchedule(resp)
}

// DeleteObject deletes the ProtectionSchedule object with the specified ID
func (objectSet *ProtectionScheduleObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(protectionSchedulePath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a ProtectionSchedule object with the given ID
func (objectSet *ProtectionScheduleObjectSet) GetObject(id string) (*nimbleos.ProtectionSchedule, error) {
	protectionScheduleObjectSetResp, err := objectSet.Client.Get(protectionSchedulePath, id, nimbleos.ProtectionSchedule{})
	if err != nil {
		return nil, err
	}

	// null check
	if protectionScheduleObjectSetResp == nil {
		return nil, nil
	}
	return protectionScheduleObjectSetResp.(*nimbleos.ProtectionSchedule), err
}

// GetObjectList returns the list of ProtectionSchedule objects
func (objectSet *ProtectionScheduleObjectSet) GetObjectList() ([]*nimbleos.ProtectionSchedule, error) {
	protectionScheduleObjectSetResp, err := objectSet.Client.List(protectionSchedulePath)
	if err != nil {
		return nil, err
	}
	return buildProtectionScheduleObjectSet(protectionScheduleObjectSetResp), err
}

// GetObjectListFromParams returns the list of ProtectionSchedule objects using the given params query info
func (objectSet *ProtectionScheduleObjectSet) GetObjectListFromParams(params *param.GetParams) ([]*nimbleos.ProtectionSchedule, error) {
	protectionScheduleObjectSetResp, err := objectSet.Client.ListFromParams(protectionSchedulePath, params)
	if err != nil {
		return nil, err
	}
	return buildProtectionScheduleObjectSet(protectionScheduleObjectSetResp), err
}

// generated function to build the appropriate response types
func buildProtectionScheduleObjectSet(response interface{}) []*nimbleos.ProtectionSchedule {
	values := reflect.ValueOf(response)
	results := make([]*nimbleos.ProtectionSchedule, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &nimbleos.ProtectionSchedule{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
