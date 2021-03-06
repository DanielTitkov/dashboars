// Code generated by entc, DO NOT EDIT.

package task

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldDisplay holds the string denoting the display field in the database.
	FieldDisplay = "display"
	// FieldSchedule holds the string denoting the schedule field in the database.
	FieldSchedule = "schedule"
	// FieldArgs holds the string denoting the args field in the database.
	FieldArgs = "args"
	// EdgeInstances holds the string denoting the instances edge name in mutations.
	EdgeInstances = "instances"
	// EdgeMetrics holds the string denoting the metrics edge name in mutations.
	EdgeMetrics = "metrics"
	// EdgeCategory holds the string denoting the category edge name in mutations.
	EdgeCategory = "category"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// Table holds the table name of the task in the database.
	Table = "tasks"
	// InstancesTable is the table that holds the instances relation/edge.
	InstancesTable = "task_instances"
	// InstancesInverseTable is the table name for the TaskInstance entity.
	// It exists in this package in order to avoid circular dependency with the "taskinstance" package.
	InstancesInverseTable = "task_instances"
	// InstancesColumn is the table column denoting the instances relation/edge.
	InstancesColumn = "task_instances"
	// MetricsTable is the table that holds the metrics relation/edge.
	MetricsTable = "metrics"
	// MetricsInverseTable is the table name for the Metric entity.
	// It exists in this package in order to avoid circular dependency with the "metric" package.
	MetricsInverseTable = "metrics"
	// MetricsColumn is the table column denoting the metrics relation/edge.
	MetricsColumn = "task_metrics"
	// CategoryTable is the table that holds the category relation/edge.
	CategoryTable = "tasks"
	// CategoryInverseTable is the table name for the TaskCategory entity.
	// It exists in this package in order to avoid circular dependency with the "taskcategory" package.
	CategoryInverseTable = "task_categories"
	// CategoryColumn is the table column denoting the category relation/edge.
	CategoryColumn = "task_category_tasks"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "task_tag_tasks"
	// TagsInverseTable is the table name for the TaskTag entity.
	// It exists in this package in order to avoid circular dependency with the "tasktag" package.
	TagsInverseTable = "task_tags"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldType,
	FieldCode,
	FieldTitle,
	FieldDescription,
	FieldActive,
	FieldDisplay,
	FieldSchedule,
	FieldArgs,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tasks"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"task_category_tasks",
}

var (
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"task_tag_id", "task_id"}
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
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
	// DefaultDisplay holds the default value on creation for the "display" field.
	DefaultDisplay bool
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeRandom  Type = "random"
	TypeFailing Type = "failing"
	TypeParser  Type = "parser"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeRandom, TypeFailing, TypeParser:
		return nil
	default:
		return fmt.Errorf("task: invalid enum value for type field: %q", _type)
	}
}
