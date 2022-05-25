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
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserthirdparty"
	"github.com/google/uuid"
)

// AppUserThirdPartyCreate is the builder for creating a AppUserThirdParty entity.
type AppUserThirdPartyCreate struct {
	config
	mutation *AppUserThirdPartyMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateAt sets the "create_at" field.
func (autpc *AppUserThirdPartyCreate) SetCreateAt(u uint32) *AppUserThirdPartyCreate {
	autpc.mutation.SetCreateAt(u)
	return autpc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (autpc *AppUserThirdPartyCreate) SetNillableCreateAt(u *uint32) *AppUserThirdPartyCreate {
	if u != nil {
		autpc.SetCreateAt(*u)
	}
	return autpc
}

// SetUpdateAt sets the "update_at" field.
func (autpc *AppUserThirdPartyCreate) SetUpdateAt(u uint32) *AppUserThirdPartyCreate {
	autpc.mutation.SetUpdateAt(u)
	return autpc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (autpc *AppUserThirdPartyCreate) SetNillableUpdateAt(u *uint32) *AppUserThirdPartyCreate {
	if u != nil {
		autpc.SetUpdateAt(*u)
	}
	return autpc
}

// SetDeleteAt sets the "delete_at" field.
func (autpc *AppUserThirdPartyCreate) SetDeleteAt(u uint32) *AppUserThirdPartyCreate {
	autpc.mutation.SetDeleteAt(u)
	return autpc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (autpc *AppUserThirdPartyCreate) SetNillableDeleteAt(u *uint32) *AppUserThirdPartyCreate {
	if u != nil {
		autpc.SetDeleteAt(*u)
	}
	return autpc
}

// SetAppID sets the "app_id" field.
func (autpc *AppUserThirdPartyCreate) SetAppID(u uuid.UUID) *AppUserThirdPartyCreate {
	autpc.mutation.SetAppID(u)
	return autpc
}

// SetUserID sets the "user_id" field.
func (autpc *AppUserThirdPartyCreate) SetUserID(u uuid.UUID) *AppUserThirdPartyCreate {
	autpc.mutation.SetUserID(u)
	return autpc
}

// SetThirdPartyUserID sets the "third_party_user_id" field.
func (autpc *AppUserThirdPartyCreate) SetThirdPartyUserID(s string) *AppUserThirdPartyCreate {
	autpc.mutation.SetThirdPartyUserID(s)
	return autpc
}

// SetThirdPartyID sets the "third_party_id" field.
func (autpc *AppUserThirdPartyCreate) SetThirdPartyID(s string) *AppUserThirdPartyCreate {
	autpc.mutation.SetThirdPartyID(s)
	return autpc
}

// SetThirdPartyUserName sets the "third_party_user_name" field.
func (autpc *AppUserThirdPartyCreate) SetThirdPartyUserName(s string) *AppUserThirdPartyCreate {
	autpc.mutation.SetThirdPartyUserName(s)
	return autpc
}

// SetThirdPartyUserAvatar sets the "third_party_user_avatar" field.
func (autpc *AppUserThirdPartyCreate) SetThirdPartyUserAvatar(s string) *AppUserThirdPartyCreate {
	autpc.mutation.SetThirdPartyUserAvatar(s)
	return autpc
}

// SetID sets the "id" field.
func (autpc *AppUserThirdPartyCreate) SetID(u uuid.UUID) *AppUserThirdPartyCreate {
	autpc.mutation.SetID(u)
	return autpc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (autpc *AppUserThirdPartyCreate) SetNillableID(u *uuid.UUID) *AppUserThirdPartyCreate {
	if u != nil {
		autpc.SetID(*u)
	}
	return autpc
}

// Mutation returns the AppUserThirdPartyMutation object of the builder.
func (autpc *AppUserThirdPartyCreate) Mutation() *AppUserThirdPartyMutation {
	return autpc.mutation
}

// Save creates the AppUserThirdParty in the database.
func (autpc *AppUserThirdPartyCreate) Save(ctx context.Context) (*AppUserThirdParty, error) {
	var (
		err  error
		node *AppUserThirdParty
	)
	if err := autpc.defaults(); err != nil {
		return nil, err
	}
	if len(autpc.hooks) == 0 {
		if err = autpc.check(); err != nil {
			return nil, err
		}
		node, err = autpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserThirdPartyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = autpc.check(); err != nil {
				return nil, err
			}
			autpc.mutation = mutation
			if node, err = autpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(autpc.hooks) - 1; i >= 0; i-- {
			if autpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = autpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, autpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (autpc *AppUserThirdPartyCreate) SaveX(ctx context.Context) *AppUserThirdParty {
	v, err := autpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (autpc *AppUserThirdPartyCreate) Exec(ctx context.Context) error {
	_, err := autpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (autpc *AppUserThirdPartyCreate) ExecX(ctx context.Context) {
	if err := autpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (autpc *AppUserThirdPartyCreate) defaults() error {
	if _, ok := autpc.mutation.CreateAt(); !ok {
		if appuserthirdparty.DefaultCreateAt == nil {
			return fmt.Errorf("ent: uninitialized appuserthirdparty.DefaultCreateAt (forgotten import ent/runtime?)")
		}
		v := appuserthirdparty.DefaultCreateAt()
		autpc.mutation.SetCreateAt(v)
	}
	if _, ok := autpc.mutation.UpdateAt(); !ok {
		if appuserthirdparty.DefaultUpdateAt == nil {
			return fmt.Errorf("ent: uninitialized appuserthirdparty.DefaultUpdateAt (forgotten import ent/runtime?)")
		}
		v := appuserthirdparty.DefaultUpdateAt()
		autpc.mutation.SetUpdateAt(v)
	}
	if _, ok := autpc.mutation.DeleteAt(); !ok {
		if appuserthirdparty.DefaultDeleteAt == nil {
			return fmt.Errorf("ent: uninitialized appuserthirdparty.DefaultDeleteAt (forgotten import ent/runtime?)")
		}
		v := appuserthirdparty.DefaultDeleteAt()
		autpc.mutation.SetDeleteAt(v)
	}
	if _, ok := autpc.mutation.ID(); !ok {
		if appuserthirdparty.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized appuserthirdparty.DefaultID (forgotten import ent/runtime?)")
		}
		v := appuserthirdparty.DefaultID()
		autpc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (autpc *AppUserThirdPartyCreate) check() error {
	if _, ok := autpc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "AppUserThirdParty.create_at"`)}
	}
	if _, ok := autpc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "AppUserThirdParty.update_at"`)}
	}
	if _, ok := autpc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "AppUserThirdParty.delete_at"`)}
	}
	if _, ok := autpc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "AppUserThirdParty.app_id"`)}
	}
	if _, ok := autpc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "AppUserThirdParty.user_id"`)}
	}
	if _, ok := autpc.mutation.ThirdPartyUserID(); !ok {
		return &ValidationError{Name: "third_party_user_id", err: errors.New(`ent: missing required field "AppUserThirdParty.third_party_user_id"`)}
	}
	if _, ok := autpc.mutation.ThirdPartyID(); !ok {
		return &ValidationError{Name: "third_party_id", err: errors.New(`ent: missing required field "AppUserThirdParty.third_party_id"`)}
	}
	if _, ok := autpc.mutation.ThirdPartyUserName(); !ok {
		return &ValidationError{Name: "third_party_user_name", err: errors.New(`ent: missing required field "AppUserThirdParty.third_party_user_name"`)}
	}
	if _, ok := autpc.mutation.ThirdPartyUserAvatar(); !ok {
		return &ValidationError{Name: "third_party_user_avatar", err: errors.New(`ent: missing required field "AppUserThirdParty.third_party_user_avatar"`)}
	}
	if v, ok := autpc.mutation.ThirdPartyUserAvatar(); ok {
		if err := appuserthirdparty.ThirdPartyUserAvatarValidator(v); err != nil {
			return &ValidationError{Name: "third_party_user_avatar", err: fmt.Errorf(`ent: validator failed for field "AppUserThirdParty.third_party_user_avatar": %w`, err)}
		}
	}
	return nil
}

func (autpc *AppUserThirdPartyCreate) sqlSave(ctx context.Context) (*AppUserThirdParty, error) {
	_node, _spec := autpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, autpc.driver, _spec); err != nil {
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

func (autpc *AppUserThirdPartyCreate) createSpec() (*AppUserThirdParty, *sqlgraph.CreateSpec) {
	var (
		_node = &AppUserThirdParty{config: autpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appuserthirdparty.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuserthirdparty.FieldID,
			},
		}
	)
	_spec.OnConflict = autpc.conflict
	if id, ok := autpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := autpc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthirdparty.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := autpc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthirdparty.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := autpc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthirdparty.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	if value, ok := autpc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserthirdparty.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := autpc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserthirdparty.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := autpc.mutation.ThirdPartyUserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserthirdparty.FieldThirdPartyUserID,
		})
		_node.ThirdPartyUserID = value
	}
	if value, ok := autpc.mutation.ThirdPartyID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserthirdparty.FieldThirdPartyID,
		})
		_node.ThirdPartyID = value
	}
	if value, ok := autpc.mutation.ThirdPartyUserName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserthirdparty.FieldThirdPartyUserName,
		})
		_node.ThirdPartyUserName = value
	}
	if value, ok := autpc.mutation.ThirdPartyUserAvatar(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserthirdparty.FieldThirdPartyUserAvatar,
		})
		_node.ThirdPartyUserAvatar = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppUserThirdParty.Create().
//		SetCreateAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppUserThirdPartyUpsert) {
//			SetCreateAt(v+v).
//		}).
//		Exec(ctx)
//
func (autpc *AppUserThirdPartyCreate) OnConflict(opts ...sql.ConflictOption) *AppUserThirdPartyUpsertOne {
	autpc.conflict = opts
	return &AppUserThirdPartyUpsertOne{
		create: autpc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppUserThirdParty.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (autpc *AppUserThirdPartyCreate) OnConflictColumns(columns ...string) *AppUserThirdPartyUpsertOne {
	autpc.conflict = append(autpc.conflict, sql.ConflictColumns(columns...))
	return &AppUserThirdPartyUpsertOne{
		create: autpc,
	}
}

type (
	// AppUserThirdPartyUpsertOne is the builder for "upsert"-ing
	//  one AppUserThirdParty node.
	AppUserThirdPartyUpsertOne struct {
		create *AppUserThirdPartyCreate
	}

	// AppUserThirdPartyUpsert is the "OnConflict" setter.
	AppUserThirdPartyUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreateAt sets the "create_at" field.
func (u *AppUserThirdPartyUpsert) SetCreateAt(v uint32) *AppUserThirdPartyUpsert {
	u.Set(appuserthirdparty.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsert) UpdateCreateAt() *AppUserThirdPartyUpsert {
	u.SetExcluded(appuserthirdparty.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *AppUserThirdPartyUpsert) AddCreateAt(v uint32) *AppUserThirdPartyUpsert {
	u.Add(appuserthirdparty.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *AppUserThirdPartyUpsert) SetUpdateAt(v uint32) *AppUserThirdPartyUpsert {
	u.Set(appuserthirdparty.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsert) UpdateUpdateAt() *AppUserThirdPartyUpsert {
	u.SetExcluded(appuserthirdparty.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *AppUserThirdPartyUpsert) AddUpdateAt(v uint32) *AppUserThirdPartyUpsert {
	u.Add(appuserthirdparty.FieldUpdateAt, v)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *AppUserThirdPartyUpsert) SetDeleteAt(v uint32) *AppUserThirdPartyUpsert {
	u.Set(appuserthirdparty.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsert) UpdateDeleteAt() *AppUserThirdPartyUpsert {
	u.SetExcluded(appuserthirdparty.FieldDeleteAt)
	return u
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *AppUserThirdPartyUpsert) AddDeleteAt(v uint32) *AppUserThirdPartyUpsert {
	u.Add(appuserthirdparty.FieldDeleteAt, v)
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppUserThirdPartyUpsert) SetAppID(v uuid.UUID) *AppUserThirdPartyUpsert {
	u.Set(appuserthirdparty.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsert) UpdateAppID() *AppUserThirdPartyUpsert {
	u.SetExcluded(appuserthirdparty.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *AppUserThirdPartyUpsert) SetUserID(v uuid.UUID) *AppUserThirdPartyUpsert {
	u.Set(appuserthirdparty.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsert) UpdateUserID() *AppUserThirdPartyUpsert {
	u.SetExcluded(appuserthirdparty.FieldUserID)
	return u
}

// SetThirdPartyUserID sets the "third_party_user_id" field.
func (u *AppUserThirdPartyUpsert) SetThirdPartyUserID(v string) *AppUserThirdPartyUpsert {
	u.Set(appuserthirdparty.FieldThirdPartyUserID, v)
	return u
}

// UpdateThirdPartyUserID sets the "third_party_user_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsert) UpdateThirdPartyUserID() *AppUserThirdPartyUpsert {
	u.SetExcluded(appuserthirdparty.FieldThirdPartyUserID)
	return u
}

// SetThirdPartyID sets the "third_party_id" field.
func (u *AppUserThirdPartyUpsert) SetThirdPartyID(v string) *AppUserThirdPartyUpsert {
	u.Set(appuserthirdparty.FieldThirdPartyID, v)
	return u
}

// UpdateThirdPartyID sets the "third_party_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsert) UpdateThirdPartyID() *AppUserThirdPartyUpsert {
	u.SetExcluded(appuserthirdparty.FieldThirdPartyID)
	return u
}

// SetThirdPartyUserName sets the "third_party_user_name" field.
func (u *AppUserThirdPartyUpsert) SetThirdPartyUserName(v string) *AppUserThirdPartyUpsert {
	u.Set(appuserthirdparty.FieldThirdPartyUserName, v)
	return u
}

// UpdateThirdPartyUserName sets the "third_party_user_name" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsert) UpdateThirdPartyUserName() *AppUserThirdPartyUpsert {
	u.SetExcluded(appuserthirdparty.FieldThirdPartyUserName)
	return u
}

// SetThirdPartyUserAvatar sets the "third_party_user_avatar" field.
func (u *AppUserThirdPartyUpsert) SetThirdPartyUserAvatar(v string) *AppUserThirdPartyUpsert {
	u.Set(appuserthirdparty.FieldThirdPartyUserAvatar, v)
	return u
}

// UpdateThirdPartyUserAvatar sets the "third_party_user_avatar" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsert) UpdateThirdPartyUserAvatar() *AppUserThirdPartyUpsert {
	u.SetExcluded(appuserthirdparty.FieldThirdPartyUserAvatar)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppUserThirdParty.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appuserthirdparty.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppUserThirdPartyUpsertOne) UpdateNewValues() *AppUserThirdPartyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appuserthirdparty.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppUserThirdParty.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppUserThirdPartyUpsertOne) Ignore() *AppUserThirdPartyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppUserThirdPartyUpsertOne) DoNothing() *AppUserThirdPartyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppUserThirdPartyCreate.OnConflict
// documentation for more info.
func (u *AppUserThirdPartyUpsertOne) Update(set func(*AppUserThirdPartyUpsert)) *AppUserThirdPartyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppUserThirdPartyUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *AppUserThirdPartyUpsertOne) SetCreateAt(v uint32) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *AppUserThirdPartyUpsertOne) AddCreateAt(v uint32) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertOne) UpdateCreateAt() *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *AppUserThirdPartyUpsertOne) SetUpdateAt(v uint32) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *AppUserThirdPartyUpsertOne) AddUpdateAt(v uint32) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertOne) UpdateUpdateAt() *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *AppUserThirdPartyUpsertOne) SetDeleteAt(v uint32) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *AppUserThirdPartyUpsertOne) AddDeleteAt(v uint32) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertOne) UpdateDeleteAt() *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateDeleteAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppUserThirdPartyUpsertOne) SetAppID(v uuid.UUID) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertOne) UpdateAppID() *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *AppUserThirdPartyUpsertOne) SetUserID(v uuid.UUID) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertOne) UpdateUserID() *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateUserID()
	})
}

// SetThirdPartyUserID sets the "third_party_user_id" field.
func (u *AppUserThirdPartyUpsertOne) SetThirdPartyUserID(v string) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetThirdPartyUserID(v)
	})
}

// UpdateThirdPartyUserID sets the "third_party_user_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertOne) UpdateThirdPartyUserID() *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateThirdPartyUserID()
	})
}

// SetThirdPartyID sets the "third_party_id" field.
func (u *AppUserThirdPartyUpsertOne) SetThirdPartyID(v string) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetThirdPartyID(v)
	})
}

// UpdateThirdPartyID sets the "third_party_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertOne) UpdateThirdPartyID() *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateThirdPartyID()
	})
}

// SetThirdPartyUserName sets the "third_party_user_name" field.
func (u *AppUserThirdPartyUpsertOne) SetThirdPartyUserName(v string) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetThirdPartyUserName(v)
	})
}

// UpdateThirdPartyUserName sets the "third_party_user_name" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertOne) UpdateThirdPartyUserName() *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateThirdPartyUserName()
	})
}

// SetThirdPartyUserAvatar sets the "third_party_user_avatar" field.
func (u *AppUserThirdPartyUpsertOne) SetThirdPartyUserAvatar(v string) *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetThirdPartyUserAvatar(v)
	})
}

// UpdateThirdPartyUserAvatar sets the "third_party_user_avatar" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertOne) UpdateThirdPartyUserAvatar() *AppUserThirdPartyUpsertOne {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateThirdPartyUserAvatar()
	})
}

// Exec executes the query.
func (u *AppUserThirdPartyUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppUserThirdPartyCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppUserThirdPartyUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppUserThirdPartyUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AppUserThirdPartyUpsertOne.ID is not supported by MySQL driver. Use AppUserThirdPartyUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppUserThirdPartyUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppUserThirdPartyCreateBulk is the builder for creating many AppUserThirdParty entities in bulk.
type AppUserThirdPartyCreateBulk struct {
	config
	builders []*AppUserThirdPartyCreate
	conflict []sql.ConflictOption
}

// Save creates the AppUserThirdParty entities in the database.
func (autpcb *AppUserThirdPartyCreateBulk) Save(ctx context.Context) ([]*AppUserThirdParty, error) {
	specs := make([]*sqlgraph.CreateSpec, len(autpcb.builders))
	nodes := make([]*AppUserThirdParty, len(autpcb.builders))
	mutators := make([]Mutator, len(autpcb.builders))
	for i := range autpcb.builders {
		func(i int, root context.Context) {
			builder := autpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppUserThirdPartyMutation)
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
					_, err = mutators[i+1].Mutate(root, autpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = autpcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, autpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, autpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (autpcb *AppUserThirdPartyCreateBulk) SaveX(ctx context.Context) []*AppUserThirdParty {
	v, err := autpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (autpcb *AppUserThirdPartyCreateBulk) Exec(ctx context.Context) error {
	_, err := autpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (autpcb *AppUserThirdPartyCreateBulk) ExecX(ctx context.Context) {
	if err := autpcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppUserThirdParty.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppUserThirdPartyUpsert) {
//			SetCreateAt(v+v).
//		}).
//		Exec(ctx)
//
func (autpcb *AppUserThirdPartyCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppUserThirdPartyUpsertBulk {
	autpcb.conflict = opts
	return &AppUserThirdPartyUpsertBulk{
		create: autpcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppUserThirdParty.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (autpcb *AppUserThirdPartyCreateBulk) OnConflictColumns(columns ...string) *AppUserThirdPartyUpsertBulk {
	autpcb.conflict = append(autpcb.conflict, sql.ConflictColumns(columns...))
	return &AppUserThirdPartyUpsertBulk{
		create: autpcb,
	}
}

// AppUserThirdPartyUpsertBulk is the builder for "upsert"-ing
// a bulk of AppUserThirdParty nodes.
type AppUserThirdPartyUpsertBulk struct {
	create *AppUserThirdPartyCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppUserThirdParty.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appuserthirdparty.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppUserThirdPartyUpsertBulk) UpdateNewValues() *AppUserThirdPartyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appuserthirdparty.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppUserThirdParty.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppUserThirdPartyUpsertBulk) Ignore() *AppUserThirdPartyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppUserThirdPartyUpsertBulk) DoNothing() *AppUserThirdPartyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppUserThirdPartyCreateBulk.OnConflict
// documentation for more info.
func (u *AppUserThirdPartyUpsertBulk) Update(set func(*AppUserThirdPartyUpsert)) *AppUserThirdPartyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppUserThirdPartyUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *AppUserThirdPartyUpsertBulk) SetCreateAt(v uint32) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *AppUserThirdPartyUpsertBulk) AddCreateAt(v uint32) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertBulk) UpdateCreateAt() *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *AppUserThirdPartyUpsertBulk) SetUpdateAt(v uint32) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *AppUserThirdPartyUpsertBulk) AddUpdateAt(v uint32) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertBulk) UpdateUpdateAt() *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *AppUserThirdPartyUpsertBulk) SetDeleteAt(v uint32) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *AppUserThirdPartyUpsertBulk) AddDeleteAt(v uint32) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertBulk) UpdateDeleteAt() *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateDeleteAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppUserThirdPartyUpsertBulk) SetAppID(v uuid.UUID) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertBulk) UpdateAppID() *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *AppUserThirdPartyUpsertBulk) SetUserID(v uuid.UUID) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertBulk) UpdateUserID() *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateUserID()
	})
}

// SetThirdPartyUserID sets the "third_party_user_id" field.
func (u *AppUserThirdPartyUpsertBulk) SetThirdPartyUserID(v string) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetThirdPartyUserID(v)
	})
}

// UpdateThirdPartyUserID sets the "third_party_user_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertBulk) UpdateThirdPartyUserID() *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateThirdPartyUserID()
	})
}

// SetThirdPartyID sets the "third_party_id" field.
func (u *AppUserThirdPartyUpsertBulk) SetThirdPartyID(v string) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetThirdPartyID(v)
	})
}

// UpdateThirdPartyID sets the "third_party_id" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertBulk) UpdateThirdPartyID() *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateThirdPartyID()
	})
}

// SetThirdPartyUserName sets the "third_party_user_name" field.
func (u *AppUserThirdPartyUpsertBulk) SetThirdPartyUserName(v string) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetThirdPartyUserName(v)
	})
}

// UpdateThirdPartyUserName sets the "third_party_user_name" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertBulk) UpdateThirdPartyUserName() *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateThirdPartyUserName()
	})
}

// SetThirdPartyUserAvatar sets the "third_party_user_avatar" field.
func (u *AppUserThirdPartyUpsertBulk) SetThirdPartyUserAvatar(v string) *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.SetThirdPartyUserAvatar(v)
	})
}

// UpdateThirdPartyUserAvatar sets the "third_party_user_avatar" field to the value that was provided on create.
func (u *AppUserThirdPartyUpsertBulk) UpdateThirdPartyUserAvatar() *AppUserThirdPartyUpsertBulk {
	return u.Update(func(s *AppUserThirdPartyUpsert) {
		s.UpdateThirdPartyUserAvatar()
	})
}

// Exec executes the query.
func (u *AppUserThirdPartyUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppUserThirdPartyCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppUserThirdPartyCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppUserThirdPartyUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
