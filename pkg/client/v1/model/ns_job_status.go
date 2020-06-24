// Copyright 2020 Hewlett Packard Enterprise Development LP
package model


/**
 * <p>Golang package for NsJobStatus Enum.</p>
 */

type NsJobStatus string 

const (
	NSJOBSTATUS_CANCELED NsJobStatus = "canceled"
	NSJOBSTATUS_PENDING NsJobStatus = "pending"
	NSJOBSTATUS_INPROGRESS NsJobStatus = "inprogress"
	NSJOBSTATUS_DONE NsJobStatus = "done"

) 