package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// MedicalEquipment holds the schema definition for the MedicalEquipment entity.
type MedicalEquipment struct {
	ent.Schema
}

// Fields of the MedicalEquipment.
func (MedicalEquipment) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("stock"),
	}
}

// Edges of the MedicalEquipment.
func (MedicalEquipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("systemequipment", Systemequipment.Type),
	}
}
