package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

// User schema.
type Medicaltype struct {
    ent.Schema
}

// Fields of the user.
func (Medicaltype) Fields() []ent.Field {
    return []ent.Field{
        field.String("Type_ID"),
        field.String("Type_name"),
    }
}
func (Medicaltype) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("Medical_type", DataSystem.Type),
    }
}