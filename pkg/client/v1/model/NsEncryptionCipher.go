/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model


/**
 * <p>Golang package for NsEncryptionCipher.</p>
 */

type NsEncryptionCipher string 

const (
	NSENCRYPTIONCIPHER_AES_256_XTS NsEncryptionCipher = "aes_256_xts"
	NSENCRYPTIONCIPHER_NONE NsEncryptionCipher = "none"

) 
