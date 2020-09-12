package schema

import (
	"time"
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

// User schema.
type DataSystem struct {
	ent.Schema
}

// Fields of the user.
func (DataSystem) Fields() []ent.Field {
	return []ent.Field{
		field.String("System_ID").
			NotEmpty(),
		field.String("Medical_ID"),
		field.String("Type_ID"),
		field.String("PHYSICIAN_ID"),
		field.Time("System_DATA").
			Default(time.Now),
			
	}
}
func (DataSystem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", UserPhysician.Type).
			Ref("User_Physician").
			Unique(),

		edge.From("ownera", Medicalequipment.Type).
			Ref("Medical_equipment").
			Unique(),

		edge.From("ownerf", Medicaltype.Type).
			Ref("Medical_type").
			Unique(),
	}
}