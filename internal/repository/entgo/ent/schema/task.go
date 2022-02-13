package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/DanielTitkov/dashboars/internal/domain"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values(
			domain.TaskTypeRandom,
			domain.TaskTypeFailing,
		).Immutable(),
		field.String("code").Unique().NotEmpty().Immutable(),
		field.String("title").Unique().NotEmpty(),
		field.String("description").Optional().MaxLen(280),
		field.Bool("active").Default(false),
		field.Bool("display").Default(false),
		field.String("schedule").Optional(),
		field.JSON("args", make(map[string]interface{})).Optional(),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("instances", TaskInstance.Type),
		edge.To("metrics", Metric.Type),
	}
}

func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
