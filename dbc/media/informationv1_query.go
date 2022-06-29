// Code generated by entc, DO NOT EDIT.

package media

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tikafog/of/dbc/media/channel"
	"github.com/tikafog/of/dbc/media/informationv1"
	"github.com/tikafog/of/dbc/media/predicate"
	"github.com/tikafog/of/dbc/media/toplist"
)

// InformationV1Query is the builder for querying InformationV1 entities.
type InformationV1Query struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.InformationV1
	// eager-loading edges.
	withTopLists *TopListQuery
	withChannel  *ChannelQuery
	modifiers    []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InformationV1Query builder.
func (iv *InformationV1Query) Where(ps ...predicate.InformationV1) *InformationV1Query {
	iv.predicates = append(iv.predicates, ps...)
	return iv
}

// Limit adds a limit step to the query.
func (iv *InformationV1Query) Limit(limit int) *InformationV1Query {
	iv.limit = &limit
	return iv
}

// Offset adds an offset step to the query.
func (iv *InformationV1Query) Offset(offset int) *InformationV1Query {
	iv.offset = &offset
	return iv
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iv *InformationV1Query) Unique(unique bool) *InformationV1Query {
	iv.unique = &unique
	return iv
}

// Order adds an order step to the query.
func (iv *InformationV1Query) Order(o ...OrderFunc) *InformationV1Query {
	iv.order = append(iv.order, o...)
	return iv
}

