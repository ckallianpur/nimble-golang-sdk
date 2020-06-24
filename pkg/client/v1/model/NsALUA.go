/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsALUA.</p>
 */

type NsALUA string 

const (
	NSALUA_LOGICAL_BLOCK_DEPENDENT NsALUA = "logical_block_dependent"
	NSALUA_OFFLINE NsALUA = "offline"
	NSALUA_ACTIVE_OPTIMIZED NsALUA = "active_optimized"
	NSALUA_STANDBY NsALUA = "standby"
	NSALUA_TRANSITIONING NsALUA = "transitioning"
	NSALUA_UNAVAILABLE NsALUA = "unavailable"
	NSALUA_INVALID NsALUA = "invalid"
	NSALUA_ACTIVE_NON_OPTIMIZED NsALUA = "active_non_optimized"

) 
