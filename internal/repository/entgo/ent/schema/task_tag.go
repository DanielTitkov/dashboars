package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// TaskTag holds the schema definition for the TaskTag entity.
type TaskTag struct {
	ent.Schema
}

// Fields of the TaskTag.
func (TaskTag) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Unique().NotEmpty(),
		field.String("description").Optional().MaxLen(280),
		field.Bool("display").Default(false),
		field.JSON("meta", make(map[string]interface{})).Optional(),
	}
}

// Edges of the TaskTag.
func (TaskTag) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("tasks", Task.Type),
	}
}

func (TaskTag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
