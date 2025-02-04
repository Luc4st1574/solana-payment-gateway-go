package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
	return []ent.Field{
		field.String("wallet_address").
			Unique().
			NotEmpty(),
		// NEW: Add access control field
		field.Bool("has_access").
			Default(false).
			Comment("Whether the user has service access"),
	}
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("matches", Matches.Type), // A user can have multiple matches
		edge.To("match_results", MatchResults.Type),
		// NEW: Add relationship to payment verifications
		edge.To("payments", PaymentVerifications.Type).
			Comment("Payment verification records for this user"),
	}
}
