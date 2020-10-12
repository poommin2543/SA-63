// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/poommin2543/app/ent/medicaltype"
	"github.com/poommin2543/app/ent/systemequipment"
)

// MedicalTypeCreate is the builder for creating a MedicalType entity.
type MedicalTypeCreate struct {
	config
	mutation *MedicalTypeMutation
	hooks    []Hook
}

// SetName sets the name field.
func (mtc *MedicalTypeCreate) SetName(s string) *MedicalTypeCreate {
	mtc.mutation.SetName(s)
	return mtc
}

// AddSystemequipmentIDs adds the systemequipment edge to Systemequipment by ids.
func (mtc *MedicalTypeCreate) AddSystemequipmentIDs(ids ...int) *MedicalTypeCreate {
	mtc.mutation.AddSystemequipmentIDs(ids...)
	return mtc
}

// AddSystemequipment adds the systemequipment edges to Systemequipment.
func (mtc *MedicalTypeCreate) AddSystemequipment(s ...*Systemequipment) *MedicalTypeCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return mtc.AddSystemequipmentIDs(ids...)
}

// Mutation returns the MedicalTypeMutation object of the builder.
func (mtc *MedicalTypeCreate) Mutation() *MedicalTypeMutation {
	return mtc.mutation
}

// Save creates the MedicalType in the database.
func (mtc *MedicalTypeCreate) Save(ctx context.Context) (*MedicalType, error) {
	if _, ok := mtc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	var (
		err  error
		node *MedicalType
	)
	if len(mtc.hooks) == 0 {
		node, err = mtc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MedicalTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mtc.mutation = mutation
			node, err = mtc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mtc.hooks) - 1; i >= 0; i-- {
			mut = mtc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mtc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mtc *MedicalTypeCreate) SaveX(ctx context.Context) *MedicalType {
	v, err := mtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mtc *MedicalTypeCreate) sqlSave(ctx context.Context) (*MedicalType, error) {
	mt, _spec := mtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mtc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	mt.ID = int(id)
	return mt, nil
}

func (mtc *MedicalTypeCreate) createSpec() (*MedicalType, *sqlgraph.CreateSpec) {
	var (
		mt    = &MedicalType{config: mtc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: medicaltype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: medicaltype.FieldID,
			},
		}
	)
	if value, ok := mtc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicaltype.FieldName,
		})
		mt.Name = value
	}
	if nodes := mtc.mutation.SystemequipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicaltype.SystemequipmentTable,
			Columns: []string{medicaltype.SystemequipmentColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return mt, _spec
}
