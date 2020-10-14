package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	
)

// Systemequipment holds the schema definition for the Systemequipment entity.
type Systemequipment struct {
	ent.Schema
}

// Fields of the Systemequipment.
func (Systemequipment) Fields() []ent.Field {
	return nil 
}

// Edges of the Systemequipment.
func (Systemequipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("physician", Physician.Type).Ref("systemequipment").Unique(),
		edge.From("medicaltype", MedicalType.Type).Ref("systemequipment").Unique(),
		edge.From("medicalequipment", MedicalEquipment.Type).Ref("systemequipment").Unique(),
	}
}
