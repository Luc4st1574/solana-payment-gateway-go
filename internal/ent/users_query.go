// Code generated by ent, DO NOT EDIT.

package ent

import (
	"bitsnake-server/internal/ent/matches"
	"bitsnake-server/internal/ent/matchresults"
	"bitsnake-server/internal/ent/paymentverifications"
	"bitsnake-server/internal/ent/predicate"
	"bitsnake-server/internal/ent/users"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UsersQuery is the builder for querying Users entities.
type UsersQuery struct {
	config
	ctx              *QueryContext
	order            []users.OrderOption
	inters           []Interceptor
	predicates       []predicate.Users
	withMatches      *MatchesQuery
	withMatchResults *MatchResultsQuery
	withPayments     *PaymentVerificationsQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UsersQuery builder.
func (uq *UsersQuery) Where(ps ...predicate.Users) *UsersQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit the number of records to be returned by this query.
func (uq *UsersQuery) Limit(limit int) *UsersQuery {
	uq.ctx.Limit = &limit
	return uq
}

// Offset to start from.
func (uq *UsersQuery) Offset(offset int) *UsersQuery {
	uq.ctx.Offset = &offset
	return uq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uq *UsersQuery) Unique(unique bool) *UsersQuery {
	uq.ctx.Unique = &unique
	return uq
}

// Order specifies how the records should be ordered.
func (uq *UsersQuery) Order(o ...users.OrderOption) *UsersQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QueryMatches chains the current query on the "matches" edge.
func (uq *UsersQuery) QueryMatches() *MatchesQuery {
	query := (&MatchesClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(users.Table, users.FieldID, selector),
			sqlgraph.To(matches.Table, matches.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, users.MatchesTable, users.MatchesColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMatchResults chains the current query on the "match_results" edge.
func (uq *UsersQuery) QueryMatchResults() *MatchResultsQuery {
	query := (&MatchResultsClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(users.Table, users.FieldID, selector),
			sqlgraph.To(matchresults.Table, matchresults.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, users.MatchResultsTable, users.MatchResultsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPayments chains the current query on the "payments" edge.
func (uq *UsersQuery) QueryPayments() *PaymentVerificationsQuery {
	query := (&PaymentVerificationsClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(users.Table, users.FieldID, selector),
			sqlgraph.To(paymentverifications.Table, paymentverifications.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, users.PaymentsTable, users.PaymentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Users entity from the query.
// Returns a *NotFoundError when no Users was found.
func (uq *UsersQuery) First(ctx context.Context) (*Users, error) {
	nodes, err := uq.Limit(1).All(setContextOp(ctx, uq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{users.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UsersQuery) FirstX(ctx context.Context) *Users {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Users ID from the query.
// Returns a *NotFoundError when no Users ID was found.
func (uq *UsersQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(1).IDs(setContextOp(ctx, uq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{users.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *UsersQuery) FirstIDX(ctx context.Context) int {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Users entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Users entity is found.
// Returns a *NotFoundError when no Users entities are found.
func (uq *UsersQuery) Only(ctx context.Context) (*Users, error) {
	nodes, err := uq.Limit(2).All(setContextOp(ctx, uq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{users.Label}
	default:
		return nil, &NotSingularError{users.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UsersQuery) OnlyX(ctx context.Context) *Users {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Users ID in the query.
// Returns a *NotSingularError when more than one Users ID is found.
// Returns a *NotFoundError when no entities are found.
func (uq *UsersQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(2).IDs(setContextOp(ctx, uq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{users.Label}
	default:
		err = &NotSingularError{users.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UsersQuery) OnlyIDX(ctx context.Context) int {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UsersSlice.
func (uq *UsersQuery) All(ctx context.Context) ([]*Users, error) {
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryAll)
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Users, *UsersQuery]()
	return withInterceptors[[]*Users](ctx, uq, qr, uq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uq *UsersQuery) AllX(ctx context.Context) []*Users {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Users IDs.
func (uq *UsersQuery) IDs(ctx context.Context) (ids []int, err error) {
	if uq.ctx.Unique == nil && uq.path != nil {
		uq.Unique(true)
	}
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryIDs)
	if err = uq.Select(users.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UsersQuery) IDsX(ctx context.Context) []int {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UsersQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryCount)
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uq, querierCount[*UsersQuery](), uq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UsersQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UsersQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryExist)
	switch _, err := uq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UsersQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UsersQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UsersQuery) Clone() *UsersQuery {
	if uq == nil {
		return nil
	}
	return &UsersQuery{
		config:           uq.config,
		ctx:              uq.ctx.Clone(),
		order:            append([]users.OrderOption{}, uq.order...),
		inters:           append([]Interceptor{}, uq.inters...),
		predicates:       append([]predicate.Users{}, uq.predicates...),
		withMatches:      uq.withMatches.Clone(),
		withMatchResults: uq.withMatchResults.Clone(),
		withPayments:     uq.withPayments.Clone(),
		// clone intermediate query.
		sql:  uq.sql.Clone(),
		path: uq.path,
	}
}

// WithMatches tells the query-builder to eager-load the nodes that are connected to
// the "matches" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UsersQuery) WithMatches(opts ...func(*MatchesQuery)) *UsersQuery {
	query := (&MatchesClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withMatches = query
	return uq
}

// WithMatchResults tells the query-builder to eager-load the nodes that are connected to
// the "match_results" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UsersQuery) WithMatchResults(opts ...func(*MatchResultsQuery)) *UsersQuery {
	query := (&MatchResultsClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withMatchResults = query
	return uq
}

// WithPayments tells the query-builder to eager-load the nodes that are connected to
// the "payments" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UsersQuery) WithPayments(opts ...func(*PaymentVerificationsQuery)) *UsersQuery {
	query := (&PaymentVerificationsClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withPayments = query
	return uq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		WalletAddress string `json:"wallet_address,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Users.Query().
//		GroupBy(users.FieldWalletAddress).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uq *UsersQuery) GroupBy(field string, fields ...string) *UsersGroupBy {
	uq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UsersGroupBy{build: uq}
	grbuild.flds = &uq.ctx.Fields
	grbuild.label = users.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		WalletAddress string `json:"wallet_address,omitempty"`
//	}
//
//	client.Users.Query().
//		Select(users.FieldWalletAddress).
//		Scan(ctx, &v)
func (uq *UsersQuery) Select(fields ...string) *UsersSelect {
	uq.ctx.Fields = append(uq.ctx.Fields, fields...)
	sbuild := &UsersSelect{UsersQuery: uq}
	sbuild.label = users.Label
	sbuild.flds, sbuild.scan = &uq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UsersSelect configured with the given aggregations.
func (uq *UsersQuery) Aggregate(fns ...AggregateFunc) *UsersSelect {
	return uq.Select().Aggregate(fns...)
}

func (uq *UsersQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uq); err != nil {
				return err
			}
		}
	}
	for _, f := range uq.ctx.Fields {
		if !users.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.sql = prev
	}
	return nil
}

func (uq *UsersQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Users, error) {
	var (
		nodes       = []*Users{}
		_spec       = uq.querySpec()
		loadedTypes = [3]bool{
			uq.withMatches != nil,
			uq.withMatchResults != nil,
			uq.withPayments != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Users).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Users{config: uq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uq.withMatches; query != nil {
		if err := uq.loadMatches(ctx, query, nodes,
			func(n *Users) { n.Edges.Matches = []*Matches{} },
			func(n *Users, e *Matches) { n.Edges.Matches = append(n.Edges.Matches, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withMatchResults; query != nil {
		if err := uq.loadMatchResults(ctx, query, nodes,
			func(n *Users) { n.Edges.MatchResults = []*MatchResults{} },
			func(n *Users, e *MatchResults) { n.Edges.MatchResults = append(n.Edges.MatchResults, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withPayments; query != nil {
		if err := uq.loadPayments(ctx, query, nodes,
			func(n *Users) { n.Edges.Payments = []*PaymentVerifications{} },
			func(n *Users, e *PaymentVerifications) { n.Edges.Payments = append(n.Edges.Payments, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uq *UsersQuery) loadMatches(ctx context.Context, query *MatchesQuery, nodes []*Users, init func(*Users), assign func(*Users, *Matches)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Users)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Matches(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(users.MatchesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.users_matches
		if fk == nil {
			return fmt.Errorf(`foreign-key "users_matches" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "users_matches" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UsersQuery) loadMatchResults(ctx context.Context, query *MatchResultsQuery, nodes []*Users, init func(*Users), assign func(*Users, *MatchResults)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Users)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.MatchResults(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(users.MatchResultsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.users_match_results
		if fk == nil {
			return fmt.Errorf(`foreign-key "users_match_results" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "users_match_results" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UsersQuery) loadPayments(ctx context.Context, query *PaymentVerificationsQuery, nodes []*Users, init func(*Users), assign func(*Users, *PaymentVerifications)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Users)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(paymentverifications.FieldUserID)
	}
	query.Where(predicate.PaymentVerifications(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(users.PaymentsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UserID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (uq *UsersQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	_spec.Node.Columns = uq.ctx.Fields
	if len(uq.ctx.Fields) > 0 {
		_spec.Unique = uq.ctx.Unique != nil && *uq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UsersQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(users.Table, users.Columns, sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt))
	_spec.From = uq.sql
	if unique := uq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uq.path != nil {
		_spec.Unique = true
	}
	if fields := uq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, users.FieldID)
		for i := range fields {
			if fields[i] != users.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *UsersQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(users.Table)
	columns := uq.ctx.Fields
	if len(columns) == 0 {
		columns = users.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uq.ctx.Unique != nil && *uq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UsersGroupBy is the group-by builder for Users entities.
type UsersGroupBy struct {
	selector
	build *UsersQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UsersGroupBy) Aggregate(fns ...AggregateFunc) *UsersGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the selector query and scans the result into the given value.
func (ugb *UsersGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ugb.build.ctx, ent.OpQueryGroupBy)
	if err := ugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UsersQuery, *UsersGroupBy](ctx, ugb.build, ugb, ugb.build.inters, v)
}

func (ugb *UsersGroupBy) sqlScan(ctx context.Context, root *UsersQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ugb.fns))
	for _, fn := range ugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ugb.flds)+len(ugb.fns))
		for _, f := range *ugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UsersSelect is the builder for selecting fields of Users entities.
type UsersSelect struct {
	*UsersQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (us *UsersSelect) Aggregate(fns ...AggregateFunc) *UsersSelect {
	us.fns = append(us.fns, fns...)
	return us
}

// Scan applies the selector query and scans the result into the given value.
func (us *UsersSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, us.ctx, ent.OpQuerySelect)
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UsersQuery, *UsersSelect](ctx, us.UsersQuery, us, us.inters, v)
}

func (us *UsersSelect) sqlScan(ctx context.Context, root *UsersQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(us.fns))
	for _, fn := range us.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*us.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
