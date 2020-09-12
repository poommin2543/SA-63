package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

// User schema.
type UserPhysician struct {
	ent.Schema
}

// Fields of the user.
func (UserPhysician) Fields() []ent.Field {
	return []ent.Field{

		field.String("PHYSICIAN_ID"),
		field.String("PHYSICIAN_NAME"),
		field.String("PHYSICIAN_EMAIL"),
	}
}
func (UserPhysician) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("User_Physician", DataSystem.Type),
	}
}
