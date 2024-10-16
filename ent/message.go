// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/lukepeltier/discordle/ent/discordmember"
	"github.com/lukepeltier/discordle/ent/message"
)

// Message is the model entity for the Message schema.
type Message struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// DateSent holds the value of the "date_sent" field.
	DateSent string `json:"date_sent,omitempty"`
	// DiscordID holds the value of the "discord_id" field.
	DiscordID string `json:"discord_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MessageQuery when eager-loading is set.
	Edges                   MessageEdges `json:"edges"`
	discord_member_messages *int
	selectValues            sql.SelectValues
}

// MessageEdges holds the relations/edges for other nodes in the graph.
type MessageEdges struct {
	// Sender holds the value of the sender edge.
	Sender *DiscordMember `json:"sender,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SenderOrErr returns the Sender value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MessageEdges) SenderOrErr() (*DiscordMember, error) {
	if e.Sender != nil {
		return e.Sender, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: discordmember.Label}
	}
	return nil, &NotLoadedError{edge: "sender"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Message) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case message.FieldID:
			values[i] = new(sql.NullInt64)
		case message.FieldText, message.FieldDateSent, message.FieldDiscordID:
			values[i] = new(sql.NullString)
		case message.ForeignKeys[0]: // discord_member_messages
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Message fields.
func (m *Message) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case message.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case message.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				m.Text = value.String
			}
		case message.FieldDateSent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field date_sent", values[i])
			} else if value.Valid {
				m.DateSent = value.String
			}
		case message.FieldDiscordID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field discord_id", values[i])
			} else if value.Valid {
				m.DiscordID = value.String
			}
		case message.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field discord_member_messages", value)
			} else if value.Valid {
				m.discord_member_messages = new(int)
				*m.discord_member_messages = int(value.Int64)
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Message.
// This includes values selected through modifiers, order, etc.
func (m *Message) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QuerySender queries the "sender" edge of the Message entity.
func (m *Message) QuerySender() *DiscordMemberQuery {
	return NewMessageClient(m.config).QuerySender(m)
}

// Update returns a builder for updating this Message.
// Note that you need to call Message.Unwrap() before calling this method if this Message
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Message) Update() *MessageUpdateOne {
	return NewMessageClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Message entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Message) Unwrap() *Message {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Message is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Message) String() string {
	var builder strings.Builder
	builder.WriteString("Message(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("text=")
	builder.WriteString(m.Text)
	builder.WriteString(", ")
	builder.WriteString("date_sent=")
	builder.WriteString(m.DateSent)
	builder.WriteString(", ")
	builder.WriteString("discord_id=")
	builder.WriteString(m.DiscordID)
	builder.WriteByte(')')
	return builder.String()
}

// Messages is a parsable slice of Message.
type Messages []*Message
