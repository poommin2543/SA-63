// Code generated by entc, DO NOT EDIT.

package physician

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/poommin/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// PHYSICIANID applies equality check predicate on the "PHYSICIAN_ID" field. It's identical to PHYSICIANIDEQ.
func PHYSICIANID(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANNAME applies equality check predicate on the "PHYSICIAN_NAME" field. It's identical to PHYSICIANNAMEEQ.
func PHYSICIANNAME(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPHYSICIANNAME), v))
	})
}

// PHYSICIANEMAIL applies equality check predicate on the "PHYSICIAN_EMAIL" field. It's identical to PHYSICIANEMAILEQ.
func PHYSICIANEMAIL(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPHYSICIANEMAIL), v))
	})
}

// PHYSICIANIDEQ applies the EQ predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDEQ(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANIDNEQ applies the NEQ predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDNEQ(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANIDIn applies the In predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDIn(vs ...string) predicate.Physician {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Physician(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPHYSICIANID), v...))
	})
}

// PHYSICIANIDNotIn applies the NotIn predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDNotIn(vs ...string) predicate.Physician {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Physician(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPHYSICIANID), v...))
	})
}

// PHYSICIANIDGT applies the GT predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDGT(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANIDGTE applies the GTE predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDGTE(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANIDLT applies the LT predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDLT(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANIDLTE applies the LTE predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDLTE(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANIDContains applies the Contains predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDContains(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANIDHasPrefix applies the HasPrefix predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDHasPrefix(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANIDHasSuffix applies the HasSuffix predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDHasSuffix(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANIDEqualFold applies the EqualFold predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDEqualFold(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANIDContainsFold applies the ContainsFold predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDContainsFold(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANNAMEEQ applies the EQ predicate on the "PHYSICIAN_NAME" field.
func PHYSICIANNAMEEQ(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPHYSICIANNAME), v))
	})
}

// PHYSICIANNAMENEQ applies the NEQ predicate on the "PHYSICIAN_NAME" field.
func PHYSICIANNAMENEQ(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPHYSICIANNAME), v))
	})
}

// PHYSICIANNAMEIn applies the In predicate on the "PHYSICIAN_NAME" field.
func PHYSICIANNAMEIn(vs ...string) predicate.Physician {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Physician(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPHYSICIANNAME), v...))
	})
}

// PHYSICIANNAMENotIn applies the NotIn predicate on the "PHYSICIAN_NAME" field.
func PHYSICIANNAMENotIn(vs ...string) predicate.Physician {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Physician(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPHYSICIANNAME), v...))
	})
}

// PHYSICIANNAMEGT applies the GT predicate on the "PHYSICIAN_NAME" field.
func PHYSICIANNAMEGT(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPHYSICIANNAME), v))
	})
}

// PHYSICIANNAMEGTE applies the GTE predicate on the "PHYSICIAN_NAME" field.
func PHYSICIANNAMEGTE(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPHYSICIANNAME), v))
	})
}

// PHYSICIANNAMELT applies the LT predicate on the "PHYSICIAN_NAME" field.
func PHYSICIANNAMELT(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPHYSICIANNAME), v))
	})
}

// PHYSICIANNAMELTE applies the LTE predicate on the "PHYSICIAN_NAME" field.
func PHYSICIANNAMELTE(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPHYSICIANNAME), v))
	})
}

// PHYSICIANNAMEContains applies the Contains predicate on the "PHYSICIAN_NAME" field.
func PHYSICIANNAMEContains(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPHYSICIANNAME), v))
	})
}

// PHYSICIANNAMEHasPrefix applies the HasPrefix predicate on the "PHYSICIAN_NAME" field.
func PHYSICIANNAMEHasPrefix(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPHYSICIANNAME), v))
	})
}

// PHYSICIANNAMEHasSuffix applies the HasSuffix predicate on the "PHYSICIAN_NAME" field.
func PHYSICIANNAMEHasSuffix(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPHYSICIANNAME), v))
	})
}

