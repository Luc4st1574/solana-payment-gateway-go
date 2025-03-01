// Code generated by ent, DO NOT EDIT.

package users

import (
	"bitsnake-server/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldID, id))
}

// WalletAddress applies equality check predicate on the "wallet_address" field. It's identical to WalletAddressEQ.
func WalletAddress(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldWalletAddress, v))
}

// HasAccess applies equality check predicate on the "has_access" field. It's identical to HasAccessEQ.
func HasAccess(v bool) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldHasAccess, v))
}

// WalletAddressEQ applies the EQ predicate on the "wallet_address" field.
func WalletAddressEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldWalletAddress, v))
}

// WalletAddressNEQ applies the NEQ predicate on the "wallet_address" field.
func WalletAddressNEQ(v string) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldWalletAddress, v))
}

// WalletAddressIn applies the In predicate on the "wallet_address" field.
func WalletAddressIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldIn(FieldWalletAddress, vs...))
}

// WalletAddressNotIn applies the NotIn predicate on the "wallet_address" field.
func WalletAddressNotIn(vs ...string) predicate.Users {
	return predicate.Users(sql.FieldNotIn(FieldWalletAddress, vs...))
}

// WalletAddressGT applies the GT predicate on the "wallet_address" field.
func WalletAddressGT(v string) predicate.Users {
	return predicate.Users(sql.FieldGT(FieldWalletAddress, v))
}

// WalletAddressGTE applies the GTE predicate on the "wallet_address" field.
func WalletAddressGTE(v string) predicate.Users {
	return predicate.Users(sql.FieldGTE(FieldWalletAddress, v))
}

// WalletAddressLT applies the LT predicate on the "wallet_address" field.
func WalletAddressLT(v string) predicate.Users {
	return predicate.Users(sql.FieldLT(FieldWalletAddress, v))
}

// WalletAddressLTE applies the LTE predicate on the "wallet_address" field.
func WalletAddressLTE(v string) predicate.Users {
	return predicate.Users(sql.FieldLTE(FieldWalletAddress, v))
}

// WalletAddressContains applies the Contains predicate on the "wallet_address" field.
func WalletAddressContains(v string) predicate.Users {
	return predicate.Users(sql.FieldContains(FieldWalletAddress, v))
}

// WalletAddressHasPrefix applies the HasPrefix predicate on the "wallet_address" field.
func WalletAddressHasPrefix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasPrefix(FieldWalletAddress, v))
}

// WalletAddressHasSuffix applies the HasSuffix predicate on the "wallet_address" field.
func WalletAddressHasSuffix(v string) predicate.Users {
	return predicate.Users(sql.FieldHasSuffix(FieldWalletAddress, v))
}

// WalletAddressEqualFold applies the EqualFold predicate on the "wallet_address" field.
func WalletAddressEqualFold(v string) predicate.Users {
	return predicate.Users(sql.FieldEqualFold(FieldWalletAddress, v))
}

// WalletAddressContainsFold applies the ContainsFold predicate on the "wallet_address" field.
func WalletAddressContainsFold(v string) predicate.Users {
	return predicate.Users(sql.FieldContainsFold(FieldWalletAddress, v))
}

// HasAccessEQ applies the EQ predicate on the "has_access" field.
func HasAccessEQ(v bool) predicate.Users {
	return predicate.Users(sql.FieldEQ(FieldHasAccess, v))
}

// HasAccessNEQ applies the NEQ predicate on the "has_access" field.
func HasAccessNEQ(v bool) predicate.Users {
	return predicate.Users(sql.FieldNEQ(FieldHasAccess, v))
}

// HasMatches applies the HasEdge predicate on the "matches" edge.
func HasMatches() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MatchesTable, MatchesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMatchesWith applies the HasEdge predicate on the "matches" edge with a given conditions (other predicates).
func HasMatchesWith(preds ...predicate.Matches) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := newMatchesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMatchResults applies the HasEdge predicate on the "match_results" edge.
func HasMatchResults() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MatchResultsTable, MatchResultsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMatchResultsWith applies the HasEdge predicate on the "match_results" edge with a given conditions (other predicates).
func HasMatchResultsWith(preds ...predicate.MatchResults) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := newMatchResultsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPayments applies the HasEdge predicate on the "payments" edge.
func HasPayments() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PaymentsTable, PaymentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPaymentsWith applies the HasEdge predicate on the "payments" edge with a given conditions (other predicates).
func HasPaymentsWith(preds ...predicate.PaymentVerifications) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := newPaymentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Users) predicate.Users {
	return predicate.Users(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Users) predicate.Users {
	return predicate.Users(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Users) predicate.Users {
	return predicate.Users(sql.NotPredicates(p))
}
