package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Armement holds the schema definition for the Armement entity.
type Armement struct {
	ent.Schema
}

// Fields of the Armement.
func (Armement) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("quantity").NotEmpty(),
	}
}

// Edges of the Armement.
func (Armement) Edges() []ent.Edge {
	return nil
}
