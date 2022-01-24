// Code generated by entc, DO NOT EDIT.

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
	hooks    []Hook
	mutation *AppUserMutation
}

// Where appends a list predicates to the AppUserUpdate builder.
func (auu *AppUserUpdate) Where(ps ...predicate.AppUser) *AppUserUpdate {
	auu.mutation.Where(ps...)
	return auu
}

// SetAppID sets the "app_id" field.
func (auu *AppUserUpdate) SetAppID(u uuid.UUID) *AppUserUpdate {
	auu.mutation.SetAppID(u)
	return auu
}

// SetEmailAddress sets the "email_address" field.
func (auu *AppUserUpdate) SetEmailAddress(s string) *AppUserUpdate {
	auu.mutation.SetEmailAddress(s)
	return auu
}

// SetPhoneNo sets the "phone_no" field.
func (auu *AppUserUpdate) SetPhoneNo(s string) *AppUserUpdate {
	auu.mutation.SetPhoneNo(s)
	return auu
}

// SetImportFromApp sets the "import_from_app" field.
func (auu *AppUserUpdate) SetImportFromApp(u uuid.UUID) *AppUserUpdate {
	auu.mutation.SetImportFromApp(u)
	return auu
}

// SetCreateAt sets the "create_at" field.
func (auu *AppUserUpdate) SetCreateAt(u uint32) *AppUserUpdate {
	auu.mutation.ResetCreateAt()
	auu.mutation.SetCreateAt(u)
	return auu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (auu *AppUserUpdate) SetNillableCreateAt(u *uint32) *AppUserUpdate {
	if u != nil {
		auu.SetCreateAt(*u)
	}
	return auu
}

// AddCreateAt adds u to the "create_at" field.
func (auu *AppUserUpdate) AddCreateAt(u int32) *AppUserUpdate {
	auu.mutation.AddCreateAt(u)
	return auu
}

// SetUpdateAt sets the "update_at" field.
func (auu *AppUserUpdate) SetUpdateAt(u uint32) *AppUserUpdate {
	auu.mutation.ResetUpdateAt()
	auu.mutation.SetUpdateAt(u)
	return auu
}

// AddUpdateAt adds u to the "update_at" field.
func (auu *AppUserUpdate) AddUpdateAt(u int32) *AppUserUpdate {
	auu.mutation.AddUpdateAt(u)
	return auu
}

// SetDeleteAt sets the "delete_at" field.
func (auu *AppUserUpdate) SetDeleteAt(u uint32) *AppUserUpdate {
	auu.mutation.ResetDeleteAt()
	auu.mutation.SetDeleteAt(u)
	return auu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (auu *AppUserUpdate) SetNillableDeleteAt(u *uint32) *AppUserUpdate {
	if u != nil {
		auu.SetDeleteAt(*u)
	}
	return auu
}

// AddDeleteAt adds u to the "delete_at" field.
func (auu *AppUserUpdate) AddDeleteAt(u int32) *AppUserUpdate {
	auu.mutation.AddDeleteAt(u)
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
	auu.defaults()
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
func (auu *AppUserUpdate) defaults() {
	if _, ok := auu.mutation.UpdateAt(); !ok {
		v := appuser.UpdateDefaultUpdateAt()
		auu.mutation.SetUpdateAt(v)
	}
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
	if value, ok := auu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
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
	if value, ok := auu.mutation.PhoneNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
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
	if value, ok := auu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldCreateAt,
		})
	}
	if value, ok := auu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldCreateAt,
		})
	}
	if value, ok := auu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldUpdateAt,
		})
	}
	if value, ok := auu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldUpdateAt,
		})
	}
	if value, ok := auu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldDeleteAt,
		})
	}
	if value, ok := auu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, auu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AppUserUpdateOne is the builder for updating a single AppUser entity.
type AppUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppUserMutation
}

// SetAppID sets the "app_id" field.
func (auuo *AppUserUpdateOne) SetAppID(u uuid.UUID) *AppUserUpdateOne {
	auuo.mutation.SetAppID(u)
	return auuo
}

// SetEmailAddress sets the "email_address" field.
func (auuo *AppUserUpdateOne) SetEmailAddress(s string) *AppUserUpdateOne {
	auuo.mutation.SetEmailAddress(s)
	return auuo
}

// SetPhoneNo sets the "phone_no" field.
func (auuo *AppUserUpdateOne) SetPhoneNo(s string) *AppUserUpdateOne {
	auuo.mutation.SetPhoneNo(s)
	return auuo
}

// SetImportFromApp sets the "import_from_app" field.
func (auuo *AppUserUpdateOne) SetImportFromApp(u uuid.UUID) *AppUserUpdateOne {
	auuo.mutation.SetImportFromApp(u)
	return auuo
}

// SetCreateAt sets the "create_at" field.
func (auuo *AppUserUpdateOne) SetCreateAt(u uint32) *AppUserUpdateOne {
	auuo.mutation.ResetCreateAt()
	auuo.mutation.SetCreateAt(u)
	return auuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (auuo *AppUserUpdateOne) SetNillableCreateAt(u *uint32) *AppUserUpdateOne {
	if u != nil {
		auuo.SetCreateAt(*u)
	}
	return auuo
}

// AddCreateAt adds u to the "create_at" field.
func (auuo *AppUserUpdateOne) AddCreateAt(u int32) *AppUserUpdateOne {
	auuo.mutation.AddCreateAt(u)
	return auuo
}

// SetUpdateAt sets the "update_at" field.
func (auuo *AppUserUpdateOne) SetUpdateAt(u uint32) *AppUserUpdateOne {
	auuo.mutation.ResetUpdateAt()
	auuo.mutation.SetUpdateAt(u)
	return auuo
}

// AddUpdateAt adds u to the "update_at" field.
func (auuo *AppUserUpdateOne) AddUpdateAt(u int32) *AppUserUpdateOne {
	auuo.mutation.AddUpdateAt(u)
	return auuo
}

// SetDeleteAt sets the "delete_at" field.
func (auuo *AppUserUpdateOne) SetDeleteAt(u uint32) *AppUserUpdateOne {
	auuo.mutation.ResetDeleteAt()
	auuo.mutation.SetDeleteAt(u)
	return auuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (auuo *AppUserUpdateOne) SetNillableDeleteAt(u *uint32) *AppUserUpdateOne {
	if u != nil {
		auuo.SetDeleteAt(*u)
	}
	return auuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (auuo *AppUserUpdateOne) AddDeleteAt(u int32) *AppUserUpdateOne {
	auuo.mutation.AddDeleteAt(u)
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
	auuo.defaults()
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
		if _, err := mut.Mutate(ctx, auuo.mutation); err != nil {
			return nil, err
		}
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
func (auuo *AppUserUpdateOne) defaults() {
	if _, ok := auuo.mutation.UpdateAt(); !ok {
		v := appuser.UpdateDefaultUpdateAt()
		auuo.mutation.SetUpdateAt(v)
	}
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
	if value, ok := auuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
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
	if value, ok := auuo.mutation.PhoneNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
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
	if value, ok := auuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldCreateAt,
		})
	}
	if value, ok := auuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldCreateAt,
		})
	}
	if value, ok := auuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldUpdateAt,
		})
	}
	if value, ok := auuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldUpdateAt,
		})
	}
	if value, ok := auuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldDeleteAt,
		})
	}
	if value, ok := auuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuser.FieldDeleteAt,
		})
	}
	_node = &AppUser{config: auuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}