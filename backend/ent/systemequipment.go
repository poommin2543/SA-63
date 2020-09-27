// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/poommin/app/ent/medicalequipment"
	"github.com/poommin/app/ent/medicaltype"
	"github.com/poommin/app/ent/physician"
	"github.com/poommin/app/ent/systemequipment"
)

// Systemequipment is the model entity for the Systemequipment schema.
type Systemequipment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SystemID holds the value of the "System_ID" field.
	SystemID string `json:"System_ID,omitempty"`
	// MedicalID holds the value of the "Medical_ID" field.
	MedicalID string `json:"Medical_ID,omitempty"`
	// TypeID holds the value of the "Type_ID" field.
	TypeID string `json:"Type_ID,omitempty"`
	// PHYSICIANID holds the value of the "PHYSICIAN_ID" field.
	PHYSICIANID string `json:"PHYSICIAN_ID,omitempty"`
	// SystemDATA holds the value of the "System_DATA" field.
	SystemDATA time.Time `json:"System_DATA,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SystemequipmentQuery when eager-loading is set.
	Edges                              SystemequipmentEdges `json:"edges"`
	medicalequipment_medical_equipment *int
	medicaltype_medical_type           *int
	physician_user_physician           *int
}

// SystemequipmentEdges holds the relations/edges for other nodes in the graph.
type SystemequipmentEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Physician
	// Ownera holds the value of the ownera edge.
	Ownera *Medicalequipment
	// Ownerf holds the value of the ownerf edge.
	Ownerf *Medicaltype
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SystemequipmentEdges) OwnerOrErr() (*Physician, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: physician.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// OwneraOrErr returns the Ownera value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SystemequipmentEdges) OwneraOrErr() (*Medicalequipment, error) {
	if e.loadedTypes[1] {
		if e.Ownera == nil {
			// The edge ownera was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: medicalequipment.Label}
		}
		return e.Ownera, nil
	}
	return nil, &NotLoadedError{edge: "ownera"}
}

// OwnerfOrErr returns the Ownerf value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SystemequipmentEdges) OwnerfOrErr() (*Medicaltype, error) {
	if e.loadedTypes[2] {
		if e.Ownerf == nil {
			// The edge ownerf was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: medicaltype.Label}
		}
		return e.Ownerf, nil
	}
	return nil, &NotLoadedError{edge: "ownerf"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Systemequipment) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // System_ID
		&sql.NullString{}, // Medical_ID
		&sql.NullString{}, // Type_ID
		&sql.NullString{}, // PHYSICIAN_ID
		&sql.NullTime{},   // System_DATA
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Systemequipment) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // medicalequipment_medical_equipment
		&sql.NullInt64{}, // medicaltype_medical_type
		&sql.NullInt64{}, // physician_user_physician
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Systemequipment fields.
func (s *Systemequipment) assignValues(values ...interface{}) error {
	if m, n := len(values), len(systemequipment.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field System_ID", values[0])
	} else if value.Valid {
		s.SystemID = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Medical_ID", values[1])
	} else if value.Valid {
		s.MedicalID = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Type_ID", values[2])
	} else if value.Valid {
		s.TypeID = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field PHYSICIAN_ID", values[3])
	} else if value.Valid {
		s.PHYSICIANID = value.String
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field System_DATA", values[4])
	} else if value.Valid {
		s.SystemDATA = value.Time
	}
	values = values[5:]
	if len(values) == len(systemequipment.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field medicalequipment_medical_equipment", value)
		} else if value.Valid {
			s.medicalequipment_medical_equipment = new(int)
			*s.medicalequipment_medical_equipment = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field medicaltype_medical_type", value)
		} else if value.Valid {
			s.medicaltype_medical_type = new(int)
			*s.medicaltype_medical_type = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field physician_user_physician", value)
		} else if value.Valid {
			s.physician_user_physician = new(int)
			*s.physician_user_physician = int(value.Int64)
		}
	}
	return nil
}

// QueryOwner queries the owner edge of the Systemequipment.
func (s *Systemequipment) QueryOwner() *PhysicianQuery {
	return (&SystemequipmentClient{config: s.config}).QueryOwner(s)
}

// QueryOwnera queries the ownera edge of the Systemequipment.
func (s *Systemequipment) QueryOwnera() *MedicalequipmentQuery {
	return (&SystemequipmentClient{config: s.config}).QueryOwnera(s)
}

// QueryOwnerf queries the ownerf edge of the Systemequipment.
func (s *Systemequipment) QueryOwnerf() *MedicaltypeQuery {
	return (&SystemequipmentClient{config: s.config}).QueryOwnerf(s)
}

// Update returns a builder for updating this Systemequipment.
// Note that, you need to call Systemequipment.Unwrap() before calling this method, if this Systemequipment
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Systemequipment) Update() *SystemequipmentUpdateOne {
	return (&SystemequipmentClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Systemequipment) Unwrap() *Systemequipment {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Systemequipment is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Systemequipment) String() string {
	var builder strings.Builder
	builder.WriteString("Systemequipment(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", System_ID=")
	builder.WriteString(s.SystemID)
	builder.WriteString(", Medical_ID=")
	builder.WriteString(s.MedicalID)
	builder.WriteString(", Type_ID=")
	builder.WriteString(s.TypeID)
	builder.WriteString(", PHYSICIAN_ID=")
	builder.WriteString(s.PHYSICIANID)
	builder.WriteString(", System_DATA=")
	builder.WriteString(s.SystemDATA.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Systemequipments is a parsable slice of Systemequipment.
type Systemequipments []*Systemequipment

func (s Systemequipments) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}