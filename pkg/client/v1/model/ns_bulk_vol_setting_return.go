/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsBulkVolSettingReturn - Return codes for setting an attribute to a list of items.
// Export NsBulkVolSettingReturnFields for advance operations like search filter etc.
var NsBulkVolSettingReturnFields *NsBulkVolSettingReturn

func init(){
		
	NsBulkVolSettingReturnFields= &NsBulkVolSettingReturn{
		
	}
}

type NsBulkVolSettingReturn struct {
   
    // Error codes for every element in a list of items.
    
	ErrorCodes []*string `json:"error_codes,omitempty"`
}
