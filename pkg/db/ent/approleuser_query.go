// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approleuser"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppRoleUserQuery is the builder for querying AppRoleUser entities.
type AppRoleUserQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppRoleUser
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppRoleUserQuery builder.
func (aruq *AppRoleUserQuery) Where(ps ...predicate.AppRoleUser) *AppRoleUserQuery {
	aruq.predicates = append(aruq.predicates, ps...)
	return aruq
}

// Limit adds a limit step to the query.
func (aruq *AppRoleUserQuery) Limit(limit int) *AppRoleUserQuery {
	aruq.limit = &limit
	return aruq
}

// Offset adds an offset step to the query.
func (aruq *AppRoleUserQuery) Offset(offset int) *AppRoleUserQuery {
	aruq.offset = &offset
	return aruq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aruq *AppRoleUserQuery) Unique(unique bool) *AppRoleUserQuery {
	aruq.unique = &unique
	return aruq
}

// Order adds an order step to the query.
func (aruq *AppRoleUserQuery) Order(o ...OrderFunc) *AppRoleUserQuery {
	aruq.order = append(aruq.order, o...)
	return aruq
}

// First returns the first AppRoleUser entity from the query.
// Returns a *NotFoundError when no AppRoleUser was found.
func (aruq *AppRoleUserQuery) First(ctx context.Context) (*AppRoleUser, error) {
	nodes, err := aruq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{approleuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aruq *AppRoleUserQuery) FirstX(ctx context.Context) *AppRoleUser {
	node, err := aruq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppRoleUser ID from the query.
// Returns a *NotFoundError when no AppRoleUser ID was found.
func (aruq *AppRoleUserQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aruq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{approleuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aruq *AppRoleUserQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := aruq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppRoleUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AppRoleUser entity is not found.
// Returns a *NotFoundError when no AppRoleUser entities are found.
func (aruq *AppRoleUserQuery) Only(ctx context.Context) (*AppRoleUser, error) {
	nodes, err := aruq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{approleuser.Label}
	default:
		return nil, &NotSingularError{approleuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aruq *AppRoleUserQuery) OnlyX(ctx context.Context) *AppRoleUser {
	node, err := aruq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppRoleUser ID in the query.
// Returns a *NotSingularError when exactly one AppRoleUser ID is not found.
// Returns a *NotFoundError when no entities are found.
func (aruq *AppRoleUserQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aruq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{approleuser.Label}
	default:
		err = &NotSingularError{approleuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aruq *AppRoleUserQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := aruq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppRoleUsers.
func (aruq *AppRoleUserQuery) All(ctx context.Context) ([]*AppRoleUser, error) {
	if err := aruq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aruq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aruq *AppRoleUserQuery) AllX(ctx context.Context) []*AppRoleUser {
	nodes, err := aruq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppRoleUser IDs.
func (aruq *AppRoleUserQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := aruq.Select(approleuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aruq *AppRoleUserQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := aruq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aruq *AppRoleUserQuery) Count(ctx context.Context) (int, error) {
	if err := aruq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aruq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aruq *AppRoleUserQuery) CountX(ctx context.Context) int {
	count, err := aruq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aruq *AppRoleUserQuery) Exist(ctx context.Context) (bool, error) {
	if err := aruq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aruq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aruq *AppRoleUserQuery) ExistX(ctx context.Context) bool {
	exist, err := aruq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppRoleUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aruq *AppRoleUserQuery) Clone() *AppRoleUserQuery {
	if aruq == nil {
		return nil
	}
	return &AppRoleUserQuery{
		config:     aruq.config,
		limit:      aruq.limit,
		offset:     aruq.offset,
		order:      append([]OrderFunc{}, aruq.order...),
		predicates: append([]predicate.AppRoleUser{}, aruq.predicates...),
		// clone intermediate query.
		sql:  aruq.sql.Clone(),
		path: aruq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AppRoleUser.Query().
//		GroupBy(approleuser.FieldAppID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (aruq *AppRoleUserQuery) GroupBy(field string, fields ...string) *AppRoleUserGroupBy {
	group := &AppRoleUserGroupBy{config: aruq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aruq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aruq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//	}
//
//	client.AppRoleUser.Query().
//		Select(approleuser.FieldAppID).
//		Scan(ctx, &v)
//
func (aruq *AppRoleUserQuery) Select(fields ...string) *AppRoleUserSelect {
	aruq.fields = append(aruq.fields, fields...)
	return &AppRoleUserSelect{AppRoleUserQuery: aruq}
}

func (aruq *AppRoleUserQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aruq.fields {
		if !approleuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aruq.path != nil {
		prev, err := aruq.path(ctx)
		if err != nil {
			return err
		}
		aruq.sql = prev
	}
	return nil
}

func (aruq *AppRoleUserQuery) sqlAll(ctx context.Context) ([]*AppRoleUser, error) {
	var (
		nodes = []*AppRoleUser{}
		_spec = aruq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AppRoleUser{config: aruq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, aruq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (aruq *AppRoleUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aruq.querySpec()
	_spec.Node.Columns = aruq.fields
	if len(aruq.fields) > 0 {
		_spec.Unique = aruq.unique != nil && *aruq.unique
	}
	return sqlgraph.CountNodes(ctx, aruq.driver, _spec)
}

func (aruq *AppRoleUserQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aruq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (aruq *AppRoleUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   approleuser.Table,
			Columns: approleuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: approleuser.FieldID,
			},
		},
		From:   aruq.sql,
		Unique: true,
	}
	if unique := aruq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aruq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, approleuser.FieldID)
		for i := range fields {
			if fields[i] != approleuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aruq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aruq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aruq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aruq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aruq *AppRoleUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aruq.driver.Dialect())
	t1 := builder.Table(approleuser.Table)
	columns := aruq.fields
	if len(columns) == 0 {
		columns = approleuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aruq.sql != nil {
		selector = aruq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aruq.unique != nil && *aruq.unique {
		selector.Distinct()
	}
	for _, p := range aruq.predicates {
		p(selector)
	}
	for _, p := range aruq.order {
		p(selector)
	}
	if offset := aruq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aruq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppRoleUserGroupBy is the group-by builder for AppRoleUser entities.
type AppRoleUserGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (arugb *AppRoleUserGroupBy) Aggregate(fns ...AggregateFunc) *AppRoleUserGroupBy {
	arugb.fns = append(arugb.fns, fns...)
	return arugb
}

// Scan applies the group-by query and scans the result into the given value.
func (arugb *AppRoleUserGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := arugb.path(ctx)
	if err != nil {
		return err
	}
	arugb.sql = query
	return arugb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (arugb *AppRoleUserGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := arugb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (arugb *AppRoleUserGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(arugb.fields) > 1 {
		return nil, errors.New("ent: AppRoleUserGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := arugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (arugb *AppRoleUserGroupBy) StringsX(ctx context.Context) []string {
	v, err := arugb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (arugb *AppRoleUserGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = arugb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approleuser.Label}
	default:
		err = fmt.Errorf("ent: AppRoleUserGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (arugb *AppRoleUserGroupBy) StringX(ctx context.Context) string {
	v, err := arugb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (arugb *AppRoleUserGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(arugb.fields) > 1 {
		return nil, errors.New("ent: AppRoleUserGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := arugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (arugb *AppRoleUserGroupBy) IntsX(ctx context.Context) []int {
	v, err := arugb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (arugb *AppRoleUserGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = arugb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approleuser.Label}
	default:
		err = fmt.Errorf("ent: AppRoleUserGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (arugb *AppRoleUserGroupBy) IntX(ctx context.Context) int {
	v, err := arugb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (arugb *AppRoleUserGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(arugb.fields) > 1 {
		return nil, errors.New("ent: AppRoleUserGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := arugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (arugb *AppRoleUserGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := arugb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (arugb *AppRoleUserGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = arugb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approleuser.Label}
	default:
		err = fmt.Errorf("ent: AppRoleUserGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (arugb *AppRoleUserGroupBy) Float64X(ctx context.Context) float64 {
	v, err := arugb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (arugb *AppRoleUserGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(arugb.fields) > 1 {
		return nil, errors.New("ent: AppRoleUserGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := arugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (arugb *AppRoleUserGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := arugb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (arugb *AppRoleUserGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = arugb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approleuser.Label}
	default:
		err = fmt.Errorf("ent: AppRoleUserGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (arugb *AppRoleUserGroupBy) BoolX(ctx context.Context) bool {
	v, err := arugb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (arugb *AppRoleUserGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range arugb.fields {
		if !approleuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := arugb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := arugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (arugb *AppRoleUserGroupBy) sqlQuery() *sql.Selector {
	selector := arugb.sql.Select()
	aggregation := make([]string, 0, len(arugb.fns))
	for _, fn := range arugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(arugb.fields)+len(arugb.fns))
		for _, f := range arugb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(arugb.fields...)...)
}

// AppRoleUserSelect is the builder for selecting fields of AppRoleUser entities.
type AppRoleUserSelect struct {
	*AppRoleUserQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (arus *AppRoleUserSelect) Scan(ctx context.Context, v interface{}) error {
	if err := arus.prepareQuery(ctx); err != nil {
		return err
	}
	arus.sql = arus.AppRoleUserQuery.sqlQuery(ctx)
	return arus.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (arus *AppRoleUserSelect) ScanX(ctx context.Context, v interface{}) {
	if err := arus.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (arus *AppRoleUserSelect) Strings(ctx context.Context) ([]string, error) {
	if len(arus.fields) > 1 {
		return nil, errors.New("ent: AppRoleUserSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := arus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (arus *AppRoleUserSelect) StringsX(ctx context.Context) []string {
	v, err := arus.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (arus *AppRoleUserSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = arus.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approleuser.Label}
	default:
		err = fmt.Errorf("ent: AppRoleUserSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (arus *AppRoleUserSelect) StringX(ctx context.Context) string {
	v, err := arus.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (arus *AppRoleUserSelect) Ints(ctx context.Context) ([]int, error) {
	if len(arus.fields) > 1 {
		return nil, errors.New("ent: AppRoleUserSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := arus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (arus *AppRoleUserSelect) IntsX(ctx context.Context) []int {
	v, err := arus.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (arus *AppRoleUserSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = arus.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approleuser.Label}
	default:
		err = fmt.Errorf("ent: AppRoleUserSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (arus *AppRoleUserSelect) IntX(ctx context.Context) int {
	v, err := arus.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (arus *AppRoleUserSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(arus.fields) > 1 {
		return nil, errors.New("ent: AppRoleUserSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := arus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (arus *AppRoleUserSelect) Float64sX(ctx context.Context) []float64 {
	v, err := arus.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (arus *AppRoleUserSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = arus.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approleuser.Label}
	default:
		err = fmt.Errorf("ent: AppRoleUserSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (arus *AppRoleUserSelect) Float64X(ctx context.Context) float64 {
	v, err := arus.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (arus *AppRoleUserSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(arus.fields) > 1 {
		return nil, errors.New("ent: AppRoleUserSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := arus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (arus *AppRoleUserSelect) BoolsX(ctx context.Context) []bool {
	v, err := arus.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (arus *AppRoleUserSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = arus.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{approleuser.Label}
	default:
		err = fmt.Errorf("ent: AppRoleUserSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (arus *AppRoleUserSelect) BoolX(ctx context.Context) bool {
	v, err := arus.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (arus *AppRoleUserSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := arus.sql.Query()
	if err := arus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
