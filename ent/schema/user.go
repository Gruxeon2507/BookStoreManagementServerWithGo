package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique(),
		field.String("password"),
		field.String("displayName"),
		field.Time("dob"),
		field.String("email").Unique(),
		field.Time("createdDate"),
		field.String("avatarPath"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("createdBy", Book.Type),
		edge.From("role", Role.Type).Ref("user"),
	}
}
