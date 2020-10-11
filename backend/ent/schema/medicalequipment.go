package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

// Medicalequipment holds the schema definition for the Medicalequipment entity.
type Medicalequipment struct {
	ent.Schema
}

// Fields of the Medicalequipment.
func (Medicalequipment) Fields() []ent.Field {
	return []ent.Field{
        field.String("Medical_ID"),
        field.String("Medical_NAME"),
        field.Int("Medical_Stock"),
	}
}

// Edges of the Medicalequipment.
func (Medicalequipment) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("Medical_equipment", Systemequipment.Type),
    }
}
