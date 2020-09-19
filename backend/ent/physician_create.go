// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/poommin/app/ent/physician"
	"github.com/poommin/app/ent/systemequipment"
)

// PhysicianCreate is the builder for creating a Physician entity.
type PhysicianCreate struct {
	config
	mutation *PhysicianMutation
	hooks    []Hook
}

// SetPHYSICIANID sets the PHYSICIAN_ID field.
func (pc *PhysicianCreate) SetPHYSICIANID(s string) *PhysicianCreate {
	pc.mutation.SetPHYSICIANID(s)
	return pc
}

// SetPHYSICIANNAME sets the PHYSICIAN_NAME field.
func (pc *PhysicianCreate) SetPHYSICIANNAME(s string) *PhysicianCreate {
	pc.mutation.SetPHYSICIANNAME(s)
	return pc
}

// SetPHYSICIANEMAIL sets the PHYSICIAN_EMAIL field.
func (pc *PhysicianCreate) SetPHYSICIANEMAIL(s string) *PhysicianCreate {
	pc.mutation.SetPHYSICIANEMAIL(s)
	return pc
}

// AddUserPhysicianIDs adds the User_Physician edge to Systemequipment by ids.
func (pc *PhysicianCreate) AddUserPhysicianIDs(ids ...int) *PhysicianCreate {
	pc.mutation.AddUserPhysicianIDs(ids...)
	return pc
}

// AddUserPhysician adds the User_Physician edges to Systemequipment.
func (pc *PhysicianCreate) AddUserPhysician(s ...*Systemequipment) *PhysicianCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddUserPhysicianIDs(ids...)
}

// Mutation returns the PhysicianMutation object of the builder.
func (pc *PhysicianCreate) Mutation() *PhysicianMutation {
	return pc.mutation
}

// Save creates the Physician in the database.
func (pc *PhysicianCreate) Save(ctx context.Context) (*Physician, error) {
	if _, ok := pc.mutation.PHYSICIANID(); !ok {
		return nil, &ValidationError{Name: "PHYSICIAN_ID", err: errors.New("ent: missing required field \"PHYSICIAN_ID\"")}
	}
	if _, ok := pc.mutation.PHYSICIANNAME(); !ok {
		return nil, &ValidationError{Name: "PHYSICIAN_NAME", err: errors.New("ent: missing required field \"PHYSICIAN_NAME\"")}
	}
	if _, ok := pc.mutation.PHYSICIANEMAIL(); !ok {
		return nil, &ValidationError{Name: "PHYSICIAN_EMAIL", err: errors.New("ent: missing required field \"PHYSICIAN_EMAIL\"")}
	}
	var (
		err  error
		node *Physician
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PhysicianMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PhysicianCreate) SaveX(ctx context.Context) *Physician {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PhysicianCreate) sqlSave(ctx context.Context) (*Physician, error) {
	ph, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	ph.ID = int(id)
	return ph, nil
}

func (pc *PhysicianCreate) createSpec() (*Physician, *sqlgraph.CreateSpec) {
	var (
		ph    = &Physician{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: physician.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: physician.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.PHYSICIANID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: physician.FieldPHYSICIANID,
		})
		ph.PHYSICIANID = value
	}
	if value, ok := pc.mutation.PHYSICIANNAME(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: physician.FieldPHYSICIANNAME,
		})
		ph.PHYSICIANNAME = value
	}
	if value, ok := pc.mutation.PHYSICIANEMAIL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: physician.FieldPHYSICIANEMAIL,
		})
		ph.PHYSICIANEMAIL = value
	}
	if nodes := pc.mutation.UserPhysicianIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   physician.UserPhysicianTable,
			Columns: []string{physician.UserPhysicianColumn},
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
	return ph, _spec
}
