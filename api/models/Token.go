package models

// Token is the entity that holds the security token
// swagger:model
type Token struct {
	// Token is a jwt token needed for authenticated requests
	// required: true
	Token string `json:"token"`
}
