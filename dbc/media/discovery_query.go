// Code generated by ent, DO NOT EDIT.

package media

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tikafog/of/dbc/media/discovery"
	"github.com/tikafog/of/dbc/media/predicate"
)

// DiscoveryQuery is the builder for querying Discovery entities.
type DiscoveryQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Discovery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DiscoveryQuery builder.
func (dq *DiscoveryQuery) Where(ps ...predicate.Discovery) *DiscoveryQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit adds a limit step to the query.
func (dq *DiscoveryQuery) Limit(limit int) *DiscoveryQuery {
	dq.limit = &limit
	return dq
}

// Offset adds an offset step to the query.
func (dq *DiscoveryQuery) Offset(offset int) *DiscoveryQuery {
	dq.offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DiscoveryQuery) Unique(unique bool) *DiscoveryQuery {
	dq.unique = &unique
	return dq
}

// Order adds an order step to the query.
func (dq *DiscoveryQuery) Order(o ...OrderFunc) *DiscoveryQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// First returns the first Discovery entity from the query.
// Returns a *NotFoundError when no Discovery was found.
func (dq *DiscoveryQuery) First(ctx context.Context) (*Discovery, error) {
	nodes, err := dq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{discovery.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DiscoveryQuery) FirstX(ctx context.Context) *Discovery {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Discovery ID from the query.
// Returns a *NotFoundError when no Discovery ID was found.
func (dq *DiscoveryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{discovery.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DiscoveryQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Discovery entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Discovery entity is found.
// Returns a *NotFoundError when no Discovery entities are found.
func (dq *DiscoveryQuery) Only(ctx context.Context) (*Discovery, error) {
	nodes, err := dq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{discovery.Label}
	default:
		return nil, &NotSingularError{discovery.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DiscoveryQuery) OnlyX(ctx context.Context) *Discovery {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Discovery ID in the query.
// Returns a *NotSingularError when more than one Discovery ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DiscoveryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{discovery.Label}
	default:
		err = &NotSingularError{discovery.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DiscoveryQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Discoveries.
func (dq *DiscoveryQuery) All(ctx context.Context) ([]*Discovery, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dq *DiscoveryQuery) AllX(ctx context.Context) []*Discovery {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Discovery IDs.
func (dq *DiscoveryQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := dq.Select(discovery.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DiscoveryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DiscoveryQuery) Count(ctx context.Context) (int, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DiscoveryQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DiscoveryQuery) Exist(ctx context.Context) (bool, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DiscoveryQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DiscoveryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DiscoveryQuery) Clone() *DiscoveryQuery {
	if dq == nil {
		return nil
	}
	return &DiscoveryQuery{
		config:     dq.config,
		limit:      dq.limit,
		offset:     dq.offset,
		order:      append([]OrderFunc{}, dq.order...),
		predicates: append([]predicate.Discovery{}, dq.predicates...),
		// clone intermediate query.
		sql:    dq.sql.Clone(),
		path:   dq.path,
		unique: dq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedUnix int64 `json:"created_unix,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Discovery.DB().
//		GroupBy(discovery.FieldCreatedUnix).
//		Aggregate(media.Count()).
//		Scan(ctx, &v)
func (dq *DiscoveryQuery) GroupBy(field string, fields ...string) *DiscoveryGroupBy {
	grbuild := &DiscoveryGroupBy{config: dq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dq.sqlQuery(ctx), nil
	}
	grbuild.label = discovery.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedUnix int64 `json:"created_unix,omitempty"`
//	}
//
//	client.Discovery.DB().
//		Select(discovery.FieldCreatedUnix).
//		Scan(ctx, &v)
func (dq *DiscoveryQuery) Select(fields ...string) *DiscoverySelect {
	dq.fields = append(dq.fields, fields...)
	selbuild := &DiscoverySelect{DiscoveryQuery: dq}
	selbuild.label = discovery.Label
	selbuild.flds, selbuild.scan = &dq.fields, selbuild.Scan
	return selbuild
}

func (dq *DiscoveryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dq.fields {
		if !discovery.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("media: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DiscoveryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Discovery, error) {
	var (
		nodes = []*Discovery{}
		_spec = dq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Discovery).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Discovery{config: dq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (dq *DiscoveryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	_spec.Node.Columns = dq.fields
	if len(dq.fields) > 0 {
		_spec.Unique = dq.unique != nil && *dq.unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DiscoveryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("media: check existence: %w", err)
	}
	return n > 0, nil
}

func (dq *DiscoveryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   discovery.Table,
			Columns: discovery.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: discovery.FieldID,
			},
		},
		From:   dq.sql,
		Unique: true,
	}
	if unique := dq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, discovery.FieldID)
		for i := range fields {
			if fields[i] != discovery.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DiscoveryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(discovery.Table)
	columns := dq.fields
	if len(columns) == 0 {
		columns = discovery.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.unique != nil && *dq.unique {
		selector.Distinct()
	}
	for _, m := range dq.modifiers {
		m(selector)
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (dq *DiscoveryQuery) ForUpdate(opts ...sql.LockOption) *DiscoveryQuery {
	if dq.driver.Dialect() == dialect.Postgres {
		dq.Unique(false)
	}
	dq.modifiers = append(dq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return dq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (dq *DiscoveryQuery) ForShare(opts ...sql.LockOption) *DiscoveryQuery {
	if dq.driver.Dialect() == dialect.Postgres {
		dq.Unique(false)
	}
	dq.modifiers = append(dq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return dq
}

// DiscoveryGroupBy is the group-by builder for Discovery entities.
type DiscoveryGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DiscoveryGroupBy) Aggregate(fns ...AggregateFunc) *DiscoveryGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dgb *DiscoveryGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dgb.path(ctx)
	if err != nil {
		return err
	}
	dgb.sql = query
	return dgb.sqlScan(ctx, v)
}

func (dgb *DiscoveryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dgb.fields {
		if !discovery.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dgb *DiscoveryGroupBy) sqlQuery() *sql.Selector {
	selector := dgb.sql.Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dgb.fields)+len(dgb.fns))
		for _, f := range dgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dgb.fields...)...)
}

// DiscoverySelect is the builder for selecting fields of Discovery entities.
type DiscoverySelect struct {
	*DiscoveryQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DiscoverySelect) Scan(ctx context.Context, v interface{}) error {
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	ds.sql = ds.DiscoveryQuery.sqlQuery(ctx)
	return ds.sqlScan(ctx, v)
}

func (ds *DiscoverySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ds.sql.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
