/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Folder - Folders are a way to group volumes, as well as a way to apply space constraints to them.
// Export FolderFields for advance operations like search filter etc.
var FolderFields *Folder

func init(){
	IDfield:= "id"
	Namefield:= "name"
	Fqnfield:= "fqn"
	FullNamefield:= "full_name"
	SearchNamefield:= "search_name"
	Descriptionfield:= "description"
	PoolNamefield:= "pool_name"
	PoolIDfield:= "pool_id"
	InheritedVolPerfpolIDfield:= "inherited_vol_perfpol_id"
	InheritedVolPerfpolNamefield:= "inherited_vol_perfpol_name"
	AppUuIDfield:= "app_uuid"
	AppserverIDfield:= "appserver_id"
	AppserverNamefield:= "appserver_name"
	FolsetIDfield:= "folset_id"
	FolsetNamefield:= "folset_name"
	TenantIDfield:= "tenant_id"
		
	FolderFields= &Folder{
		ID: &IDfield,
		Name: &Namefield,
		Fqn: &Fqnfield,
		FullName: &FullNamefield,
		SearchName: &SearchNamefield,
		Description: &Descriptionfield,
		PoolName: &PoolNamefield,
		PoolID: &PoolIDfield,
		InheritedVolPerfpolID: &InheritedVolPerfpolIDfield,
		InheritedVolPerfpolName: &InheritedVolPerfpolNamefield,
		AppUuID: &AppUuIDfield,
		AppserverID: &AppserverIDfield,
		AppserverName: &AppserverNamefield,
		FolsetID: &FolsetIDfield,
		FolsetName: &FolsetNamefield,
		TenantID: &TenantIDfield,
		
	}
}