// PHYSICIANNAMEEqualFold applies the EqualFold predicate on the "PHYSICIAN_NAME" field.
func PHYSICIANNAMEEqualFold(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPHYSICIANNAME), v))
	})
}

// PHYSICIANNAMEContainsFold applies the ContainsFold predicate on the "PHYSICIAN_NAME" field.
func PHYSICIANNAMEContainsFold(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPHYSICIANNAME), v))
	})
}

// PHYSICIANEMAILEQ applies the EQ predicate on the "PHYSICIAN_EMAIL" field.
func PHYSICIANEMAILEQ(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPHYSICIANEMAIL), v))
	})
}

// PHYSICIANEMAILNEQ applies the NEQ predicate on the "PHYSICIAN_EMAIL" field.
func PHYSICIANEMAILNEQ(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPHYSICIANEMAIL), v))
	})
}

// PHYSICIANEMAILIn applies the In predicate on the "PHYSICIAN_EMAIL" field.
func PHYSICIANEMAILIn(vs ...string) predicate.Physician {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Physician(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPHYSICIANEMAIL), v...))
	})
}

// PHYSICIANEMAILNotIn applies the NotIn predicate on the "PHYSICIAN_EMAIL" field.
func PHYSICIANEMAILNotIn(vs ...string) predicate.Physician {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Physician(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPHYSICIANEMAIL), v...))
	})
}

// PHYSICIANEMAILGT applies the GT predicate on the "PHYSICIAN_EMAIL" field.
func PHYSICIANEMAILGT(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPHYSICIANEMAIL), v))
	})
}

// PHYSICIANEMAILGTE applies the GTE predicate on the "PHYSICIAN_EMAIL" field.
func PHYSICIANEMAILGTE(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPHYSICIANEMAIL), v))
	})
}

// PHYSICIANEMAILLT applies the LT predicate on the "PHYSICIAN_EMAIL" field.
func PHYSICIANEMAILLT(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPHYSICIANEMAIL), v))
	})
}

// PHYSICIANEMAILLTE applies the LTE predicate on the "PHYSICIAN_EMAIL" field.
func PHYSICIANEMAILLTE(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPHYSICIANEMAIL), v))
	})
}

// PHYSICIANEMAILContains applies the Contains predicate on the "PHYSICIAN_EMAIL" field.
func PHYSICIANEMAILContains(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPHYSICIANEMAIL), v))
	})
}

// PHYSICIANEMAILHasPrefix applies the HasPrefix predicate on the "PHYSICIAN_EMAIL" field.
func PHYSICIANEMAILHasPrefix(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPHYSICIANEMAIL), v))
	})
}

// PHYSICIANEMAILHasSuffix applies the HasSuffix predicate on the "PHYSICIAN_EMAIL" field.
func PHYSICIANEMAILHasSuffix(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPHYSICIANEMAIL), v))
	})
}

// PHYSICIANEMAILEqualFold applies the EqualFold predicate on the "PHYSICIAN_EMAIL" field.
func PHYSICIANEMAILEqualFold(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPHYSICIANEMAIL), v))
	})
}

// PHYSICIANEMAILContainsFold applies the ContainsFold predicate on the "PHYSICIAN_EMAIL" field.
func PHYSICIANEMAILContainsFold(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPHYSICIANEMAIL), v))
	})
}

// HasUserPhysician applies the HasEdge predicate on the "User_Physician" edge.
func HasUserPhysician() predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserPhysicianTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserPhysicianTable, UserPhysicianColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserPhysicianWith applies the HasEdge predicate on the "User_Physician" edge with a given conditions (other predicates).
func HasUserPhysicianWith(preds ...predicate.Systemequipment) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserPhysicianInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserPhysicianTable, UserPhysicianColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Physician) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Physician) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Physician) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		p(s.Not())
	})
}