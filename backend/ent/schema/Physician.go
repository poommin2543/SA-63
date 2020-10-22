package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Physician holds the schema definition for the Physician entity.
type Physician struct {
	ent.Schema
}

// Fields of the Physician.
func (Physician) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.String("email").NotEmpty().Unique(),
	}
}

// Edges of the Physician.
func (Physician) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("systemequipment", Systemequipment.Type),
	}
}
