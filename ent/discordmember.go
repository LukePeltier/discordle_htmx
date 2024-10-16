// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/lukepeltier/discordle/ent/discordmember"
)

// DiscordMember is the model entity for the DiscordMember schema.
type DiscordMember struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Nicknames holds the value of the "nicknames" field.
	Nicknames []string `json:"nicknames,omitempty"`
	// DiscordID holds the value of the "discord_id" field.
	DiscordID string `json:"discord_id,omitempty"`
	// Blacklisted holds the value of the "blacklisted" field.
	Blacklisted bool `json:"blacklisted,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DiscordMemberQuery when eager-loading is set.
	Edges        DiscordMemberEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DiscordMemberEdges holds the relations/edges for other nodes in the graph.
type DiscordMemberEdges struct {
	// Messages holds the value of the messages edge.
	Messages []*Message `json:"messages,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MessagesOrErr returns the Messages value or an error if the edge
// was not loaded in eager-loading.
func (e DiscordMemberEdges) MessagesOrErr() ([]*Message, error) {
	if e.loadedTypes[0] {
		return e.Messages, nil
	}
	return nil, &NotLoadedError{edge: "messages"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DiscordMember) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case discordmember.FieldNicknames:
			values[i] = new([]byte)
		case discordmember.FieldBlacklisted:
			values[i] = new(sql.NullBool)
		case discordmember.FieldID:
			values[i] = new(sql.NullInt64)
		case discordmember.FieldUsername, discordmember.FieldDiscordID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DiscordMember fields.
func (dm *DiscordMember) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case discordmember.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dm.ID = int(value.Int64)
		case discordmember.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				dm.Username = value.String
			}
		case discordmember.FieldNicknames:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field nicknames", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &dm.Nicknames); err != nil {
					return fmt.Errorf("unmarshal field nicknames: %w", err)
				}
			}
		case discordmember.FieldDiscordID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field discord_id", values[i])
			} else if value.Valid {
				dm.DiscordID = value.String
			}
		case discordmember.FieldBlacklisted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field blacklisted", values[i])
			} else if value.Valid {
				dm.Blacklisted = value.Bool
			}
		default:
			dm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DiscordMember.
// This includes values selected through modifiers, order, etc.
func (dm *DiscordMember) Value(name string) (ent.Value, error) {
	return dm.selectValues.Get(name)
}

// QueryMessages queries the "messages" edge of the DiscordMember entity.
func (dm *DiscordMember) QueryMessages() *MessageQuery {
	return NewDiscordMemberClient(dm.config).QueryMessages(dm)
}

// Update returns a builder for updating this DiscordMember.
// Note that you need to call DiscordMember.Unwrap() before calling this method if this DiscordMember
// was returned from a transaction, and the transaction was committed or rolled back.
func (dm *DiscordMember) Update() *DiscordMemberUpdateOne {
	return NewDiscordMemberClient(dm.config).UpdateOne(dm)
}

// Unwrap unwraps the DiscordMember entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dm *DiscordMember) Unwrap() *DiscordMember {
	_tx, ok := dm.config.driver.(*txDriver)
	if !ok {
		panic("ent: DiscordMember is not a transactional entity")
	}
	dm.config.driver = _tx.drv
	return dm
}

// String implements the fmt.Stringer.
func (dm *DiscordMember) String() string {
	var builder strings.Builder
	builder.WriteString("DiscordMember(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dm.ID))
	builder.WriteString("username=")
	builder.WriteString(dm.Username)
	builder.WriteString(", ")
	builder.WriteString("nicknames=")
	builder.WriteString(fmt.Sprintf("%v", dm.Nicknames))
	builder.WriteString(", ")
	builder.WriteString("discord_id=")
	builder.WriteString(dm.DiscordID)
	builder.WriteString(", ")
	builder.WriteString("blacklisted=")
	builder.WriteString(fmt.Sprintf("%v", dm.Blacklisted))
	builder.WriteByte(')')
	return builder.String()
}

// DiscordMembers is a parsable slice of DiscordMember.
type DiscordMembers []*DiscordMember
