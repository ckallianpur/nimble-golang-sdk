/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsSeverityLevel.</p>
 */

type NsSeverityLevel string 

const (
	NSSEVERITYLEVEL_CRITICAL NsSeverityLevel = "critical"
	NSSEVERITYLEVEL_WARNING NsSeverityLevel = "warning"
	NSSEVERITYLEVEL_INFO NsSeverityLevel = "info"
	NSSEVERITYLEVEL_NOTICE NsSeverityLevel = "notice"

) 
