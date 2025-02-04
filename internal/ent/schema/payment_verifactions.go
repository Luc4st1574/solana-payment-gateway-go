package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PaymentVerifications holds the schema definition for the PaymentVerifications entity.
type PaymentVerifications struct {
	ent.Schema
}

// Fields of the PaymentVerifications.
func (PaymentVerifications) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").
			Comment("Foreign key to the user who made the payment"),

		field.String("wallet_address").
			MaxLen(255).
			NotEmpty().
			Comment("Wallet address where payment was received"),

		field.Float("amount").
			GoType(float64(0)).
			SchemaType(map[string]string{
				"postgres": "numeric(20,8)",
			}).
			Comment("Amount received in SOL"),

		field.String("transaction_hash").
			MaxLen(88).
			Optional().
			Comment("Transaction hash from Solana blockchain"),

		field.Enum("status").
			Values("pending", "verified", "failed").
			Default("pending").
			Comment("Current verification status"),

		field.Bool("access_granted").
			Default(false).
			Comment("Whether access was granted after verification"),

		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("Timestamp when the record was created"),
	}
}

// Edges of the PaymentVerifications.
func (PaymentVerifications) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Users.Type).
			Ref("payments").
			Field("user_id").
			Unique().
			Required().
			Comment("User who made the payment"),
	}
}

// Indexes of the PaymentVerifications.
func (PaymentVerifications) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("wallet_address"),
		index.Fields("status"),
		index.Fields("created_at"),
		index.Fields("transaction_hash").Unique(),
	}
}

// Mixin for common fields
func (PaymentVerifications) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Mixin{}, // If you have a base mixin for common fields like ID, timestamps etc.
	}
}
