// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
)

// AppUserSecretDelete is the builder for deleting a AppUserSecret entity.
type AppUserSecretDelete struct {
	config
	hooks    []Hook
	mutation *AppUserSecretMutation
}

// Where appends a list predicates to the AppUserSecretDelete builder.
func (ausd *AppUserSecretDelete) Where(ps ...predicate.AppUserSecret) *AppUserSecretDelete {
	ausd.mutation.Where(ps...)
	return ausd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ausd *AppUserSecretDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ausd.hooks) == 0 {
		affected, err = ausd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserSecretMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ausd.mutation = mutation
			affected, err = ausd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ausd.hooks) - 1; i >= 0; i-- {
			if ausd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ausd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ausd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ausd *AppUserSecretDelete) ExecX(ctx context.Context) int {
	n, err := ausd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ausd *AppUserSecretDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: appusersecret.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appusersecret.FieldID,
			},
		},
	}
	if ps := ausd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ausd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// AppUserSecretDeleteOne is the builder for deleting a single AppUserSecret entity.
type AppUserSecretDeleteOne struct {
	ausd *AppUserSecretDelete
}

// Exec executes the deletion query.
func (ausdo *AppUserSecretDeleteOne) Exec(ctx context.Context) error {
	n, err := ausdo.ausd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appusersecret.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ausdo *AppUserSecretDeleteOne) ExecX(ctx context.Context) {
	ausdo.ausd.ExecX(ctx)
}
