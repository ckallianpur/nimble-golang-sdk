// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// Disk - Disks are used for storing user data.
// Export DiskFields for advance operations like search filter etc.
var DiskFields *Disk

func init() {
	IDfield := "id"
	Serialfield := "serial"
	Pathfield := "path"
	ShelfSerialfield := "shelf_serial"
	ShelfLocationfield := "shelf_location"
	ShelfIdfield := "shelf_id"
	Modelfield := "model"
	Vendorfield := "vendor"
	FirmwareVersionfield := "firmware_version"
	DiskInternalStat1field := "disk_internal_stat1"
	ArrayNamefield := "array_name"
	ArrayIdfield := "array_id"

	DiskFields = &Disk{
		ID:                &IDfield,
		Serial:            &Serialfield,
		Path:              &Pathfield,
		ShelfSerial:       &ShelfSerialfield,
		ShelfLocation:     &ShelfLocationfield,
		ShelfId:           &ShelfIdfield,
		Model:             &Modelfield,
		Vendor:            &Vendorfield,
		FirmwareVersion:   &FirmwareVersionfield,
		DiskInternalStat1: &DiskInternalStat1field,
		ArrayName:         &ArrayNamefield,
		ArrayId:           &ArrayIdfield,
	}
}

type Disk struct {
	// ID - ID of disk.
	ID *string `json:"id,omitempty"`
	// IsDfc - Is disk part of dual flash carrier.
	IsDfc *bool `json:"is_dfc,omitempty"`
	// Serial - Disk serial number(N/A if empty).
	Serial *string `json:"serial,omitempty"`
	// Path - Disk SCSI device path.
	Path *string `json:"path,omitempty"`
	// ShelfSerial - Serial number of the shelf the disk is attached to.
	ShelfSerial *string `json:"shelf_serial,omitempty"`
	// ShelfLocation - Identifies the controller, port, and chain position of the shelf the disk belongs to.
	ShelfLocation *string `json:"shelf_location,omitempty"`
	// ShelfId - Identifies the physical shelf the disk belongs to.
	ShelfId *string `json:"shelf_id,omitempty"`
	// ShelfLocationId - Identifies the position shelf the disk belongs to, as coded integer.
	ShelfLocationId *int64 `json:"shelf_location_id,omitempty"`
	// VshelfId - Identifies the local shelf ID the disk belongs to.
	VshelfId *int64 `json:"vshelf_id,omitempty"`
	// Slot - Disk slot number.
	Slot *int64 `json:"slot,omitempty"`
	// Bank - Disk bank number.
	Bank *int64 `json:"bank,omitempty"`
	// Model - Disk model name.
	Model *string `json:"model,omitempty"`
	// Vendor - Vendor name of the disk manufacturer.
	Vendor *string `json:"vendor,omitempty"`
	// FirmwareVersion - Firmware version on the disk.
	FirmwareVersion *string `json:"firmware_version,omitempty"`
	// Hba - HBA ID the disk is connected to.
	Hba *int64 `json:"hba,omitempty"`
	// Port - HBA port number the disk is connected to.
	Port *int64 `json:"port,omitempty"`
	// Size - Disk size in bytes.
	Size *int64 `json:"size,omitempty"`
	// State - Disk hardware state.
	State *NsDiskState `json:"state,omitempty"`
	// Type - Type of disk (HDD, SSD, N/A).
	Type *NsDiskType `json:"type,omitempty"`
	// BlockType - Native block type of the disk.
	BlockType *NsDiskBlockType `json:"block_type,omitempty"`
	// RaidId - Raid ID.
	RaidId *int64 `json:"raid_id,omitempty"`
	// RaidResyncPercent - Percentage RAID rebuild completed on this disk.
	RaidResyncPercent *int64 `json:"raid_resync_percent,omitempty"`
	// RaidResyncCurrentSpeed - Current RAID rebuild speed (bytes/sec).
	RaidResyncCurrentSpeed *int64 `json:"raid_resync_current_speed,omitempty"`
	// RaidResyncAverageSpeed - Average RAID rebuild speed (bytes/sec).
	RaidResyncAverageSpeed *int64 `json:"raid_resync_average_speed,omitempty"`
	// RaidState - RAID status for the disk (N/A, okay, resynchronizing, spare, faulty).
	RaidState *NsDiskRaidState `json:"raid_state,omitempty"`
	// DiskInternalStat1 - Internal disk statistic 1.
	DiskInternalStat1 *string `json:"disk_internal_stat1,omitempty"`
	// SmartAttributeList - S.M.A.R.T. attributes for the disk.
	SmartAttributeList []*NsDiskSmartAttribute `json:"smart_attribute_list,omitempty"`
	// DiskOp - The intended operation to be performed on the specified disk.
	DiskOp *NsDiskOp `json:"disk_op,omitempty"`
	// Force - Forcibly add a disk.
	Force *bool `json:"force,omitempty"`
	// ArrayName - Name of array the disk belongs to.
	ArrayName *string `json:"array_name,omitempty"`
	// ArrayId - ID of array the disk belongs to.
	ArrayId *string `json:"array_id,omitempty"`
	// PartialResponseOk - Indicate that it is okay to provide partially available response.
	PartialResponseOk *bool `json:"partial_response_ok,omitempty"`
}
