// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/destinies"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/predicate"
	"github.com/Milkado/api-challenge-jornada-milhas/ent/testimonies"
)

// DestiniesQuery is the builder for querying Destinies entities.
type DestiniesQuery struct {
	config
	ctx             *QueryContext
	order           []destinies.OrderOption
	inters          []Interceptor
	predicates      []predicate.Destinies
	withTestimonies *TestimoniesQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DestiniesQuery builder.
func (dq *DestiniesQuery) Where(ps ...predicate.Destinies) *DestiniesQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DestiniesQuery) Limit(limit int) *DestiniesQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DestiniesQuery) Offset(offset int) *DestiniesQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DestiniesQuery) Unique(unique bool) *DestiniesQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DestiniesQuery) Order(o ...destinies.OrderOption) *DestiniesQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryTestimonies chains the current query on the "testimonies" edge.
func (dq *DestiniesQuery) QueryTestimonies() *TestimoniesQuery {
	query := (&TestimoniesClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(destinies.Table, destinies.FieldID, selector),
			sqlgraph.To(testimonies.Table, testimonies.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, destinies.TestimoniesTable, destinies.TestimoniesColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Destinies entity from the query.
// Returns a *NotFoundError when no Destinies was found.
func (dq *DestiniesQuery) First(ctx context.Context) (*Destinies, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{destinies.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DestiniesQuery) FirstX(ctx context.Context) *Destinies {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Destinies ID from the query.
// Returns a *NotFoundError when no Destinies ID was found.
func (dq *DestiniesQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{destinies.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DestiniesQuery) FirstIDX(ctx context.Context) int {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Destinies entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Destinies entity is found.
// Returns a *NotFoundError when no Destinies entities are found.
func (dq *DestiniesQuery) Only(ctx context.Context) (*Destinies, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{destinies.Label}
	default:
		return nil, &NotSingularError{destinies.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DestiniesQuery) OnlyX(ctx context.Context) *Destinies {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Destinies ID in the query.
// Returns a *NotSingularError when more than one Destinies ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DestiniesQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{destinies.Label}
	default:
		err = &NotSingularError{destinies.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DestiniesQuery) OnlyIDX(ctx context.Context) int {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DestiniesSlice.
func (dq *DestiniesQuery) All(ctx context.Context) ([]*Destinies, error) {
	ctx = setContextOp(ctx, dq.ctx, "All")
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Destinies, *DestiniesQuery]()
	return withInterceptors[[]*Destinies](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DestiniesQuery) AllX(ctx context.Context) []*Destinies {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Destinies IDs.
func (dq *DestiniesQuery) IDs(ctx context.Context) (ids []int, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, "IDs")
	if err = dq.Select(destinies.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DestiniesQuery) IDsX(ctx context.Context) []int {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DestiniesQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, "Count")
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DestiniesQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DestiniesQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DestiniesQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, "Exist")
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DestiniesQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DestiniesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DestiniesQuery) Clone() *DestiniesQuery {
	if dq == nil {
		return nil
	}
	return &DestiniesQuery{
		config:          dq.config,
		ctx:             dq.ctx.Clone(),
		order:           append([]destinies.OrderOption{}, dq.order...),
		inters:          append([]Interceptor{}, dq.inters...),
		predicates:      append([]predicate.Destinies{}, dq.predicates...),
		withTestimonies: dq.withTestimonies.Clone(),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// WithTestimonies tells the query-builder to eager-load the nodes that are connected to
// the "testimonies" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DestiniesQuery) WithTestimonies(opts ...func(*TestimoniesQuery)) *DestiniesQuery {
	query := (&TestimoniesClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withTestimonies = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Destinies.Query().
//		GroupBy(destinies.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DestiniesQuery) GroupBy(field string, fields ...string) *DestiniesGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DestiniesGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = destinies.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Destinies.Query().
//		Select(destinies.FieldName).
//		Scan(ctx, &v)
func (dq *DestiniesQuery) Select(fields ...string) *DestiniesSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DestiniesSelect{DestiniesQuery: dq}
	sbuild.label = destinies.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DestiniesSelect configured with the given aggregations.
func (dq *DestiniesQuery) Aggregate(fns ...AggregateFunc) *DestiniesSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DestiniesQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !destinies.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
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

func (dq *DestiniesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Destinies, error) {
	var (
		nodes       = []*Destinies{}
		_spec       = dq.querySpec()
		loadedTypes = [1]bool{
			dq.withTestimonies != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Destinies).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Destinies{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
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
	if query := dq.withTestimonies; query != nil {
		if err := dq.loadTestimonies(ctx, query, nodes,
			func(n *Destinies) { n.Edges.Testimonies = []*Testimonies{} },
			func(n *Destinies, e *Testimonies) { n.Edges.Testimonies = append(n.Edges.Testimonies, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DestiniesQuery) loadTestimonies(ctx context.Context, query *TestimoniesQuery, nodes []*Destinies, init func(*Destinies), assign func(*Destinies, *Testimonies)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Destinies)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(testimonies.FieldDestinyID)
	}
	query.Where(predicate.Testimonies(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(destinies.TestimoniesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.DestinyID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "destiny_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (dq *DestiniesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DestiniesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(destinies.Table, destinies.Columns, sqlgraph.NewFieldSpec(destinies.FieldID, field.TypeInt))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, destinies.FieldID)
		for i := range fields {
			if fields[i] != destinies.FieldID {
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
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
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

func (dq *DestiniesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(destinies.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = destinies.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DestiniesGroupBy is the group-by builder for Destinies entities.
type DestiniesGroupBy struct {
	selector
	build *DestiniesQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DestiniesGroupBy) Aggregate(fns ...AggregateFunc) *DestiniesGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DestiniesGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, "GroupBy")
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DestiniesQuery, *DestiniesGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DestiniesGroupBy) sqlScan(ctx context.Context, root *DestiniesQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DestiniesSelect is the builder for selecting fields of Destinies entities.
type DestiniesSelect struct {
	*DestiniesQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DestiniesSelect) Aggregate(fns ...AggregateFunc) *DestiniesSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DestiniesSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, "Select")
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DestiniesQuery, *DestiniesSelect](ctx, ds.DestiniesQuery, ds, ds.inters, v)
}

func (ds *DestiniesSelect) sqlScan(ctx context.Context, root *DestiniesQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
