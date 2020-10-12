// Code generated by entc, DO NOT EDIT.

package medicaltype

const (
	// Label holds the string label denoting the medicaltype type in the database.
	Label = "medical_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeSystemequipment holds the string denoting the systemequipment edge name in mutations.
	EdgeSystemequipment = "systemequipment"

	// Table holds the table name of the medicaltype in the database.
	Table = "medical_types"
	// SystemequipmentTable is the table the holds the systemequipment relation/edge.
	SystemequipmentTable = "systemequipments"
	// SystemequipmentInverseTable is the table name for the Systemequipment entity.
	// It exists in this package in order to avoid circular dependency with the "systemequipment" package.
	SystemequipmentInverseTable = "systemequipments"
	// SystemequipmentColumn is the table column denoting the systemequipment relation/edge.
	SystemequipmentColumn = "medicaltype_id"
)

// Columns holds all SQL columns for medicaltype fields.
var Columns = []string{
	FieldID,
	FieldName,
}
