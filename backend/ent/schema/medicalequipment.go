package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)
// User schema.
type Medicalequipment  struct {
    ent.Schema
}

// Fields of the user.
func (Medicalequipment ) Fields() []ent.Field {
    return []ent.Field{
        field.String("Medical_ID"),
        field.String("Medical_NAME"),
        field.Int("Medical_Stock"),
       
    }
}
func (Medicalequipment) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("Medical_equipment", DataSystem.Type),
    }
}