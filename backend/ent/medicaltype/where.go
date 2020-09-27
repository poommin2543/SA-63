// Code generated by entc, DO NOT EDIT.

package medicaltype

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/poommin/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TypeID applies equality check predicate on the "Type_ID" field. It's identical to TypeIDEQ.
func TypeID(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTypeID), v))
	})
}

// TypeName applies equality check predicate on the "Type_name" field. It's identical to TypeNameEQ.
func TypeName(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTypeName), v))
	})
}

// TypeIDEQ applies the EQ predicate on the "Type_ID" field.
func TypeIDEQ(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTypeID), v))
	})
}

// TypeIDNEQ applies the NEQ predicate on the "Type_ID" field.
func TypeIDNEQ(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTypeID), v))
	})
}

// TypeIDIn applies the In predicate on the "Type_ID" field.
func TypeIDIn(vs ...string) predicate.Medicaltype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Medicaltype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTypeID), v...))
	})
}

// TypeIDNotIn applies the NotIn predicate on the "Type_ID" field.
func TypeIDNotIn(vs ...string) predicate.Medicaltype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Medicaltype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTypeID), v...))
	})
}

// TypeIDGT applies the GT predicate on the "Type_ID" field.
func TypeIDGT(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTypeID), v))
	})
}

// TypeIDGTE applies the GTE predicate on the "Type_ID" field.
func TypeIDGTE(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTypeID), v))
	})
}

// TypeIDLT applies the LT predicate on the "Type_ID" field.
func TypeIDLT(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTypeID), v))
	})
}

// TypeIDLTE applies the LTE predicate on the "Type_ID" field.
func TypeIDLTE(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTypeID), v))
	})
}

// TypeIDContains applies the Contains predicate on the "Type_ID" field.
func TypeIDContains(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTypeID), v))
	})
}

// TypeIDHasPrefix applies the HasPrefix predicate on the "Type_ID" field.
func TypeIDHasPrefix(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTypeID), v))
	})
}

// TypeIDHasSuffix applies the HasSuffix predicate on the "Type_ID" field.
func TypeIDHasSuffix(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTypeID), v))
	})
}

// TypeIDEqualFold applies the EqualFold predicate on the "Type_ID" field.
func TypeIDEqualFold(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTypeID), v))
	})
}

// TypeIDContainsFold applies the ContainsFold predicate on the "Type_ID" field.
func TypeIDContainsFold(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTypeID), v))
	})
}

// TypeNameEQ applies the EQ predicate on the "Type_name" field.
func TypeNameEQ(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTypeName), v))
	})
}

// TypeNameNEQ applies the NEQ predicate on the "Type_name" field.
func TypeNameNEQ(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTypeName), v))
	})
}

// TypeNameIn applies the In predicate on the "Type_name" field.
func TypeNameIn(vs ...string) predicate.Medicaltype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Medicaltype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTypeName), v...))
	})
}

// TypeNameNotIn applies the NotIn predicate on the "Type_name" field.
func TypeNameNotIn(vs ...string) predicate.Medicaltype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Medicaltype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTypeName), v...))
	})
}

// TypeNameGT applies the GT predicate on the "Type_name" field.
func TypeNameGT(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTypeName), v))
	})
}

// TypeNameGTE applies the GTE predicate on the "Type_name" field.
func TypeNameGTE(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTypeName), v))
	})
}

// TypeNameLT applies the LT predicate on the "Type_name" field.
func TypeNameLT(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTypeName), v))
	})
}

// TypeNameLTE applies the LTE predicate on the "Type_name" field.
func TypeNameLTE(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTypeName), v))
	})
}

// TypeNameContains applies the Contains predicate on the "Type_name" field.
func TypeNameContains(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTypeName), v))
	})
}

// TypeNameHasPrefix applies the HasPrefix predicate on the "Type_name" field.
func TypeNameHasPrefix(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTypeName), v))
	})
}

// TypeNameHasSuffix applies the HasSuffix predicate on the "Type_name" field.
func TypeNameHasSuffix(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTypeName), v))
	})
}

// TypeNameEqualFold applies the EqualFold predicate on the "Type_name" field.
func TypeNameEqualFold(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTypeName), v))
	})
}

// TypeNameContainsFold applies the ContainsFold predicate on the "Type_name" field.
func TypeNameContainsFold(v string) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTypeName), v))
	})
}

// HasMedicalType applies the HasEdge predicate on the "Medical_type" edge.
func HasMedicalType() predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MedicalTypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MedicalTypeTable, MedicalTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMedicalTypeWith applies the HasEdge predicate on the "Medical_type" edge with a given conditions (other predicates).
func HasMedicalTypeWith(preds ...predicate.Systemequipment) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MedicalTypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MedicalTypeTable, MedicalTypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Medicaltype) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Medicaltype) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
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
func Not(p predicate.Medicaltype) predicate.Medicaltype {
	return predicate.Medicaltype(func(s *sql.Selector) {
		p(s.Not())
	})
}