package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Spacecraft holds the schema definition for the Spacecraft entity.
type Spacecraft struct {
	ent.Schema
}

// Fields of the Spacecraft.
func (Spacecraft) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("class").NotEmpty(),
		field.Uint32("crew"),
		field.String("image").Optional(),
		field.Float("value"),
		field.String("status").NotEmpty(),
		field.Bool("deleted").Default(false),
	}
}

// Edges of the Spacecraft.
func (Spacecraft) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("armament", Armement.Type).StorageKey(edge.Column("spacecraft_id")),
	}
}

// Indexes of the Spacecraft.
func (Spacecraft) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "class", "status").Unique(),
	}
}
