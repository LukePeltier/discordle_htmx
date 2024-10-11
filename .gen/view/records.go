//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package view

import (
	"github.com/go-jet/jet/v2/sqlite"
)

var Records = newRecordsTable("", "records", "")

type recordsTable struct {
	sqlite.Table

	// Columns
	Username       sqlite.ColumnString
	CorrectGuesses sqlite.ColumnString
	ID             sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type RecordsTable struct {
	recordsTable

	EXCLUDED recordsTable
}

// AS creates new RecordsTable with assigned alias
func (a RecordsTable) AS(alias string) *RecordsTable {
	return newRecordsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new RecordsTable with assigned schema name
func (a RecordsTable) FromSchema(schemaName string) *RecordsTable {
	return newRecordsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new RecordsTable with assigned table prefix
func (a RecordsTable) WithPrefix(prefix string) *RecordsTable {
	return newRecordsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new RecordsTable with assigned table suffix
func (a RecordsTable) WithSuffix(suffix string) *RecordsTable {
	return newRecordsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newRecordsTable(schemaName, tableName, alias string) *RecordsTable {
	return &RecordsTable{
		recordsTable: newRecordsTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newRecordsTableImpl("", "excluded", ""),
	}
}

func newRecordsTableImpl(schemaName, tableName, alias string) recordsTable {
	var (
		UsernameColumn       = sqlite.StringColumn("username")
		CorrectGuessesColumn = sqlite.StringColumn("correct_guesses")
		IDColumn             = sqlite.StringColumn("id")
		allColumns           = sqlite.ColumnList{UsernameColumn, CorrectGuessesColumn, IDColumn}
		mutableColumns       = sqlite.ColumnList{UsernameColumn, CorrectGuessesColumn, IDColumn}
	)

	return recordsTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		Username:       UsernameColumn,
		CorrectGuesses: CorrectGuessesColumn,
		ID:             IDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
