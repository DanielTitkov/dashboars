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
	// EdgeDimensions holds the string denoting the dimensions edge name in mutations.
	EdgeDimensions = "dimensions"
	// EdgeTaskInstance holds the string denoting the task_instance edge name in mutations.
	EdgeTaskInstance = "task_instance"
	// EdgeMetric holds the string denoting the metric edge name in mutations.
	EdgeMetric = "metric"
	// Table holds the table name of the item in the database.
	Table = "items"
	// DimensionsTable is the table that holds the dimensions relation/edge. The primary key declared below.
	DimensionsTable = "item_dimensions"
	// DimensionsInverseTable is the table name for the Dimension entity.
	// It exists in this package in order to avoid circular dependency with the "dimension" package.
	DimensionsInverseTable = "dimensions"
	// TaskInstanceTable is the table that holds the task_instance relation/edge.
	TaskInstanceTable = "items"
	// TaskInstanceInverseTable is the table name for the TaskInstance entity.
	// It exists in this package in order to avoid circular dependency with the "taskinstance" package.
	TaskInstanceInverseTable = "task_instances"
	// TaskInstanceColumn is the table column denoting the task_instance relation/edge.
	TaskInstanceColumn = "task_instance_items"
	// MetricTable is the table that holds the metric relation/edge.
	MetricTable = "items"
	// MetricInverseTable is the table name for the Metric entity.
	// It exists in this package in order to avoid circular dependency with the "metric" package.
	MetricInverseTable = "metrics"
	// MetricColumn is the table column denoting the metric relation/edge.
	MetricColumn = "metric_items"
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
	"metric_items",
	"task_instance_items",
}

var (
	// DimensionsPrimaryKey and DimensionsColumn2 are the table columns denoting the
	// primary key for the dimensions relation (M2M).
	DimensionsPrimaryKey = []string{"item_id", "dimension_id"}
)

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
