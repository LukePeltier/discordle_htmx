package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Blacklist struct {
	ent.Schema
}

func (Blacklist) Fields() []ent.Field {
	return []ent.Field{
		field.String("bad").NotEmpty().Unique(),
	}
}
