package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.Float("value"),
		field.Time("timestamp").Default(time.Now),
		field.JSON("meta", make(map[string]interface{})).Optional(),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		// belongs to
		edge.From("task_instance", TaskInstance.Type).Ref("items").Unique().Required(),
		edge.From("metric", Metric.Type).Ref("items").Unique().Required(),
	}
}

func (Item) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