// QueryTopLists chains the current query on the "top_lists" edge.
func (iv *InformationV1Query) QueryTopLists() *TopListQuery {
	query := &TopListQuery{config: iv.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iv.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iv.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(informationv1.Table, informationv1.FieldID, selector),
			sqlgraph.To(toplist.Table, toplist.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, informationv1.TopListsTable, informationv1.TopListsColumn),
		)
		fromU = sqlgraph.SetNeighbors(iv.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChannel chains the current query on the "channel" edge.
func (iv *InformationV1Query) QueryChannel() *ChannelQuery {
	query := &ChannelQuery{config: iv.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iv.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iv.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(informationv1.Table, informationv1.FieldID, selector),
			sqlgraph.To(channel.Table, channel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, informationv1.ChannelTable, informationv1.ChannelColumn),
		)
		fromU = sqlgraph.SetNeighbors(iv.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first InformationV1 entity from the query.
// Returns a *NotFoundError when no InformationV1 was found.
func (iv *InformationV1Query) First(ctx context.Context) (*InformationV1, error) {
	nodes, err := iv.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{informationv1.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iv *InformationV1Query) FirstX(ctx context.Context) *InformationV1 {
	node, err := iv.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first InformationV1 ID from the query.
// Returns a *NotFoundError when no InformationV1 ID was found.
func (iv *InformationV1Query) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iv.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{informationv1.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iv *InformationV1Query) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := iv.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single InformationV1 entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one InformationV1 entity is found.
// Returns a *NotFoundError when no InformationV1 entities are found.
func (iv *InformationV1Query) Only(ctx context.Context) (*InformationV1, error) {
	nodes, err := iv.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{informationv1.Label}
	default:
		return nil, &NotSingularError{informationv1.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iv *InformationV1Query) OnlyX(ctx context.Context) *InformationV1 {
	node, err := iv.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only InformationV1 ID in the query.
// Returns a *NotSingularError when more than one InformationV1 ID is found.
// Returns a *NotFoundError when no entities are found.
func (iv *InformationV1Query) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iv.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{informationv1.Label}
	default:
		err = &NotSingularError{informationv1.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iv *InformationV1Query) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := iv.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of InformationV1s.
func (iv *InformationV1Query) All(ctx context.Context) ([]*InformationV1, error) {
	if err := iv.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return iv.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (iv *InformationV1Query) AllX(ctx context.Context) []*InformationV1 {
	nodes, err := iv.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of InformationV1 IDs.
func (iv *InformationV1Query) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := iv.Select(informationv1.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iv *InformationV1Query) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := iv.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iv *InformationV1Query) Count(ctx context.Context) (int, error) {
	if err := iv.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return iv.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (iv *InformationV1Query) CountX(ctx context.Context) int {
	count, err := iv.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iv *InformationV1Query) Exist(ctx context.Context) (bool, error) {
	if err := iv.prepareQuery(ctx); err != nil {
		return false, err
	}
	return iv.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (iv *InformationV1Query) ExistX(ctx context.Context) bool {
	exist, err := iv.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InformationV1Query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iv *InformationV1Query) Clone() *InformationV1Query {
	if iv == nil {
		return nil
	}
	return &InformationV1Query{
		config:       iv.config,
		limit:        iv.limit,
		offset:       iv.offset,
		order:        append([]OrderFunc{}, iv.order...),
		predicates:   append([]predicate.InformationV1{}, iv.predicates...),
		withTopLists: iv.withTopLists.Clone(),
		withChannel:  iv.withChannel.Clone(),
		// clone intermediate query.
		sql:    iv.sql.Clone(),
		path:   iv.path,
		unique: iv.unique,
	}
}

// WithTopLists tells the query-builder to eager-load the nodes that are connected to
// the "top_lists" edge. The optional arguments are used to configure the query builder of the edge.
func (iv *InformationV1Query) WithTopLists(opts ...func(*TopListQuery)) *InformationV1Query {
	query := &TopListQuery{config: iv.config}
	for _, opt := range opts {
		opt(query)
	}
	iv.withTopLists = query
	return iv
}

// WithChannel tells the query-builder to eager-load the nodes that are connected to
// the "channel" edge. The optional arguments are used to configure the query builder of the edge.
func (iv *InformationV1Query) WithChannel(opts ...func(*ChannelQuery)) *InformationV1Query {
	query := &ChannelQuery{config: iv.config}
	for _, opt := range opts {
		opt(query)
	}
	iv.withChannel = query
	return iv
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
//	client.InformationV1.Query().
//		GroupBy(informationv1.FieldCreatedUnix).
//		Aggregate(media.Count()).
//		Scan(ctx, &v)
//
func (iv *InformationV1Query) GroupBy(field string, fields ...string) *InformationV1GroupBy {
	group := &InformationV1GroupBy{config: iv.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := iv.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return iv.sqlQuery(ctx), nil
	}
	return group
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
//	client.InformationV1.Query().
//		Select(informationv1.FieldCreatedUnix).
//		Scan(ctx, &v)
//
func (iv *InformationV1Query) Select(fields ...string) *InformationV1Select {
	iv.fields = append(iv.fields, fields...)
	return &InformationV1Select{InformationV1Query: iv}
}

func (iv *InformationV1Query) prepareQuery(ctx context.Context) error {
	for _, f := range iv.fields {
		if !informationv1.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("media: invalid field %q for query", f)}
		}
	}
	if iv.path != nil {
		prev, err := iv.path(ctx)
		if err != nil {
			return err
		}
		iv.sql = prev
	}
	return nil
}

func (iv *InformationV1Query) sqlAll(ctx context.Context) ([]*InformationV1, error) {
	var (
		nodes       = []*InformationV1{}
		_spec       = iv.querySpec()
		loadedTypes = [2]bool{
			iv.withTopLists != nil,
			iv.withChannel != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &InformationV1{config: iv.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("media: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(iv.modifiers) > 0 {
		_spec.Modifiers = iv.modifiers
	}
	if err := sqlgraph.QueryNodes(ctx, iv.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := iv.withTopLists; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*InformationV1)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.TopLists = []*TopList{}
		}
		query.Where(predicate.TopList(func(s *sql.Selector) {
			s.Where(sql.InValues(informationv1.TopListsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.InformationID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "information_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.TopLists = append(node.Edges.TopLists, n)
		}
	}

	if query := iv.withChannel; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*InformationV1)
		for i := range nodes {
			fk := nodes[i].ChannelID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(channel.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "channel_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Channel = n
			}
		}
	}

	return nodes, nil
}

func (iv *InformationV1Query) sqlCount(ctx context.Context) (int, error) {
	_spec := iv.querySpec()
	if len(iv.modifiers) > 0 {
		_spec.Modifiers = iv.modifiers
	}
	_spec.Node.Columns = iv.fields
	if len(iv.fields) > 0 {
		_spec.Unique = iv.unique != nil && *iv.unique
	}
	return sqlgraph.CountNodes(ctx, iv.driver, _spec)
}

func (iv *InformationV1Query) sqlExist(ctx context.Context) (bool, error) {
	n, err := iv.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("media: check existence: %w", err)
	}
	return n > 0, nil
}

func (iv *InformationV1Query) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   informationv1.Table,
			Columns: informationv1.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: informationv1.FieldID,
			},
		},
		From:   iv.sql,
		Unique: true,
	}
	if unique := iv.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := iv.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, informationv1.FieldID)
		for i := range fields {
			if fields[i] != informationv1.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iv.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iv.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iv.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iv.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iv *InformationV1Query) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iv.driver.Dialect())
	t1 := builder.Table(informationv1.Table)
	columns := iv.fields
	if len(columns) == 0 {
		columns = informationv1.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iv.sql != nil {
		selector = iv.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iv.unique != nil && *iv.unique {
		selector.Distinct()
	}
	for _, m := range iv.modifiers {
		m(selector)
	}
	for _, p := range iv.predicates {
		p(selector)
	}
	for _, p := range iv.order {
		p(selector)
	}
	if offset := iv.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iv.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (iv *InformationV1Query) ForUpdate(opts ...sql.LockOption) *InformationV1Query {
	if iv.driver.Dialect() == dialect.Postgres {
		iv.Unique(false)
	}
	iv.modifiers = append(iv.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return iv
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (iv *InformationV1Query) ForShare(opts ...sql.LockOption) *InformationV1Query {
	if iv.driver.Dialect() == dialect.Postgres {
		iv.Unique(false)
	}
	iv.modifiers = append(iv.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return iv
}

// InformationV1GroupBy is the group-by builder for InformationV1 entities.
type InformationV1GroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ivb *InformationV1GroupBy) Aggregate(fns ...AggregateFunc) *InformationV1GroupBy {
	ivb.fns = append(ivb.fns, fns...)
	return ivb
}

// Scan applies the group-by query and scans the result into the given value.
func (ivb *InformationV1GroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ivb.path(ctx)
	if err != nil {
		return err
	}
	ivb.sql = query
	return ivb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ivb *InformationV1GroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ivb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ivb *InformationV1GroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ivb.fields) > 1 {
		return nil, errors.New("media: InformationV1GroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ivb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ivb *InformationV1GroupBy) StringsX(ctx context.Context) []string {
	v, err := ivb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ivb *InformationV1GroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ivb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{informationv1.Label}
	default:
		err = fmt.Errorf("media: InformationV1GroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ivb *InformationV1GroupBy) StringX(ctx context.Context) string {
	v, err := ivb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ivb *InformationV1GroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ivb.fields) > 1 {
		return nil, errors.New("media: InformationV1GroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ivb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ivb *InformationV1GroupBy) IntsX(ctx context.Context) []int {
	v, err := ivb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ivb *InformationV1GroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ivb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{informationv1.Label}
	default:
		err = fmt.Errorf("media: InformationV1GroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ivb *InformationV1GroupBy) IntX(ctx context.Context) int {
	v, err := ivb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ivb *InformationV1GroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ivb.fields) > 1 {
		return nil, errors.New("media: InformationV1GroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ivb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ivb *InformationV1GroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ivb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ivb *InformationV1GroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ivb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{informationv1.Label}
	default:
		err = fmt.Errorf("media: InformationV1GroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ivb *InformationV1GroupBy) Float64X(ctx context.Context) float64 {
	v, err := ivb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ivb *InformationV1GroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ivb.fields) > 1 {
		return nil, errors.New("media: InformationV1GroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ivb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ivb *InformationV1GroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ivb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ivb *InformationV1GroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ivb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{informationv1.Label}
	default:
		err = fmt.Errorf("media: InformationV1GroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ivb *InformationV1GroupBy) BoolX(ctx context.Context) bool {
	v, err := ivb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ivb *InformationV1GroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ivb.fields {
		if !informationv1.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ivb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ivb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ivb *InformationV1GroupBy) sqlQuery() *sql.Selector {
	selector := ivb.sql.Select()
	aggregation := make([]string, 0, len(ivb.fns))
	for _, fn := range ivb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ivb.fields)+len(ivb.fns))
		for _, f := range ivb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ivb.fields...)...)
}

// InformationV1Select is the builder for selecting fields of InformationV1 entities.
type InformationV1Select struct {
	*InformationV1Query
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (iv *InformationV1Select) Scan(ctx context.Context, v interface{}) error {
	if err := iv.prepareQuery(ctx); err != nil {
		return err
	}
	iv.sql = iv.InformationV1Query.sqlQuery(ctx)
	return iv.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (iv *InformationV1Select) ScanX(ctx context.Context, v interface{}) {
	if err := iv.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (iv *InformationV1Select) Strings(ctx context.Context) ([]string, error) {
	if len(iv.fields) > 1 {
		return nil, errors.New("media: InformationV1Select.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := iv.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (iv *InformationV1Select) StringsX(ctx context.Context) []string {
	v, err := iv.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (iv *InformationV1Select) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = iv.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{informationv1.Label}
	default:
		err = fmt.Errorf("media: InformationV1Select.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (iv *InformationV1Select) StringX(ctx context.Context) string {
	v, err := iv.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (iv *InformationV1Select) Ints(ctx context.Context) ([]int, error) {
	if len(iv.fields) > 1 {
		return nil, errors.New("media: InformationV1Select.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := iv.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (iv *InformationV1Select) IntsX(ctx context.Context) []int {
	v, err := iv.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (iv *InformationV1Select) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = iv.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{informationv1.Label}
	default:
		err = fmt.Errorf("media: InformationV1Select.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (iv *InformationV1Select) IntX(ctx context.Context) int {
	v, err := iv.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (iv *InformationV1Select) Float64s(ctx context.Context) ([]float64, error) {
	if len(iv.fields) > 1 {
		return nil, errors.New("media: InformationV1Select.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := iv.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (iv *InformationV1Select) Float64sX(ctx context.Context) []float64 {
	v, err := iv.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (iv *InformationV1Select) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = iv.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{informationv1.Label}
	default:
		err = fmt.Errorf("media: InformationV1Select.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (iv *InformationV1Select) Float64X(ctx context.Context) float64 {
	v, err := iv.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (iv *InformationV1Select) Bools(ctx context.Context) ([]bool, error) {
	if len(iv.fields) > 1 {
		return nil, errors.New("media: InformationV1Select.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := iv.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (iv *InformationV1Select) BoolsX(ctx context.Context) []bool {
	v, err := iv.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (iv *InformationV1Select) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = iv.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{informationv1.Label}
	default:
		err = fmt.Errorf("media: InformationV1Select.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (iv *InformationV1Select) BoolX(ctx context.Context) bool {
	v, err := iv.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (iv *InformationV1Select) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := iv.sql.Query()
	if err := iv.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
