package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type DiscordMember struct {
	ent.Schema
}

func (DiscordMember) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().NotEmpty(),
		field.JSON("nicknames", []string{}),
		field.String("discord_id"),
		field.Bool("blacklisted").Default(false),
	}
}

func (DiscordMember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("messages", Message.Type),
	}
}
