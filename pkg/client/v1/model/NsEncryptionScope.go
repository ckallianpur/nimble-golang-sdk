/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsEncryptionScope.</p>
 */

type NsEncryptionScope string 

const (
	NSENCRYPTIONSCOPE_VOLUME NsEncryptionScope = "volume"
	NSENCRYPTIONSCOPE_POOL NsEncryptionScope = "pool"
	NSENCRYPTIONSCOPE_NONE NsEncryptionScope = "none"
	NSENCRYPTIONSCOPE_GROUP NsEncryptionScope = "group"

) 
