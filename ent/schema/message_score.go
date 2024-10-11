package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type MessageScore struct {
	ent.Schema
}

func (MessageScore) Fields() []ent.Field {
	return []ent.Field{
		field.Uint16("score"),
	}
}

func (MessageScore) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("message", Message.Type).Unique(),
	}
}
