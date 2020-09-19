// Code generated by entc, DO NOT EDIT.

package systemequipment

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/poommin/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// SystemID applies equality check predicate on the "System_ID" field. It's identical to SystemIDEQ.
func SystemID(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSystemID), v))
	})
}

// MedicalID applies equality check predicate on the "Medical_ID" field. It's identical to MedicalIDEQ.
func MedicalID(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMedicalID), v))
	})
}

// TypeID applies equality check predicate on the "Type_ID" field. It's identical to TypeIDEQ.
func TypeID(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTypeID), v))
	})
}

// PHYSICIANID applies equality check predicate on the "PHYSICIAN_ID" field. It's identical to PHYSICIANIDEQ.
func PHYSICIANID(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPHYSICIANID), v))
	})
}

// SystemDATA applies equality check predicate on the "System_DATA" field. It's identical to SystemDATAEQ.
func SystemDATA(v time.Time) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSystemDATA), v))
	})
}

// SystemIDEQ applies the EQ predicate on the "System_ID" field.
func SystemIDEQ(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSystemID), v))
	})
}

// SystemIDNEQ applies the NEQ predicate on the "System_ID" field.
func SystemIDNEQ(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSystemID), v))
	})
}

// SystemIDIn applies the In predicate on the "System_ID" field.
func SystemIDIn(vs ...string) predicate.Systemequipment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Systemequipment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSystemID), v...))
	})
}

// SystemIDNotIn applies the NotIn predicate on the "System_ID" field.
func SystemIDNotIn(vs ...string) predicate.Systemequipment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Systemequipment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSystemID), v...))
	})
}

// SystemIDGT applies the GT predicate on the "System_ID" field.
func SystemIDGT(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSystemID), v))
	})
}

// SystemIDGTE applies the GTE predicate on the "System_ID" field.
func SystemIDGTE(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSystemID), v))
	})
}

// SystemIDLT applies the LT predicate on the "System_ID" field.
func SystemIDLT(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSystemID), v))
	})
}

// SystemIDLTE applies the LTE predicate on the "System_ID" field.
func SystemIDLTE(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSystemID), v))
	})
}

// SystemIDContains applies the Contains predicate on the "System_ID" field.
func SystemIDContains(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSystemID), v))
	})
}

// SystemIDHasPrefix applies the HasPrefix predicate on the "System_ID" field.
func SystemIDHasPrefix(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSystemID), v))
	})
}

// SystemIDHasSuffix applies the HasSuffix predicate on the "System_ID" field.
func SystemIDHasSuffix(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSystemID), v))
	})
}

// SystemIDEqualFold applies the EqualFold predicate on the "System_ID" field.
func SystemIDEqualFold(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSystemID), v))
	})
}

// SystemIDContainsFold applies the ContainsFold predicate on the "System_ID" field.
func SystemIDContainsFold(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSystemID), v))
	})
}

// MedicalIDEQ applies the EQ predicate on the "Medical_ID" field.
func MedicalIDEQ(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMedicalID), v))
	})
}

// MedicalIDNEQ applies the NEQ predicate on the "Medical_ID" field.
func MedicalIDNEQ(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMedicalID), v))
	})
}

// MedicalIDIn applies the In predicate on the "Medical_ID" field.
func MedicalIDIn(vs ...string) predicate.Systemequipment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Systemequipment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMedicalID), v...))
	})
}

// MedicalIDNotIn applies the NotIn predicate on the "Medical_ID" field.
func MedicalIDNotIn(vs ...string) predicate.Systemequipment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Systemequipment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMedicalID), v...))
	})
}

// MedicalIDGT applies the GT predicate on the "Medical_ID" field.
func MedicalIDGT(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMedicalID), v))
	})
}

// MedicalIDGTE applies the GTE predicate on the "Medical_ID" field.
func MedicalIDGTE(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMedicalID), v))
	})
}

// MedicalIDLT applies the LT predicate on the "Medical_ID" field.
func MedicalIDLT(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMedicalID), v))
	})
}

// MedicalIDLTE applies the LTE predicate on the "Medical_ID" field.
func MedicalIDLTE(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMedicalID), v))
	})
}

// MedicalIDContains applies the Contains predicate on the "Medical_ID" field.
func MedicalIDContains(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMedicalID), v))
	})
}

// MedicalIDHasPrefix applies the HasPrefix predicate on the "Medical_ID" field.
func MedicalIDHasPrefix(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMedicalID), v))
	})
}

// MedicalIDHasSuffix applies the HasSuffix predicate on the "Medical_ID" field.
func MedicalIDHasSuffix(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMedicalID), v))
	})
}

// MedicalIDEqualFold applies the EqualFold predicate on the "Medical_ID" field.
func MedicalIDEqualFold(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMedicalID), v))
	})
}

// MedicalIDContainsFold applies the ContainsFold predicate on the "Medical_ID" field.
func MedicalIDContainsFold(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMedicalID), v))
	})
}

// TypeIDEQ applies the EQ predicate on the "Type_ID" field.
func TypeIDEQ(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTypeID), v))
	})
}

