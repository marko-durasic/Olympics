package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CombinedSport holds the schema definition for the CombinedSport entity.
type CombinedSport struct {
	ent.Schema
}

// Fields of the CombinedSport.
func (CombinedSport) Fields() []ent.Field {
	return []ent.Field{
		field.String("country"),
		field.String("sport"),
		field.Int("gold"),
		field.Int("silver"),
		field.Int("bronze"),
		field.Int("points"),
		field.Int("total_score"),
		field.Int("population"),
		field.Float("per_capita"),
	}
}

// Edges of the CombinedSport.
func (CombinedSport) Edges() []ent.Edge {
	return nil
}
