/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// ProtectionSchedule - Manage protection schedules used in protection templates.
// Export ProtectionScheduleFields for advance operations like search filter etc.
var ProtectionScheduleFields *ProtectionSchedule

func init(){
	IDfield:= "id"
	Namefield:= "name"
	Descriptionfield:= "description"
	VolcollOrProttmplIDfield:= "volcoll_or_prottmpl_id"
	Daysfield:= "days"
	DownstreamPartnerfield:= "downstream_partner"
	DownstreamPartnerNamefield:= "downstream_partner_name"
	DownstreamPartnerIDfield:= "downstream_partner_id"
	UpstreamPartnerNamefield:= "upstream_partner_name"
	UpstreamPartnerIDfield:= "upstream_partner_id"
	LastReplicatedSnapcollNamefield:= "last_replicated_snapcoll_name"
	LastReplicatedSnapcollIDfield:= "last_replicated_snapcoll_id"
	SchedOwnerIDfield:= "sched_owner_id"
	SchedOwnerNamefield:= "sched_owner_name"
		
	ProtectionScheduleFields= &ProtectionSchedule{
		ID: &IDfield,
		Name: &Namefield,
		Description: &Descriptionfield,
		VolcollOrProttmplID: &VolcollOrProttmplIDfield,
		Days: &Daysfield,
		DownstreamPartner: &DownstreamPartnerfield,
		DownstreamPartnerName: &DownstreamPartnerNamefield,
		DownstreamPartnerID: &DownstreamPartnerIDfield,
		UpstreamPartnerName: &UpstreamPartnerNamefield,
		UpstreamPartnerID: &UpstreamPartnerIDfield,
		LastReplicatedSnapcollName: &LastReplicatedSnapcollNamefield,
		LastReplicatedSnapcollID: &LastReplicatedSnapcollIDfield,
		SchedOwnerID: &SchedOwnerIDfield,
		SchedOwnerName: &SchedOwnerNamefield,
		
	}
}

