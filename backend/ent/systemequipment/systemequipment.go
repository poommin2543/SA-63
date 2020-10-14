// Code generated by entc, DO NOT EDIT.

package systemequipment

const (
	// Label holds the string label denoting the systemequipment type in the database.
	Label = "systemequipment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"

	// EdgePhysician holds the string denoting the physician edge name in mutations.
	EdgePhysician = "physician"
	// EdgeMedicaltype holds the string denoting the medicaltype edge name in mutations.
	EdgeMedicaltype = "medicaltype"
	// EdgeMedicalequipment holds the string denoting the medicalequipment edge name in mutations.
	EdgeMedicalequipment = "medicalequipment"

	// Table holds the table name of the systemequipment in the database.
	Table = "systemequipments"
	// PhysicianTable is the table the holds the physician relation/edge.
	PhysicianTable = "systemequipments"
	// PhysicianInverseTable is the table name for the Physician entity.
	// It exists in this package in order to avoid circular dependency with the "physician" package.
	PhysicianInverseTable = "physicians"
	// PhysicianColumn is the table column denoting the physician relation/edge.
	PhysicianColumn = "physician_systemequipment"
	// MedicaltypeTable is the table the holds the medicaltype relation/edge.
	MedicaltypeTable = "systemequipments"
	// MedicaltypeInverseTable is the table name for the MedicalType entity.
	// It exists in this package in order to avoid circular dependency with the "medicaltype" package.
	MedicaltypeInverseTable = "medical_types"
	// MedicaltypeColumn is the table column denoting the medicaltype relation/edge.
	MedicaltypeColumn = "medical_type_systemequipment"
	// MedicalequipmentTable is the table the holds the medicalequipment relation/edge.
	MedicalequipmentTable = "systemequipments"
	// MedicalequipmentInverseTable is the table name for the MedicalEquipment entity.
	// It exists in this package in order to avoid circular dependency with the "medicalequipment" package.
	MedicalequipmentInverseTable = "medical_equipments"
	// MedicalequipmentColumn is the table column denoting the medicalequipment relation/edge.
	MedicalequipmentColumn = "medical_equipment_systemequipment"
)

// Columns holds all SQL columns for systemequipment fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Systemequipment type.
var ForeignKeys = []string{
	"medical_equipment_systemequipment",
	"medical_type_systemequipment",
	"physician_systemequipment",
}
