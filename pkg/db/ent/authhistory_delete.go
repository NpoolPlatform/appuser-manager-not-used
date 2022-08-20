// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/authhistory"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
)

// AuthHistoryDelete is the builder for deleting a AuthHistory entity.
type AuthHistoryDelete struct {
	config
	hooks    []Hook
	mutation *AuthHistoryMutation
}

// Where appends a list predicates to the AuthHistoryDelete builder.
func (ahd *AuthHistoryDelete) Where(ps ...predicate.AuthHistory) *AuthHistoryDelete {
	ahd.mutation.Where(ps...)
	return ahd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ahd *AuthHistoryDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ahd.hooks) == 0 {
		affected, err = ahd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthHistoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ahd.mutation = mutation
			affected, err = ahd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ahd.hooks) - 1; i >= 0; i-- {
			if ahd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ahd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ahd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ahd *AuthHistoryDelete) ExecX(ctx context.Context) int {
	n, err := ahd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ahd *AuthHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: authhistory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: authhistory.FieldID,
			},
		},
	}
	if ps := ahd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ahd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// AuthHistoryDeleteOne is the builder for deleting a single AuthHistory entity.
type AuthHistoryDeleteOne struct {
	ahd *AuthHistoryDelete
}

// Exec executes the deletion query.
func (ahdo *AuthHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := ahdo.ahd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{authhistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ahdo *AuthHistoryDeleteOne) ExecX(ctx context.Context) {
	ahdo.ahd.ExecX(ctx)
}
