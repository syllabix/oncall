// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSlackID holds the string denoting the slack_id field in the database.
	FieldSlackID = "slack_id"
	// FieldSlackHandle holds the string denoting the slack_handle field in the database.
	FieldSlackHandle = "slack_handle"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldAvatarURL holds the string denoting the avatar_url field in the database.
	FieldAvatarURL = "avatar_url"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeShifts holds the string denoting the shifts edge name in mutations.
	EdgeShifts = "shifts"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ShiftsTable is the table that holds the shifts relation/edge.
	ShiftsTable = "shifts"
	// ShiftsInverseTable is the table name for the Shift entity.
	// It exists in this package in order to avoid circular dependency with the "shift" package.
	ShiftsInverseTable = "shifts"
	// ShiftsColumn is the table column denoting the shifts relation/edge.
	ShiftsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldSlackID,
	FieldSlackHandle,
	FieldEmail,
	FieldFirstName,
	FieldLastName,
	FieldAvatarURL,
	FieldDisplayName,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)