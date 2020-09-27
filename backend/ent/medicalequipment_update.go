// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/poommin/app/ent/medicalequipment"
	"github.com/poommin/app/ent/predicate"
	"github.com/poommin/app/ent/systemequipment"
)

// MedicalequipmentUpdate is the builder for updating Medicalequipment entities.
type MedicalequipmentUpdate struct {
	config
	hooks      []Hook
	mutation   *MedicalequipmentMutation
	predicates []predicate.Medicalequipment
}

// Where adds a new predicate for the builder.
func (mu *MedicalequipmentUpdate) Where(ps ...predicate.Medicalequipment) *MedicalequipmentUpdate {
	mu.predicates = append(mu.predicates, ps...)
	return mu
}

// SetMedicalID sets the Medical_ID field.
func (mu *MedicalequipmentUpdate) SetMedicalID(s string) *MedicalequipmentUpdate {
	mu.mutation.SetMedicalID(s)
	return mu
}

// SetMedicalNAME sets the Medical_NAME field.
func (mu *MedicalequipmentUpdate) SetMedicalNAME(s string) *MedicalequipmentUpdate {
	mu.mutation.SetMedicalNAME(s)
	return mu
}

// SetMedicalStock sets the Medical_Stock field.
func (mu *MedicalequipmentUpdate) SetMedicalStock(i int) *MedicalequipmentUpdate {
	mu.mutation.ResetMedicalStock()
	mu.mutation.SetMedicalStock(i)
	return mu
}

// AddMedicalStock adds i to Medical_Stock.
func (mu *MedicalequipmentUpdate) AddMedicalStock(i int) *MedicalequipmentUpdate {
	mu.mutation.AddMedicalStock(i)
	return mu
}

// AddMedicalEquipmentIDs adds the Medical_equipment edge to Systemequipment by ids.
func (mu *MedicalequipmentUpdate) AddMedicalEquipmentIDs(ids ...int) *MedicalequipmentUpdate {
	mu.mutation.AddMedicalEquipmentIDs(ids...)
	return mu
}

// AddMedicalEquipment adds the Medical_equipment edges to Systemequipment.
func (mu *MedicalequipmentUpdate) AddMedicalEquipment(s ...*Systemequipment) *MedicalequipmentUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return mu.AddMedicalEquipmentIDs(ids...)
}

// Mutation returns the MedicalequipmentMutation object of the builder.
func (mu *MedicalequipmentUpdate) Mutation() *MedicalequipmentMutation {
	return mu.mutation
}

// RemoveMedicalEquipmentIDs removes the Medical_equipment edge to Systemequipment by ids.
func (mu *MedicalequipmentUpdate) RemoveMedicalEquipmentIDs(ids ...int) *MedicalequipmentUpdate {
	mu.mutation.RemoveMedicalEquipmentIDs(ids...)
	return mu
}

