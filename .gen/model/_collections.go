//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Collections struct {
	ID         string `sql:"primary_key"`
	System     bool
	Type       string
	Name       string
	Schema     string
	Indexes    string
	ListRule   *string
	ViewRule   *string
	CreateRule *string
	UpdateRule *string
	DeleteRule *string
	Options    string
	Created    string
	Updated    string
}
