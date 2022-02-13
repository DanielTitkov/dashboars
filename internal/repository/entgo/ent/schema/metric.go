package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Metric holds the schema definition for the Metric entity.
type Metric struct {
	ent.Schema
}

// Fields of the Metric.
func (Metric) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("description").Optional().MaxLen(280),
		field.JSON("meta", make(map[string]interface{})).Optional(),
	}
}

// Edges of the Metric.
func (Metric) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("items", Item.Type),
		// belongs to
		edge.From("task", Task.Type).Ref("metrics").Unique().Required(),
	}
}

func (Metric) Indexes() []ent.Index {
	return []ent.Index{
		// unique title for task
		index.Fields("title").Edges("task").Unique(),
	}
}

func (Metric) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