type Folder struct {
   
    // Identifier for the folder.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of folder.
    
 	Name *string `json:"name,omitempty"`
   
    // Fully qualified name of folder in the pool.
    
 	Fqn *string `json:"fqn,omitempty"`
   
    // Fully qualified name of folder in the group.
    
 	FullName *string `json:"full_name,omitempty"`
   
    // Name of folder used for object search.
    
 	SearchName *string `json:"search_name,omitempty"`
   
    // Text description of folder.
    
 	Description *string `json:"description,omitempty"`
   
    // Name of the pool where the folder resides.
    
 	PoolName *string `json:"pool_name,omitempty"`
   
    // ID of the pool where the folder resides.
    
 	PoolID *string `json:"pool_id,omitempty"`
   
    // Indicates whether the folder has a limit.
    
 	LimitBytesSpecified *bool `json:"limit_bytes_specified,omitempty"`
   
    // Folder limit size in bytes. By default, a folder (except SMIS and VVol types) does not have a limit. If limit_bytes is not specified when a folder is created, or if limit_bytes is set to the largest possible 64-bit signed integer (9223372036854775807), then the folder has no limit. Otherwise, a limit smaller than the capacity of the pool can be set. On output, if the folder has a limit, the limit_bytes_specified attribute will be true and limit_bytes will be the limit. If the folder does not have a limit, the limit_bytes_specified attribute will be false and limit_bytes will be interpreted based on the value of the usage_valid attribute. If the usage_valid attribute is true, limits_byte will be the capacity of the pool. Otherwise, limits_bytes is not meaningful and can be null. SMIS and VVol folders require a size limit. This attribute is superseded by limit_size_bytes.
    
   	LimitBytes *int64 `json:"limit_bytes,omitempty"`
   
    // Folder size limit in bytes. If limit_size_bytes is not specified when a folder is created, or if limit_size_bytes is set to -1, then the folder has no limit. Otherwise, a limit smaller than the capacity of the pool can be set. Folders with an agent_type of 'smis' or 'vvol' must have a size limit.
    
  	LimitSizeBytes  *int64 `json:"limit_size_bytes,omitempty"`
   
    // Limit on the provisioned size of volumes in a folder. If provisioned_limit_size_bytes is not specified when a folder is created, or if provisioned_limit_size_bytes is set to -1, then the folder has no provisioned size limit.
    
  	ProvisionedLimitSizeBytes  *int64 `json:"provisioned_limit_size_bytes,omitempty"`
   
    // Amount of space to consider as overdraft range for this folder as a percentage of folder used limit. Valid values are from 0% - 200%. This is the limit above the folder usage limit beyond which enforcement action(volume offline/non-writable) is issued.
    
   	OverdraftLimitPct *int64 `json:"overdraft_limit_pct,omitempty"`
   
    // Capacity of the folder in bytes. If the folder's size has a usage limit, capacity_bytes will be the folder's usage limit. If the folder's size does not have a usage limit, capacity_bytes will be the pool's capacity. This field is meaningful only when the usage_valid attribute is true.
    
   	CapacityBytes *int64 `json:"capacity_bytes,omitempty"`
   
    // Free space in the folder in bytes. If the folder has a usage limit, free_space_bytes will be the folder's free space (the folder's usage limit minus the folder's space usage). If the folder does not have a usage limit, free_space_bytes will be the pool's free space. This field is meaningful only when the usage_valid attribute is true.
    
   	FreeSpaceBytes *int64 `json:"free_space_bytes,omitempty"`
   
    // Sum of provisioned size of volumes in the folder.
    
   	ProvisionedBytes *int64 `json:"provisioned_bytes,omitempty"`
   
    // Sum of mapped usage and snapshot uncompressed usage of volumes in the folder.
    
   	UsageBytes *int64 `json:"usage_bytes,omitempty"`
   
    // Sum of mapped usage of volumes in the folder.
    
   	VolumeMappedBytes *int64 `json:"volume_mapped_bytes,omitempty"`
   
    // Indicate whether the space usage attributes of folder are valid.
    
 	UsageValID *bool `json:"usage_valid,omitempty"`
   
    // External management agent type.
    
   	AgentType *NsFolderAgentType `json:"agent_type,omitempty"`
   
    // Identifier of the default performance policy for a newly created volume.
    
 	InheritedVolPerfpolID *string `json:"inherited_vol_perfpol_id,omitempty"`
   
    // Name of the default performance policy for a newly created volume.
    
 	InheritedVolPerfpolName *string `json:"inherited_vol_perfpol_name,omitempty"`
   
    // Unused reserve of volumes in the folder in bytes. This field is meaningful only when the usage_valid attribute is true.
    
   	UnusedReserveBytes *int64 `json:"unused_reserve_bytes,omitempty"`
   
    // Unused reserve of snapshots of volumes in the folder in bytes. This field is meaningful only when the usage_valid attribute is true.
    
   	UnusedSnapReserveBytes *int64 `json:"unused_snap_reserve_bytes,omitempty"`
   
    // Compressed usage of volumes in the folder. This field is meaningful only when the usage_valid attribute is true.
    
   	CompressedVolUsageBytes *int64 `json:"compressed_vol_usage_bytes,omitempty"`
   
    // Compressed usage of snapshots in the folder. This field is meaningful only when the usage_valid attribute is true.
    
   	CompressedSnapUsageBytes *int64 `json:"compressed_snap_usage_bytes,omitempty"`
   
    // Uncompressed usage of volumes in the folder. This field is meaningful only when the usage_valid attribute is true.
    
   	UncompressedVolUsageBytes *int64 `json:"uncompressed_vol_usage_bytes,omitempty"`
   
    // Uncompressed usage of snapshots in the folder. This field is meaningful only when the usage_valid attribute is true.
    
   	UncompressedSnapUsageBytes *int64 `json:"uncompressed_snap_usage_bytes,omitempty"`
   
    // Compression ratio of volumes in the folder. This field is meaningful only when the usage_valid attribute is true.
    
  	VolCompressionRatio *float32 `json:"vol_compression_ratio,omitempty"`
   
    // Compression ratio of snapshots in the folder. This field is meaningful only when the usage_valid attribute is true.
    
  	SnapCompressionRatio *float32 `json:"snap_compression_ratio,omitempty"`
   
    // Compression savings for the folder expressed as ratio. This field is meaningful only when the usage_valid attribute is true.
    
  	CompressionRatio *float32 `json:"compression_ratio,omitempty"`
   
    // Time when this folder was created.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Time when this folder was last modified.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
   
    // Number of snapshots inside the folder. This attribute is deprecated and has no meaningful value.
    
   	NumSnaps *int64 `json:"num_snaps,omitempty"`
   
    // Number of snapshot collections inside the folder. This attribute is deprecated and has no meaningful value.
    
   	NumSnapcolls *int64 `json:"num_snapcolls,omitempty"`
   
    // Application identifier of the folder.
    
 	AppUuID *string `json:"app_uuid,omitempty"`
   
    // List of volumes contained by the folder.
    
   	VolumeList []*NsVolumeSummary `json:"volume_list,omitempty"`
   
    // Identifier of the application server associated with the folder.
    
 	AppserverID *string `json:"appserver_id,omitempty"`
   
    // Name of the application server associated with the folder.
    
 	AppserverName *string `json:"appserver_name,omitempty"`
   
    // Identifier of the folder set associated with the folder. Only VVol folder can be associated with the folder set. The folder and the containing folder set must be associated with the same application server.
    
 	FolsetID *string `json:"folset_id,omitempty"`
   
    // Name of the folder set associated with the folder. Only VVol folder can be associated with the folder set. The folder and the containing folder set must be associated with the same application server.
    
 	FolsetName *string `json:"folset_name,omitempty"`
   
    // IOPS limit for this folder. If limit_iops is not specified when a folder is created, or if limit_iops is set to -1, then the folder has no IOPS limit. IOPS limit should be in range [256, 4294967294] or -1 for unlimited.
    
  	LimitIops  *int64 `json:"limit_iops,omitempty"`
   
    // Throughput limit for this folder in MB/s. If limit_mbps is not specified when a folder is created, or if limit_mbps is set to -1, then the folder has no throughput limit. MBPS limit should be in range [1, 4294967294] or -1 for unlimited.
    
  	LimitMbps  *int64 `json:"limit_mbps,omitempty"`
   
    // Access protocol of the folder. This attribute is used by the VASA Provider to determine the access protocol of the bind request. If not specified in the creation request, it will be the access protocol supported by the group. If the group supports multiple protocols, the default will be Fibre Channel. This field is meaningful only to VVol folder.
    
   	AccessProtocol *NsAccessProtocol `json:"access_protocol,omitempty"`
   
    // Tenant ID of the folder. This is used to determine what tenant context the folder belongs to.
    
 	TenantID *string `json:"tenant_id,omitempty"`
}
