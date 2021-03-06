// Code generated by entc, DO NOT EDIT.

package metric

import (
	"time"
)

const (
	// Label holds the string label denoting the metric type in the database.
	Label = "metric"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDisplayTitle holds the string denoting the display_title field in the database.
	FieldDisplayTitle = "display_title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldMeta holds the string denoting the meta field in the database.
	FieldMeta = "meta"
	// EdgeItems holds the string denoting the items edge name in mutations.
	EdgeItems = "items"
	// EdgeTask holds the string denoting the task edge name in mutations.
	EdgeTask = "task"
	// Table holds the table name of the metric in the database.
	Table = "metrics"
	// ItemsTable is the table that holds the items relation/edge.
	ItemsTable = "items"
	// ItemsInverseTable is the table name for the Item entity.
	// It exists in this package in order to avoid circular dependency with the "item" package.
	ItemsInverseTable = "items"
	// ItemsColumn is the table column denoting the items relation/edge.
	ItemsColumn = "metric_items"
	// TaskTable is the table that holds the task relation/edge.
	TaskTable = "metrics"
	// TaskInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TaskInverseTable = "tasks"
	// TaskColumn is the table column denoting the task relation/edge.
	TaskColumn = "task_metrics"
)

// Columns holds all SQL columns for metric fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldTitle,
	FieldDisplayTitle,
	FieldDescription,
	FieldMeta,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "metrics"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"task_metrics",
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
)
