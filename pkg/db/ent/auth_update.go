// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/auth"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AuthUpdate is the builder for updating Auth entities.
type AuthUpdate struct {
	config
	hooks     []Hook
	mutation  *AuthMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AuthUpdate builder.
func (au *AuthUpdate) Where(ps ...predicate.Auth) *AuthUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AuthUpdate) SetCreatedAt(u uint32) *AuthUpdate {
	au.mutation.ResetCreatedAt()
	au.mutation.SetCreatedAt(u)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AuthUpdate) SetNillableCreatedAt(u *uint32) *AuthUpdate {
	if u != nil {
		au.SetCreatedAt(*u)
	}
	return au
}

// AddCreatedAt adds u to the "created_at" field.
func (au *AuthUpdate) AddCreatedAt(u int32) *AuthUpdate {
	au.mutation.AddCreatedAt(u)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AuthUpdate) SetUpdatedAt(u uint32) *AuthUpdate {
	au.mutation.ResetUpdatedAt()
	au.mutation.SetUpdatedAt(u)
	return au
}

// AddUpdatedAt adds u to the "updated_at" field.
func (au *AuthUpdate) AddUpdatedAt(u int32) *AuthUpdate {
	au.mutation.AddUpdatedAt(u)
	return au
}

// SetDeletedAt sets the "deleted_at" field.
func (au *AuthUpdate) SetDeletedAt(u uint32) *AuthUpdate {
	au.mutation.ResetDeletedAt()
	au.mutation.SetDeletedAt(u)
	return au
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (au *AuthUpdate) SetNillableDeletedAt(u *uint32) *AuthUpdate {
	if u != nil {
		au.SetDeletedAt(*u)
	}
	return au
}

// AddDeletedAt adds u to the "deleted_at" field.
func (au *AuthUpdate) AddDeletedAt(u int32) *AuthUpdate {
	au.mutation.AddDeletedAt(u)
	return au
}

// SetAppID sets the "app_id" field.
func (au *AuthUpdate) SetAppID(u uuid.UUID) *AuthUpdate {
	au.mutation.SetAppID(u)
	return au
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (au *AuthUpdate) SetNillableAppID(u *uuid.UUID) *AuthUpdate {
	if u != nil {
		au.SetAppID(*u)
	}
	return au
}

// ClearAppID clears the value of the "app_id" field.
func (au *AuthUpdate) ClearAppID() *AuthUpdate {
	au.mutation.ClearAppID()
	return au
}

// SetRoleID sets the "role_id" field.
func (au *AuthUpdate) SetRoleID(u uuid.UUID) *AuthUpdate {
	au.mutation.SetRoleID(u)
	return au
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (au *AuthUpdate) SetNillableRoleID(u *uuid.UUID) *AuthUpdate {
	if u != nil {
		au.SetRoleID(*u)
	}
	return au
}

// ClearRoleID clears the value of the "role_id" field.
func (au *AuthUpdate) ClearRoleID() *AuthUpdate {
	au.mutation.ClearRoleID()
	return au
}

// SetUserID sets the "user_id" field.
func (au *AuthUpdate) SetUserID(u uuid.UUID) *AuthUpdate {
	au.mutation.SetUserID(u)
	return au
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (au *AuthUpdate) SetNillableUserID(u *uuid.UUID) *AuthUpdate {
	if u != nil {
		au.SetUserID(*u)
	}
	return au
}

// ClearUserID clears the value of the "user_id" field.
func (au *AuthUpdate) ClearUserID() *AuthUpdate {
	au.mutation.ClearUserID()
	return au
}

// SetResource sets the "resource" field.
func (au *AuthUpdate) SetResource(s string) *AuthUpdate {
	au.mutation.SetResource(s)
	return au
}

// SetNillableResource sets the "resource" field if the given value is not nil.
func (au *AuthUpdate) SetNillableResource(s *string) *AuthUpdate {
	if s != nil {
		au.SetResource(*s)
	}
	return au
}

// ClearResource clears the value of the "resource" field.
func (au *AuthUpdate) ClearResource() *AuthUpdate {
	au.mutation.ClearResource()
	return au
}

// SetMethod sets the "method" field.
func (au *AuthUpdate) SetMethod(s string) *AuthUpdate {
	au.mutation.SetMethod(s)
	return au
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (au *AuthUpdate) SetNillableMethod(s *string) *AuthUpdate {
	if s != nil {
		au.SetMethod(*s)
	}
	return au
}

// ClearMethod clears the value of the "method" field.
func (au *AuthUpdate) ClearMethod() *AuthUpdate {
	au.mutation.ClearMethod()
	return au
}

// Mutation returns the AuthMutation object of the builder.
func (au *AuthUpdate) Mutation() *AuthMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AuthUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := au.defaults(); err != nil {
		return 0, err
	}
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AuthUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AuthUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AuthUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AuthUpdate) defaults() error {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		if auth.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized auth.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := auth.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (au *AuthUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AuthUpdate {
	au.modifiers = append(au.modifiers, modifiers...)
	return au
}

func (au *AuthUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   auth.Table,
			Columns: auth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: auth.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: auth.FieldCreatedAt,
		})
	}
	if value, ok := au.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: auth.FieldCreatedAt,
		})
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: auth.FieldUpdatedAt,
		})
	}
	if value, ok := au.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: auth.FieldUpdatedAt,
		})
	}
	if value, ok := au.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: auth.FieldDeletedAt,
		})
	}
	if value, ok := au.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: auth.FieldDeletedAt,
		})
	}
	if value, ok := au.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: auth.FieldAppID,
		})
	}
	if au.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: auth.FieldAppID,
		})
	}
	if value, ok := au.mutation.RoleID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: auth.FieldRoleID,
		})
	}
	if au.mutation.RoleIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: auth.FieldRoleID,
		})
	}
	if value, ok := au.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: auth.FieldUserID,
		})
	}
	if au.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: auth.FieldUserID,
		})
	}
	if value, ok := au.mutation.Resource(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: auth.FieldResource,
		})
	}
	if au.mutation.ResourceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: auth.FieldResource,
		})
	}
	if value, ok := au.mutation.Method(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: auth.FieldMethod,
		})
	}
	if au.mutation.MethodCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: auth.FieldMethod,
		})
	}
	_spec.Modifiers = au.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{auth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AuthUpdateOne is the builder for updating a single Auth entity.
type AuthUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AuthMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (auo *AuthUpdateOne) SetCreatedAt(u uint32) *AuthUpdateOne {
	auo.mutation.ResetCreatedAt()
	auo.mutation.SetCreatedAt(u)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AuthUpdateOne) SetNillableCreatedAt(u *uint32) *AuthUpdateOne {
	if u != nil {
		auo.SetCreatedAt(*u)
	}
	return auo
}

// AddCreatedAt adds u to the "created_at" field.
func (auo *AuthUpdateOne) AddCreatedAt(u int32) *AuthUpdateOne {
	auo.mutation.AddCreatedAt(u)
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AuthUpdateOne) SetUpdatedAt(u uint32) *AuthUpdateOne {
	auo.mutation.ResetUpdatedAt()
	auo.mutation.SetUpdatedAt(u)
	return auo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (auo *AuthUpdateOne) AddUpdatedAt(u int32) *AuthUpdateOne {
	auo.mutation.AddUpdatedAt(u)
	return auo
}

// SetDeletedAt sets the "deleted_at" field.
func (auo *AuthUpdateOne) SetDeletedAt(u uint32) *AuthUpdateOne {
	auo.mutation.ResetDeletedAt()
	auo.mutation.SetDeletedAt(u)
	return auo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (auo *AuthUpdateOne) SetNillableDeletedAt(u *uint32) *AuthUpdateOne {
	if u != nil {
		auo.SetDeletedAt(*u)
	}
	return auo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (auo *AuthUpdateOne) AddDeletedAt(u int32) *AuthUpdateOne {
	auo.mutation.AddDeletedAt(u)
	return auo
}

// SetAppID sets the "app_id" field.
func (auo *AuthUpdateOne) SetAppID(u uuid.UUID) *AuthUpdateOne {
	auo.mutation.SetAppID(u)
	return auo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (auo *AuthUpdateOne) SetNillableAppID(u *uuid.UUID) *AuthUpdateOne {
	if u != nil {
		auo.SetAppID(*u)
	}
	return auo
}

// ClearAppID clears the value of the "app_id" field.
func (auo *AuthUpdateOne) ClearAppID() *AuthUpdateOne {
	auo.mutation.ClearAppID()
	return auo
}

// SetRoleID sets the "role_id" field.
func (auo *AuthUpdateOne) SetRoleID(u uuid.UUID) *AuthUpdateOne {
	auo.mutation.SetRoleID(u)
	return auo
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (auo *AuthUpdateOne) SetNillableRoleID(u *uuid.UUID) *AuthUpdateOne {
	if u != nil {
		auo.SetRoleID(*u)
	}
	return auo
}

// ClearRoleID clears the value of the "role_id" field.
func (auo *AuthUpdateOne) ClearRoleID() *AuthUpdateOne {
	auo.mutation.ClearRoleID()
	return auo
}

// SetUserID sets the "user_id" field.
func (auo *AuthUpdateOne) SetUserID(u uuid.UUID) *AuthUpdateOne {
	auo.mutation.SetUserID(u)
	return auo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (auo *AuthUpdateOne) SetNillableUserID(u *uuid.UUID) *AuthUpdateOne {
	if u != nil {
		auo.SetUserID(*u)
	}
	return auo
}

// ClearUserID clears the value of the "user_id" field.
func (auo *AuthUpdateOne) ClearUserID() *AuthUpdateOne {
	auo.mutation.ClearUserID()
	return auo
}

// SetResource sets the "resource" field.
func (auo *AuthUpdateOne) SetResource(s string) *AuthUpdateOne {
	auo.mutation.SetResource(s)
	return auo
}

// SetNillableResource sets the "resource" field if the given value is not nil.
func (auo *AuthUpdateOne) SetNillableResource(s *string) *AuthUpdateOne {
	if s != nil {
		auo.SetResource(*s)
	}
	return auo
}

// ClearResource clears the value of the "resource" field.
func (auo *AuthUpdateOne) ClearResource() *AuthUpdateOne {
	auo.mutation.ClearResource()
	return auo
}

// SetMethod sets the "method" field.
func (auo *AuthUpdateOne) SetMethod(s string) *AuthUpdateOne {
	auo.mutation.SetMethod(s)
	return auo
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (auo *AuthUpdateOne) SetNillableMethod(s *string) *AuthUpdateOne {
	if s != nil {
		auo.SetMethod(*s)
	}
	return auo
}

// ClearMethod clears the value of the "method" field.
func (auo *AuthUpdateOne) ClearMethod() *AuthUpdateOne {
	auo.mutation.ClearMethod()
	return auo
}

// Mutation returns the AuthMutation object of the builder.
func (auo *AuthUpdateOne) Mutation() *AuthMutation {
	return auo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AuthUpdateOne) Select(field string, fields ...string) *AuthUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Auth entity.
func (auo *AuthUpdateOne) Save(ctx context.Context) (*Auth, error) {
	var (
		err  error
		node *Auth
	)
	if err := auo.defaults(); err != nil {
		return nil, err
	}
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, auo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Auth)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AuthMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AuthUpdateOne) SaveX(ctx context.Context) *Auth {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AuthUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AuthUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AuthUpdateOne) defaults() error {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		if auth.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized auth.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := auth.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auo *AuthUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AuthUpdateOne {
	auo.modifiers = append(auo.modifiers, modifiers...)
	return auo
}

func (auo *AuthUpdateOne) sqlSave(ctx context.Context) (_node *Auth, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   auth.Table,
			Columns: auth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: auth.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Auth.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, auth.FieldID)
		for _, f := range fields {
			if !auth.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != auth.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: auth.FieldCreatedAt,
		})
	}
	if value, ok := auo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: auth.FieldCreatedAt,
		})
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: auth.FieldUpdatedAt,
		})
	}
	if value, ok := auo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: auth.FieldUpdatedAt,
		})
	}
	if value, ok := auo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: auth.FieldDeletedAt,
		})
	}
	if value, ok := auo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: auth.FieldDeletedAt,
		})
	}
	if value, ok := auo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: auth.FieldAppID,
		})
	}
	if auo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: auth.FieldAppID,
		})
	}
	if value, ok := auo.mutation.RoleID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: auth.FieldRoleID,
		})
	}
	if auo.mutation.RoleIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: auth.FieldRoleID,
		})
	}
	if value, ok := auo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: auth.FieldUserID,
		})
	}
	if auo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: auth.FieldUserID,
		})
	}
	if value, ok := auo.mutation.Resource(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: auth.FieldResource,
		})
	}
	if auo.mutation.ResourceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: auth.FieldResource,
		})
	}
	if value, ok := auo.mutation.Method(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: auth.FieldMethod,
		})
	}
	if auo.mutation.MethodCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: auth.FieldMethod,
		})
	}
	_spec.Modifiers = auo.modifiers
	_node = &Auth{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{auth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
