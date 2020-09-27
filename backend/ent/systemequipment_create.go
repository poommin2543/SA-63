// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/poommin/app/ent/medicalequipment"
	"github.com/poommin/app/ent/medicaltype"
	"github.com/poommin/app/ent/physician"
	"github.com/poommin/app/ent/systemequipment"
)

// SystemequipmentCreate is the builder for creating a Systemequipment entity.
type SystemequipmentCreate struct {
	config
	mutation *SystemequipmentMutation
	hooks    []Hook
}

// SetSystemID sets the System_ID field.
func (sc *SystemequipmentCreate) SetSystemID(s string) *SystemequipmentCreate {
	sc.mutation.SetSystemID(s)
	return sc
}

// SetMedicalID sets the Medical_ID field.
func (sc *SystemequipmentCreate) SetMedicalID(s string) *SystemequipmentCreate {
	sc.mutation.SetMedicalID(s)
	return sc
}

// SetTypeID sets the Type_ID field.
func (sc *SystemequipmentCreate) SetTypeID(s string) *SystemequipmentCreate {
	sc.mutation.SetTypeID(s)
	return sc
}

// SetPHYSICIANID sets the PHYSICIAN_ID field.
func (sc *SystemequipmentCreate) SetPHYSICIANID(s string) *SystemequipmentCreate {
	sc.mutation.SetPHYSICIANID(s)
	return sc
}

// SetSystemDATA sets the System_DATA field.
func (sc *SystemequipmentCreate) SetSystemDATA(t time.Time) *SystemequipmentCreate {
	sc.mutation.SetSystemDATA(t)
	return sc
}

// SetNillableSystemDATA sets the System_DATA field if the given value is not nil.
func (sc *SystemequipmentCreate) SetNillableSystemDATA(t *time.Time) *SystemequipmentCreate {
	if t != nil {
		sc.SetSystemDATA(*t)
	}
	return sc
}

// SetOwnerID sets the owner edge to Physician by id.
func (sc *SystemequipmentCreate) SetOwnerID(id int) *SystemequipmentCreate {
	sc.mutation.SetOwnerID(id)
	return sc
}

// SetNillableOwnerID sets the owner edge to Physician by id if the given value is not nil.
func (sc *SystemequipmentCreate) SetNillableOwnerID(id *int) *SystemequipmentCreate {
	if id != nil {
		sc = sc.SetOwnerID(*id)
	}
	return sc
}

// SetOwner sets the owner edge to Physician.
func (sc *SystemequipmentCreate) SetOwner(p *Physician) *SystemequipmentCreate {
	return sc.SetOwnerID(p.ID)
}

// SetOwneraID sets the ownera edge to Medicalequipment by id.
func (sc *SystemequipmentCreate) SetOwneraID(id int) *SystemequipmentCreate {
	sc.mutation.SetOwneraID(id)
	return sc
}

// SetNillableOwneraID sets the ownera edge to Medicalequipment by id if the given value is not nil.
func (sc *SystemequipmentCreate) SetNillableOwneraID(id *int) *SystemequipmentCreate {
	if id != nil {
		sc = sc.SetOwneraID(*id)
	}
	return sc
}

// SetOwnera sets the ownera edge to Medicalequipment.
func (sc *SystemequipmentCreate) SetOwnera(m *Medicalequipment) *SystemequipmentCreate {
	return sc.SetOwneraID(m.ID)
}

// SetOwnerfID sets the ownerf edge to Medicaltype by id.
func (sc *SystemequipmentCreate) SetOwnerfID(id int) *SystemequipmentCreate {
	sc.mutation.SetOwnerfID(id)
	return sc
}

// SetNillableOwnerfID sets the ownerf edge to Medicaltype by id if the given value is not nil.
func (sc *SystemequipmentCreate) SetNillableOwnerfID(id *int) *SystemequipmentCreate {
	if id != nil {
		sc = sc.SetOwnerfID(*id)
	}
	return sc
}

// SetOwnerf sets the ownerf edge to Medicaltype.
func (sc *SystemequipmentCreate) SetOwnerf(m *Medicaltype) *SystemequipmentCreate {
	return sc.SetOwnerfID(m.ID)
}

// Mutation returns the SystemequipmentMutation object of the builder.
func (sc *SystemequipmentCreate) Mutation() *SystemequipmentMutation {
	return sc.mutation
}

// Save creates the Systemequipment in the database.
func (sc *SystemequipmentCreate) Save(ctx context.Context) (*Systemequipment, error) {
	if _, ok := sc.mutation.SystemID(); !ok {
		return nil, &ValidationError{Name: "System_ID", err: errors.New("ent: missing required field \"System_ID\"")}
	}
	if v, ok := sc.mutation.SystemID(); ok {
		if err := systemequipment.SystemIDValidator(v); err != nil {
			return nil, &ValidationError{Name: "System_ID", err: fmt.Errorf("ent: validator failed for field \"System_ID\": %w", err)}
		}
	}
	if _, ok := sc.mutation.MedicalID(); !ok {
		return nil, &ValidationError{Name: "Medical_ID", err: errors.New("ent: missing required field \"Medical_ID\"")}
	}
	if _, ok := sc.mutation.TypeID(); !ok {
		return nil, &ValidationError{Name: "Type_ID", err: errors.New("ent: missing required field \"Type_ID\"")}
	}
	if _, ok := sc.mutation.PHYSICIANID(); !ok {
		return nil, &ValidationError{Name: "PHYSICIAN_ID", err: errors.New("ent: missing required field \"PHYSICIAN_ID\"")}
	}
	if _, ok := sc.mutation.SystemDATA(); !ok {
		v := systemequipment.DefaultSystemDATA()
		sc.mutation.SetSystemDATA(v)
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
	if value, ok := sc.mutation.SystemID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemequipment.FieldSystemID,
		})
		s.SystemID = value
	}
	if value, ok := sc.mutation.MedicalID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemequipment.FieldMedicalID,
		})
		s.MedicalID = value
	}
	if value, ok := sc.mutation.TypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemequipment.FieldTypeID,
		})
		s.TypeID = value
	}
	if value, ok := sc.mutation.PHYSICIANID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: systemequipment.FieldPHYSICIANID,
		})
		s.PHYSICIANID = value
	}
	if value, ok := sc.mutation.SystemDATA(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: systemequipment.FieldSystemDATA,
		})
		s.SystemDATA = value
	}
	if nodes := sc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   systemequipment.OwnerTable,
			Columns: []string{systemequipment.OwnerColumn},
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
	if nodes := sc.mutation.OwneraIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   systemequipment.OwneraTable,
			Columns: []string{systemequipment.OwneraColumn},
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
	if nodes := sc.mutation.OwnerfIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   systemequipment.OwnerfTable,
			Columns: []string{systemequipment.OwnerfColumn},
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
	return s, _spec
}