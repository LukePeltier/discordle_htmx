//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/sqlite"
)

var Admins = newAdminsTable("", "_admins", "")

type adminsTable struct {
	sqlite.Table

	// Columns
	ID              sqlite.ColumnString
	Avatar          sqlite.ColumnInteger
	Email           sqlite.ColumnString
	TokenKey        sqlite.ColumnString
	PasswordHash    sqlite.ColumnString
	LastResetSentAt sqlite.ColumnString
	Created         sqlite.ColumnString
	Updated         sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type AdminsTable struct {
	adminsTable

	EXCLUDED adminsTable
}

// AS creates new AdminsTable with assigned alias
func (a AdminsTable) AS(alias string) *AdminsTable {
	return newAdminsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AdminsTable with assigned schema name
func (a AdminsTable) FromSchema(schemaName string) *AdminsTable {
	return newAdminsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AdminsTable with assigned table prefix
func (a AdminsTable) WithPrefix(prefix string) *AdminsTable {
	return newAdminsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AdminsTable with assigned table suffix
func (a AdminsTable) WithSuffix(suffix string) *AdminsTable {
	return newAdminsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAdminsTable(schemaName, tableName, alias string) *AdminsTable {
	return &AdminsTable{
		adminsTable: newAdminsTableImpl(schemaName, tableName, alias),
		EXCLUDED:    newAdminsTableImpl("", "excluded", ""),
	}
}

func newAdminsTableImpl(schemaName, tableName, alias string) adminsTable {
	var (
		IDColumn              = sqlite.StringColumn("id")
		AvatarColumn          = sqlite.IntegerColumn("avatar")
		EmailColumn           = sqlite.StringColumn("email")
		TokenKeyColumn        = sqlite.StringColumn("tokenKey")
		PasswordHashColumn    = sqlite.StringColumn("passwordHash")
		LastResetSentAtColumn = sqlite.StringColumn("lastResetSentAt")
		CreatedColumn         = sqlite.StringColumn("created")
		UpdatedColumn         = sqlite.StringColumn("updated")
		allColumns            = sqlite.ColumnList{IDColumn, AvatarColumn, EmailColumn, TokenKeyColumn, PasswordHashColumn, LastResetSentAtColumn, CreatedColumn, UpdatedColumn}
		mutableColumns        = sqlite.ColumnList{AvatarColumn, EmailColumn, TokenKeyColumn, PasswordHashColumn, LastResetSentAtColumn, CreatedColumn, UpdatedColumn}
	)

	return adminsTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:              IDColumn,
		Avatar:          AvatarColumn,
		Email:           EmailColumn,
		TokenKey:        TokenKeyColumn,
		PasswordHash:    PasswordHashColumn,
		LastResetSentAt: LastResetSentAtColumn,
		Created:         CreatedColumn,
		Updated:         UpdatedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
