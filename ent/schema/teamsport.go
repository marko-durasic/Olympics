package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TeamSport holds the schema definition for the TeamSport entity.
type TeamSport struct {
	ent.Schema
}

// Fields of the TeamSport.
func (TeamSport) Fields() []ent.Field {
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

// Edges of the TeamSport.
func (TeamSport) Edges() []ent.Edge {
	return nil
}
