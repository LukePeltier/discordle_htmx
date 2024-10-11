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

var Guesses = newGuessesTable("", "guesses", "")

type guessesTable struct {
	sqlite.Table

	// Columns
	Created     sqlite.ColumnString
	GuessedTime sqlite.ColumnString
	ID          sqlite.ColumnString
	Message     sqlite.ColumnString
	StartTime   sqlite.ColumnString
	Updated     sqlite.ColumnString
	Winner      sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type GuessesTable struct {
	guessesTable

	EXCLUDED guessesTable
}

// AS creates new GuessesTable with assigned alias
func (a GuessesTable) AS(alias string) *GuessesTable {
	return newGuessesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new GuessesTable with assigned schema name
func (a GuessesTable) FromSchema(schemaName string) *GuessesTable {
	return newGuessesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new GuessesTable with assigned table prefix
func (a GuessesTable) WithPrefix(prefix string) *GuessesTable {
	return newGuessesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new GuessesTable with assigned table suffix
func (a GuessesTable) WithSuffix(suffix string) *GuessesTable {
	return newGuessesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newGuessesTable(schemaName, tableName, alias string) *GuessesTable {
	return &GuessesTable{
		guessesTable: newGuessesTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newGuessesTableImpl("", "excluded", ""),
	}
}

func newGuessesTableImpl(schemaName, tableName, alias string) guessesTable {
	var (
		CreatedColumn     = sqlite.StringColumn("created")
		GuessedTimeColumn = sqlite.StringColumn("guessed_time")
		IDColumn          = sqlite.StringColumn("id")
		MessageColumn     = sqlite.StringColumn("message")
		StartTimeColumn   = sqlite.StringColumn("start_time")
		UpdatedColumn     = sqlite.StringColumn("updated")
		WinnerColumn      = sqlite.StringColumn("winner")
		allColumns        = sqlite.ColumnList{CreatedColumn, GuessedTimeColumn, IDColumn, MessageColumn, StartTimeColumn, UpdatedColumn, WinnerColumn}
		mutableColumns    = sqlite.ColumnList{CreatedColumn, GuessedTimeColumn, MessageColumn, StartTimeColumn, UpdatedColumn, WinnerColumn}
	)

	return guessesTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		Created:     CreatedColumn,
		GuessedTime: GuessedTimeColumn,
		ID:          IDColumn,
		Message:     MessageColumn,
		StartTime:   StartTimeColumn,
		Updated:     UpdatedColumn,
		Winner:      WinnerColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
