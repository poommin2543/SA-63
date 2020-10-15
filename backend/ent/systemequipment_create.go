// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/poommin2543/app/ent/medicalequipment"
	"github.com/poommin2543/app/ent/medicaltype"
	"github.com/poommin2543/app/ent/physician"
	"github.com/poommin2543/app/ent/systemequipment"
)

// SystemequipmentCreate is the builder for creating a Systemequipment entity.
type SystemequipmentCreate struct {
	config
	mutation *SystemequipmentMutation
	hooks    []Hook
}

// SetAddedtime sets the addedtime field.
func (sc *SystemequipmentCreate) SetAddedtime(s string) *SystemequipmentCreate {
	sc.mutation.SetAddedtime(s)
	return sc
}

// SetPhysicianID sets the physician edge to Physician by id.
func (sc *SystemequipmentCreate) SetPhysicianID(id int) *SystemequipmentCreate {
	sc.mutation.SetPhysicianID(id)
	return sc
}

// SetNillablePhysicianID sets the physician edge to Physician by id if the given value is not nil.
func (sc *SystemequipmentCreate) SetNillablePhysicianID(id *int) *SystemequipmentCreate {
	if id != nil {
		sc = sc.SetPhysicianID(*id)
	}
	return sc
}

// SetPhysician sets the physician edge to Physician.
func (sc *SystemequipmentCreate) SetPhysician(p *Physician) *SystemequipmentCreate {
	return sc.SetPhysicianID(p.ID)
}

// SetMedicaltypeID sets the medicaltype edge to MedicalType by id.
func (sc *SystemequipmentCreate) SetMedicaltypeID(id int) *SystemequipmentCreate {
	sc.mutation.SetMedicaltypeID(id)
	return sc
}

// SetNillableMedicaltypeID sets the medicaltype edge to MedicalType by id if the given value is not nil.
func (sc *SystemequipmentCreate) SetNillableMedicaltypeID(id *int) *SystemequipmentCreate {
	if id != nil {
		sc = sc.SetMedicaltypeID(*id)
	}
	return sc
}

// SetMedicaltype sets the medicaltype edge to MedicalType.
func (sc *SystemequipmentCreate) SetMedicaltype(m *MedicalType) *SystemequipmentCreate {
	return sc.SetMedicaltypeID(m.ID)
}

// SetMedicalequipmentID sets the medicalequipment edge to MedicalEquipment by id.
func (sc *SystemequipmentCreate) SetMedicalequipmentID(id int) *SystemequipmentCreate {
	sc.mutation.SetMedicalequipmentID(id)
	return sc
}

// SetNillableMedicalequipmentID sets the medicalequipment edge to MedicalEquipment by id if the given value is not nil.
func (sc *SystemequipmentCreate) SetNillableMedicalequipmentID(id *int) *SystemequipmentCreate {
	if id != nil {
		sc = sc.SetMedicalequipmentID(*id)
	}
	return sc
}

// SetMedicalequipment sets the medicalequipment edge to MedicalEquipment.
func (sc *SystemequipmentCreate) SetMedicalequipment(m *MedicalEquipment) *SystemequipmentCreate {
	return sc.SetMedicalequipmentID(m.ID)
}

// Mutation returns the SystemequipmentMutation object of the builder.
func (sc *SystemequipmentCreate) Mutation() *SystemequipmentMutation {
	return sc.mutation
}

// Save creates the Systemequipment in the database.
func (sc *SystemequipmentCreate) Save(ctx context.Context) (*Systemequipment, error) {
	if _, ok := sc.mutation.Addedtime(); !ok {
		return nil, &ValidationError{Name: "addedtime", err: errors.New("ent: missing required field \"addedtime\"")}
	}
	var (
		err  error
		node *Systemequipment
	)
	if len(sc.hooks) == 0 {
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SystemequipmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SystemequipmentCreate) SaveX(ctx context.Context) *Systemequipment {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *SystemequipmentCreate) sqlSave(ctx context.Context) (*Systemequipment, error) {
	s, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}

func (sc *SystemequipmentCreate) createSpec() (*Systemequipment, *sqlgraph.CreateSpec) {
	var (
		s     = &Systemequipment{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: systemequipment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: systemequipment.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Addedtime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemequipment.FieldAddedtime,
		})
		s.Addedtime = value
	}
	if nodes := sc.mutation.PhysicianIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   systemequipment.PhysicianTable,
			Columns: []string{systemequipment.PhysicianColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: physician.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.MedicaltypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   systemequipment.MedicaltypeTable,
			Columns: []string{systemequipment.MedicaltypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicaltype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.MedicalequipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   systemequipment.MedicalequipmentTable,
			Columns: []string{systemequipment.MedicalequipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicalequipment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return s, _spec
}
