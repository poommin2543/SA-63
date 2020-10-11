package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

// Medicaltype holds the schema definition for the Medicaltype entity.
type Medicaltype struct {
	ent.Schema
}

// Fields of the Medicaltype.
func (Medicaltype) Fields() []ent.Field {
	return []ent.Field{
        field.String("Type_ID"),
        field.String("Type_name"),
    }
}

// Edges of the Medicaltype.
func (Medicaltype) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("Medical_type", Systemequipment.Type),
    }
}
