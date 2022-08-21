// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/kyc"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// KycQuery is the builder for querying Kyc entities.
type KycQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Kyc
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the KycQuery builder.
func (kq *KycQuery) Where(ps ...predicate.Kyc) *KycQuery {
	kq.predicates = append(kq.predicates, ps...)
	return kq
}

// Limit adds a limit step to the query.
func (kq *KycQuery) Limit(limit int) *KycQuery {
	kq.limit = &limit
	return kq
}

// Offset adds an offset step to the query.
func (kq *KycQuery) Offset(offset int) *KycQuery {
	kq.offset = &offset
	return kq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (kq *KycQuery) Unique(unique bool) *KycQuery {
	kq.unique = &unique
	return kq
}

// Order adds an order step to the query.
func (kq *KycQuery) Order(o ...OrderFunc) *KycQuery {
	kq.order = append(kq.order, o...)
	return kq
}

// First returns the first Kyc entity from the query.
// Returns a *NotFoundError when no Kyc was found.
func (kq *KycQuery) First(ctx context.Context) (*Kyc, error) {
	nodes, err := kq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{kyc.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (kq *KycQuery) FirstX(ctx context.Context) *Kyc {
	node, err := kq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Kyc ID from the query.
// Returns a *NotFoundError when no Kyc ID was found.
func (kq *KycQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = kq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{kyc.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (kq *KycQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := kq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Kyc entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Kyc entity is found.
// Returns a *NotFoundError when no Kyc entities are found.
func (kq *KycQuery) Only(ctx context.Context) (*Kyc, error) {
	nodes, err := kq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{kyc.Label}
	default:
		return nil, &NotSingularError{kyc.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (kq *KycQuery) OnlyX(ctx context.Context) *Kyc {
	node, err := kq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Kyc ID in the query.
// Returns a *NotSingularError when more than one Kyc ID is found.
// Returns a *NotFoundError when no entities are found.
func (kq *KycQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = kq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{kyc.Label}
	default:
		err = &NotSingularError{kyc.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (kq *KycQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := kq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Kycs.
func (kq *KycQuery) All(ctx context.Context) ([]*Kyc, error) {
	if err := kq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return kq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (kq *KycQuery) AllX(ctx context.Context) []*Kyc {
	nodes, err := kq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Kyc IDs.
func (kq *KycQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := kq.Select(kyc.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (kq *KycQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := kq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (kq *KycQuery) Count(ctx context.Context) (int, error) {
	if err := kq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return kq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (kq *KycQuery) CountX(ctx context.Context) int {
	count, err := kq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (kq *KycQuery) Exist(ctx context.Context) (bool, error) {
	if err := kq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return kq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (kq *KycQuery) ExistX(ctx context.Context) bool {
	exist, err := kq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the KycQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (kq *KycQuery) Clone() *KycQuery {
	if kq == nil {
		return nil
	}
	return &KycQuery{
		config:     kq.config,
		limit:      kq.limit,
		offset:     kq.offset,
		order:      append([]OrderFunc{}, kq.order...),
		predicates: append([]predicate.Kyc{}, kq.predicates...),
		// clone intermediate query.
		sql:    kq.sql.Clone(),
		path:   kq.path,
		unique: kq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Kyc.Query().
//		GroupBy(kyc.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (kq *KycQuery) GroupBy(field string, fields ...string) *KycGroupBy {
	grbuild := &KycGroupBy{config: kq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := kq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return kq.sqlQuery(ctx), nil
	}
	grbuild.label = kyc.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.Kyc.Query().
//		Select(kyc.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (kq *KycQuery) Select(fields ...string) *KycSelect {
	kq.fields = append(kq.fields, fields...)
	selbuild := &KycSelect{KycQuery: kq}
	selbuild.label = kyc.Label
	selbuild.flds, selbuild.scan = &kq.fields, selbuild.Scan
	return selbuild
}

func (kq *KycQuery) prepareQuery(ctx context.Context) error {
	for _, f := range kq.fields {
		if !kyc.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if kq.path != nil {
		prev, err := kq.path(ctx)
		if err != nil {
			return err
		}
		kq.sql = prev
	}
	if kyc.Policy == nil {
		return errors.New("ent: uninitialized kyc.Policy (forgotten import ent/runtime?)")
	}
	if err := kyc.Policy.EvalQuery(ctx, kq); err != nil {
		return err
	}
	return nil
}

func (kq *KycQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Kyc, error) {
	var (
		nodes = []*Kyc{}
		_spec = kq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Kyc).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Kyc{config: kq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(kq.modifiers) > 0 {
		_spec.Modifiers = kq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, kq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (kq *KycQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := kq.querySpec()
	if len(kq.modifiers) > 0 {
		_spec.Modifiers = kq.modifiers
	}
	_spec.Node.Columns = kq.fields
	if len(kq.fields) > 0 {
		_spec.Unique = kq.unique != nil && *kq.unique
	}
	return sqlgraph.CountNodes(ctx, kq.driver, _spec)
}

func (kq *KycQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := kq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (kq *KycQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kyc.Table,
			Columns: kyc.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: kyc.FieldID,
			},
		},
		From:   kq.sql,
		Unique: true,
	}
	if unique := kq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := kq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, kyc.FieldID)
		for i := range fields {
			if fields[i] != kyc.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := kq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := kq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := kq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := kq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (kq *KycQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(kq.driver.Dialect())
	t1 := builder.Table(kyc.Table)
	columns := kq.fields
	if len(columns) == 0 {
		columns = kyc.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if kq.sql != nil {
		selector = kq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if kq.unique != nil && *kq.unique {
		selector.Distinct()
	}
	for _, m := range kq.modifiers {
		m(selector)
	}
	for _, p := range kq.predicates {
		p(selector)
	}
	for _, p := range kq.order {
		p(selector)
	}
	if offset := kq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := kq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (kq *KycQuery) ForUpdate(opts ...sql.LockOption) *KycQuery {
	if kq.driver.Dialect() == dialect.Postgres {
		kq.Unique(false)
	}
	kq.modifiers = append(kq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return kq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (kq *KycQuery) ForShare(opts ...sql.LockOption) *KycQuery {
	if kq.driver.Dialect() == dialect.Postgres {
		kq.Unique(false)
	}
	kq.modifiers = append(kq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return kq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (kq *KycQuery) Modify(modifiers ...func(s *sql.Selector)) *KycSelect {
	kq.modifiers = append(kq.modifiers, modifiers...)
	return kq.Select()
}

// KycGroupBy is the group-by builder for Kyc entities.
type KycGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (kgb *KycGroupBy) Aggregate(fns ...AggregateFunc) *KycGroupBy {
	kgb.fns = append(kgb.fns, fns...)
	return kgb
}

// Scan applies the group-by query and scans the result into the given value.
func (kgb *KycGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := kgb.path(ctx)
	if err != nil {
		return err
	}
	kgb.sql = query
	return kgb.sqlScan(ctx, v)
}

func (kgb *KycGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range kgb.fields {
		if !kyc.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := kgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := kgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (kgb *KycGroupBy) sqlQuery() *sql.Selector {
	selector := kgb.sql.Select()
	aggregation := make([]string, 0, len(kgb.fns))
	for _, fn := range kgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(kgb.fields)+len(kgb.fns))
		for _, f := range kgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(kgb.fields...)...)
}

// KycSelect is the builder for selecting fields of Kyc entities.
type KycSelect struct {
	*KycQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ks *KycSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ks.prepareQuery(ctx); err != nil {
		return err
	}
	ks.sql = ks.KycQuery.sqlQuery(ctx)
	return ks.sqlScan(ctx, v)
}

func (ks *KycSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ks.sql.Query()
	if err := ks.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ks *KycSelect) Modify(modifiers ...func(s *sql.Selector)) *KycSelect {
	ks.modifiers = append(ks.modifiers, modifiers...)
	return ks
}