type ProtectionSchedule struct {
   
    // Identifier for protection schedule.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of snapshot schedule to create.
    
 	Name *string `json:"name,omitempty"`
   
    // Description of the schedule.
    
 	Description *string `json:"description,omitempty"`
   
    // Type of the protection policy this schedule is attached to. Valid values are protection_template and volume_collection.
    
   	VolcollOrProttmplType *NsProtectionPolicyType `json:"volcoll_or_prottmpl_type,omitempty"`
   
    // Identifier of the protection policy (protection_template or volume_collection) in which this protection schedule is attached to.
    
 	VolcollOrProttmplID *string `json:"volcoll_or_prottmpl_id,omitempty"`
   
    // Repeat interval for snapshots with respect to the period_unit.  For example, a value of 2 with the 'period_unit' of 'hours' results in one snapshot every 2 hours.
    
   	Period *int64 `json:"period,omitempty"`
   
    // Time unit over which to take the number of snapshots specified in 'period'. For example, a value of 'days' with a 'period' of '1' results in one snapshot every day.
    
   	PeriodUnit *NsPeriodUnit `json:"period_unit,omitempty"`
   
    // Time of day when snapshot should be taken. In case repeat frequency specifies more than one snapshot in a day then the until_time option specifies until what time of day to take snapshots.
    
   	AtTime *int64 `json:"at_time,omitempty"`
   
    // Time of day to stop taking snapshots. Applicable only when repeat frequency specifies more than one snapshot in a day.
    
   	UntilTime *int64 `json:"until_time,omitempty"`
   
    // Specifies which days snapshots should be taken.
    
 	Days *string `json:"days,omitempty"`
   
    // Number of snapshots to retain. If replication is enabled on this schedule the array will always retain the latest replicated snapshot, which may exceed the specified retention value. This is necessary to ensure efficient replication performance.
    
   	NumRetain *int64 `json:"num_retain,omitempty"`
   
    // Specifies the partner name if snapshots created by this schedule should be replicated.
    
 	DownstreamPartner *string `json:"downstream_partner,omitempty"`
   
    // Specifies the partner name if snapshots created by this schedule should be replicated.
    
 	DownstreamPartnerName *string `json:"downstream_partner_name,omitempty"`
   
    // Specifies the partner ID if snapshots created by this schedule should be replicated. In an update operation, if snapshots should be replicated, set this attribute to the ID of the replication partner. If snapshots should not be replicated, set this attribute to the empty string.
    
 	DownstreamPartnerID *string `json:"downstream_partner_id,omitempty"`
   
    // Specifies the partner name from which snapshots created by this schedule are replicated.
    
 	UpstreamPartnerName *string `json:"upstream_partner_name,omitempty"`
   
    // Specifies the partner ID from which snapshots created by this schedule are replicated.
    
 	UpstreamPartnerID *string `json:"upstream_partner_id,omitempty"`
   
    // Specifies which snapshots should be replicated. If snapshots are replicated and this option is not specified, every snapshot is replicated.
    
   	ReplicateEvery *int64 `json:"replicate_every,omitempty"`
   
    // Number of snapshots to retain on the replica.
    
   	NumRetainReplica *int64 `json:"num_retain_replica,omitempty"`
   
    // Replication alert threshold in seconds. If the replication of a snapshot takes more than this amount of time to complete an alert will be generated. Enter 0 to disable this alert.
    
   	ReplAlertThres *int64 `json:"repl_alert_thres,omitempty"`
   
    // Run verification tool on snapshot created by this schedule. This option can only be used with snapshot schedules of a protection template that has application synchronization. The tool used to verify snapshot depends on the type of application. For example, if application synchronization is VSS and the application ID is Exchange, eseutil tool is run on the snapshots. If verification fails, the logs are not truncated.
    
 	SnapVerify *bool `json:"snap_verify,omitempty"`
   
    // Skip consistency check for database files on snapshots created by this schedule. This option only applies to snapshot schedules of a protection template with application synchronization set to VSS, application ID set to MS Exchange 2010 or later w/DAG, this schedule's snap_verify option set to yes, and its disable_appsync option set to false. Skipping consistency checks is only recommended if each database in a DAG has multiple copies.
    
 	SkipDbConsistencyCheck *bool `json:"skip_db_consistency_check,omitempty"`
   
    // Disables application synchronized snapshots and creates crash consistent snapshots instead.
    
 	DisableAppsync *bool `json:"disable_appsync,omitempty"`
   
    // Normal schedules have internal timers which drive snapshot creation. An externally driven schedule has no internal timers. All snapshot activity is driven by an external trigger. In other words, these schedules are used only for externally driven manual snapshots.
    
   	ScheduleType *NsScheduleType `json:"schedule_type,omitempty"`
   
    // A schedule is active only if it is owned by the same owner as the volume collection. Only active schedules of a volume collection participate in the creation of snapshots and replication.
    
 	Active *bool `json:"active,omitempty"`
   
    // Time when this protection schedule was created.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Time when this protection schedule was last modified.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
   
    // Time when the timing of the protection schedule was last modified.
    
   	LastModSchedTime *int64 `json:"last_mod_sched_time,omitempty"`
   
    // Specifies the name of last replicated snapshot collection.
    
 	LastReplicatedSnapcollName *string `json:"last_replicated_snapcoll_name,omitempty"`
   
    // Specifies the snapshot collection ID of last replicated snapshot collection.
    
 	LastReplicatedSnapcollID *string `json:"last_replicated_snapcoll_id,omitempty"`
   
    // Time when last snapshot collection was replicated.
    
   	LastReplicatedAtTime *int64 `json:"last_replicated_at_time,omitempty"`
   
    // Time when last snapshot was taken.
    
   	LastSnapTime *int64 `json:"last_snap_time,omitempty"`
   
    // Time when next snapshot will be taken.
    
   	NextSnapTime *int64 `json:"next_snap_time,omitempty"`
   
    // Time when next snapshot will be replicated.
    
   	NextReplSnapTime *int64 `json:"next_repl_snap_time,omitempty"`
   
    // This is only used by custom read handler for internal calculations.
    
   	SnapCounter *int64 `json:"snap_counter,omitempty"`
   
    // Identifier of the group that owns this schedule.
    
 	SchedOwnerID *string `json:"sched_owner_id,omitempty"`
   
    // Name of the group that owns this schedule.
    
 	SchedOwnerName *string `json:"sched_owner_name,omitempty"`
   
    // The last timing configutation changed.
    
   	LastConfigChangeTime *int64 `json:"last_config_change_time,omitempty"`
   
    // The list of the replication status of volumes undergoing replication.
    
   	VolStatusList []*NsReplVolStatus `json:"vol_status_list,omitempty"`
   
    // A list of the replication status of volumes undergoing synchronous replication.
    
   	SyncReplVolStatusList []*NsSyncReplVolStatus `json:"sync_repl_vol_status_list,omitempty"`
   
    // Break synchronous replication for the specified volume collection and present downstream volumes to host(s). Downstream volumes in the volume collection will be set to online and presented to the host(s) using new serial and LUN numbers. No changes will be made to the upstream volumes, their serial and LUN numbers, and their online state. The existing ACLs on the upstream volumes will be copied to the downstream volumes. Use this in conjunction with an empty downstream_partner_id. This unconfigures synchronous replication when the partner is removed from the last replicating schedule in the specified volume collection and presents the downstream volumes to host(s). Host(s) will need to be configured to access the new volumes with the newly assigned serial and LUN numbers. Use this option to expose downstream volumes in a synchronously replicated volume collection to host(s) only when the upstream partner is confirmed to be down and there is no communication between partners. Do not execute this operation if a previous Group Management Service takeover has been performed on a different array. Do not perform a subsequent Group Management Service takeover on a different array as it will lead to irreconcilable conflicts. This limitation is cleared once the Group management service backup array has successfully synchronized after reconnection.
    
 	UseDownstreamForDr *bool `json:"use_downstream_for_dr,omitempty"`
}
