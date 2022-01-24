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
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/genesisuser"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// GenesisUserQuery is the builder for querying GenesisUser entities.
type GenesisUserQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GenesisUser
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GenesisUserQuery builder.
func (guq *GenesisUserQuery) Where(ps ...predicate.GenesisUser) *GenesisUserQuery {
	guq.predicates = append(guq.predicates, ps...)
	return guq
}

// Limit adds a limit step to the query.
func (guq *GenesisUserQuery) Limit(limit int) *GenesisUserQuery {
	guq.limit = &limit
	return guq
}

// Offset adds an offset step to the query.
func (guq *GenesisUserQuery) Offset(offset int) *GenesisUserQuery {
	guq.offset = &offset
	return guq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (guq *GenesisUserQuery) Unique(unique bool) *GenesisUserQuery {
	guq.unique = &unique
	return guq
}

// Order adds an order step to the query.
func (guq *GenesisUserQuery) Order(o ...OrderFunc) *GenesisUserQuery {
	guq.order = append(guq.order, o...)
	return guq
}

// First returns the first GenesisUser entity from the query.
// Returns a *NotFoundError when no GenesisUser was found.
func (guq *GenesisUserQuery) First(ctx context.Context) (*GenesisUser, error) {
	nodes, err := guq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{genesisuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (guq *GenesisUserQuery) FirstX(ctx context.Context) *GenesisUser {
	node, err := guq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GenesisUser ID from the query.
// Returns a *NotFoundError when no GenesisUser ID was found.
func (guq *GenesisUserQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = guq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{genesisuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (guq *GenesisUserQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := guq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GenesisUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one GenesisUser entity is not found.
// Returns a *NotFoundError when no GenesisUser entities are found.
func (guq *GenesisUserQuery) Only(ctx context.Context) (*GenesisUser, error) {
	nodes, err := guq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{genesisuser.Label}
	default:
		return nil, &NotSingularError{genesisuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (guq *GenesisUserQuery) OnlyX(ctx context.Context) *GenesisUser {
	node, err := guq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GenesisUser ID in the query.
// Returns a *NotSingularError when exactly one GenesisUser ID is not found.
// Returns a *NotFoundError when no entities are found.
func (guq *GenesisUserQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = guq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{genesisuser.Label}
	default:
		err = &NotSingularError{genesisuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (guq *GenesisUserQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := guq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GenesisUsers.
func (guq *GenesisUserQuery) All(ctx context.Context) ([]*GenesisUser, error) {
	if err := guq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return guq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (guq *GenesisUserQuery) AllX(ctx context.Context) []*GenesisUser {
	nodes, err := guq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GenesisUser IDs.
func (guq *GenesisUserQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := guq.Select(genesisuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (guq *GenesisUserQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := guq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (guq *GenesisUserQuery) Count(ctx context.Context) (int, error) {
	if err := guq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return guq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (guq *GenesisUserQuery) CountX(ctx context.Context) int {
	count, err := guq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (guq *GenesisUserQuery) Exist(ctx context.Context) (bool, error) {
	if err := guq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return guq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (guq *GenesisUserQuery) ExistX(ctx context.Context) bool {
	exist, err := guq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GenesisUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (guq *GenesisUserQuery) Clone() *GenesisUserQuery {
	if guq == nil {
		return nil
	}
	return &GenesisUserQuery{
		config:     guq.config,
		limit:      guq.limit,
		offset:     guq.offset,
		order:      append([]OrderFunc{}, guq.order...),
		predicates: append([]predicate.GenesisUser{}, guq.predicates...),
		// clone intermediate query.
		sql:  guq.sql.Clone(),
		path: guq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GenesisUser.Query().
//		GroupBy(genesisuser.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (guq *GenesisUserQuery) GroupBy(field string, fields ...string) *GenesisUserGroupBy {
	group := &GenesisUserGroupBy{config: guq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := guq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return guq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//	}
//
//	client.GenesisUser.Query().
//		Select(genesisuser.FieldUserID).
//		Scan(ctx, &v)
//
func (guq *GenesisUserQuery) Select(fields ...string) *GenesisUserSelect {
	guq.fields = append(guq.fields, fields...)
	return &GenesisUserSelect{GenesisUserQuery: guq}
}

func (guq *GenesisUserQuery) prepareQuery(ctx context.Context) error {
	for _, f := range guq.fields {
		if !genesisuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if guq.path != nil {
		prev, err := guq.path(ctx)
		if err != nil {
			return err
		}
		guq.sql = prev
	}
	return nil
}

func (guq *GenesisUserQuery) sqlAll(ctx context.Context) ([]*GenesisUser, error) {
	var (
		nodes = []*GenesisUser{}
		_spec = guq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &GenesisUser{config: guq.config}
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
	if err := sqlgraph.QueryNodes(ctx, guq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (guq *GenesisUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := guq.querySpec()
	_spec.Node.Columns = guq.fields
	if len(guq.fields) > 0 {
		_spec.Unique = guq.unique != nil && *guq.unique
	}
	return sqlgraph.CountNodes(ctx, guq.driver, _spec)
}

func (guq *GenesisUserQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := guq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (guq *GenesisUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   genesisuser.Table,
			Columns: genesisuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: genesisuser.FieldID,
			},
		},
		From:   guq.sql,
		Unique: true,
	}
	if unique := guq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := guq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, genesisuser.FieldID)
		for i := range fields {
			if fields[i] != genesisuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := guq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := guq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := guq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := guq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (guq *GenesisUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(guq.driver.Dialect())
	t1 := builder.Table(genesisuser.Table)
	columns := guq.fields
	if len(columns) == 0 {
		columns = genesisuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if guq.sql != nil {
		selector = guq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if guq.unique != nil && *guq.unique {
		selector.Distinct()
	}
	for _, p := range guq.predicates {
		p(selector)
	}
	for _, p := range guq.order {
		p(selector)
	}
	if offset := guq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := guq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GenesisUserGroupBy is the group-by builder for GenesisUser entities.
type GenesisUserGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gugb *GenesisUserGroupBy) Aggregate(fns ...AggregateFunc) *GenesisUserGroupBy {
	gugb.fns = append(gugb.fns, fns...)
	return gugb
}

// Scan applies the group-by query and scans the result into the given value.
func (gugb *GenesisUserGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gugb.path(ctx)
	if err != nil {
		return err
	}
	gugb.sql = query
	return gugb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gugb *GenesisUserGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gugb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (gugb *GenesisUserGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gugb.fields) > 1 {
		return nil, errors.New("ent: GenesisUserGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gugb *GenesisUserGroupBy) StringsX(ctx context.Context) []string {
	v, err := gugb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gugb *GenesisUserGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gugb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{genesisuser.Label}
	default:
		err = fmt.Errorf("ent: GenesisUserGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gugb *GenesisUserGroupBy) StringX(ctx context.Context) string {
	v, err := gugb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (gugb *GenesisUserGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gugb.fields) > 1 {
		return nil, errors.New("ent: GenesisUserGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gugb *GenesisUserGroupBy) IntsX(ctx context.Context) []int {
	v, err := gugb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gugb *GenesisUserGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gugb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{genesisuser.Label}
	default:
		err = fmt.Errorf("ent: GenesisUserGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gugb *GenesisUserGroupBy) IntX(ctx context.Context) int {
	v, err := gugb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (gugb *GenesisUserGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gugb.fields) > 1 {
		return nil, errors.New("ent: GenesisUserGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gugb *GenesisUserGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gugb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gugb *GenesisUserGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gugb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{genesisuser.Label}
	default:
		err = fmt.Errorf("ent: GenesisUserGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gugb *GenesisUserGroupBy) Float64X(ctx context.Context) float64 {
	v, err := gugb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (gugb *GenesisUserGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gugb.fields) > 1 {
		return nil, errors.New("ent: GenesisUserGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gugb *GenesisUserGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gugb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gugb *GenesisUserGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gugb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{genesisuser.Label}
	default:
		err = fmt.Errorf("ent: GenesisUserGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gugb *GenesisUserGroupBy) BoolX(ctx context.Context) bool {
	v, err := gugb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gugb *GenesisUserGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gugb.fields {
		if !genesisuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gugb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gugb *GenesisUserGroupBy) sqlQuery() *sql.Selector {
	selector := gugb.sql.Select()
	aggregation := make([]string, 0, len(gugb.fns))
	for _, fn := range gugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(gugb.fields)+len(gugb.fns))
		for _, f := range gugb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(gugb.fields...)...)
}

// GenesisUserSelect is the builder for selecting fields of GenesisUser entities.
type GenesisUserSelect struct {
	*GenesisUserQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gus *GenesisUserSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gus.prepareQuery(ctx); err != nil {
		return err
	}
	gus.sql = gus.GenesisUserQuery.sqlQuery(ctx)
	return gus.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gus *GenesisUserSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gus.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gus *GenesisUserSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gus.fields) > 1 {
		return nil, errors.New("ent: GenesisUserSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gus *GenesisUserSelect) StringsX(ctx context.Context) []string {
	v, err := gus.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gus *GenesisUserSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gus.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{genesisuser.Label}
	default:
		err = fmt.Errorf("ent: GenesisUserSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gus *GenesisUserSelect) StringX(ctx context.Context) string {
	v, err := gus.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gus *GenesisUserSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gus.fields) > 1 {
		return nil, errors.New("ent: GenesisUserSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gus *GenesisUserSelect) IntsX(ctx context.Context) []int {
	v, err := gus.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gus *GenesisUserSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gus.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{genesisuser.Label}
	default:
		err = fmt.Errorf("ent: GenesisUserSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gus *GenesisUserSelect) IntX(ctx context.Context) int {
	v, err := gus.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gus *GenesisUserSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gus.fields) > 1 {
		return nil, errors.New("ent: GenesisUserSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gus *GenesisUserSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gus.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gus *GenesisUserSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gus.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{genesisuser.Label}
	default:
		err = fmt.Errorf("ent: GenesisUserSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gus *GenesisUserSelect) Float64X(ctx context.Context) float64 {
	v, err := gus.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gus *GenesisUserSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gus.fields) > 1 {
		return nil, errors.New("ent: GenesisUserSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gus *GenesisUserSelect) BoolsX(ctx context.Context) []bool {
	v, err := gus.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gus *GenesisUserSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gus.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{genesisuser.Label}
	default:
		err = fmt.Errorf("ent: GenesisUserSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gus *GenesisUserSelect) BoolX(ctx context.Context) bool {
	v, err := gus.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gus *GenesisUserSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gus.sql.Query()
	if err := gus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}