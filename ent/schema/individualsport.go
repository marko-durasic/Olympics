package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// IndividualSport holds the schema definition for the IndividualSport entity.
type IndividualSport struct {
	ent.Schema
}

// Fields of the IndividualSport.
func (IndividualSport) Fields() []ent.Field {
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

// Edges of the IndividualSport.
func (IndividualSport) Edges() []ent.Edge {
	return nil
}
