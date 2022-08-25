// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/kyc"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// KycUpdate is the builder for updating Kyc entities.
type KycUpdate struct {
	config
	hooks     []Hook
	mutation  *KycMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the KycUpdate builder.
func (ku *KycUpdate) Where(ps ...predicate.Kyc) *KycUpdate {
	ku.mutation.Where(ps...)
	return ku
}

// SetCreatedAt sets the "created_at" field.
func (ku *KycUpdate) SetCreatedAt(u uint32) *KycUpdate {
	ku.mutation.ResetCreatedAt()
	ku.mutation.SetCreatedAt(u)
	return ku
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ku *KycUpdate) SetNillableCreatedAt(u *uint32) *KycUpdate {
	if u != nil {
		ku.SetCreatedAt(*u)
	}
	return ku
}

// AddCreatedAt adds u to the "created_at" field.
func (ku *KycUpdate) AddCreatedAt(u int32) *KycUpdate {
	ku.mutation.AddCreatedAt(u)
	return ku
}

// SetUpdatedAt sets the "updated_at" field.
func (ku *KycUpdate) SetUpdatedAt(u uint32) *KycUpdate {
	ku.mutation.ResetUpdatedAt()
	ku.mutation.SetUpdatedAt(u)
	return ku
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ku *KycUpdate) AddUpdatedAt(u int32) *KycUpdate {
	ku.mutation.AddUpdatedAt(u)
	return ku
}

// SetDeletedAt sets the "deleted_at" field.
func (ku *KycUpdate) SetDeletedAt(u uint32) *KycUpdate {
	ku.mutation.ResetDeletedAt()
	ku.mutation.SetDeletedAt(u)
	return ku
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ku *KycUpdate) SetNillableDeletedAt(u *uint32) *KycUpdate {
	if u != nil {
		ku.SetDeletedAt(*u)
	}
	return ku
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ku *KycUpdate) AddDeletedAt(u int32) *KycUpdate {
	ku.mutation.AddDeletedAt(u)
	return ku
}

// SetAppID sets the "app_id" field.
func (ku *KycUpdate) SetAppID(u uuid.UUID) *KycUpdate {
	ku.mutation.SetAppID(u)
	return ku
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (ku *KycUpdate) SetNillableAppID(u *uuid.UUID) *KycUpdate {
	if u != nil {
		ku.SetAppID(*u)
	}
	return ku
}

// ClearAppID clears the value of the "app_id" field.
func (ku *KycUpdate) ClearAppID() *KycUpdate {
	ku.mutation.ClearAppID()
	return ku
}

// SetUserID sets the "user_id" field.
func (ku *KycUpdate) SetUserID(u uuid.UUID) *KycUpdate {
	ku.mutation.SetUserID(u)
	return ku
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ku *KycUpdate) SetNillableUserID(u *uuid.UUID) *KycUpdate {
	if u != nil {
		ku.SetUserID(*u)
	}
	return ku
}

// ClearUserID clears the value of the "user_id" field.
func (ku *KycUpdate) ClearUserID() *KycUpdate {
	ku.mutation.ClearUserID()
	return ku
}

// SetDocumentType sets the "document_type" field.
func (ku *KycUpdate) SetDocumentType(s string) *KycUpdate {
	ku.mutation.SetDocumentType(s)
	return ku
}

// SetNillableDocumentType sets the "document_type" field if the given value is not nil.
func (ku *KycUpdate) SetNillableDocumentType(s *string) *KycUpdate {
	if s != nil {
		ku.SetDocumentType(*s)
	}
	return ku
}

// ClearDocumentType clears the value of the "document_type" field.
func (ku *KycUpdate) ClearDocumentType() *KycUpdate {
	ku.mutation.ClearDocumentType()
	return ku
}

// SetIDNumber sets the "id_number" field.
func (ku *KycUpdate) SetIDNumber(s string) *KycUpdate {
	ku.mutation.SetIDNumber(s)
	return ku
}

// SetNillableIDNumber sets the "id_number" field if the given value is not nil.
func (ku *KycUpdate) SetNillableIDNumber(s *string) *KycUpdate {
	if s != nil {
		ku.SetIDNumber(*s)
	}
	return ku
}

// ClearIDNumber clears the value of the "id_number" field.
func (ku *KycUpdate) ClearIDNumber() *KycUpdate {
	ku.mutation.ClearIDNumber()
	return ku
}

// SetFrontImg sets the "front_img" field.
func (ku *KycUpdate) SetFrontImg(s string) *KycUpdate {
	ku.mutation.SetFrontImg(s)
	return ku
}

// SetNillableFrontImg sets the "front_img" field if the given value is not nil.
func (ku *KycUpdate) SetNillableFrontImg(s *string) *KycUpdate {
	if s != nil {
		ku.SetFrontImg(*s)
	}
	return ku
}

// ClearFrontImg clears the value of the "front_img" field.
func (ku *KycUpdate) ClearFrontImg() *KycUpdate {
	ku.mutation.ClearFrontImg()
	return ku
}

// SetBackImg sets the "back_img" field.
func (ku *KycUpdate) SetBackImg(s string) *KycUpdate {
	ku.mutation.SetBackImg(s)
	return ku
}

// SetNillableBackImg sets the "back_img" field if the given value is not nil.
func (ku *KycUpdate) SetNillableBackImg(s *string) *KycUpdate {
	if s != nil {
		ku.SetBackImg(*s)
	}
	return ku
}

// ClearBackImg clears the value of the "back_img" field.
func (ku *KycUpdate) ClearBackImg() *KycUpdate {
	ku.mutation.ClearBackImg()
	return ku
}

// SetSelfieImg sets the "selfie_img" field.
func (ku *KycUpdate) SetSelfieImg(s string) *KycUpdate {
	ku.mutation.SetSelfieImg(s)
	return ku
}

// SetNillableSelfieImg sets the "selfie_img" field if the given value is not nil.
func (ku *KycUpdate) SetNillableSelfieImg(s *string) *KycUpdate {
	if s != nil {
		ku.SetSelfieImg(*s)
	}
	return ku
}

// ClearSelfieImg clears the value of the "selfie_img" field.
func (ku *KycUpdate) ClearSelfieImg() *KycUpdate {
	ku.mutation.ClearSelfieImg()
	return ku
}

// SetEntityType sets the "entity_type" field.
func (ku *KycUpdate) SetEntityType(s string) *KycUpdate {
	ku.mutation.SetEntityType(s)
	return ku
}

// SetNillableEntityType sets the "entity_type" field if the given value is not nil.
func (ku *KycUpdate) SetNillableEntityType(s *string) *KycUpdate {
	if s != nil {
		ku.SetEntityType(*s)
	}
	return ku
}

// ClearEntityType clears the value of the "entity_type" field.
func (ku *KycUpdate) ClearEntityType() *KycUpdate {
	ku.mutation.ClearEntityType()
	return ku
}

// SetReviewID sets the "review_id" field.
func (ku *KycUpdate) SetReviewID(u uuid.UUID) *KycUpdate {
	ku.mutation.SetReviewID(u)
	return ku
}

// SetNillableReviewID sets the "review_id" field if the given value is not nil.
func (ku *KycUpdate) SetNillableReviewID(u *uuid.UUID) *KycUpdate {
	if u != nil {
		ku.SetReviewID(*u)
	}
	return ku
}

// ClearReviewID clears the value of the "review_id" field.
func (ku *KycUpdate) ClearReviewID() *KycUpdate {
	ku.mutation.ClearReviewID()
	return ku
}

// SetState sets the "state" field.
func (ku *KycUpdate) SetState(s string) *KycUpdate {
	ku.mutation.SetState(s)
	return ku
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ku *KycUpdate) SetNillableState(s *string) *KycUpdate {
	if s != nil {
		ku.SetState(*s)
	}
	return ku
}

// ClearState clears the value of the "state" field.
func (ku *KycUpdate) ClearState() *KycUpdate {
	ku.mutation.ClearState()
	return ku
}

// Mutation returns the KycMutation object of the builder.
func (ku *KycUpdate) Mutation() *KycMutation {
	return ku.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ku *KycUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ku.defaults(); err != nil {
		return 0, err
	}
	if len(ku.hooks) == 0 {
		affected, err = ku.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KycMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ku.mutation = mutation
			affected, err = ku.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ku.hooks) - 1; i >= 0; i-- {
			if ku.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ku.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ku.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ku *KycUpdate) SaveX(ctx context.Context) int {
	affected, err := ku.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ku *KycUpdate) Exec(ctx context.Context) error {
	_, err := ku.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ku *KycUpdate) ExecX(ctx context.Context) {
	if err := ku.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ku *KycUpdate) defaults() error {
	if _, ok := ku.mutation.UpdatedAt(); !ok {
		if kyc.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized kyc.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := kyc.UpdateDefaultUpdatedAt()
		ku.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ku *KycUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *KycUpdate {
	ku.modifiers = append(ku.modifiers, modifiers...)
	return ku
}

func (ku *KycUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kyc.Table,
			Columns: kyc.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: kyc.FieldID,
			},
		},
	}
	if ps := ku.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ku.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldCreatedAt,
		})
	}
	if value, ok := ku.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldCreatedAt,
		})
	}
	if value, ok := ku.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldUpdatedAt,
		})
	}
	if value, ok := ku.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldUpdatedAt,
		})
	}
	if value, ok := ku.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldDeletedAt,
		})
	}
	if value, ok := ku.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldDeletedAt,
		})
	}
	if value, ok := ku.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: kyc.FieldAppID,
		})
	}
	if ku.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: kyc.FieldAppID,
		})
	}
	if value, ok := ku.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: kyc.FieldUserID,
		})
	}
	if ku.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: kyc.FieldUserID,
		})
	}
	if value, ok := ku.mutation.DocumentType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldDocumentType,
		})
	}
	if ku.mutation.DocumentTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: kyc.FieldDocumentType,
		})
	}
	if value, ok := ku.mutation.IDNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldIDNumber,
		})
	}
	if ku.mutation.IDNumberCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: kyc.FieldIDNumber,
		})
	}
	if value, ok := ku.mutation.FrontImg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldFrontImg,
		})
	}
	if ku.mutation.FrontImgCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: kyc.FieldFrontImg,
		})
	}
	if value, ok := ku.mutation.BackImg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldBackImg,
		})
	}
	if ku.mutation.BackImgCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: kyc.FieldBackImg,
		})
	}
	if value, ok := ku.mutation.SelfieImg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldSelfieImg,
		})
	}
	if ku.mutation.SelfieImgCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: kyc.FieldSelfieImg,
		})
	}
	if value, ok := ku.mutation.EntityType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldEntityType,
		})
	}
	if ku.mutation.EntityTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: kyc.FieldEntityType,
		})
	}
	if value, ok := ku.mutation.ReviewID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: kyc.FieldReviewID,
		})
	}
	if ku.mutation.ReviewIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: kyc.FieldReviewID,
		})
	}
	if value, ok := ku.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldState,
		})
	}
	if ku.mutation.StateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: kyc.FieldState,
		})
	}
	_spec.Modifiers = ku.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, ku.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{kyc.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// KycUpdateOne is the builder for updating a single Kyc entity.
type KycUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *KycMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (kuo *KycUpdateOne) SetCreatedAt(u uint32) *KycUpdateOne {
	kuo.mutation.ResetCreatedAt()
	kuo.mutation.SetCreatedAt(u)
	return kuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (kuo *KycUpdateOne) SetNillableCreatedAt(u *uint32) *KycUpdateOne {
	if u != nil {
		kuo.SetCreatedAt(*u)
	}
	return kuo
}

// AddCreatedAt adds u to the "created_at" field.
func (kuo *KycUpdateOne) AddCreatedAt(u int32) *KycUpdateOne {
	kuo.mutation.AddCreatedAt(u)
	return kuo
}

// SetUpdatedAt sets the "updated_at" field.
func (kuo *KycUpdateOne) SetUpdatedAt(u uint32) *KycUpdateOne {
	kuo.mutation.ResetUpdatedAt()
	kuo.mutation.SetUpdatedAt(u)
	return kuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (kuo *KycUpdateOne) AddUpdatedAt(u int32) *KycUpdateOne {
	kuo.mutation.AddUpdatedAt(u)
	return kuo
}

// SetDeletedAt sets the "deleted_at" field.
func (kuo *KycUpdateOne) SetDeletedAt(u uint32) *KycUpdateOne {
	kuo.mutation.ResetDeletedAt()
	kuo.mutation.SetDeletedAt(u)
	return kuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (kuo *KycUpdateOne) SetNillableDeletedAt(u *uint32) *KycUpdateOne {
	if u != nil {
		kuo.SetDeletedAt(*u)
	}
	return kuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (kuo *KycUpdateOne) AddDeletedAt(u int32) *KycUpdateOne {
	kuo.mutation.AddDeletedAt(u)
	return kuo
}

// SetAppID sets the "app_id" field.
func (kuo *KycUpdateOne) SetAppID(u uuid.UUID) *KycUpdateOne {
	kuo.mutation.SetAppID(u)
	return kuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (kuo *KycUpdateOne) SetNillableAppID(u *uuid.UUID) *KycUpdateOne {
	if u != nil {
		kuo.SetAppID(*u)
	}
	return kuo
}

// ClearAppID clears the value of the "app_id" field.
func (kuo *KycUpdateOne) ClearAppID() *KycUpdateOne {
	kuo.mutation.ClearAppID()
	return kuo
}

// SetUserID sets the "user_id" field.
func (kuo *KycUpdateOne) SetUserID(u uuid.UUID) *KycUpdateOne {
	kuo.mutation.SetUserID(u)
	return kuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (kuo *KycUpdateOne) SetNillableUserID(u *uuid.UUID) *KycUpdateOne {
	if u != nil {
		kuo.SetUserID(*u)
	}
	return kuo
}

// ClearUserID clears the value of the "user_id" field.
func (kuo *KycUpdateOne) ClearUserID() *KycUpdateOne {
	kuo.mutation.ClearUserID()
	return kuo
}

// SetDocumentType sets the "document_type" field.
func (kuo *KycUpdateOne) SetDocumentType(s string) *KycUpdateOne {
	kuo.mutation.SetDocumentType(s)
	return kuo
}

// SetNillableDocumentType sets the "document_type" field if the given value is not nil.
func (kuo *KycUpdateOne) SetNillableDocumentType(s *string) *KycUpdateOne {
	if s != nil {
		kuo.SetDocumentType(*s)
	}
	return kuo
}

// ClearDocumentType clears the value of the "document_type" field.
func (kuo *KycUpdateOne) ClearDocumentType() *KycUpdateOne {
	kuo.mutation.ClearDocumentType()
	return kuo
}

// SetIDNumber sets the "id_number" field.
func (kuo *KycUpdateOne) SetIDNumber(s string) *KycUpdateOne {
	kuo.mutation.SetIDNumber(s)
	return kuo
}

// SetNillableIDNumber sets the "id_number" field if the given value is not nil.
func (kuo *KycUpdateOne) SetNillableIDNumber(s *string) *KycUpdateOne {
	if s != nil {
		kuo.SetIDNumber(*s)
	}
	return kuo
}

// ClearIDNumber clears the value of the "id_number" field.
func (kuo *KycUpdateOne) ClearIDNumber() *KycUpdateOne {
	kuo.mutation.ClearIDNumber()
	return kuo
}

// SetFrontImg sets the "front_img" field.
func (kuo *KycUpdateOne) SetFrontImg(s string) *KycUpdateOne {
	kuo.mutation.SetFrontImg(s)
	return kuo
}

// SetNillableFrontImg sets the "front_img" field if the given value is not nil.
func (kuo *KycUpdateOne) SetNillableFrontImg(s *string) *KycUpdateOne {
	if s != nil {
		kuo.SetFrontImg(*s)
	}
	return kuo
}

// ClearFrontImg clears the value of the "front_img" field.
func (kuo *KycUpdateOne) ClearFrontImg() *KycUpdateOne {
	kuo.mutation.ClearFrontImg()
	return kuo
}

// SetBackImg sets the "back_img" field.
func (kuo *KycUpdateOne) SetBackImg(s string) *KycUpdateOne {
	kuo.mutation.SetBackImg(s)
	return kuo
}

// SetNillableBackImg sets the "back_img" field if the given value is not nil.
func (kuo *KycUpdateOne) SetNillableBackImg(s *string) *KycUpdateOne {
	if s != nil {
		kuo.SetBackImg(*s)
	}
	return kuo
}

// ClearBackImg clears the value of the "back_img" field.
func (kuo *KycUpdateOne) ClearBackImg() *KycUpdateOne {
	kuo.mutation.ClearBackImg()
	return kuo
}

// SetSelfieImg sets the "selfie_img" field.
func (kuo *KycUpdateOne) SetSelfieImg(s string) *KycUpdateOne {
	kuo.mutation.SetSelfieImg(s)
	return kuo
}

// SetNillableSelfieImg sets the "selfie_img" field if the given value is not nil.
func (kuo *KycUpdateOne) SetNillableSelfieImg(s *string) *KycUpdateOne {
	if s != nil {
		kuo.SetSelfieImg(*s)
	}
	return kuo
}

// ClearSelfieImg clears the value of the "selfie_img" field.
func (kuo *KycUpdateOne) ClearSelfieImg() *KycUpdateOne {
	kuo.mutation.ClearSelfieImg()
	return kuo
}

// SetEntityType sets the "entity_type" field.
func (kuo *KycUpdateOne) SetEntityType(s string) *KycUpdateOne {
	kuo.mutation.SetEntityType(s)
	return kuo
}

// SetNillableEntityType sets the "entity_type" field if the given value is not nil.
func (kuo *KycUpdateOne) SetNillableEntityType(s *string) *KycUpdateOne {
	if s != nil {
		kuo.SetEntityType(*s)
	}
	return kuo
}

// ClearEntityType clears the value of the "entity_type" field.
func (kuo *KycUpdateOne) ClearEntityType() *KycUpdateOne {
	kuo.mutation.ClearEntityType()
	return kuo
}

// SetReviewID sets the "review_id" field.
func (kuo *KycUpdateOne) SetReviewID(u uuid.UUID) *KycUpdateOne {
	kuo.mutation.SetReviewID(u)
	return kuo
}

// SetNillableReviewID sets the "review_id" field if the given value is not nil.
func (kuo *KycUpdateOne) SetNillableReviewID(u *uuid.UUID) *KycUpdateOne {
	if u != nil {
		kuo.SetReviewID(*u)
	}
	return kuo
}

// ClearReviewID clears the value of the "review_id" field.
func (kuo *KycUpdateOne) ClearReviewID() *KycUpdateOne {
	kuo.mutation.ClearReviewID()
	return kuo
}

// SetState sets the "state" field.
func (kuo *KycUpdateOne) SetState(s string) *KycUpdateOne {
	kuo.mutation.SetState(s)
	return kuo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (kuo *KycUpdateOne) SetNillableState(s *string) *KycUpdateOne {
	if s != nil {
		kuo.SetState(*s)
	}
	return kuo
}

// ClearState clears the value of the "state" field.
func (kuo *KycUpdateOne) ClearState() *KycUpdateOne {
	kuo.mutation.ClearState()
	return kuo
}

// Mutation returns the KycMutation object of the builder.
func (kuo *KycUpdateOne) Mutation() *KycMutation {
	return kuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (kuo *KycUpdateOne) Select(field string, fields ...string) *KycUpdateOne {
	kuo.fields = append([]string{field}, fields...)
	return kuo
}

// Save executes the query and returns the updated Kyc entity.
func (kuo *KycUpdateOne) Save(ctx context.Context) (*Kyc, error) {
	var (
		err  error
		node *Kyc
	)
	if err := kuo.defaults(); err != nil {
		return nil, err
	}
	if len(kuo.hooks) == 0 {
		node, err = kuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KycMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			kuo.mutation = mutation
			node, err = kuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kuo.hooks) - 1; i >= 0; i-- {
			if kuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, kuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Kyc)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from KycMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (kuo *KycUpdateOne) SaveX(ctx context.Context) *Kyc {
	node, err := kuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (kuo *KycUpdateOne) Exec(ctx context.Context) error {
	_, err := kuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kuo *KycUpdateOne) ExecX(ctx context.Context) {
	if err := kuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kuo *KycUpdateOne) defaults() error {
	if _, ok := kuo.mutation.UpdatedAt(); !ok {
		if kyc.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized kyc.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := kyc.UpdateDefaultUpdatedAt()
		kuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (kuo *KycUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *KycUpdateOne {
	kuo.modifiers = append(kuo.modifiers, modifiers...)
	return kuo
}

func (kuo *KycUpdateOne) sqlSave(ctx context.Context) (_node *Kyc, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kyc.Table,
			Columns: kyc.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: kyc.FieldID,
			},
		},
	}
	id, ok := kuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Kyc.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := kuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, kyc.FieldID)
		for _, f := range fields {
			if !kyc.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != kyc.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := kuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldCreatedAt,
		})
	}
	if value, ok := kuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldCreatedAt,
		})
	}
	if value, ok := kuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldUpdatedAt,
		})
	}
	if value, ok := kuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldUpdatedAt,
		})
	}
	if value, ok := kuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldDeletedAt,
		})
	}
	if value, ok := kuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldDeletedAt,
		})
	}
	if value, ok := kuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: kyc.FieldAppID,
		})
	}
	if kuo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: kyc.FieldAppID,
		})
	}
	if value, ok := kuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: kyc.FieldUserID,
		})
	}
	if kuo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: kyc.FieldUserID,
		})
	}
	if value, ok := kuo.mutation.DocumentType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldDocumentType,
		})
	}
	if kuo.mutation.DocumentTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: kyc.FieldDocumentType,
		})
	}
	if value, ok := kuo.mutation.IDNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldIDNumber,
		})
	}
	if kuo.mutation.IDNumberCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: kyc.FieldIDNumber,
		})
	}
	if value, ok := kuo.mutation.FrontImg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldFrontImg,
		})
	}
	if kuo.mutation.FrontImgCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: kyc.FieldFrontImg,
		})
	}
	if value, ok := kuo.mutation.BackImg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldBackImg,
		})
	}
	if kuo.mutation.BackImgCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: kyc.FieldBackImg,
		})
	}
	if value, ok := kuo.mutation.SelfieImg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldSelfieImg,
		})
	}
	if kuo.mutation.SelfieImgCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: kyc.FieldSelfieImg,
		})
	}
	if value, ok := kuo.mutation.EntityType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldEntityType,
		})
	}
	if kuo.mutation.EntityTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: kyc.FieldEntityType,
		})
	}
	if value, ok := kuo.mutation.ReviewID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: kyc.FieldReviewID,
		})
	}
	if kuo.mutation.ReviewIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: kyc.FieldReviewID,
		})
	}
	if value, ok := kuo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldState,
		})
	}
	if kuo.mutation.StateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: kyc.FieldState,
		})
	}
	_spec.Modifiers = kuo.modifiers
	_node = &Kyc{config: kuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, kuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{kyc.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
