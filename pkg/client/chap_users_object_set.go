// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
	"reflect"
)

// Manage Challenge-Response Handshake Authentication Protocol (CHAP) user accounts. CHAP users are one method of access control for iSCSI initiators. Each CHAP user has a CHAP
// password, sometimes called a CHAP secret. The CHAP passwords must match on the array and on the iSCSI initiator in order for the array to authenicate the initiator and allow it
// access. The CHAP user information must exist on both the array and the iSCSI initiator. Target authentication gives security only for the specific iSCSI target.
const (
	chapUserPath = "chap_users"
)

// ChapUserObjectSet
type ChapUserObjectSet struct {
	Client *GroupMgmtClient
}

// CreateObject creates a new ChapUser object
func (objectSet *ChapUserObjectSet) CreateObject(payload *model.ChapUser) (*model.ChapUser, error) {
	newPayload, err := model.EncodeChapUser(payload)
	chapUserObjectSetResp, err := objectSet.Client.Post(chapUserPath, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if chapUserObjectSetResp == nil {
		return nil, nil
	}

	return model.DecodeChapUser(chapUserObjectSetResp)
}

// UpdateObject Modify existing ChapUser object
func (objectSet *ChapUserObjectSet) UpdateObject(id string, payload *model.ChapUser) (*model.ChapUser, error) {
	newPayload, err := model.EncodeChapUser(payload)
	chapUserObjectSetResp, err := objectSet.Client.Put(chapUserPath, id, newPayload)
	if err != nil {
		return nil, err
	}

	// null check
	if chapUserObjectSetResp == nil {
		return nil, nil
	}
	return model.DecodeChapUser(chapUserObjectSetResp)
}

// DeleteObject deletes the ChapUser object with the specified ID
func (objectSet *ChapUserObjectSet) DeleteObject(id string) error {
	err := objectSet.Client.Delete(chapUserPath, id)
	if err != nil {
		return err
	}
	return nil
}

// GetObject returns a ChapUser object with the given ID
func (objectSet *ChapUserObjectSet) GetObject(id string) (*model.ChapUser, error) {
	chapUserObjectSetResp, err := objectSet.Client.Get(chapUserPath, id, model.ChapUser{})
	if err != nil {
		return nil, err
	}

	// null check
	if chapUserObjectSetResp == nil {
		return nil, nil
	}
	return chapUserObjectSetResp.(*model.ChapUser), err
}

// GetObjectList returns the list of ChapUser objects
func (objectSet *ChapUserObjectSet) GetObjectList() ([]*model.ChapUser, error) {
	chapUserObjectSetResp, err := objectSet.Client.List(chapUserPath)
	if err != nil {
		return nil, err
	}
	return buildChapUserObjectSet(chapUserObjectSetResp), err
}

// GetObjectListFromParams returns the list of ChapUser objects using the given params query info
func (objectSet *ChapUserObjectSet) GetObjectListFromParams(params *util.GetParams) ([]*model.ChapUser, error) {
	chapUserObjectSetResp, err := objectSet.Client.ListFromParams(chapUserPath, params)
	if err != nil {
		return nil, err
	}
	return buildChapUserObjectSet(chapUserObjectSetResp), err
}

// generated function to build the appropriate response types
func buildChapUserObjectSet(response interface{}) []*model.ChapUser {
	values := reflect.ValueOf(response)
	results := make([]*model.ChapUser, values.Len())

	for i := 0; i < values.Len(); i++ {
		value := &model.ChapUser{}
		jsonutil.Decode(values.Index(i).Interface(), value)
		results[i] = value
	}

	return results
}
