// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/marko-durasic/olympics/ent/predicate"
	"github.com/marko-durasic/olympics/ent/teamsport"
)

// TeamSportQuery is the builder for querying TeamSport entities.
type TeamSportQuery struct {
	config
	ctx        *QueryContext
	order      []teamsport.OrderOption
	inters     []Interceptor
	predicates []predicate.TeamSport
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TeamSportQuery builder.
func (tsq *TeamSportQuery) Where(ps ...predicate.TeamSport) *TeamSportQuery {
	tsq.predicates = append(tsq.predicates, ps...)
	return tsq
}

// Limit the number of records to be returned by this query.
func (tsq *TeamSportQuery) Limit(limit int) *TeamSportQuery {
	tsq.ctx.Limit = &limit
	return tsq
}

// Offset to start from.
func (tsq *TeamSportQuery) Offset(offset int) *TeamSportQuery {
	tsq.ctx.Offset = &offset
	return tsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tsq *TeamSportQuery) Unique(unique bool) *TeamSportQuery {
	tsq.ctx.Unique = &unique
	return tsq
}

// Order specifies how the records should be ordered.
func (tsq *TeamSportQuery) Order(o ...teamsport.OrderOption) *TeamSportQuery {
	tsq.order = append(tsq.order, o...)
	return tsq
}

// First returns the first TeamSport entity from the query.
// Returns a *NotFoundError when no TeamSport was found.
func (tsq *TeamSportQuery) First(ctx context.Context) (*TeamSport, error) {
	nodes, err := tsq.Limit(1).All(setContextOp(ctx, tsq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{teamsport.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tsq *TeamSportQuery) FirstX(ctx context.Context) *TeamSport {
	node, err := tsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TeamSport ID from the query.
// Returns a *NotFoundError when no TeamSport ID was found.
func (tsq *TeamSportQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tsq.Limit(1).IDs(setContextOp(ctx, tsq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{teamsport.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tsq *TeamSportQuery) FirstIDX(ctx context.Context) int {
	id, err := tsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TeamSport entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TeamSport entity is found.
// Returns a *NotFoundError when no TeamSport entities are found.
func (tsq *TeamSportQuery) Only(ctx context.Context) (*TeamSport, error) {
	nodes, err := tsq.Limit(2).All(setContextOp(ctx, tsq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{teamsport.Label}
	default:
		return nil, &NotSingularError{teamsport.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tsq *TeamSportQuery) OnlyX(ctx context.Context) *TeamSport {
	node, err := tsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TeamSport ID in the query.
// Returns a *NotSingularError when more than one TeamSport ID is found.
// Returns a *NotFoundError when no entities are found.
func (tsq *TeamSportQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tsq.Limit(2).IDs(setContextOp(ctx, tsq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{teamsport.Label}
	default:
		err = &NotSingularError{teamsport.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tsq *TeamSportQuery) OnlyIDX(ctx context.Context) int {
	id, err := tsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TeamSports.
func (tsq *TeamSportQuery) All(ctx context.Context) ([]*TeamSport, error) {
	ctx = setContextOp(ctx, tsq.ctx, ent.OpQueryAll)
	if err := tsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TeamSport, *TeamSportQuery]()
	return withInterceptors[[]*TeamSport](ctx, tsq, qr, tsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tsq *TeamSportQuery) AllX(ctx context.Context) []*TeamSport {
	nodes, err := tsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TeamSport IDs.
func (tsq *TeamSportQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tsq.ctx.Unique == nil && tsq.path != nil {
		tsq.Unique(true)
	}
	ctx = setContextOp(ctx, tsq.ctx, ent.OpQueryIDs)
	if err = tsq.Select(teamsport.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tsq *TeamSportQuery) IDsX(ctx context.Context) []int {
	ids, err := tsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tsq *TeamSportQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tsq.ctx, ent.OpQueryCount)
	if err := tsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tsq, querierCount[*TeamSportQuery](), tsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tsq *TeamSportQuery) CountX(ctx context.Context) int {
	count, err := tsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tsq *TeamSportQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tsq.ctx, ent.OpQueryExist)
	switch _, err := tsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tsq *TeamSportQuery) ExistX(ctx context.Context) bool {
	exist, err := tsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TeamSportQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tsq *TeamSportQuery) Clone() *TeamSportQuery {
	if tsq == nil {
		return nil
	}
	return &TeamSportQuery{
		config:     tsq.config,
		ctx:        tsq.ctx.Clone(),
		order:      append([]teamsport.OrderOption{}, tsq.order...),
		inters:     append([]Interceptor{}, tsq.inters...),
		predicates: append([]predicate.TeamSport{}, tsq.predicates...),
		// clone intermediate query.
		sql:  tsq.sql.Clone(),
		path: tsq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Country string `json:"country,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TeamSport.Query().
//		GroupBy(teamsport.FieldCountry).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tsq *TeamSportQuery) GroupBy(field string, fields ...string) *TeamSportGroupBy {
	tsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TeamSportGroupBy{build: tsq}
	grbuild.flds = &tsq.ctx.Fields
	grbuild.label = teamsport.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Country string `json:"country,omitempty"`
//	}
//
//	client.TeamSport.Query().
//		Select(teamsport.FieldCountry).
//		Scan(ctx, &v)
func (tsq *TeamSportQuery) Select(fields ...string) *TeamSportSelect {
	tsq.ctx.Fields = append(tsq.ctx.Fields, fields...)
	sbuild := &TeamSportSelect{TeamSportQuery: tsq}
	sbuild.label = teamsport.Label
	sbuild.flds, sbuild.scan = &tsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TeamSportSelect configured with the given aggregations.
func (tsq *TeamSportQuery) Aggregate(fns ...AggregateFunc) *TeamSportSelect {
	return tsq.Select().Aggregate(fns...)
}

func (tsq *TeamSportQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tsq); err != nil {
				return err
			}
		}
	}
	for _, f := range tsq.ctx.Fields {
		if !teamsport.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tsq.path != nil {
		prev, err := tsq.path(ctx)
		if err != nil {
			return err
		}
		tsq.sql = prev
	}
	return nil
}

func (tsq *TeamSportQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TeamSport, error) {
	var (
		nodes = []*TeamSport{}
		_spec = tsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TeamSport).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TeamSport{config: tsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tsq *TeamSportQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tsq.querySpec()
	_spec.Node.Columns = tsq.ctx.Fields
	if len(tsq.ctx.Fields) > 0 {
		_spec.Unique = tsq.ctx.Unique != nil && *tsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tsq.driver, _spec)
}

func (tsq *TeamSportQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(teamsport.Table, teamsport.Columns, sqlgraph.NewFieldSpec(teamsport.FieldID, field.TypeInt))
	_spec.From = tsq.sql
	if unique := tsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tsq.path != nil {
		_spec.Unique = true
	}
	if fields := tsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, teamsport.FieldID)
		for i := range fields {
			if fields[i] != teamsport.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tsq *TeamSportQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tsq.driver.Dialect())
	t1 := builder.Table(teamsport.Table)
	columns := tsq.ctx.Fields
	if len(columns) == 0 {
		columns = teamsport.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tsq.sql != nil {
		selector = tsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tsq.ctx.Unique != nil && *tsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tsq.predicates {
		p(selector)
	}
	for _, p := range tsq.order {
		p(selector)
	}
	if offset := tsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TeamSportGroupBy is the group-by builder for TeamSport entities.
type TeamSportGroupBy struct {
	selector
	build *TeamSportQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tsgb *TeamSportGroupBy) Aggregate(fns ...AggregateFunc) *TeamSportGroupBy {
	tsgb.fns = append(tsgb.fns, fns...)
	return tsgb
}

// Scan applies the selector query and scans the result into the given value.
func (tsgb *TeamSportGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tsgb.build.ctx, ent.OpQueryGroupBy)
	if err := tsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TeamSportQuery, *TeamSportGroupBy](ctx, tsgb.build, tsgb, tsgb.build.inters, v)
}

func (tsgb *TeamSportGroupBy) sqlScan(ctx context.Context, root *TeamSportQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tsgb.fns))
	for _, fn := range tsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tsgb.flds)+len(tsgb.fns))
		for _, f := range *tsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TeamSportSelect is the builder for selecting fields of TeamSport entities.
type TeamSportSelect struct {
	*TeamSportQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tss *TeamSportSelect) Aggregate(fns ...AggregateFunc) *TeamSportSelect {
	tss.fns = append(tss.fns, fns...)
	return tss
}

// Scan applies the selector query and scans the result into the given value.
func (tss *TeamSportSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tss.ctx, ent.OpQuerySelect)
	if err := tss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TeamSportQuery, *TeamSportSelect](ctx, tss.TeamSportQuery, tss, tss.inters, v)
}

func (tss *TeamSportSelect) sqlScan(ctx context.Context, root *TeamSportQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tss.fns))
	for _, fn := range tss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
