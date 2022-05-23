// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserthird"
	"github.com/google/uuid"
)

// AppUserThirdCreate is the builder for creating a AppUserThird entity.
type AppUserThirdCreate struct {
	config
	mutation *AppUserThirdMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateAt sets the "create_at" field.
func (autc *AppUserThirdCreate) SetCreateAt(u uint32) *AppUserThirdCreate {
	autc.mutation.SetCreateAt(u)
	return autc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (autc *AppUserThirdCreate) SetNillableCreateAt(u *uint32) *AppUserThirdCreate {
	if u != nil {
		autc.SetCreateAt(*u)
	}
	return autc
}

// SetUpdateAt sets the "update_at" field.
func (autc *AppUserThirdCreate) SetUpdateAt(u uint32) *AppUserThirdCreate {
	autc.mutation.SetUpdateAt(u)
	return autc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (autc *AppUserThirdCreate) SetNillableUpdateAt(u *uint32) *AppUserThirdCreate {
	if u != nil {
		autc.SetUpdateAt(*u)
	}
	return autc
}

// SetDeleteAt sets the "delete_at" field.
func (autc *AppUserThirdCreate) SetDeleteAt(u uint32) *AppUserThirdCreate {
	autc.mutation.SetDeleteAt(u)
	return autc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (autc *AppUserThirdCreate) SetNillableDeleteAt(u *uint32) *AppUserThirdCreate {
	if u != nil {
		autc.SetDeleteAt(*u)
	}
	return autc
}

// SetAppID sets the "app_id" field.
func (autc *AppUserThirdCreate) SetAppID(u uuid.UUID) *AppUserThirdCreate {
	autc.mutation.SetAppID(u)
	return autc
}

// SetUserID sets the "user_id" field.
func (autc *AppUserThirdCreate) SetUserID(u uuid.UUID) *AppUserThirdCreate {
	autc.mutation.SetUserID(u)
	return autc
}

// SetThirdUserID sets the "third_user_id" field.
func (autc *AppUserThirdCreate) SetThirdUserID(s string) *AppUserThirdCreate {
	autc.mutation.SetThirdUserID(s)
	return autc
}

// SetThird sets the "third" field.
func (autc *AppUserThirdCreate) SetThird(s string) *AppUserThirdCreate {
	autc.mutation.SetThird(s)
	return autc
}

// SetThirdID sets the "third_id" field.
func (autc *AppUserThirdCreate) SetThirdID(s string) *AppUserThirdCreate {
	autc.mutation.SetThirdID(s)
	return autc
}

// SetThirdUserName sets the "third_user_name" field.
func (autc *AppUserThirdCreate) SetThirdUserName(s string) *AppUserThirdCreate {
	autc.mutation.SetThirdUserName(s)
	return autc
}

// SetThirdUserAvatar sets the "third_user_avatar" field.
func (autc *AppUserThirdCreate) SetThirdUserAvatar(s string) *AppUserThirdCreate {
	autc.mutation.SetThirdUserAvatar(s)
	return autc
}

// SetID sets the "id" field.
func (autc *AppUserThirdCreate) SetID(u uuid.UUID) *AppUserThirdCreate {
	autc.mutation.SetID(u)
	return autc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (autc *AppUserThirdCreate) SetNillableID(u *uuid.UUID) *AppUserThirdCreate {
	if u != nil {
		autc.SetID(*u)
	}
	return autc
}

// Mutation returns the AppUserThirdMutation object of the builder.
func (autc *AppUserThirdCreate) Mutation() *AppUserThirdMutation {
	return autc.mutation
}

// Save creates the AppUserThird in the database.
func (autc *AppUserThirdCreate) Save(ctx context.Context) (*AppUserThird, error) {
	var (
		err  error
		node *AppUserThird
	)
	if err := autc.defaults(); err != nil {
		return nil, err
	}
	if len(autc.hooks) == 0 {
		if err = autc.check(); err != nil {
			return nil, err
		}
		node, err = autc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserThirdMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = autc.check(); err != nil {
				return nil, err
			}
			autc.mutation = mutation
			if node, err = autc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(autc.hooks) - 1; i >= 0; i-- {
			if autc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = autc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, autc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (autc *AppUserThirdCreate) SaveX(ctx context.Context) *AppUserThird {
	v, err := autc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (autc *AppUserThirdCreate) Exec(ctx context.Context) error {
	_, err := autc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (autc *AppUserThirdCreate) ExecX(ctx context.Context) {
	if err := autc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (autc *AppUserThirdCreate) defaults() error {
	if _, ok := autc.mutation.CreateAt(); !ok {
		if appuserthird.DefaultCreateAt == nil {
			return fmt.Errorf("ent: uninitialized appuserthird.DefaultCreateAt (forgotten import ent/runtime?)")
		}
		v := appuserthird.DefaultCreateAt()
		autc.mutation.SetCreateAt(v)
	}
	if _, ok := autc.mutation.UpdateAt(); !ok {
		if appuserthird.DefaultUpdateAt == nil {
			return fmt.Errorf("ent: uninitialized appuserthird.DefaultUpdateAt (forgotten import ent/runtime?)")
		}
		v := appuserthird.DefaultUpdateAt()
		autc.mutation.SetUpdateAt(v)
	}
	if _, ok := autc.mutation.DeleteAt(); !ok {
		if appuserthird.DefaultDeleteAt == nil {
			return fmt.Errorf("ent: uninitialized appuserthird.DefaultDeleteAt (forgotten import ent/runtime?)")
		}
		v := appuserthird.DefaultDeleteAt()
		autc.mutation.SetDeleteAt(v)
	}
	if _, ok := autc.mutation.ID(); !ok {
		if appuserthird.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized appuserthird.DefaultID (forgotten import ent/runtime?)")
		}
		v := appuserthird.DefaultID()
		autc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (autc *AppUserThirdCreate) check() error {
	if _, ok := autc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "AppUserThird.create_at"`)}
	}
	if _, ok := autc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "AppUserThird.update_at"`)}
	}
	if _, ok := autc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "AppUserThird.delete_at"`)}
	}
	if _, ok := autc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "AppUserThird.app_id"`)}
	}
	if _, ok := autc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "AppUserThird.user_id"`)}
	}
	if _, ok := autc.mutation.ThirdUserID(); !ok {
		return &ValidationError{Name: "third_user_id", err: errors.New(`ent: missing required field "AppUserThird.third_user_id"`)}
	}
	if _, ok := autc.mutation.Third(); !ok {
		return &ValidationError{Name: "third", err: errors.New(`ent: missing required field "AppUserThird.third"`)}
	}
	if _, ok := autc.mutation.ThirdID(); !ok {
		return &ValidationError{Name: "third_id", err: errors.New(`ent: missing required field "AppUserThird.third_id"`)}
	}
	if _, ok := autc.mutation.ThirdUserName(); !ok {
		return &ValidationError{Name: "third_user_name", err: errors.New(`ent: missing required field "AppUserThird.third_user_name"`)}
	}
	if _, ok := autc.mutation.ThirdUserAvatar(); !ok {
		return &ValidationError{Name: "third_user_avatar", err: errors.New(`ent: missing required field "AppUserThird.third_user_avatar"`)}
	}
	if v, ok := autc.mutation.ThirdUserAvatar(); ok {
		if err := appuserthird.ThirdUserAvatarValidator(v); err != nil {
			return &ValidationError{Name: "third_user_avatar", err: fmt.Errorf(`ent: validator failed for field "AppUserThird.third_user_avatar": %w`, err)}
		}
	}
	return nil
}

func (autc *AppUserThirdCreate) sqlSave(ctx context.Context) (*AppUserThird, error) {
	_node, _spec := autc.createSpec()
	if err := sqlgraph.CreateNode(ctx, autc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
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

func (autc *AppUserThirdCreate) createSpec() (*AppUserThird, *sqlgraph.CreateSpec) {
	var (
		_node = &AppUserThird{config: autc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appuserthird.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuserthird.FieldID,
			},
		}
	)
	_spec.OnConflict = autc.conflict
	if id, ok := autc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := autc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthird.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := autc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthird.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := autc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthird.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	if value, ok := autc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserthird.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := autc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserthird.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := autc.mutation.ThirdUserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserthird.FieldThirdUserID,
		})
		_node.ThirdUserID = value
	}
	if value, ok := autc.mutation.Third(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserthird.FieldThird,
		})
		_node.Third = value
	}
	if value, ok := autc.mutation.ThirdID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserthird.FieldThirdID,
		})
		_node.ThirdID = value
	}
	if value, ok := autc.mutation.ThirdUserName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserthird.FieldThirdUserName,
		})
		_node.ThirdUserName = value
	}
	if value, ok := autc.mutation.ThirdUserAvatar(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserthird.FieldThirdUserAvatar,
		})
		_node.ThirdUserAvatar = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppUserThird.Create().
//		SetCreateAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppUserThirdUpsert) {
//			SetCreateAt(v+v).
//		}).
//		Exec(ctx)
//
func (autc *AppUserThirdCreate) OnConflict(opts ...sql.ConflictOption) *AppUserThirdUpsertOne {
	autc.conflict = opts
	return &AppUserThirdUpsertOne{
		create: autc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppUserThird.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (autc *AppUserThirdCreate) OnConflictColumns(columns ...string) *AppUserThirdUpsertOne {
	autc.conflict = append(autc.conflict, sql.ConflictColumns(columns...))
	return &AppUserThirdUpsertOne{
		create: autc,
	}
}

type (
	// AppUserThirdUpsertOne is the builder for "upsert"-ing
	//  one AppUserThird node.
	AppUserThirdUpsertOne struct {
		create *AppUserThirdCreate
	}

	// AppUserThirdUpsert is the "OnConflict" setter.
	AppUserThirdUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreateAt sets the "create_at" field.
func (u *AppUserThirdUpsert) SetCreateAt(v uint32) *AppUserThirdUpsert {
	u.Set(appuserthird.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppUserThirdUpsert) UpdateCreateAt() *AppUserThirdUpsert {
	u.SetExcluded(appuserthird.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *AppUserThirdUpsert) AddCreateAt(v uint32) *AppUserThirdUpsert {
	u.Add(appuserthird.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *AppUserThirdUpsert) SetUpdateAt(v uint32) *AppUserThirdUpsert {
	u.Set(appuserthird.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppUserThirdUpsert) UpdateUpdateAt() *AppUserThirdUpsert {
	u.SetExcluded(appuserthird.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *AppUserThirdUpsert) AddUpdateAt(v uint32) *AppUserThirdUpsert {
	u.Add(appuserthird.FieldUpdateAt, v)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *AppUserThirdUpsert) SetDeleteAt(v uint32) *AppUserThirdUpsert {
	u.Set(appuserthird.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *AppUserThirdUpsert) UpdateDeleteAt() *AppUserThirdUpsert {
	u.SetExcluded(appuserthird.FieldDeleteAt)
	return u
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *AppUserThirdUpsert) AddDeleteAt(v uint32) *AppUserThirdUpsert {
	u.Add(appuserthird.FieldDeleteAt, v)
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppUserThirdUpsert) SetAppID(v uuid.UUID) *AppUserThirdUpsert {
	u.Set(appuserthird.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppUserThirdUpsert) UpdateAppID() *AppUserThirdUpsert {
	u.SetExcluded(appuserthird.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *AppUserThirdUpsert) SetUserID(v uuid.UUID) *AppUserThirdUpsert {
	u.Set(appuserthird.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppUserThirdUpsert) UpdateUserID() *AppUserThirdUpsert {
	u.SetExcluded(appuserthird.FieldUserID)
	return u
}

// SetThirdUserID sets the "third_user_id" field.
func (u *AppUserThirdUpsert) SetThirdUserID(v string) *AppUserThirdUpsert {
	u.Set(appuserthird.FieldThirdUserID, v)
	return u
}

// UpdateThirdUserID sets the "third_user_id" field to the value that was provided on create.
func (u *AppUserThirdUpsert) UpdateThirdUserID() *AppUserThirdUpsert {
	u.SetExcluded(appuserthird.FieldThirdUserID)
	return u
}

// SetThird sets the "third" field.
func (u *AppUserThirdUpsert) SetThird(v string) *AppUserThirdUpsert {
	u.Set(appuserthird.FieldThird, v)
	return u
}

// UpdateThird sets the "third" field to the value that was provided on create.
func (u *AppUserThirdUpsert) UpdateThird() *AppUserThirdUpsert {
	u.SetExcluded(appuserthird.FieldThird)
	return u
}

// SetThirdID sets the "third_id" field.
func (u *AppUserThirdUpsert) SetThirdID(v string) *AppUserThirdUpsert {
	u.Set(appuserthird.FieldThirdID, v)
	return u
}

// UpdateThirdID sets the "third_id" field to the value that was provided on create.
func (u *AppUserThirdUpsert) UpdateThirdID() *AppUserThirdUpsert {
	u.SetExcluded(appuserthird.FieldThirdID)
	return u
}

// SetThirdUserName sets the "third_user_name" field.
func (u *AppUserThirdUpsert) SetThirdUserName(v string) *AppUserThirdUpsert {
	u.Set(appuserthird.FieldThirdUserName, v)
	return u
}

// UpdateThirdUserName sets the "third_user_name" field to the value that was provided on create.
func (u *AppUserThirdUpsert) UpdateThirdUserName() *AppUserThirdUpsert {
	u.SetExcluded(appuserthird.FieldThirdUserName)
	return u
}

// SetThirdUserAvatar sets the "third_user_avatar" field.
func (u *AppUserThirdUpsert) SetThirdUserAvatar(v string) *AppUserThirdUpsert {
	u.Set(appuserthird.FieldThirdUserAvatar, v)
	return u
}

// UpdateThirdUserAvatar sets the "third_user_avatar" field to the value that was provided on create.
func (u *AppUserThirdUpsert) UpdateThirdUserAvatar() *AppUserThirdUpsert {
	u.SetExcluded(appuserthird.FieldThirdUserAvatar)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppUserThird.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appuserthird.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppUserThirdUpsertOne) UpdateNewValues() *AppUserThirdUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appuserthird.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppUserThird.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppUserThirdUpsertOne) Ignore() *AppUserThirdUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppUserThirdUpsertOne) DoNothing() *AppUserThirdUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppUserThirdCreate.OnConflict
// documentation for more info.
func (u *AppUserThirdUpsertOne) Update(set func(*AppUserThirdUpsert)) *AppUserThirdUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppUserThirdUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *AppUserThirdUpsertOne) SetCreateAt(v uint32) *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *AppUserThirdUpsertOne) AddCreateAt(v uint32) *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppUserThirdUpsertOne) UpdateCreateAt() *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *AppUserThirdUpsertOne) SetUpdateAt(v uint32) *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *AppUserThirdUpsertOne) AddUpdateAt(v uint32) *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppUserThirdUpsertOne) UpdateUpdateAt() *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *AppUserThirdUpsertOne) SetDeleteAt(v uint32) *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *AppUserThirdUpsertOne) AddDeleteAt(v uint32) *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *AppUserThirdUpsertOne) UpdateDeleteAt() *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.UpdateDeleteAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppUserThirdUpsertOne) SetAppID(v uuid.UUID) *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppUserThirdUpsertOne) UpdateAppID() *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *AppUserThirdUpsertOne) SetUserID(v uuid.UUID) *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppUserThirdUpsertOne) UpdateUserID() *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.UpdateUserID()
	})
}

// SetThirdUserID sets the "third_user_id" field.
func (u *AppUserThirdUpsertOne) SetThirdUserID(v string) *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.SetThirdUserID(v)
	})
}

// UpdateThirdUserID sets the "third_user_id" field to the value that was provided on create.
func (u *AppUserThirdUpsertOne) UpdateThirdUserID() *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.UpdateThirdUserID()
	})
}

// SetThird sets the "third" field.
func (u *AppUserThirdUpsertOne) SetThird(v string) *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.SetThird(v)
	})
}

// UpdateThird sets the "third" field to the value that was provided on create.
func (u *AppUserThirdUpsertOne) UpdateThird() *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.UpdateThird()
	})
}

// SetThirdID sets the "third_id" field.
func (u *AppUserThirdUpsertOne) SetThirdID(v string) *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.SetThirdID(v)
	})
}

// UpdateThirdID sets the "third_id" field to the value that was provided on create.
func (u *AppUserThirdUpsertOne) UpdateThirdID() *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.UpdateThirdID()
	})
}

// SetThirdUserName sets the "third_user_name" field.
func (u *AppUserThirdUpsertOne) SetThirdUserName(v string) *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.SetThirdUserName(v)
	})
}

// UpdateThirdUserName sets the "third_user_name" field to the value that was provided on create.
func (u *AppUserThirdUpsertOne) UpdateThirdUserName() *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.UpdateThirdUserName()
	})
}

// SetThirdUserAvatar sets the "third_user_avatar" field.
func (u *AppUserThirdUpsertOne) SetThirdUserAvatar(v string) *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.SetThirdUserAvatar(v)
	})
}

// UpdateThirdUserAvatar sets the "third_user_avatar" field to the value that was provided on create.
func (u *AppUserThirdUpsertOne) UpdateThirdUserAvatar() *AppUserThirdUpsertOne {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.UpdateThirdUserAvatar()
	})
}

// Exec executes the query.
func (u *AppUserThirdUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppUserThirdCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppUserThirdUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppUserThirdUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AppUserThirdUpsertOne.ID is not supported by MySQL driver. Use AppUserThirdUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppUserThirdUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppUserThirdCreateBulk is the builder for creating many AppUserThird entities in bulk.
type AppUserThirdCreateBulk struct {
	config
	builders []*AppUserThirdCreate
	conflict []sql.ConflictOption
}

// Save creates the AppUserThird entities in the database.
func (autcb *AppUserThirdCreateBulk) Save(ctx context.Context) ([]*AppUserThird, error) {
	specs := make([]*sqlgraph.CreateSpec, len(autcb.builders))
	nodes := make([]*AppUserThird, len(autcb.builders))
	mutators := make([]Mutator, len(autcb.builders))
	for i := range autcb.builders {
		func(i int, root context.Context) {
			builder := autcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppUserThirdMutation)
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
					_, err = mutators[i+1].Mutate(root, autcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = autcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, autcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
		if _, err := mutators[0].Mutate(ctx, autcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (autcb *AppUserThirdCreateBulk) SaveX(ctx context.Context) []*AppUserThird {
	v, err := autcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (autcb *AppUserThirdCreateBulk) Exec(ctx context.Context) error {
	_, err := autcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (autcb *AppUserThirdCreateBulk) ExecX(ctx context.Context) {
	if err := autcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppUserThird.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppUserThirdUpsert) {
//			SetCreateAt(v+v).
//		}).
//		Exec(ctx)
//
func (autcb *AppUserThirdCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppUserThirdUpsertBulk {
	autcb.conflict = opts
	return &AppUserThirdUpsertBulk{
		create: autcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppUserThird.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (autcb *AppUserThirdCreateBulk) OnConflictColumns(columns ...string) *AppUserThirdUpsertBulk {
	autcb.conflict = append(autcb.conflict, sql.ConflictColumns(columns...))
	return &AppUserThirdUpsertBulk{
		create: autcb,
	}
}

// AppUserThirdUpsertBulk is the builder for "upsert"-ing
// a bulk of AppUserThird nodes.
type AppUserThirdUpsertBulk struct {
	create *AppUserThirdCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppUserThird.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appuserthird.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppUserThirdUpsertBulk) UpdateNewValues() *AppUserThirdUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appuserthird.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppUserThird.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppUserThirdUpsertBulk) Ignore() *AppUserThirdUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppUserThirdUpsertBulk) DoNothing() *AppUserThirdUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppUserThirdCreateBulk.OnConflict
// documentation for more info.
func (u *AppUserThirdUpsertBulk) Update(set func(*AppUserThirdUpsert)) *AppUserThirdUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppUserThirdUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *AppUserThirdUpsertBulk) SetCreateAt(v uint32) *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *AppUserThirdUpsertBulk) AddCreateAt(v uint32) *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppUserThirdUpsertBulk) UpdateCreateAt() *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *AppUserThirdUpsertBulk) SetUpdateAt(v uint32) *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *AppUserThirdUpsertBulk) AddUpdateAt(v uint32) *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppUserThirdUpsertBulk) UpdateUpdateAt() *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *AppUserThirdUpsertBulk) SetDeleteAt(v uint32) *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *AppUserThirdUpsertBulk) AddDeleteAt(v uint32) *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *AppUserThirdUpsertBulk) UpdateDeleteAt() *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.UpdateDeleteAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppUserThirdUpsertBulk) SetAppID(v uuid.UUID) *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppUserThirdUpsertBulk) UpdateAppID() *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *AppUserThirdUpsertBulk) SetUserID(v uuid.UUID) *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppUserThirdUpsertBulk) UpdateUserID() *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.UpdateUserID()
	})
}

// SetThirdUserID sets the "third_user_id" field.
func (u *AppUserThirdUpsertBulk) SetThirdUserID(v string) *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.SetThirdUserID(v)
	})
}

// UpdateThirdUserID sets the "third_user_id" field to the value that was provided on create.
func (u *AppUserThirdUpsertBulk) UpdateThirdUserID() *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.UpdateThirdUserID()
	})
}

// SetThird sets the "third" field.
func (u *AppUserThirdUpsertBulk) SetThird(v string) *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.SetThird(v)
	})
}

// UpdateThird sets the "third" field to the value that was provided on create.
func (u *AppUserThirdUpsertBulk) UpdateThird() *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.UpdateThird()
	})
}

// SetThirdID sets the "third_id" field.
func (u *AppUserThirdUpsertBulk) SetThirdID(v string) *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.SetThirdID(v)
	})
}

// UpdateThirdID sets the "third_id" field to the value that was provided on create.
func (u *AppUserThirdUpsertBulk) UpdateThirdID() *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.UpdateThirdID()
	})
}

// SetThirdUserName sets the "third_user_name" field.
func (u *AppUserThirdUpsertBulk) SetThirdUserName(v string) *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.SetThirdUserName(v)
	})
}

// UpdateThirdUserName sets the "third_user_name" field to the value that was provided on create.
func (u *AppUserThirdUpsertBulk) UpdateThirdUserName() *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.UpdateThirdUserName()
	})
}

// SetThirdUserAvatar sets the "third_user_avatar" field.
func (u *AppUserThirdUpsertBulk) SetThirdUserAvatar(v string) *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.SetThirdUserAvatar(v)
	})
}

// UpdateThirdUserAvatar sets the "third_user_avatar" field to the value that was provided on create.
func (u *AppUserThirdUpsertBulk) UpdateThirdUserAvatar() *AppUserThirdUpsertBulk {
	return u.Update(func(s *AppUserThirdUpsert) {
		s.UpdateThirdUserAvatar()
	})
}

// Exec executes the query.
func (u *AppUserThirdUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppUserThirdCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppUserThirdCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppUserThirdUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
