// Code generated by ent, DO NOT EDIT.

package matches

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the matches type in the database.
	Label = "matches"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGameHashID holds the string denoting the game_hash_id field in the database.
	FieldGameHashID = "game_hash_id"
	// FieldExpirationDate holds the string denoting the expiration_date field in the database.
	FieldExpirationDate = "expiration_date"
	// FieldWalletAddress holds the string denoting the wallet_address field in the database.
	FieldWalletAddress = "wallet_address"
	// FieldTransactionHash holds the string denoting the transaction_hash field in the database.
	FieldTransactionHash = "transaction_hash"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeMatchResults holds the string denoting the match_results edge name in mutations.
	EdgeMatchResults = "match_results"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the matches in the database.
	Table = "matches"
	// MatchResultsTable is the table that holds the match_results relation/edge.
	MatchResultsTable = "match_results"
	// MatchResultsInverseTable is the table name for the MatchResults entity.
	// It exists in this package in order to avoid circular dependency with the "matchresults" package.
	MatchResultsInverseTable = "match_results"
	// MatchResultsColumn is the table column denoting the match_results relation/edge.
	MatchResultsColumn = "matches_match_results"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "matches"
	// UserInverseTable is the table name for the Users entity.
	// It exists in this package in order to avoid circular dependency with the "users" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "users_matches"
)

// Columns holds all SQL columns for matches fields.
var Columns = []string{
	FieldID,
	FieldGameHashID,
	FieldExpirationDate,
	FieldWalletAddress,
	FieldTransactionHash,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "matches"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"users_matches",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// GameHashIDValidator is a validator for the "game_hash_id" field. It is called by the builders before save.
	GameHashIDValidator func(string) error
	// WalletAddressValidator is a validator for the "wallet_address" field. It is called by the builders before save.
	WalletAddressValidator func(string) error
	// TransactionHashValidator is a validator for the "transaction_hash" field. It is called by the builders before save.
	TransactionHashValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the Matches queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByGameHashID orders the results by the game_hash_id field.
func ByGameHashID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGameHashID, opts...).ToFunc()
}

// ByExpirationDate orders the results by the expiration_date field.
func ByExpirationDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpirationDate, opts...).ToFunc()
}

// ByWalletAddress orders the results by the wallet_address field.
func ByWalletAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWalletAddress, opts...).ToFunc()
}

// ByTransactionHash orders the results by the transaction_hash field.
func ByTransactionHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTransactionHash, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByMatchResultsCount orders the results by match_results count.
func ByMatchResultsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMatchResultsStep(), opts...)
	}
}

// ByMatchResults orders the results by match_results terms.
func ByMatchResults(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMatchResultsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newMatchResultsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MatchResultsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MatchResultsTable, MatchResultsColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
