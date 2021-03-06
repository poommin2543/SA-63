// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/poommin2543/app/ent/medicaltype"
	"github.com/poommin2543/app/ent/predicate"
	"github.com/poommin2543/app/ent/systemequipment"
)

// MedicalTypeQuery is the builder for querying MedicalType entities.
type MedicalTypeQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.MedicalType
	// eager-loading edges.
	withSystemequipment *SystemequipmentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (mtq *MedicalTypeQuery) Where(ps ...predicate.MedicalType) *MedicalTypeQuery {
	mtq.predicates = append(mtq.predicates, ps...)
	return mtq
}

// Limit adds a limit step to the query.
func (mtq *MedicalTypeQuery) Limit(limit int) *MedicalTypeQuery {
	mtq.limit = &limit
	return mtq
}

// Offset adds an offset step to the query.
func (mtq *MedicalTypeQuery) Offset(offset int) *MedicalTypeQuery {
	mtq.offset = &offset
	return mtq
}

// Order adds an order step to the query.
func (mtq *MedicalTypeQuery) Order(o ...OrderFunc) *MedicalTypeQuery {
	mtq.order = append(mtq.order, o...)
	return mtq
}

// QuerySystemequipment chains the current query on the systemequipment edge.
func (mtq *MedicalTypeQuery) QuerySystemequipment() *SystemequipmentQuery {
	query := &SystemequipmentQuery{config: mtq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(medicaltype.Table, medicaltype.FieldID, mtq.sqlQuery()),
			sqlgraph.To(systemequipment.Table, systemequipment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, medicaltype.SystemequipmentTable, medicaltype.SystemequipmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(mtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MedicalType entity in the query. Returns *NotFoundError when no medicaltype was found.
func (mtq *MedicalTypeQuery) First(ctx context.Context) (*MedicalType, error) {
	mts, err := mtq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(mts) == 0 {
		return nil, &NotFoundError{medicaltype.Label}
	}
	return mts[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mtq *MedicalTypeQuery) FirstX(ctx context.Context) *MedicalType {
	mt, err := mtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return mt
}

// FirstID returns the first MedicalType id in the query. Returns *NotFoundError when no id was found.
func (mtq *MedicalTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mtq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{medicaltype.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (mtq *MedicalTypeQuery) FirstXID(ctx context.Context) int {
	id, err := mtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only MedicalType entity in the query, returns an error if not exactly one entity was returned.
func (mtq *MedicalTypeQuery) Only(ctx context.Context) (*MedicalType, error) {
	mts, err := mtq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(mts) {
	case 1:
		return mts[0], nil
	case 0:
		return nil, &NotFoundError{medicaltype.Label}
	default:
		return nil, &NotSingularError{medicaltype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mtq *MedicalTypeQuery) OnlyX(ctx context.Context) *MedicalType {
	mt, err := mtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return mt
}

// OnlyID returns the only MedicalType id in the query, returns an error if not exactly one id was returned.
func (mtq *MedicalTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mtq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{medicaltype.Label}
	default:
		err = &NotSingularError{medicaltype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mtq *MedicalTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := mtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MedicalTypes.
func (mtq *MedicalTypeQuery) All(ctx context.Context) ([]*MedicalType, error) {
	if err := mtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mtq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mtq *MedicalTypeQuery) AllX(ctx context.Context) []*MedicalType {
	mts, err := mtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return mts
}

// IDs executes the query and returns a list of MedicalType ids.
func (mtq *MedicalTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mtq.Select(medicaltype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mtq *MedicalTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := mtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mtq *MedicalTypeQuery) Count(ctx context.Context) (int, error) {
	if err := mtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mtq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mtq *MedicalTypeQuery) CountX(ctx context.Context) int {
	count, err := mtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mtq *MedicalTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := mtq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mtq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mtq *MedicalTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := mtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mtq *MedicalTypeQuery) Clone() *MedicalTypeQuery {
	return &MedicalTypeQuery{
		config:     mtq.config,
		limit:      mtq.limit,
		offset:     mtq.offset,
		order:      append([]OrderFunc{}, mtq.order...),
		unique:     append([]string{}, mtq.unique...),
		predicates: append([]predicate.MedicalType{}, mtq.predicates...),
		// clone intermediate query.
		sql:  mtq.sql.Clone(),
		path: mtq.path,
	}
}

//  WithSystemequipment tells the query-builder to eager-loads the nodes that are connected to
// the "systemequipment" edge. The optional arguments used to configure the query builder of the edge.
func (mtq *MedicalTypeQuery) WithSystemequipment(opts ...func(*SystemequipmentQuery)) *MedicalTypeQuery {
	query := &SystemequipmentQuery{config: mtq.config}
	for _, opt := range opts {
		opt(query)
	}
	mtq.withSystemequipment = query
	return mtq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MedicalType.Query().
//		GroupBy(medicaltype.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mtq *MedicalTypeQuery) GroupBy(field string, fields ...string) *MedicalTypeGroupBy {
	group := &MedicalTypeGroupBy{config: mtq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mtq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.MedicalType.Query().
//		Select(medicaltype.FieldName).
//		Scan(ctx, &v)
//
func (mtq *MedicalTypeQuery) Select(field string, fields ...string) *MedicalTypeSelect {
	selector := &MedicalTypeSelect{config: mtq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mtq.sqlQuery(), nil
	}
	return selector
}

func (mtq *MedicalTypeQuery) prepareQuery(ctx context.Context) error {
	if mtq.path != nil {
		prev, err := mtq.path(ctx)
		if err != nil {
			return err
		}
		mtq.sql = prev
	}
	return nil
}

func (mtq *MedicalTypeQuery) sqlAll(ctx context.Context) ([]*MedicalType, error) {
	var (
		nodes       = []*MedicalType{}
		_spec       = mtq.querySpec()
		loadedTypes = [1]bool{
			mtq.withSystemequipment != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &MedicalType{config: mtq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, mtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := mtq.withSystemequipment; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*MedicalType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Systemequipment(func(s *sql.Selector) {
			s.Where(sql.InValues(medicaltype.SystemequipmentColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.medical_type_systemequipment
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "medical_type_systemequipment" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "medical_type_systemequipment" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Systemequipment = append(node.Edges.Systemequipment, n)
		}
	}

	return nodes, nil
}

func (mtq *MedicalTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mtq.querySpec()
	return sqlgraph.CountNodes(ctx, mtq.driver, _spec)
}

func (mtq *MedicalTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mtq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (mtq *MedicalTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   medicaltype.Table,
			Columns: medicaltype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: medicaltype.FieldID,
			},
		},
		From:   mtq.sql,
		Unique: true,
	}
	if ps := mtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mtq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mtq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mtq *MedicalTypeQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(mtq.driver.Dialect())
	t1 := builder.Table(medicaltype.Table)
	selector := builder.Select(t1.Columns(medicaltype.Columns...)...).From(t1)
	if mtq.sql != nil {
		selector = mtq.sql
		selector.Select(selector.Columns(medicaltype.Columns...)...)
	}
	for _, p := range mtq.predicates {
		p(selector)
	}
	for _, p := range mtq.order {
		p(selector)
	}
	if offset := mtq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mtq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MedicalTypeGroupBy is the builder for group-by MedicalType entities.
type MedicalTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mtgb *MedicalTypeGroupBy) Aggregate(fns ...AggregateFunc) *MedicalTypeGroupBy {
	mtgb.fns = append(mtgb.fns, fns...)
	return mtgb
}

// Scan applies the group-by query and scan the result into the given value.
func (mtgb *MedicalTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mtgb.path(ctx)
	if err != nil {
		return err
	}
	mtgb.sql = query
	return mtgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mtgb *MedicalTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mtgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (mtgb *MedicalTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mtgb.fields) > 1 {
		return nil, errors.New("ent: MedicalTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mtgb *MedicalTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := mtgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (mtgb *MedicalTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mtgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicaltype.Label}
	default:
		err = fmt.Errorf("ent: MedicalTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mtgb *MedicalTypeGroupBy) StringX(ctx context.Context) string {
	v, err := mtgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (mtgb *MedicalTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mtgb.fields) > 1 {
		return nil, errors.New("ent: MedicalTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mtgb *MedicalTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := mtgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (mtgb *MedicalTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mtgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicaltype.Label}
	default:
		err = fmt.Errorf("ent: MedicalTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mtgb *MedicalTypeGroupBy) IntX(ctx context.Context) int {
	v, err := mtgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (mtgb *MedicalTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mtgb.fields) > 1 {
		return nil, errors.New("ent: MedicalTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mtgb *MedicalTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mtgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (mtgb *MedicalTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mtgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicaltype.Label}
	default:
		err = fmt.Errorf("ent: MedicalTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mtgb *MedicalTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mtgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (mtgb *MedicalTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mtgb.fields) > 1 {
		return nil, errors.New("ent: MedicalTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mtgb *MedicalTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mtgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (mtgb *MedicalTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mtgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicaltype.Label}
	default:
		err = fmt.Errorf("ent: MedicalTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mtgb *MedicalTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := mtgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mtgb *MedicalTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mtgb.sqlQuery().Query()
	if err := mtgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mtgb *MedicalTypeGroupBy) sqlQuery() *sql.Selector {
	selector := mtgb.sql
	columns := make([]string, 0, len(mtgb.fields)+len(mtgb.fns))
	columns = append(columns, mtgb.fields...)
	for _, fn := range mtgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(mtgb.fields...)
}

// MedicalTypeSelect is the builder for select fields of MedicalType entities.
type MedicalTypeSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (mts *MedicalTypeSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := mts.path(ctx)
	if err != nil {
		return err
	}
	mts.sql = query
	return mts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mts *MedicalTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (mts *MedicalTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mts.fields) > 1 {
		return nil, errors.New("ent: MedicalTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mts *MedicalTypeSelect) StringsX(ctx context.Context) []string {
	v, err := mts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (mts *MedicalTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicaltype.Label}
	default:
		err = fmt.Errorf("ent: MedicalTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mts *MedicalTypeSelect) StringX(ctx context.Context) string {
	v, err := mts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (mts *MedicalTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mts.fields) > 1 {
		return nil, errors.New("ent: MedicalTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mts *MedicalTypeSelect) IntsX(ctx context.Context) []int {
	v, err := mts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (mts *MedicalTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicaltype.Label}
	default:
		err = fmt.Errorf("ent: MedicalTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mts *MedicalTypeSelect) IntX(ctx context.Context) int {
	v, err := mts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (mts *MedicalTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mts.fields) > 1 {
		return nil, errors.New("ent: MedicalTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mts *MedicalTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (mts *MedicalTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicaltype.Label}
	default:
		err = fmt.Errorf("ent: MedicalTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mts *MedicalTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := mts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (mts *MedicalTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mts.fields) > 1 {
		return nil, errors.New("ent: MedicalTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mts *MedicalTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := mts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (mts *MedicalTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicaltype.Label}
	default:
		err = fmt.Errorf("ent: MedicalTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mts *MedicalTypeSelect) BoolX(ctx context.Context) bool {
	v, err := mts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mts *MedicalTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mts.sqlQuery().Query()
	if err := mts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mts *MedicalTypeSelect) sqlQuery() sql.Querier {
	selector := mts.sql
	selector.Select(selector.Columns(mts.fields...)...)
	return selector
}
