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
	"github.com/poommin2543/app/ent/medicalequipment"
	"github.com/poommin2543/app/ent/predicate"
	"github.com/poommin2543/app/ent/systemequipment"
)

// MedicalequipmentQuery is the builder for querying Medicalequipment entities.
type MedicalequipmentQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Medicalequipment
	// eager-loading edges.
	withMedicalEquipment *SystemequipmentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (mq *MedicalequipmentQuery) Where(ps ...predicate.Medicalequipment) *MedicalequipmentQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit adds a limit step to the query.
func (mq *MedicalequipmentQuery) Limit(limit int) *MedicalequipmentQuery {
	mq.limit = &limit
	return mq
}

// Offset adds an offset step to the query.
func (mq *MedicalequipmentQuery) Offset(offset int) *MedicalequipmentQuery {
	mq.offset = &offset
	return mq
}

// Order adds an order step to the query.
func (mq *MedicalequipmentQuery) Order(o ...OrderFunc) *MedicalequipmentQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryMedicalEquipment chains the current query on the Medical_equipment edge.
func (mq *MedicalequipmentQuery) QueryMedicalEquipment() *SystemequipmentQuery {
	query := &SystemequipmentQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(medicalequipment.Table, medicalequipment.FieldID, mq.sqlQuery()),
			sqlgraph.To(systemequipment.Table, systemequipment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, medicalequipment.MedicalEquipmentTable, medicalequipment.MedicalEquipmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Medicalequipment entity in the query. Returns *NotFoundError when no medicalequipment was found.
func (mq *MedicalequipmentQuery) First(ctx context.Context) (*Medicalequipment, error) {
	ms, err := mq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ms) == 0 {
		return nil, &NotFoundError{medicalequipment.Label}
	}
	return ms[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MedicalequipmentQuery) FirstX(ctx context.Context) *Medicalequipment {
	m, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return m
}

// FirstID returns the first Medicalequipment id in the query. Returns *NotFoundError when no id was found.
func (mq *MedicalequipmentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{medicalequipment.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (mq *MedicalequipmentQuery) FirstXID(ctx context.Context) int {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Medicalequipment entity in the query, returns an error if not exactly one entity was returned.
func (mq *MedicalequipmentQuery) Only(ctx context.Context) (*Medicalequipment, error) {
	ms, err := mq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ms) {
	case 1:
		return ms[0], nil
	case 0:
		return nil, &NotFoundError{medicalequipment.Label}
	default:
		return nil, &NotSingularError{medicalequipment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MedicalequipmentQuery) OnlyX(ctx context.Context) *Medicalequipment {
	m, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return m
}

// OnlyID returns the only Medicalequipment id in the query, returns an error if not exactly one id was returned.
func (mq *MedicalequipmentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{medicalequipment.Label}
	default:
		err = &NotSingularError{medicalequipment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MedicalequipmentQuery) OnlyIDX(ctx context.Context) int {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Medicalequipments.
func (mq *MedicalequipmentQuery) All(ctx context.Context) ([]*Medicalequipment, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mq *MedicalequipmentQuery) AllX(ctx context.Context) []*Medicalequipment {
	ms, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ms
}

// IDs executes the query and returns a list of Medicalequipment ids.
func (mq *MedicalequipmentQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mq.Select(medicalequipment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MedicalequipmentQuery) IDsX(ctx context.Context) []int {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MedicalequipmentQuery) Count(ctx context.Context) (int, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MedicalequipmentQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MedicalequipmentQuery) Exist(ctx context.Context) (bool, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MedicalequipmentQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MedicalequipmentQuery) Clone() *MedicalequipmentQuery {
	return &MedicalequipmentQuery{
		config:     mq.config,
		limit:      mq.limit,
		offset:     mq.offset,
		order:      append([]OrderFunc{}, mq.order...),
		unique:     append([]string{}, mq.unique...),
		predicates: append([]predicate.Medicalequipment{}, mq.predicates...),
		// clone intermediate query.
		sql:  mq.sql.Clone(),
		path: mq.path,
	}
}

//  WithMedicalEquipment tells the query-builder to eager-loads the nodes that are connected to
// the "Medical_equipment" edge. The optional arguments used to configure the query builder of the edge.
func (mq *MedicalequipmentQuery) WithMedicalEquipment(opts ...func(*SystemequipmentQuery)) *MedicalequipmentQuery {
	query := &SystemequipmentQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withMedicalEquipment = query
	return mq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		MedicalID string `json:"Medical_ID,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Medicalequipment.Query().
//		GroupBy(medicalequipment.FieldMedicalID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mq *MedicalequipmentQuery) GroupBy(field string, fields ...string) *MedicalequipmentGroupBy {
	group := &MedicalequipmentGroupBy{config: mq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		MedicalID string `json:"Medical_ID,omitempty"`
//	}
//
//	client.Medicalequipment.Query().
//		Select(medicalequipment.FieldMedicalID).
//		Scan(ctx, &v)
//
func (mq *MedicalequipmentQuery) Select(field string, fields ...string) *MedicalequipmentSelect {
	selector := &MedicalequipmentSelect{config: mq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mq.sqlQuery(), nil
	}
	return selector
}

func (mq *MedicalequipmentQuery) prepareQuery(ctx context.Context) error {
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MedicalequipmentQuery) sqlAll(ctx context.Context) ([]*Medicalequipment, error) {
	var (
		nodes       = []*Medicalequipment{}
		_spec       = mq.querySpec()
		loadedTypes = [1]bool{
			mq.withMedicalEquipment != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Medicalequipment{config: mq.config}
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
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := mq.withMedicalEquipment; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Medicalequipment)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Systemequipment(func(s *sql.Selector) {
			s.Where(sql.InValues(medicalequipment.MedicalEquipmentColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.medicalequipment_medical_equipment
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "medicalequipment_medical_equipment" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "medicalequipment_medical_equipment" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.MedicalEquipment = append(node.Edges.MedicalEquipment, n)
		}
	}

	return nodes, nil
}

func (mq *MedicalequipmentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MedicalequipmentQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (mq *MedicalequipmentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   medicalequipment.Table,
			Columns: medicalequipment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: medicalequipment.FieldID,
			},
		},
		From:   mq.sql,
		Unique: true,
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MedicalequipmentQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(medicalequipment.Table)
	selector := builder.Select(t1.Columns(medicalequipment.Columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(medicalequipment.Columns...)...)
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MedicalequipmentGroupBy is the builder for group-by Medicalequipment entities.
type MedicalequipmentGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MedicalequipmentGroupBy) Aggregate(fns ...AggregateFunc) *MedicalequipmentGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the group-by query and scan the result into the given value.
func (mgb *MedicalequipmentGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mgb.path(ctx)
	if err != nil {
		return err
	}
	mgb.sql = query
	return mgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mgb *MedicalequipmentGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (mgb *MedicalequipmentGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MedicalequipmentGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mgb *MedicalequipmentGroupBy) StringsX(ctx context.Context) []string {
	v, err := mgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (mgb *MedicalequipmentGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalequipment.Label}
	default:
		err = fmt.Errorf("ent: MedicalequipmentGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mgb *MedicalequipmentGroupBy) StringX(ctx context.Context) string {
	v, err := mgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (mgb *MedicalequipmentGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MedicalequipmentGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mgb *MedicalequipmentGroupBy) IntsX(ctx context.Context) []int {
	v, err := mgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (mgb *MedicalequipmentGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalequipment.Label}
	default:
		err = fmt.Errorf("ent: MedicalequipmentGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mgb *MedicalequipmentGroupBy) IntX(ctx context.Context) int {
	v, err := mgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (mgb *MedicalequipmentGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MedicalequipmentGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mgb *MedicalequipmentGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (mgb *MedicalequipmentGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalequipment.Label}
	default:
		err = fmt.Errorf("ent: MedicalequipmentGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mgb *MedicalequipmentGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (mgb *MedicalequipmentGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MedicalequipmentGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mgb *MedicalequipmentGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (mgb *MedicalequipmentGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalequipment.Label}
	default:
		err = fmt.Errorf("ent: MedicalequipmentGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mgb *MedicalequipmentGroupBy) BoolX(ctx context.Context) bool {
	v, err := mgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mgb *MedicalequipmentGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mgb.sqlQuery().Query()
	if err := mgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mgb *MedicalequipmentGroupBy) sqlQuery() *sql.Selector {
	selector := mgb.sql
	columns := make([]string, 0, len(mgb.fields)+len(mgb.fns))
	columns = append(columns, mgb.fields...)
	for _, fn := range mgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(mgb.fields...)
}

// MedicalequipmentSelect is the builder for select fields of Medicalequipment entities.
type MedicalequipmentSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ms *MedicalequipmentSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ms.path(ctx)
	if err != nil {
		return err
	}
	ms.sql = query
	return ms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ms *MedicalequipmentSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ms *MedicalequipmentSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MedicalequipmentSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ms *MedicalequipmentSelect) StringsX(ctx context.Context) []string {
	v, err := ms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ms *MedicalequipmentSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalequipment.Label}
	default:
		err = fmt.Errorf("ent: MedicalequipmentSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ms *MedicalequipmentSelect) StringX(ctx context.Context) string {
	v, err := ms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ms *MedicalequipmentSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MedicalequipmentSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ms *MedicalequipmentSelect) IntsX(ctx context.Context) []int {
	v, err := ms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ms *MedicalequipmentSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalequipment.Label}
	default:
		err = fmt.Errorf("ent: MedicalequipmentSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ms *MedicalequipmentSelect) IntX(ctx context.Context) int {
	v, err := ms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ms *MedicalequipmentSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MedicalequipmentSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ms *MedicalequipmentSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ms *MedicalequipmentSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalequipment.Label}
	default:
		err = fmt.Errorf("ent: MedicalequipmentSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ms *MedicalequipmentSelect) Float64X(ctx context.Context) float64 {
	v, err := ms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ms *MedicalequipmentSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MedicalequipmentSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ms *MedicalequipmentSelect) BoolsX(ctx context.Context) []bool {
	v, err := ms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ms *MedicalequipmentSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{medicalequipment.Label}
	default:
		err = fmt.Errorf("ent: MedicalequipmentSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ms *MedicalequipmentSelect) BoolX(ctx context.Context) bool {
	v, err := ms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ms *MedicalequipmentSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ms.sqlQuery().Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ms *MedicalequipmentSelect) sqlQuery() sql.Querier {
	selector := ms.sql
	selector.Select(selector.Columns(ms.fields...)...)
	return selector
}
