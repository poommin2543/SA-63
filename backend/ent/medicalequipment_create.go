// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/poommin/app/ent/medicalequipment"
	"github.com/poommin/app/ent/systemequipment"
)

// MedicalequipmentCreate is the builder for creating a Medicalequipment entity.
type MedicalequipmentCreate struct {
	config
	mutation *MedicalequipmentMutation
	hooks    []Hook
}

// SetMedicalID sets the Medical_ID field.
func (mc *MedicalequipmentCreate) SetMedicalID(s string) *MedicalequipmentCreate {
	mc.mutation.SetMedicalID(s)
	return mc
}

// SetMedicalNAME sets the Medical_NAME field.
func (mc *MedicalequipmentCreate) SetMedicalNAME(s string) *MedicalequipmentCreate {
	mc.mutation.SetMedicalNAME(s)
	return mc
}

// SetMedicalStock sets the Medical_Stock field.
func (mc *MedicalequipmentCreate) SetMedicalStock(i int) *MedicalequipmentCreate {
	mc.mutation.SetMedicalStock(i)
	return mc
}

// AddMedicalEquipmentIDs adds the Medical_equipment edge to Systemequipment by ids.
func (mc *MedicalequipmentCreate) AddMedicalEquipmentIDs(ids ...int) *MedicalequipmentCreate {
	mc.mutation.AddMedicalEquipmentIDs(ids...)
	return mc
}

// AddMedicalEquipment adds the Medical_equipment edges to Systemequipment.
func (mc *MedicalequipmentCreate) AddMedicalEquipment(s ...*Systemequipment) *MedicalequipmentCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return mc.AddMedicalEquipmentIDs(ids...)
}

// Mutation returns the MedicalequipmentMutation object of the builder.
func (mc *MedicalequipmentCreate) Mutation() *MedicalequipmentMutation {
	return mc.mutation
}

// Save creates the Medicalequipment in the database.
func (mc *MedicalequipmentCreate) Save(ctx context.Context) (*Medicalequipment, error) {
	if _, ok := mc.mutation.MedicalID(); !ok {
		return nil, &ValidationError{Name: "Medical_ID", err: errors.New("ent: missing required field \"Medical_ID\"")}
	}
	if _, ok := mc.mutation.MedicalNAME(); !ok {
		return nil, &ValidationError{Name: "Medical_NAME", err: errors.New("ent: missing required field \"Medical_NAME\"")}
	}
	if _, ok := mc.mutation.MedicalStock(); !ok {
		return nil, &ValidationError{Name: "Medical_Stock", err: errors.New("ent: missing required field \"Medical_Stock\"")}
	}
	var (
		err  error
		node *Medicalequipment
	)
	if len(mc.hooks) == 0 {
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MedicalequipmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mc.mutation = mutation
			node, err = mc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MedicalequipmentCreate) SaveX(ctx context.Context) *Medicalequipment {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mc *MedicalequipmentCreate) sqlSave(ctx context.Context) (*Medicalequipment, error) {
	m, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	m.ID = int(id)
	return m, nil
}

func (mc *MedicalequipmentCreate) createSpec() (*Medicalequipment, *sqlgraph.CreateSpec) {
	var (
		m     = &Medicalequipment{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: medicalequipment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: medicalequipment.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.MedicalID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicalequipment.FieldMedicalID,
		})
		m.MedicalID = value
	}
	if value, ok := mc.mutation.MedicalNAME(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicalequipment.FieldMedicalNAME,
		})
		m.MedicalNAME = value
	}
	if value, ok := mc.mutation.MedicalStock(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: medicalequipment.FieldMedicalStock,
		})
		m.MedicalStock = value
	}
	if nodes := mc.mutation.MedicalEquipmentIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return m, _spec
}