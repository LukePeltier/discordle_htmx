//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Admins struct {
	ID              string `sql:"primary_key"`
	Avatar          int32
	Email           string
	TokenKey        string
	PasswordHash    string
	LastResetSentAt string
	Created         string
	Updated         string
}
