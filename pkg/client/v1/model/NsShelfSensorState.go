/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsShelfSensorState.</p>
 */

type NsShelfSensorState string 

const (
	NSSHELFSENSORSTATE_MISSING NsShelfSensorState = "Missing"
	NSSHELFSENSORSTATE_FAILED NsShelfSensorState = "Failed"
	NSSHELFSENSORSTATE_OK NsShelfSensorState = "OK"
	NSSHELFSENSORSTATE_ALERTED NsShelfSensorState = "Alerted"

) 
