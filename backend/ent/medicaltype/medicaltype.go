// Code generated by entc, DO NOT EDIT.

package medicaltype

const (
	// Label holds the string label denoting the medicaltype type in the database.
	Label = "medicaltype"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTypeID holds the string denoting the type_id field in the database.
	FieldTypeID = "type_id"
	// FieldTypeName holds the string denoting the type_name field in the database.
	FieldTypeName = "type_name"

	// EdgeMedicalType holds the string denoting the medical_type edge name in mutations.
	EdgeMedicalType = "Medical_type"

	// Table holds the table name of the medicaltype in the database.
	Table = "medicaltypes"
	// MedicalTypeTable is the table the holds the Medical_type relation/edge.
	MedicalTypeTable = "systemequipments"
	// MedicalTypeInverseTable is the table name for the Systemequipment entity.
	// It exists in this package in order to avoid circular dependency with the "systemequipment" package.
	MedicalTypeInverseTable = "systemequipments"
	// MedicalTypeColumn is the table column denoting the Medical_type relation/edge.
	MedicalTypeColumn = "medicaltype_medical_type"
)

// Columns holds all SQL columns for medicaltype fields.
var Columns = []string{
	FieldID,
	FieldTypeID,
	FieldTypeName,
}
