// Code generated by entc, DO NOT EDIT.

package item

import (
	"time"
)

const (
	// Label holds the string label denoting the item type in the database.
	Label = "item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// FieldMeta holds the string denoting the meta field in the database.
	FieldMeta = "meta"
	// EdgeTaskInstance holds the string denoting the task_instance edge name in mutations.
	EdgeTaskInstance = "task_instance"
	// Table holds the table name of the item in the database.
	Table = "items"
	// TaskInstanceTable is the table that holds the task_instance relation/edge.
	TaskInstanceTable = "items"
	// TaskInstanceInverseTable is the table name for the TaskInstance entity.
	// It exists in this package in order to avoid circular dependency with the "taskinstance" package.
	TaskInstanceInverseTable = "task_instances"
	// TaskInstanceColumn is the table column denoting the task_instance relation/edge.
	TaskInstanceColumn = "task_instance_items"
)

// Columns holds all SQL columns for item fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldValue,
	FieldTimestamp,
	FieldMeta,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "items"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"task_instance_items",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultTimestamp holds the default value on creation for the "timestamp" field.
	DefaultTimestamp func() time.Time
)
