// Code generated by ent, DO NOT EDIT.

package ent

import (
	"bitsnake-server/internal/ent/matchresults"
	"bitsnake-server/internal/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MatchResultsDelete is the builder for deleting a MatchResults entity.
type MatchResultsDelete struct {
	config
	hooks    []Hook
	mutation *MatchResultsMutation
}

// Where appends a list predicates to the MatchResultsDelete builder.
func (mrd *MatchResultsDelete) Where(ps ...predicate.MatchResults) *MatchResultsDelete {
	mrd.mutation.Where(ps...)
	return mrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mrd *MatchResultsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, mrd.sqlExec, mrd.mutation, mrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mrd *MatchResultsDelete) ExecX(ctx context.Context) int {
	n, err := mrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mrd *MatchResultsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(matchresults.Table, sqlgraph.NewFieldSpec(matchresults.FieldID, field.TypeInt))
	if ps := mrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mrd.mutation.done = true
	return affected, err
}

// MatchResultsDeleteOne is the builder for deleting a single MatchResults entity.
type MatchResultsDeleteOne struct {
	mrd *MatchResultsDelete
}

// Where appends a list predicates to the MatchResultsDelete builder.
func (mrdo *MatchResultsDeleteOne) Where(ps ...predicate.MatchResults) *MatchResultsDeleteOne {
	mrdo.mrd.mutation.Where(ps...)
	return mrdo
}

// Exec executes the deletion query.
func (mrdo *MatchResultsDeleteOne) Exec(ctx context.Context) error {
	n, err := mrdo.mrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{matchresults.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mrdo *MatchResultsDeleteOne) ExecX(ctx context.Context) {
	if err := mrdo.Exec(ctx); err != nil {
		panic(err)
	}
}