// RemoveMedicalEquipment removes Medical_equipment edges to Systemequipment.
func (mu *MedicalequipmentUpdate) RemoveMedicalEquipment(s ...*Systemequipment) *MedicalequipmentUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return mu.RemoveMedicalEquipmentIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (mu *MedicalequipmentUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MedicalequipmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MedicalequipmentUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MedicalequipmentUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MedicalequipmentUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MedicalequipmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   medicalequipment.Table,
			Columns: medicalequipment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: medicalequipment.FieldID,
			},
		},
	}
	if ps := mu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.MedicalID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicalequipment.FieldMedicalID,
		})
	}
	if value, ok := mu.mutation.MedicalNAME(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicalequipment.FieldMedicalNAME,
		})
	}
	if value, ok := mu.mutation.MedicalStock(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: medicalequipment.FieldMedicalStock,
		})
	}
	if value, ok := mu.mutation.AddedMedicalStock(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: medicalequipment.FieldMedicalStock,
		})
	}
	if nodes := mu.mutation.RemovedMedicalEquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicalequipment.MedicalEquipmentTable,
			Columns: []string{medicalequipment.MedicalEquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: systemequipment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MedicalEquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicalequipment.MedicalEquipmentTable,
			Columns: []string{medicalequipment.MedicalEquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: systemequipment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{medicalequipment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MedicalequipmentUpdateOne is the builder for updating a single Medicalequipment entity.
type MedicalequipmentUpdateOne struct {
	config
	hooks    []Hook
	mutation *MedicalequipmentMutation
}

// SetMedicalID sets the Medical_ID field.
func (muo *MedicalequipmentUpdateOne) SetMedicalID(s string) *MedicalequipmentUpdateOne {
	muo.mutation.SetMedicalID(s)
	return muo
}

// SetMedicalNAME sets the Medical_NAME field.
func (muo *MedicalequipmentUpdateOne) SetMedicalNAME(s string) *MedicalequipmentUpdateOne {
	muo.mutation.SetMedicalNAME(s)
	return muo
}

// SetMedicalStock sets the Medical_Stock field.
func (muo *MedicalequipmentUpdateOne) SetMedicalStock(i int) *MedicalequipmentUpdateOne {
	muo.mutation.ResetMedicalStock()
	muo.mutation.SetMedicalStock(i)
	return muo
}

// AddMedicalStock adds i to Medical_Stock.
func (muo *MedicalequipmentUpdateOne) AddMedicalStock(i int) *MedicalequipmentUpdateOne {
	muo.mutation.AddMedicalStock(i)
	return muo
}

// AddMedicalEquipmentIDs adds the Medical_equipment edge to Systemequipment by ids.
func (muo *MedicalequipmentUpdateOne) AddMedicalEquipmentIDs(ids ...int) *MedicalequipmentUpdateOne {
	muo.mutation.AddMedicalEquipmentIDs(ids...)
	return muo
}

// AddMedicalEquipment adds the Medical_equipment edges to Systemequipment.
func (muo *MedicalequipmentUpdateOne) AddMedicalEquipment(s ...*Systemequipment) *MedicalequipmentUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return muo.AddMedicalEquipmentIDs(ids...)
}

// Mutation returns the MedicalequipmentMutation object of the builder.
func (muo *MedicalequipmentUpdateOne) Mutation() *MedicalequipmentMutation {
	return muo.mutation
}

// RemoveMedicalEquipmentIDs removes the Medical_equipment edge to Systemequipment by ids.
func (muo *MedicalequipmentUpdateOne) RemoveMedicalEquipmentIDs(ids ...int) *MedicalequipmentUpdateOne {
	muo.mutation.RemoveMedicalEquipmentIDs(ids...)
	return muo
}

// RemoveMedicalEquipment removes Medical_equipment edges to Systemequipment.
func (muo *MedicalequipmentUpdateOne) RemoveMedicalEquipment(s ...*Systemequipment) *MedicalequipmentUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return muo.RemoveMedicalEquipmentIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (muo *MedicalequipmentUpdateOne) Save(ctx context.Context) (*Medicalequipment, error) {

	var (
		err  error
		node *Medicalequipment
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MedicalequipmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MedicalequipmentUpdateOne) SaveX(ctx context.Context) *Medicalequipment {
	m, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return m
}

// Exec executes the query on the entity.
func (muo *MedicalequipmentUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MedicalequipmentUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MedicalequipmentUpdateOne) sqlSave(ctx context.Context) (m *Medicalequipment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   medicalequipment.Table,
			Columns: medicalequipment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: medicalequipment.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Medicalequipment.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := muo.mutation.MedicalID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicalequipment.FieldMedicalID,
		})
	}
	if value, ok := muo.mutation.MedicalNAME(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicalequipment.FieldMedicalNAME,
		})
	}
	if value, ok := muo.mutation.MedicalStock(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: medicalequipment.FieldMedicalStock,
		})
	}
	if value, ok := muo.mutation.AddedMedicalStock(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: medicalequipment.FieldMedicalStock,
		})
	}
	if nodes := muo.mutation.RemovedMedicalEquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicalequipment.MedicalEquipmentTable,
			Columns: []string{medicalequipment.MedicalEquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: systemequipment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MedicalEquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicalequipment.MedicalEquipmentTable,
			Columns: []string{medicalequipment.MedicalEquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: systemequipment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	m = &Medicalequipment{config: muo.config}
	_spec.Assign = m.assignValues
	_spec.ScanValues = m.scanValues()
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{medicalequipment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return m, nil
}