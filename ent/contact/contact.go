// Code generated by entc, DO NOT EDIT.

package contact

import (
	"time"
)

const (
	// Label holds the string label denoting the contact type in the database.
	Label = "contact"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// Table holds the table name of the contact in the database.
	Table = "contacts"
)

// Columns holds all SQL columns for contact fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPhone,
	FieldCreateTime,
	FieldUpdateTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
)
