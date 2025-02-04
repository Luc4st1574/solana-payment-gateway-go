package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Matches holds the schema definition for the Matches entity.
type Matches struct {
	ent.Schema
}

// Fields of the Matches.
func (Matches) Fields() []ent.Field {
	return []ent.Field{
		field.String("game_hash_id").
			Unique().
			NotEmpty(), // Unique GameHashID for each session
		field.Time("expiration_date").
			Optional(), // Expiration date for the session
		field.String("wallet_address").
			NotEmpty(), // User's wallet address
		field.String("transaction_hash").
			NotEmpty(), // Transaction hash from payment verification
		field.Time("created_at").
			Default(time.Now), // Align with DEFAULT NOW() in PostgreSQL
	}
}

// Edges of the Matches.
func (Matches) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("match_results", MatchResults.Type),
		edge.From("user", Users.Type).
			Ref("matches").
			Unique().
			Required(), // Match belongs to a single user
	}
}
