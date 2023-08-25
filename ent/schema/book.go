package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("authorName"),
		field.Bool("isApproved").Default(false),
		field.String("coverPath"),
		field.String("pdfPath"),
		field.Int("createdBy").Optional(),
		field.Float("price"),
		field.Int("noSale").Default(0),
		field.Int("noView").Default(0),
		field.String("description"),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("createdBy").Unique().Field("createdBy"),
		edge.To("category", Category.Type),
	}
}
