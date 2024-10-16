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

var Externalauths = newExternalauthsTable("", "_externalAuths", "")

type externalauthsTable struct {
	sqlite.Table

	// Columns
	ID           sqlite.ColumnString
	CollectionId sqlite.ColumnString
	RecordId     sqlite.ColumnString
	Provider     sqlite.ColumnString
	ProviderId   sqlite.ColumnString
	Created      sqlite.ColumnString
	Updated      sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type ExternalauthsTable struct {
	externalauthsTable

	EXCLUDED externalauthsTable
}

// AS creates new ExternalauthsTable with assigned alias
func (a ExternalauthsTable) AS(alias string) *ExternalauthsTable {
	return newExternalauthsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ExternalauthsTable with assigned schema name
func (a ExternalauthsTable) FromSchema(schemaName string) *ExternalauthsTable {
	return newExternalauthsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ExternalauthsTable with assigned table prefix
func (a ExternalauthsTable) WithPrefix(prefix string) *ExternalauthsTable {
	return newExternalauthsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ExternalauthsTable with assigned table suffix
func (a ExternalauthsTable) WithSuffix(suffix string) *ExternalauthsTable {
	return newExternalauthsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newExternalauthsTable(schemaName, tableName, alias string) *ExternalauthsTable {
	return &ExternalauthsTable{
		externalauthsTable: newExternalauthsTableImpl(schemaName, tableName, alias),
		EXCLUDED:           newExternalauthsTableImpl("", "excluded", ""),
	}
}

func newExternalauthsTableImpl(schemaName, tableName, alias string) externalauthsTable {
	var (
		IDColumn           = sqlite.StringColumn("id")
		CollectionIdColumn = sqlite.StringColumn("collectionId")
		RecordIdColumn     = sqlite.StringColumn("recordId")
		ProviderColumn     = sqlite.StringColumn("provider")
		ProviderIdColumn   = sqlite.StringColumn("providerId")
		CreatedColumn      = sqlite.StringColumn("created")
		UpdatedColumn      = sqlite.StringColumn("updated")
		allColumns         = sqlite.ColumnList{IDColumn, CollectionIdColumn, RecordIdColumn, ProviderColumn, ProviderIdColumn, CreatedColumn, UpdatedColumn}
		mutableColumns     = sqlite.ColumnList{CollectionIdColumn, RecordIdColumn, ProviderColumn, ProviderIdColumn, CreatedColumn, UpdatedColumn}
	)

	return externalauthsTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		CollectionId: CollectionIdColumn,
		RecordId:     RecordIdColumn,
		Provider:     ProviderColumn,
		ProviderId:   ProviderIdColumn,
		Created:      CreatedColumn,
		Updated:      UpdatedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
