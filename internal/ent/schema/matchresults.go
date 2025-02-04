package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MatchResults holds the schema definition for the MatchResults entity.
type MatchResults struct {
	ent.Schema
}

// Fields of the MatchResults.
func (MatchResults) Fields() []ent.Field {
	return []ent.Field{
		field.Int("kills").
			Default(0),
		field.Bool("is_winner").
			Default(false),
		field.Float("reward_amount").
			Optional().
			Nillable(),
	}
}

// Edges of the MatchResults.
func (MatchResults) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Users.Type).
			Ref("match_results").
			Unique().
			Required(),
		edge.From("match", Matches.Type).
			Ref("match_results").
			Unique().
			Required(),
	}
}
