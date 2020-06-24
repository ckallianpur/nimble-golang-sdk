/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// UserPolicy - Manages the password policies configured for the group.
// Export UserPolicyFields for advance operations like search filter etc.
var UserPolicyFields *UserPolicy

func init(){
	IDfield:= "id"
		
	UserPolicyFields= &UserPolicy{
		ID: &IDfield,
		
	}
}

type UserPolicy struct {
   
    // Identifier for the security policy.
    
 	ID *string `json:"id,omitempty"`
   
    // Number of authentication attempts allowed before the user account is locked. Allowed range is [1, 10] inclusive. '0' indicates no limit.
    
  	AllowedAttempts  *int64 `json:"allowed_attempts,omitempty"`
   
    // Minimum length for user passwords. Allowed range is [8, 255] inclusive.
    
  	MinLength  *int64 `json:"min_length,omitempty"`
   
    // Number of uppercase characters required in user passwords. Allowed range is [0, 255] inclusive.
    
  	Upper  *int64 `json:"upper,omitempty"`
   
    // Number of lowercase characters required in user passwords. Allowed range is [0, 255] inclusive.
    
  	Lower  *int64 `json:"lower,omitempty"`
   
    // Number of numerical characters required in user passwords. Allowed range is [0, 255] inclusive.
    
  	Digit  *int64 `json:"digit,omitempty"`
   
    // Number of special characters required in user passwords. Allowed range is [0, 255] inclusive.
    
  	Special  *int64 `json:"special,omitempty"`
   
    // Number of characters that must be different from the previous password. Allowed range is [1, 255] inclusive.
    
  	PreviousDiff  *int64 `json:"previous_diff,omitempty"`
   
    // Number of times that a password must change before you can reuse a previous password. Allowed range is [1,20] inclusive.
    
  	NoReuse  *int64 `json:"no_reuse,omitempty"`
   
    // Maximum number of sessions allowed for a group. Allowed range is [10, 1000] inclusive. '0' indicates no limit.
    
  	MaxSessions  *int64 `json:"max_sessions,omitempty"`
}