// TypeIDNEQ applies the NEQ predicate on the "Type_ID" field.
func TypeIDNEQ(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTypeID), v))
	})
}

// TypeIDIn applies the In predicate on the "Type_ID" field.
func TypeIDIn(vs ...string) predicate.Systemequipment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Systemequipment(func(s *sql.Selector) {
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
func TypeIDNotIn(vs ...string) predicate.Systemequipment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Systemequipment(func(s *sql.Selector) {
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
func TypeIDGT(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTypeID), v))
	})
}

// TypeIDGTE applies the GTE predicate on the "Type_ID" field.
func TypeIDGTE(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTypeID), v))
	})
}

// TypeIDLT applies the LT predicate on the "Type_ID" field.
func TypeIDLT(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTypeID), v))
	})
}

// TypeIDLTE applies the LTE predicate on the "Type_ID" field.
func TypeIDLTE(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTypeID), v))
	})
}

// TypeIDContains applies the Contains predicate on the "Type_ID" field.
func TypeIDContains(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTypeID), v))
	})
}

// TypeIDHasPrefix applies the HasPrefix predicate on the "Type_ID" field.
func TypeIDHasPrefix(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTypeID), v))
	})
}

// TypeIDHasSuffix applies the HasSuffix predicate on the "Type_ID" field.
func TypeIDHasSuffix(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTypeID), v))
	})
}

// TypeIDEqualFold applies the EqualFold predicate on the "Type_ID" field.
func TypeIDEqualFold(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTypeID), v))
	})
}

// TypeIDContainsFold applies the ContainsFold predicate on the "Type_ID" field.
func TypeIDContainsFold(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTypeID), v))
	})
}

// PHYSICIANIDEQ applies the EQ predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDEQ(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANIDNEQ applies the NEQ predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDNEQ(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANIDIn applies the In predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDIn(vs ...string) predicate.Systemequipment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Systemequipment(func(s *sql.Selector) {
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
func PHYSICIANIDNotIn(vs ...string) predicate.Systemequipment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Systemequipment(func(s *sql.Selector) {
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
func PHYSICIANIDGT(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANIDGTE applies the GTE predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDGTE(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANIDLT applies the LT predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDLT(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANIDLTE applies the LTE predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDLTE(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANIDContains applies the Contains predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDContains(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANIDHasPrefix applies the HasPrefix predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDHasPrefix(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANIDHasSuffix applies the HasSuffix predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDHasSuffix(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANIDEqualFold applies the EqualFold predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDEqualFold(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPHYSICIANID), v))
	})
}

// PHYSICIANIDContainsFold applies the ContainsFold predicate on the "PHYSICIAN_ID" field.
func PHYSICIANIDContainsFold(v string) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPHYSICIANID), v))
	})
}

// SystemDATAEQ applies the EQ predicate on the "System_DATA" field.
func SystemDATAEQ(v time.Time) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSystemDATA), v))
	})
}

// SystemDATANEQ applies the NEQ predicate on the "System_DATA" field.
func SystemDATANEQ(v time.Time) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSystemDATA), v))
	})
}

// SystemDATAIn applies the In predicate on the "System_DATA" field.
func SystemDATAIn(vs ...time.Time) predicate.Systemequipment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Systemequipment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSystemDATA), v...))
	})
}

// SystemDATANotIn applies the NotIn predicate on the "System_DATA" field.
func SystemDATANotIn(vs ...time.Time) predicate.Systemequipment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Systemequipment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSystemDATA), v...))
	})
}

// SystemDATAGT applies the GT predicate on the "System_DATA" field.
func SystemDATAGT(v time.Time) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSystemDATA), v))
	})
}

// SystemDATAGTE applies the GTE predicate on the "System_DATA" field.
func SystemDATAGTE(v time.Time) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSystemDATA), v))
	})
}

// SystemDATALT applies the LT predicate on the "System_DATA" field.
func SystemDATALT(v time.Time) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSystemDATA), v))
	})
}

// SystemDATALTE applies the LTE predicate on the "System_DATA" field.
func SystemDATALTE(v time.Time) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSystemDATA), v))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Physician) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOwnera applies the HasEdge predicate on the "ownera" edge.
func HasOwnera() predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwneraTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwneraTable, OwneraColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwneraWith applies the HasEdge predicate on the "ownera" edge with a given conditions (other predicates).
func HasOwneraWith(preds ...predicate.Medicalequipment) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwneraInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwneraTable, OwneraColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOwnerf applies the HasEdge predicate on the "ownerf" edge.
func HasOwnerf() predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerfTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerfTable, OwnerfColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerfWith applies the HasEdge predicate on the "ownerf" edge with a given conditions (other predicates).
func HasOwnerfWith(preds ...predicate.Medicaltype) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerfInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerfTable, OwnerfColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Systemequipment) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Systemequipment) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
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
func Not(p predicate.Systemequipment) predicate.Systemequipment {
	return predicate.Systemequipment(func(s *sql.Selector) {
		p(s.Not())
	})
}
