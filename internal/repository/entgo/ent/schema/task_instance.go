package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// TaskInstance holds the schema definition for the TaskInstance entity.
type TaskInstance struct {
	ent.Schema
}

// Fields of the TaskInstance.
func (TaskInstance) Fields() []ent.Field {
	return []ent.Field{
		field.Time("start_time").Default(time.Now),
		field.Time("end_time").Optional(),
		field.Int("attempt").Immutable().Default(0),
		field.Bool("success").Nillable().Optional(),
		field.Bool("running").Default(false),
		field.String("error").Nillable().Optional(),
		field.JSON("meta", make(map[string]interface{})).Optional(),
	}
}

// Edges of the TaskInstance.
func (TaskInstance) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("items", Item.Type),
		// belongs to
		edge.From("task", Task.Type).Ref("instances").Unique().Required(),
	}
}

func (TaskInstance) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
