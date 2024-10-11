package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Message struct {
	ent.Schema
}

func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").NotEmpty(),
		field.String("date_sent").NotEmpty(),
		field.String("discord_id"),
	}
}

func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sender", DiscordMember.Type).Ref("messages").Unique(),
	}
}
