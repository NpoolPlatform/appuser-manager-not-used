// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppUserUpdate is the builder for updating AppUser entities.
type AppUserUpdate struct {
	config
	hooks     []Hook
	mutation  *AppUserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AppUserUpdate builder.
func (auu *AppUserUpdate) Where(ps ...predicate.AppUser) *AppUserUpdate {
	auu.mutation.Where(ps...)
	return auu
}

// SetCreatedAt sets the "created_at" field.
func (auu *AppUserUpdate) SetCreatedAt(u uint32) *AppUserUpdate {
	auu.mutation.ResetCreatedAt()
	auu.mutation.SetCreatedAt(u)
	return auu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auu *AppUserUpdate) SetNillableCreatedAt(u *uint32) *AppUserUpdate {
	if u != nil {
		auu.SetCreatedAt(*u)
	}
	return auu
}

// AddCreatedAt adds u to the "created_at" field.
func (auu *AppUserUpdate) AddCreatedAt(u int32) *AppUserUpdate {
	auu.mutation.AddCreatedAt(u)
	return auu
}

// SetUpdatedAt sets the "updated_at" field.
func (auu *AppUserUpdate) SetUpdatedAt(u uint32) *AppUserUpdate {
	auu.mutation.ResetUpdatedAt()
	auu.mutation.SetUpdatedAt(u)
	return auu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (auu *AppUserUpdate) AddUpdatedAt(u int32) *AppUserUpdate {
	auu.mutation.AddUpdatedAt(u)
	return auu
}

// SetDeletedAt sets the "deleted_at" field.
func (auu *AppUserUpdate) SetDeletedAt(u uint32) *AppUserUpdate {
	auu.mutation.ResetDeletedAt()
	auu.mutation.SetDeletedAt(u)
	return auu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (auu *AppUserUpdate) SetNillableDeletedAt(u *uint32) *AppUserUpdate {
	if u != nil {
		auu.SetDeletedAt(*u)
	}
	return auu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (auu *AppUserUpdate) AddDeletedAt(u int32) *AppUserUpdate {
	auu.mutation.AddDeletedAt(u)
	return auu
}

// SetAppID sets the "app_id" field.
func (auu *AppUserUpdate) SetAppID(u uuid.UUID) *AppUserUpdate {
	auu.mutation.SetAppID(u)
	return auu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (auu *AppUserUpdate) SetNillableAppID(u *uuid.UUID) *AppUserUpdate {
	if u != nil {
		auu.SetAppID(*u)
	}
	return auu
}

// ClearAppID clears the value of the "app_id" field.
func (auu *AppUserUpdate) ClearAppID() *AppUserUpdate {
	auu.mutation.ClearAppID()
	return auu
}

// SetEmailAddress sets the "email_address" field.
func (auu *AppUserUpdate) SetEmailAddress(s string) *AppUserUpdate {
	auu.mutation.SetEmailAddress(s)
	return auu
}

// SetNillableEmailAddress sets the "email_address" field if the given value is not nil.
func (auu *AppUserUpdate) SetNillableEmailAddress(s *string) *AppUserUpdate {
	if s != nil {
		auu.SetEmailAddress(*s)
	}
	return auu
}

// ClearEmailAddress clears the value of the "email_address" field.
func (auu *AppUserUpdate) ClearEmailAddress() *AppUserUpdate {
	auu.mutation.ClearEmailAddress()
	return auu
}

// SetPhoneNo sets the "phone_no" field.
func (auu *AppUserUpdate) SetPhoneNo(s string) *AppUserUpdate {
	auu.mutation.SetPhoneNo(s)
	return auu
}

// SetNillablePhoneNo sets the "phone_no" field if the given value is not nil.
func (auu *AppUserUpdate) SetNillablePhoneNo(s *string) *AppUserUpdate {
	if s != nil {
		auu.SetPhoneNo(*s)
	}
	return auu
}

// ClearPhoneNo clears the value of the "phone_no" field.
func (auu *AppUserUpdate) ClearPhoneNo() *AppUserUpdate {
	auu.mutation.ClearPhoneNo()
	return auu
}

// SetImportFromApp sets the "import_from_app" field.
func (auu *AppUserUpdate) SetImportFromApp(u uuid.UUID) *AppUserUpdate {
	auu.mutation.SetImportFromApp(u)
	return auu
}

// SetNillableImportFromApp sets the "import_from_app" field if the given value is not nil.
func (auu *AppUserUpdate) SetNillableImportFromApp(u *uuid.UUID) *AppUserUpdate {
	if u != nil {
		auu.SetImportFromApp(*u)
	}
	return auu
}

// ClearImportFromApp clears the value of the "import_from_app" field.
func (auu *AppUserUpdate) ClearImportFromApp() *AppUserUpdate {
	auu.mutation.ClearImportFromApp()
	return auu
}

// Mutation returns the AppUserMutation object of the builder.
func (auu *AppUserUpdate) Mutation() *AppUserMutation {
	return auu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (auu *AppUserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := auu.defaults(); err != nil {
		return 0, err
	}
	if len(auu.hooks) == 0 {
		affected, err = auu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auu.mutation = mutation
			affected, err = auu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(auu.hooks) - 1; i >= 0; i-- {
			if auu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (auu *AppUserUpdate) SaveX(ctx context.Context) int {
	affected, err := auu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (auu *AppUserUpdate) Exec(ctx context.Context) error {
	_, err := auu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auu *AppUserUpdate) ExecX(ctx context.Context) {
	if err := auu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auu *AppUserUpdate) defaults() error {
	if _, ok := auu.mutation.UpdatedAt(); !ok {
		if appuser.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appuser.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appuser.UpdateDefaultUpdatedAt()
		auu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auu *AppUserUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppUserUpdate {
	auu.modifiers = append(auu.modifiers, modifiers...)
	return auu
}

func (auu *AppUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appuser.Table,
			Columns: appuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuser.FieldID,
			},
		},
	}
	if ps := auu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldCreatedAt,
		})
	}
	if value, ok := auu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldCreatedAt,
		})
	}
	if value, ok := auu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldUpdatedAt,
		})
	}
	if value, ok := auu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldUpdatedAt,
		})
	}
	if value, ok := auu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldDeletedAt,
		})
	}
	if value, ok := auu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldDeletedAt,
		})
	}
	if value, ok := auu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuser.FieldAppID,
		})
	}
	if auu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appuser.FieldAppID,
		})
	}
	if value, ok := auu.mutation.EmailAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuser.FieldEmailAddress,
		})
	}
	if auu.mutation.EmailAddressCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appuser.FieldEmailAddress,
		})
	}
	if value, ok := auu.mutation.PhoneNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuser.FieldPhoneNo,
		})
	}
	if auu.mutation.PhoneNoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appuser.FieldPhoneNo,
		})
	}
	if value, ok := auu.mutation.ImportFromApp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuser.FieldImportFromApp,
		})
	}
	if auu.mutation.ImportFromAppCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appuser.FieldImportFromApp,
		})
	}
	_spec.Modifiers = auu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, auu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AppUserUpdateOne is the builder for updating a single AppUser entity.
type AppUserUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AppUserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (auuo *AppUserUpdateOne) SetCreatedAt(u uint32) *AppUserUpdateOne {
	auuo.mutation.ResetCreatedAt()
	auuo.mutation.SetCreatedAt(u)
	return auuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auuo *AppUserUpdateOne) SetNillableCreatedAt(u *uint32) *AppUserUpdateOne {
	if u != nil {
		auuo.SetCreatedAt(*u)
	}
	return auuo
}

// AddCreatedAt adds u to the "created_at" field.
func (auuo *AppUserUpdateOne) AddCreatedAt(u int32) *AppUserUpdateOne {
	auuo.mutation.AddCreatedAt(u)
	return auuo
}

// SetUpdatedAt sets the "updated_at" field.
func (auuo *AppUserUpdateOne) SetUpdatedAt(u uint32) *AppUserUpdateOne {
	auuo.mutation.ResetUpdatedAt()
	auuo.mutation.SetUpdatedAt(u)
	return auuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (auuo *AppUserUpdateOne) AddUpdatedAt(u int32) *AppUserUpdateOne {
	auuo.mutation.AddUpdatedAt(u)
	return auuo
}

// SetDeletedAt sets the "deleted_at" field.
func (auuo *AppUserUpdateOne) SetDeletedAt(u uint32) *AppUserUpdateOne {
	auuo.mutation.ResetDeletedAt()
	auuo.mutation.SetDeletedAt(u)
	return auuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (auuo *AppUserUpdateOne) SetNillableDeletedAt(u *uint32) *AppUserUpdateOne {
	if u != nil {
		auuo.SetDeletedAt(*u)
	}
	return auuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (auuo *AppUserUpdateOne) AddDeletedAt(u int32) *AppUserUpdateOne {
	auuo.mutation.AddDeletedAt(u)
	return auuo
}

// SetAppID sets the "app_id" field.
func (auuo *AppUserUpdateOne) SetAppID(u uuid.UUID) *AppUserUpdateOne {
	auuo.mutation.SetAppID(u)
	return auuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (auuo *AppUserUpdateOne) SetNillableAppID(u *uuid.UUID) *AppUserUpdateOne {
	if u != nil {
		auuo.SetAppID(*u)
	}
	return auuo
}

// ClearAppID clears the value of the "app_id" field.
func (auuo *AppUserUpdateOne) ClearAppID() *AppUserUpdateOne {
	auuo.mutation.ClearAppID()
	return auuo
}

// SetEmailAddress sets the "email_address" field.
func (auuo *AppUserUpdateOne) SetEmailAddress(s string) *AppUserUpdateOne {
	auuo.mutation.SetEmailAddress(s)
	return auuo
}

// SetNillableEmailAddress sets the "email_address" field if the given value is not nil.
func (auuo *AppUserUpdateOne) SetNillableEmailAddress(s *string) *AppUserUpdateOne {
	if s != nil {
		auuo.SetEmailAddress(*s)
	}
	return auuo
}

// ClearEmailAddress clears the value of the "email_address" field.
func (auuo *AppUserUpdateOne) ClearEmailAddress() *AppUserUpdateOne {
	auuo.mutation.ClearEmailAddress()
	return auuo
}

// SetPhoneNo sets the "phone_no" field.
func (auuo *AppUserUpdateOne) SetPhoneNo(s string) *AppUserUpdateOne {
	auuo.mutation.SetPhoneNo(s)
	return auuo
}

// SetNillablePhoneNo sets the "phone_no" field if the given value is not nil.
func (auuo *AppUserUpdateOne) SetNillablePhoneNo(s *string) *AppUserUpdateOne {
	if s != nil {
		auuo.SetPhoneNo(*s)
	}
	return auuo
}

// ClearPhoneNo clears the value of the "phone_no" field.
func (auuo *AppUserUpdateOne) ClearPhoneNo() *AppUserUpdateOne {
	auuo.mutation.ClearPhoneNo()
	return auuo
}

// SetImportFromApp sets the "import_from_app" field.
func (auuo *AppUserUpdateOne) SetImportFromApp(u uuid.UUID) *AppUserUpdateOne {
	auuo.mutation.SetImportFromApp(u)
	return auuo
}

// SetNillableImportFromApp sets the "import_from_app" field if the given value is not nil.
func (auuo *AppUserUpdateOne) SetNillableImportFromApp(u *uuid.UUID) *AppUserUpdateOne {
	if u != nil {
		auuo.SetImportFromApp(*u)
	}
	return auuo
}

// ClearImportFromApp clears the value of the "import_from_app" field.
func (auuo *AppUserUpdateOne) ClearImportFromApp() *AppUserUpdateOne {
	auuo.mutation.ClearImportFromApp()
	return auuo
}

// Mutation returns the AppUserMutation object of the builder.
func (auuo *AppUserUpdateOne) Mutation() *AppUserMutation {
	return auuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auuo *AppUserUpdateOne) Select(field string, fields ...string) *AppUserUpdateOne {
	auuo.fields = append([]string{field}, fields...)
	return auuo
}

// Save executes the query and returns the updated AppUser entity.
func (auuo *AppUserUpdateOne) Save(ctx context.Context) (*AppUser, error) {
	var (
		err  error
		node *AppUser
	)
	if err := auuo.defaults(); err != nil {
		return nil, err
	}
	if len(auuo.hooks) == 0 {
		node, err = auuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auuo.mutation = mutation
			node, err = auuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auuo.hooks) - 1; i >= 0; i-- {
			if auuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, auuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppUser)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppUserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auuo *AppUserUpdateOne) SaveX(ctx context.Context) *AppUser {
	node, err := auuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auuo *AppUserUpdateOne) Exec(ctx context.Context) error {
	_, err := auuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auuo *AppUserUpdateOne) ExecX(ctx context.Context) {
	if err := auuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auuo *AppUserUpdateOne) defaults() error {
	if _, ok := auuo.mutation.UpdatedAt(); !ok {
		if appuser.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appuser.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appuser.UpdateDefaultUpdatedAt()
		auuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auuo *AppUserUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppUserUpdateOne {
	auuo.modifiers = append(auuo.modifiers, modifiers...)
	return auuo
}

func (auuo *AppUserUpdateOne) sqlSave(ctx context.Context) (_node *AppUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appuser.Table,
			Columns: appuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuser.FieldID,
			},
		},
	}
	id, ok := auuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appuser.FieldID)
		for _, f := range fields {
			if !appuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldCreatedAt,
		})
	}
	if value, ok := auuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldCreatedAt,
		})
	}
	if value, ok := auuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldUpdatedAt,
		})
	}
	if value, ok := auuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldUpdatedAt,
		})
	}
	if value, ok := auuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldDeletedAt,
		})
	}
	if value, ok := auuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldDeletedAt,
		})
	}
	if value, ok := auuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuser.FieldAppID,
		})
	}
	if auuo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appuser.FieldAppID,
		})
	}
	if value, ok := auuo.mutation.EmailAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuser.FieldEmailAddress,
		})
	}
	if auuo.mutation.EmailAddressCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appuser.FieldEmailAddress,
		})
	}
	if value, ok := auuo.mutation.PhoneNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuser.FieldPhoneNo,
		})
	}
	if auuo.mutation.PhoneNoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appuser.FieldPhoneNo,
		})
	}
	if value, ok := auuo.mutation.ImportFromApp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuser.FieldImportFromApp,
		})
	}
	if auuo.mutation.ImportFromAppCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appuser.FieldImportFromApp,
		})
	}
	_spec.Modifiers = auuo.modifiers
	_node = &AppUser{config: auuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
