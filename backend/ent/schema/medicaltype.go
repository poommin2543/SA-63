package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// MedicalType holds the schema definition for the MedicalType entity.
type MedicalType struct {
	ent.Schema
}

// Fields of the MedicalType.
func (MedicalType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the MedicalType.
func (MedicalType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("systemequipment", Systemequipment.Type).StorageKey(edge.Column("medicaltype_id")),
	}
}
