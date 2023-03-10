// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserUUID holds the string denoting the user_uuid field in the database.
	FieldUserUUID = "user_uuid"
	// FieldUserEmail holds the string denoting the user_email field in the database.
	FieldUserEmail = "user_email"
	// FieldUserPhoneNumber holds the string denoting the user_phone_number field in the database.
	FieldUserPhoneNumber = "user_phone_number"
	// FieldUserPassword holds the string denoting the user_password field in the database.
	FieldUserPassword = "user_password"
	// FieldUserBirthday holds the string denoting the user_birthday field in the database.
	FieldUserBirthday = "user_birthday"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUserUUID,
	FieldUserEmail,
	FieldUserPhoneNumber,
	FieldUserPassword,
	FieldUserBirthday,
	FieldCreatedAt,
	FieldUpdatedAt,
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
	// DefaultUserUUID holds the default value on creation for the "user_uuid" field.
	DefaultUserUUID func() uuid.UUID
	// DefaultUserEmail holds the default value on creation for the "user_email" field.
	DefaultUserEmail string
	// DefaultUserPhoneNumber holds the default value on creation for the "user_phone_number" field.
	DefaultUserPhoneNumber string
	// DefaultUserPassword holds the default value on creation for the "user_password" field.
	DefaultUserPassword string
	// DefaultUserBirthday holds the default value on creation for the "user_birthday" field.
	DefaultUserBirthday string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)
