// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusercontrol"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppUserControlUpdate is the builder for updating AppUserControl entities.
type AppUserControlUpdate struct {
	config
	hooks     []Hook
	mutation  *AppUserControlMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AppUserControlUpdate builder.
func (aucu *AppUserControlUpdate) Where(ps ...predicate.AppUserControl) *AppUserControlUpdate {
	aucu.mutation.Where(ps...)
	return aucu
}

// SetCreatedAt sets the "created_at" field.
func (aucu *AppUserControlUpdate) SetCreatedAt(u uint32) *AppUserControlUpdate {
	aucu.mutation.ResetCreatedAt()
	aucu.mutation.SetCreatedAt(u)
	return aucu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aucu *AppUserControlUpdate) SetNillableCreatedAt(u *uint32) *AppUserControlUpdate {
	if u != nil {
		aucu.SetCreatedAt(*u)
	}
	return aucu
}

// AddCreatedAt adds u to the "created_at" field.
func (aucu *AppUserControlUpdate) AddCreatedAt(u int32) *AppUserControlUpdate {
	aucu.mutation.AddCreatedAt(u)
	return aucu
}

// SetUpdatedAt sets the "updated_at" field.
func (aucu *AppUserControlUpdate) SetUpdatedAt(u uint32) *AppUserControlUpdate {
	aucu.mutation.ResetUpdatedAt()
	aucu.mutation.SetUpdatedAt(u)
	return aucu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (aucu *AppUserControlUpdate) AddUpdatedAt(u int32) *AppUserControlUpdate {
	aucu.mutation.AddUpdatedAt(u)
	return aucu
}

// SetDeletedAt sets the "deleted_at" field.
func (aucu *AppUserControlUpdate) SetDeletedAt(u uint32) *AppUserControlUpdate {
	aucu.mutation.ResetDeletedAt()
	aucu.mutation.SetDeletedAt(u)
	return aucu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (aucu *AppUserControlUpdate) SetNillableDeletedAt(u *uint32) *AppUserControlUpdate {
	if u != nil {
		aucu.SetDeletedAt(*u)
	}
	return aucu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (aucu *AppUserControlUpdate) AddDeletedAt(u int32) *AppUserControlUpdate {
	aucu.mutation.AddDeletedAt(u)
	return aucu
}

// SetAppID sets the "app_id" field.
func (aucu *AppUserControlUpdate) SetAppID(u uuid.UUID) *AppUserControlUpdate {
	aucu.mutation.SetAppID(u)
	return aucu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (aucu *AppUserControlUpdate) SetNillableAppID(u *uuid.UUID) *AppUserControlUpdate {
	if u != nil {
		aucu.SetAppID(*u)
	}
	return aucu
}

// ClearAppID clears the value of the "app_id" field.
func (aucu *AppUserControlUpdate) ClearAppID() *AppUserControlUpdate {
	aucu.mutation.ClearAppID()
	return aucu
}

// SetUserID sets the "user_id" field.
func (aucu *AppUserControlUpdate) SetUserID(u uuid.UUID) *AppUserControlUpdate {
	aucu.mutation.SetUserID(u)
	return aucu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (aucu *AppUserControlUpdate) SetNillableUserID(u *uuid.UUID) *AppUserControlUpdate {
	if u != nil {
		aucu.SetUserID(*u)
	}
	return aucu
}

// ClearUserID clears the value of the "user_id" field.
func (aucu *AppUserControlUpdate) ClearUserID() *AppUserControlUpdate {
	aucu.mutation.ClearUserID()
	return aucu
}

// SetSigninVerifyByGoogleAuthentication sets the "signin_verify_by_google_authentication" field.
func (aucu *AppUserControlUpdate) SetSigninVerifyByGoogleAuthentication(b bool) *AppUserControlUpdate {
	aucu.mutation.SetSigninVerifyByGoogleAuthentication(b)
	return aucu
}

// SetNillableSigninVerifyByGoogleAuthentication sets the "signin_verify_by_google_authentication" field if the given value is not nil.
func (aucu *AppUserControlUpdate) SetNillableSigninVerifyByGoogleAuthentication(b *bool) *AppUserControlUpdate {
	if b != nil {
		aucu.SetSigninVerifyByGoogleAuthentication(*b)
	}
	return aucu
}

// ClearSigninVerifyByGoogleAuthentication clears the value of the "signin_verify_by_google_authentication" field.
func (aucu *AppUserControlUpdate) ClearSigninVerifyByGoogleAuthentication() *AppUserControlUpdate {
	aucu.mutation.ClearSigninVerifyByGoogleAuthentication()
	return aucu
}

// SetGoogleAuthenticationVerified sets the "google_authentication_verified" field.
func (aucu *AppUserControlUpdate) SetGoogleAuthenticationVerified(b bool) *AppUserControlUpdate {
	aucu.mutation.SetGoogleAuthenticationVerified(b)
	return aucu
}

// SetNillableGoogleAuthenticationVerified sets the "google_authentication_verified" field if the given value is not nil.
func (aucu *AppUserControlUpdate) SetNillableGoogleAuthenticationVerified(b *bool) *AppUserControlUpdate {
	if b != nil {
		aucu.SetGoogleAuthenticationVerified(*b)
	}
	return aucu
}

// ClearGoogleAuthenticationVerified clears the value of the "google_authentication_verified" field.
func (aucu *AppUserControlUpdate) ClearGoogleAuthenticationVerified() *AppUserControlUpdate {
	aucu.mutation.ClearGoogleAuthenticationVerified()
	return aucu
}

// Mutation returns the AppUserControlMutation object of the builder.
func (aucu *AppUserControlUpdate) Mutation() *AppUserControlMutation {
	return aucu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aucu *AppUserControlUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := aucu.defaults(); err != nil {
		return 0, err
	}
	if len(aucu.hooks) == 0 {
		affected, err = aucu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserControlMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aucu.mutation = mutation
			affected, err = aucu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aucu.hooks) - 1; i >= 0; i-- {
			if aucu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aucu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aucu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (aucu *AppUserControlUpdate) SaveX(ctx context.Context) int {
	affected, err := aucu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aucu *AppUserControlUpdate) Exec(ctx context.Context) error {
	_, err := aucu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aucu *AppUserControlUpdate) ExecX(ctx context.Context) {
	if err := aucu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aucu *AppUserControlUpdate) defaults() error {
	if _, ok := aucu.mutation.UpdatedAt(); !ok {
		if appusercontrol.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appusercontrol.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appusercontrol.UpdateDefaultUpdatedAt()
		aucu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (aucu *AppUserControlUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppUserControlUpdate {
	aucu.modifiers = append(aucu.modifiers, modifiers...)
	return aucu
}

func (aucu *AppUserControlUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appusercontrol.Table,
			Columns: appusercontrol.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appusercontrol.FieldID,
			},
		},
	}
	if ps := aucu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aucu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusercontrol.FieldCreatedAt,
		})
	}
	if value, ok := aucu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusercontrol.FieldCreatedAt,
		})
	}
	if value, ok := aucu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusercontrol.FieldUpdatedAt,
		})
	}
	if value, ok := aucu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusercontrol.FieldUpdatedAt,
		})
	}
	if value, ok := aucu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusercontrol.FieldDeletedAt,
		})
	}
	if value, ok := aucu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusercontrol.FieldDeletedAt,
		})
	}
	if value, ok := aucu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appusercontrol.FieldAppID,
		})
	}
	if aucu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appusercontrol.FieldAppID,
		})
	}
	if value, ok := aucu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appusercontrol.FieldUserID,
		})
	}
	if aucu.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appusercontrol.FieldUserID,
		})
	}
	if value, ok := aucu.mutation.SigninVerifyByGoogleAuthentication(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: appusercontrol.FieldSigninVerifyByGoogleAuthentication,
		})
	}
	if aucu.mutation.SigninVerifyByGoogleAuthenticationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: appusercontrol.FieldSigninVerifyByGoogleAuthentication,
		})
	}
	if value, ok := aucu.mutation.GoogleAuthenticationVerified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: appusercontrol.FieldGoogleAuthenticationVerified,
		})
	}
	if aucu.mutation.GoogleAuthenticationVerifiedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: appusercontrol.FieldGoogleAuthenticationVerified,
		})
	}
	_spec.Modifiers = aucu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, aucu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appusercontrol.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AppUserControlUpdateOne is the builder for updating a single AppUserControl entity.
type AppUserControlUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AppUserControlMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (aucuo *AppUserControlUpdateOne) SetCreatedAt(u uint32) *AppUserControlUpdateOne {
	aucuo.mutation.ResetCreatedAt()
	aucuo.mutation.SetCreatedAt(u)
	return aucuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aucuo *AppUserControlUpdateOne) SetNillableCreatedAt(u *uint32) *AppUserControlUpdateOne {
	if u != nil {
		aucuo.SetCreatedAt(*u)
	}
	return aucuo
}

// AddCreatedAt adds u to the "created_at" field.
func (aucuo *AppUserControlUpdateOne) AddCreatedAt(u int32) *AppUserControlUpdateOne {
	aucuo.mutation.AddCreatedAt(u)
	return aucuo
}

// SetUpdatedAt sets the "updated_at" field.
func (aucuo *AppUserControlUpdateOne) SetUpdatedAt(u uint32) *AppUserControlUpdateOne {
	aucuo.mutation.ResetUpdatedAt()
	aucuo.mutation.SetUpdatedAt(u)
	return aucuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (aucuo *AppUserControlUpdateOne) AddUpdatedAt(u int32) *AppUserControlUpdateOne {
	aucuo.mutation.AddUpdatedAt(u)
	return aucuo
}

// SetDeletedAt sets the "deleted_at" field.
func (aucuo *AppUserControlUpdateOne) SetDeletedAt(u uint32) *AppUserControlUpdateOne {
	aucuo.mutation.ResetDeletedAt()
	aucuo.mutation.SetDeletedAt(u)
	return aucuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (aucuo *AppUserControlUpdateOne) SetNillableDeletedAt(u *uint32) *AppUserControlUpdateOne {
	if u != nil {
		aucuo.SetDeletedAt(*u)
	}
	return aucuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (aucuo *AppUserControlUpdateOne) AddDeletedAt(u int32) *AppUserControlUpdateOne {
	aucuo.mutation.AddDeletedAt(u)
	return aucuo
}

// SetAppID sets the "app_id" field.
func (aucuo *AppUserControlUpdateOne) SetAppID(u uuid.UUID) *AppUserControlUpdateOne {
	aucuo.mutation.SetAppID(u)
	return aucuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (aucuo *AppUserControlUpdateOne) SetNillableAppID(u *uuid.UUID) *AppUserControlUpdateOne {
	if u != nil {
		aucuo.SetAppID(*u)
	}
	return aucuo
}

// ClearAppID clears the value of the "app_id" field.
func (aucuo *AppUserControlUpdateOne) ClearAppID() *AppUserControlUpdateOne {
	aucuo.mutation.ClearAppID()
	return aucuo
}

// SetUserID sets the "user_id" field.
func (aucuo *AppUserControlUpdateOne) SetUserID(u uuid.UUID) *AppUserControlUpdateOne {
	aucuo.mutation.SetUserID(u)
	return aucuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (aucuo *AppUserControlUpdateOne) SetNillableUserID(u *uuid.UUID) *AppUserControlUpdateOne {
	if u != nil {
		aucuo.SetUserID(*u)
	}
	return aucuo
}

// ClearUserID clears the value of the "user_id" field.
func (aucuo *AppUserControlUpdateOne) ClearUserID() *AppUserControlUpdateOne {
	aucuo.mutation.ClearUserID()
	return aucuo
}

// SetSigninVerifyByGoogleAuthentication sets the "signin_verify_by_google_authentication" field.
func (aucuo *AppUserControlUpdateOne) SetSigninVerifyByGoogleAuthentication(b bool) *AppUserControlUpdateOne {
	aucuo.mutation.SetSigninVerifyByGoogleAuthentication(b)
	return aucuo
}

// SetNillableSigninVerifyByGoogleAuthentication sets the "signin_verify_by_google_authentication" field if the given value is not nil.
func (aucuo *AppUserControlUpdateOne) SetNillableSigninVerifyByGoogleAuthentication(b *bool) *AppUserControlUpdateOne {
	if b != nil {
		aucuo.SetSigninVerifyByGoogleAuthentication(*b)
	}
	return aucuo
}

// ClearSigninVerifyByGoogleAuthentication clears the value of the "signin_verify_by_google_authentication" field.
func (aucuo *AppUserControlUpdateOne) ClearSigninVerifyByGoogleAuthentication() *AppUserControlUpdateOne {
	aucuo.mutation.ClearSigninVerifyByGoogleAuthentication()
	return aucuo
}

// SetGoogleAuthenticationVerified sets the "google_authentication_verified" field.
func (aucuo *AppUserControlUpdateOne) SetGoogleAuthenticationVerified(b bool) *AppUserControlUpdateOne {
	aucuo.mutation.SetGoogleAuthenticationVerified(b)
	return aucuo
}

// SetNillableGoogleAuthenticationVerified sets the "google_authentication_verified" field if the given value is not nil.
func (aucuo *AppUserControlUpdateOne) SetNillableGoogleAuthenticationVerified(b *bool) *AppUserControlUpdateOne {
	if b != nil {
		aucuo.SetGoogleAuthenticationVerified(*b)
	}
	return aucuo
}

// ClearGoogleAuthenticationVerified clears the value of the "google_authentication_verified" field.
func (aucuo *AppUserControlUpdateOne) ClearGoogleAuthenticationVerified() *AppUserControlUpdateOne {
	aucuo.mutation.ClearGoogleAuthenticationVerified()
	return aucuo
}

// Mutation returns the AppUserControlMutation object of the builder.
func (aucuo *AppUserControlUpdateOne) Mutation() *AppUserControlMutation {
	return aucuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aucuo *AppUserControlUpdateOne) Select(field string, fields ...string) *AppUserControlUpdateOne {
	aucuo.fields = append([]string{field}, fields...)
	return aucuo
}

// Save executes the query and returns the updated AppUserControl entity.
func (aucuo *AppUserControlUpdateOne) Save(ctx context.Context) (*AppUserControl, error) {
	var (
		err  error
		node *AppUserControl
	)
	if err := aucuo.defaults(); err != nil {
		return nil, err
	}
	if len(aucuo.hooks) == 0 {
		node, err = aucuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserControlMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aucuo.mutation = mutation
			node, err = aucuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aucuo.hooks) - 1; i >= 0; i-- {
			if aucuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aucuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, aucuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppUserControl)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppUserControlMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aucuo *AppUserControlUpdateOne) SaveX(ctx context.Context) *AppUserControl {
	node, err := aucuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aucuo *AppUserControlUpdateOne) Exec(ctx context.Context) error {
	_, err := aucuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aucuo *AppUserControlUpdateOne) ExecX(ctx context.Context) {
	if err := aucuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aucuo *AppUserControlUpdateOne) defaults() error {
	if _, ok := aucuo.mutation.UpdatedAt(); !ok {
		if appusercontrol.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appusercontrol.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appusercontrol.UpdateDefaultUpdatedAt()
		aucuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (aucuo *AppUserControlUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppUserControlUpdateOne {
	aucuo.modifiers = append(aucuo.modifiers, modifiers...)
	return aucuo
}

func (aucuo *AppUserControlUpdateOne) sqlSave(ctx context.Context) (_node *AppUserControl, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appusercontrol.Table,
			Columns: appusercontrol.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appusercontrol.FieldID,
			},
		},
	}
	id, ok := aucuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppUserControl.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aucuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appusercontrol.FieldID)
		for _, f := range fields {
			if !appusercontrol.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appusercontrol.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aucuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aucuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusercontrol.FieldCreatedAt,
		})
	}
	if value, ok := aucuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusercontrol.FieldCreatedAt,
		})
	}
	if value, ok := aucuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusercontrol.FieldUpdatedAt,
		})
	}
	if value, ok := aucuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusercontrol.FieldUpdatedAt,
		})
	}
	if value, ok := aucuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusercontrol.FieldDeletedAt,
		})
	}
	if value, ok := aucuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appusercontrol.FieldDeletedAt,
		})
	}
	if value, ok := aucuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appusercontrol.FieldAppID,
		})
	}
	if aucuo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appusercontrol.FieldAppID,
		})
	}
	if value, ok := aucuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appusercontrol.FieldUserID,
		})
	}
	if aucuo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appusercontrol.FieldUserID,
		})
	}
	if value, ok := aucuo.mutation.SigninVerifyByGoogleAuthentication(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: appusercontrol.FieldSigninVerifyByGoogleAuthentication,
		})
	}
	if aucuo.mutation.SigninVerifyByGoogleAuthenticationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: appusercontrol.FieldSigninVerifyByGoogleAuthentication,
		})
	}
	if value, ok := aucuo.mutation.GoogleAuthenticationVerified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: appusercontrol.FieldGoogleAuthenticationVerified,
		})
	}
	if aucuo.mutation.GoogleAuthenticationVerifiedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: appusercontrol.FieldGoogleAuthenticationVerified,
		})
	}
	_spec.Modifiers = aucuo.modifiers
	_node = &AppUserControl{config: aucuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aucuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appusercontrol.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
