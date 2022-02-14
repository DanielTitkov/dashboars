package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Dimension holds the schema definition for the Dimension entity.
type Dimension struct {
	ent.Schema
}

// Fields of the Dimension.
func (Dimension) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("value").NotEmpty(),
		field.JSON("meta", make(map[string]interface{})).Optional(),
	}
}

// Edges of the Dimension.
func (Dimension) Edges() []ent.Edge {
	return []ent.Edge{
		// m2m
		edge.From("item", Item.Type).Ref("dimensions"),
	}
}

func (Dimension) Indexes() []ent.Index {
	return []ent.Index{
		// unique title for item
		index.Fields("title", "value").Unique(),
	}
}

func (Dimension) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}
