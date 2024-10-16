// Code generated by ent, DO NOT EDIT.

package blacklist

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the blacklist type in the database.
	Label = "blacklist"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBad holds the string denoting the bad field in the database.
	FieldBad = "bad"
	// Table holds the table name of the blacklist in the database.
	Table = "blacklists"
)

// Columns holds all SQL columns for blacklist fields.
var Columns = []string{
	FieldID,
	FieldBad,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// BadValidator is a validator for the "bad" field. It is called by the builders before save.
	BadValidator func(string) error
)

// OrderOption defines the ordering options for the Blacklist queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBad orders the results by the bad field.
func ByBad(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBad, opts...).ToFunc()
}
