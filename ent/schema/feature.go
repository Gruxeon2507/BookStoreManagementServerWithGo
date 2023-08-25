package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Feature holds the schema definition for the Feature entity.
type Feature struct {
	ent.Schema
}

// Fields of the Feature.
func (Feature) Fields() []ent.Field {
	return []ent.Field{
		field.String("featureName"),
		field.String("featureUrl"),
	}
}

// Edges of the Feature.
func (Feature) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", Role.Type).Ref("feature"),
	}
}
