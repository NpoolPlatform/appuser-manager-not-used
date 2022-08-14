// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approleuser"
	"github.com/google/uuid"
)

// AppRoleUserCreate is the builder for creating a AppRoleUser entity.
type AppRoleUserCreate struct {
	config
	mutation *AppRoleUserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (aruc *AppRoleUserCreate) SetCreatedAt(u uint32) *AppRoleUserCreate {
	aruc.mutation.SetCreatedAt(u)
	return aruc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aruc *AppRoleUserCreate) SetNillableCreatedAt(u *uint32) *AppRoleUserCreate {
	if u != nil {
		aruc.SetCreatedAt(*u)
	}
	return aruc
}

// SetUpdatedAt sets the "updated_at" field.
func (aruc *AppRoleUserCreate) SetUpdatedAt(u uint32) *AppRoleUserCreate {
	aruc.mutation.SetUpdatedAt(u)
	return aruc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (aruc *AppRoleUserCreate) SetNillableUpdatedAt(u *uint32) *AppRoleUserCreate {
	if u != nil {
		aruc.SetUpdatedAt(*u)
	}
	return aruc
}

// SetDeletedAt sets the "deleted_at" field.
func (aruc *AppRoleUserCreate) SetDeletedAt(u uint32) *AppRoleUserCreate {
	aruc.mutation.SetDeletedAt(u)
	return aruc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (aruc *AppRoleUserCreate) SetNillableDeletedAt(u *uint32) *AppRoleUserCreate {
	if u != nil {
		aruc.SetDeletedAt(*u)
	}
	return aruc
}

// SetAppID sets the "app_id" field.
func (aruc *AppRoleUserCreate) SetAppID(u uuid.UUID) *AppRoleUserCreate {
	aruc.mutation.SetAppID(u)
	return aruc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (aruc *AppRoleUserCreate) SetNillableAppID(u *uuid.UUID) *AppRoleUserCreate {
	if u != nil {
		aruc.SetAppID(*u)
	}
	return aruc
}

// SetRoleID sets the "role_id" field.
func (aruc *AppRoleUserCreate) SetRoleID(u uuid.UUID) *AppRoleUserCreate {
	aruc.mutation.SetRoleID(u)
	return aruc
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (aruc *AppRoleUserCreate) SetNillableRoleID(u *uuid.UUID) *AppRoleUserCreate {
	if u != nil {
		aruc.SetRoleID(*u)
	}
	return aruc
}

// SetUserID sets the "user_id" field.
func (aruc *AppRoleUserCreate) SetUserID(u uuid.UUID) *AppRoleUserCreate {
	aruc.mutation.SetUserID(u)
	return aruc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (aruc *AppRoleUserCreate) SetNillableUserID(u *uuid.UUID) *AppRoleUserCreate {
	if u != nil {
		aruc.SetUserID(*u)
	}
	return aruc
}

// SetID sets the "id" field.
func (aruc *AppRoleUserCreate) SetID(u uuid.UUID) *AppRoleUserCreate {
	aruc.mutation.SetID(u)
	return aruc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (aruc *AppRoleUserCreate) SetNillableID(u *uuid.UUID) *AppRoleUserCreate {
	if u != nil {
		aruc.SetID(*u)
	}
	return aruc
}

// Mutation returns the AppRoleUserMutation object of the builder.
func (aruc *AppRoleUserCreate) Mutation() *AppRoleUserMutation {
	return aruc.mutation
}

// Save creates the AppRoleUser in the database.
func (aruc *AppRoleUserCreate) Save(ctx context.Context) (*AppRoleUser, error) {
	var (
		err  error
		node *AppRoleUser
	)
	if err := aruc.defaults(); err != nil {
		return nil, err
	}
	if len(aruc.hooks) == 0 {
		if err = aruc.check(); err != nil {
			return nil, err
		}
		node, err = aruc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppRoleUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = aruc.check(); err != nil {
				return nil, err
			}
			aruc.mutation = mutation
			if node, err = aruc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(aruc.hooks) - 1; i >= 0; i-- {
			if aruc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aruc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, aruc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppRoleUser)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppRoleUserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (aruc *AppRoleUserCreate) SaveX(ctx context.Context) *AppRoleUser {
	v, err := aruc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aruc *AppRoleUserCreate) Exec(ctx context.Context) error {
	_, err := aruc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aruc *AppRoleUserCreate) ExecX(ctx context.Context) {
	if err := aruc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aruc *AppRoleUserCreate) defaults() error {
	if _, ok := aruc.mutation.CreatedAt(); !ok {
		if approleuser.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized approleuser.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := approleuser.DefaultCreatedAt()
		aruc.mutation.SetCreatedAt(v)
	}
	if _, ok := aruc.mutation.UpdatedAt(); !ok {
		if approleuser.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized approleuser.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := approleuser.DefaultUpdatedAt()
		aruc.mutation.SetUpdatedAt(v)
	}
	if _, ok := aruc.mutation.DeletedAt(); !ok {
		if approleuser.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized approleuser.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := approleuser.DefaultDeletedAt()
		aruc.mutation.SetDeletedAt(v)
	}
	if _, ok := aruc.mutation.AppID(); !ok {
		if approleuser.DefaultAppID == nil {
			return fmt.Errorf("ent: uninitialized approleuser.DefaultAppID (forgotten import ent/runtime?)")
		}
		v := approleuser.DefaultAppID()
		aruc.mutation.SetAppID(v)
	}
	if _, ok := aruc.mutation.RoleID(); !ok {
		if approleuser.DefaultRoleID == nil {
			return fmt.Errorf("ent: uninitialized approleuser.DefaultRoleID (forgotten import ent/runtime?)")
		}
		v := approleuser.DefaultRoleID()
		aruc.mutation.SetRoleID(v)
	}
	if _, ok := aruc.mutation.UserID(); !ok {
		if approleuser.DefaultUserID == nil {
			return fmt.Errorf("ent: uninitialized approleuser.DefaultUserID (forgotten import ent/runtime?)")
		}
		v := approleuser.DefaultUserID()
		aruc.mutation.SetUserID(v)
	}
	if _, ok := aruc.mutation.ID(); !ok {
		if approleuser.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized approleuser.DefaultID (forgotten import ent/runtime?)")
		}
		v := approleuser.DefaultID()
		aruc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (aruc *AppRoleUserCreate) check() error {
	if _, ok := aruc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppRoleUser.created_at"`)}
	}
	if _, ok := aruc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AppRoleUser.updated_at"`)}
	}
	if _, ok := aruc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "AppRoleUser.deleted_at"`)}
	}
	return nil
}

func (aruc *AppRoleUserCreate) sqlSave(ctx context.Context) (*AppRoleUser, error) {
	_node, _spec := aruc.createSpec()
	if err := sqlgraph.CreateNode(ctx, aruc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (aruc *AppRoleUserCreate) createSpec() (*AppRoleUser, *sqlgraph.CreateSpec) {
	var (
		_node = &AppRoleUser{config: aruc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: approleuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: approleuser.FieldID,
			},
		}
	)
	_spec.OnConflict = aruc.conflict
	if id, ok := aruc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := aruc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: approleuser.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := aruc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: approleuser.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := aruc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: approleuser.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := aruc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: approleuser.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := aruc.mutation.RoleID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: approleuser.FieldRoleID,
		})
		_node.RoleID = value
	}
	if value, ok := aruc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: approleuser.FieldUserID,
		})
		_node.UserID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppRoleUser.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppRoleUserUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (aruc *AppRoleUserCreate) OnConflict(opts ...sql.ConflictOption) *AppRoleUserUpsertOne {
	aruc.conflict = opts
	return &AppRoleUserUpsertOne{
		create: aruc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppRoleUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (aruc *AppRoleUserCreate) OnConflictColumns(columns ...string) *AppRoleUserUpsertOne {
	aruc.conflict = append(aruc.conflict, sql.ConflictColumns(columns...))
	return &AppRoleUserUpsertOne{
		create: aruc,
	}
}

type (
	// AppRoleUserUpsertOne is the builder for "upsert"-ing
	//  one AppRoleUser node.
	AppRoleUserUpsertOne struct {
		create *AppRoleUserCreate
	}

	// AppRoleUserUpsert is the "OnConflict" setter.
	AppRoleUserUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *AppRoleUserUpsert) SetCreatedAt(v uint32) *AppRoleUserUpsert {
	u.Set(approleuser.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppRoleUserUpsert) UpdateCreatedAt() *AppRoleUserUpsert {
	u.SetExcluded(approleuser.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppRoleUserUpsert) AddCreatedAt(v uint32) *AppRoleUserUpsert {
	u.Add(approleuser.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppRoleUserUpsert) SetUpdatedAt(v uint32) *AppRoleUserUpsert {
	u.Set(approleuser.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppRoleUserUpsert) UpdateUpdatedAt() *AppRoleUserUpsert {
	u.SetExcluded(approleuser.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppRoleUserUpsert) AddUpdatedAt(v uint32) *AppRoleUserUpsert {
	u.Add(approleuser.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppRoleUserUpsert) SetDeletedAt(v uint32) *AppRoleUserUpsert {
	u.Set(approleuser.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppRoleUserUpsert) UpdateDeletedAt() *AppRoleUserUpsert {
	u.SetExcluded(approleuser.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppRoleUserUpsert) AddDeletedAt(v uint32) *AppRoleUserUpsert {
	u.Add(approleuser.FieldDeletedAt, v)
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppRoleUserUpsert) SetAppID(v uuid.UUID) *AppRoleUserUpsert {
	u.Set(approleuser.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppRoleUserUpsert) UpdateAppID() *AppRoleUserUpsert {
	u.SetExcluded(approleuser.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppRoleUserUpsert) ClearAppID() *AppRoleUserUpsert {
	u.SetNull(approleuser.FieldAppID)
	return u
}

// SetRoleID sets the "role_id" field.
func (u *AppRoleUserUpsert) SetRoleID(v uuid.UUID) *AppRoleUserUpsert {
	u.Set(approleuser.FieldRoleID, v)
	return u
}

// UpdateRoleID sets the "role_id" field to the value that was provided on create.
func (u *AppRoleUserUpsert) UpdateRoleID() *AppRoleUserUpsert {
	u.SetExcluded(approleuser.FieldRoleID)
	return u
}

// ClearRoleID clears the value of the "role_id" field.
func (u *AppRoleUserUpsert) ClearRoleID() *AppRoleUserUpsert {
	u.SetNull(approleuser.FieldRoleID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *AppRoleUserUpsert) SetUserID(v uuid.UUID) *AppRoleUserUpsert {
	u.Set(approleuser.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppRoleUserUpsert) UpdateUserID() *AppRoleUserUpsert {
	u.SetExcluded(approleuser.FieldUserID)
	return u
}

// ClearUserID clears the value of the "user_id" field.
func (u *AppRoleUserUpsert) ClearUserID() *AppRoleUserUpsert {
	u.SetNull(approleuser.FieldUserID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppRoleUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(approleuser.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppRoleUserUpsertOne) UpdateNewValues() *AppRoleUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(approleuser.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppRoleUser.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppRoleUserUpsertOne) Ignore() *AppRoleUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppRoleUserUpsertOne) DoNothing() *AppRoleUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppRoleUserCreate.OnConflict
// documentation for more info.
func (u *AppRoleUserUpsertOne) Update(set func(*AppRoleUserUpsert)) *AppRoleUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppRoleUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppRoleUserUpsertOne) SetCreatedAt(v uint32) *AppRoleUserUpsertOne {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppRoleUserUpsertOne) AddCreatedAt(v uint32) *AppRoleUserUpsertOne {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppRoleUserUpsertOne) UpdateCreatedAt() *AppRoleUserUpsertOne {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppRoleUserUpsertOne) SetUpdatedAt(v uint32) *AppRoleUserUpsertOne {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppRoleUserUpsertOne) AddUpdatedAt(v uint32) *AppRoleUserUpsertOne {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppRoleUserUpsertOne) UpdateUpdatedAt() *AppRoleUserUpsertOne {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppRoleUserUpsertOne) SetDeletedAt(v uint32) *AppRoleUserUpsertOne {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppRoleUserUpsertOne) AddDeletedAt(v uint32) *AppRoleUserUpsertOne {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppRoleUserUpsertOne) UpdateDeletedAt() *AppRoleUserUpsertOne {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppRoleUserUpsertOne) SetAppID(v uuid.UUID) *AppRoleUserUpsertOne {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppRoleUserUpsertOne) UpdateAppID() *AppRoleUserUpsertOne {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppRoleUserUpsertOne) ClearAppID() *AppRoleUserUpsertOne {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.ClearAppID()
	})
}

// SetRoleID sets the "role_id" field.
func (u *AppRoleUserUpsertOne) SetRoleID(v uuid.UUID) *AppRoleUserUpsertOne {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.SetRoleID(v)
	})
}

// UpdateRoleID sets the "role_id" field to the value that was provided on create.
func (u *AppRoleUserUpsertOne) UpdateRoleID() *AppRoleUserUpsertOne {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.UpdateRoleID()
	})
}

// ClearRoleID clears the value of the "role_id" field.
func (u *AppRoleUserUpsertOne) ClearRoleID() *AppRoleUserUpsertOne {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.ClearRoleID()
	})
}

// SetUserID sets the "user_id" field.
func (u *AppRoleUserUpsertOne) SetUserID(v uuid.UUID) *AppRoleUserUpsertOne {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppRoleUserUpsertOne) UpdateUserID() *AppRoleUserUpsertOne {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *AppRoleUserUpsertOne) ClearUserID() *AppRoleUserUpsertOne {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.ClearUserID()
	})
}

// Exec executes the query.
func (u *AppRoleUserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppRoleUserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppRoleUserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppRoleUserUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AppRoleUserUpsertOne.ID is not supported by MySQL driver. Use AppRoleUserUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppRoleUserUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppRoleUserCreateBulk is the builder for creating many AppRoleUser entities in bulk.
type AppRoleUserCreateBulk struct {
	config
	builders []*AppRoleUserCreate
	conflict []sql.ConflictOption
}

// Save creates the AppRoleUser entities in the database.
func (arucb *AppRoleUserCreateBulk) Save(ctx context.Context) ([]*AppRoleUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(arucb.builders))
	nodes := make([]*AppRoleUser, len(arucb.builders))
	mutators := make([]Mutator, len(arucb.builders))
	for i := range arucb.builders {
		func(i int, root context.Context) {
			builder := arucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppRoleUserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, arucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = arucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, arucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, arucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (arucb *AppRoleUserCreateBulk) SaveX(ctx context.Context) []*AppRoleUser {
	v, err := arucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arucb *AppRoleUserCreateBulk) Exec(ctx context.Context) error {
	_, err := arucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arucb *AppRoleUserCreateBulk) ExecX(ctx context.Context) {
	if err := arucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppRoleUser.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppRoleUserUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (arucb *AppRoleUserCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppRoleUserUpsertBulk {
	arucb.conflict = opts
	return &AppRoleUserUpsertBulk{
		create: arucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppRoleUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (arucb *AppRoleUserCreateBulk) OnConflictColumns(columns ...string) *AppRoleUserUpsertBulk {
	arucb.conflict = append(arucb.conflict, sql.ConflictColumns(columns...))
	return &AppRoleUserUpsertBulk{
		create: arucb,
	}
}

// AppRoleUserUpsertBulk is the builder for "upsert"-ing
// a bulk of AppRoleUser nodes.
type AppRoleUserUpsertBulk struct {
	create *AppRoleUserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppRoleUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(approleuser.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppRoleUserUpsertBulk) UpdateNewValues() *AppRoleUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(approleuser.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppRoleUser.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppRoleUserUpsertBulk) Ignore() *AppRoleUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppRoleUserUpsertBulk) DoNothing() *AppRoleUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppRoleUserCreateBulk.OnConflict
// documentation for more info.
func (u *AppRoleUserUpsertBulk) Update(set func(*AppRoleUserUpsert)) *AppRoleUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppRoleUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppRoleUserUpsertBulk) SetCreatedAt(v uint32) *AppRoleUserUpsertBulk {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppRoleUserUpsertBulk) AddCreatedAt(v uint32) *AppRoleUserUpsertBulk {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppRoleUserUpsertBulk) UpdateCreatedAt() *AppRoleUserUpsertBulk {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppRoleUserUpsertBulk) SetUpdatedAt(v uint32) *AppRoleUserUpsertBulk {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppRoleUserUpsertBulk) AddUpdatedAt(v uint32) *AppRoleUserUpsertBulk {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppRoleUserUpsertBulk) UpdateUpdatedAt() *AppRoleUserUpsertBulk {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppRoleUserUpsertBulk) SetDeletedAt(v uint32) *AppRoleUserUpsertBulk {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppRoleUserUpsertBulk) AddDeletedAt(v uint32) *AppRoleUserUpsertBulk {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppRoleUserUpsertBulk) UpdateDeletedAt() *AppRoleUserUpsertBulk {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppRoleUserUpsertBulk) SetAppID(v uuid.UUID) *AppRoleUserUpsertBulk {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppRoleUserUpsertBulk) UpdateAppID() *AppRoleUserUpsertBulk {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppRoleUserUpsertBulk) ClearAppID() *AppRoleUserUpsertBulk {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.ClearAppID()
	})
}

// SetRoleID sets the "role_id" field.
func (u *AppRoleUserUpsertBulk) SetRoleID(v uuid.UUID) *AppRoleUserUpsertBulk {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.SetRoleID(v)
	})
}

// UpdateRoleID sets the "role_id" field to the value that was provided on create.
func (u *AppRoleUserUpsertBulk) UpdateRoleID() *AppRoleUserUpsertBulk {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.UpdateRoleID()
	})
}

// ClearRoleID clears the value of the "role_id" field.
func (u *AppRoleUserUpsertBulk) ClearRoleID() *AppRoleUserUpsertBulk {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.ClearRoleID()
	})
}

// SetUserID sets the "user_id" field.
func (u *AppRoleUserUpsertBulk) SetUserID(v uuid.UUID) *AppRoleUserUpsertBulk {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppRoleUserUpsertBulk) UpdateUserID() *AppRoleUserUpsertBulk {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *AppRoleUserUpsertBulk) ClearUserID() *AppRoleUserUpsertBulk {
	return u.Update(func(s *AppRoleUserUpsert) {
		s.ClearUserID()
	})
}

// Exec executes the query.
func (u *AppRoleUserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppRoleUserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppRoleUserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppRoleUserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
