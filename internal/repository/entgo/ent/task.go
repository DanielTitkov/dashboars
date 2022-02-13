// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/task"
)

// Task is the model entity for the Task schema.
type Task struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Type holds the value of the "type" field.
	Type task.Type `json:"type,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// Display holds the value of the "display" field.
	Display bool `json:"display,omitempty"`
	// Schedule holds the value of the "schedule" field.
	Schedule string `json:"schedule,omitempty"`
	// Args holds the value of the "args" field.
	Args map[string]interface{} `json:"args,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TaskQuery when eager-loading is set.
	Edges TaskEdges `json:"edges"`
}

// TaskEdges holds the relations/edges for other nodes in the graph.
type TaskEdges struct {
	// Instances holds the value of the instances edge.
	Instances []*TaskInstance `json:"instances,omitempty"`
	// Metrics holds the value of the metrics edge.
	Metrics []*Metric `json:"metrics,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// InstancesOrErr returns the Instances value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) InstancesOrErr() ([]*TaskInstance, error) {
	if e.loadedTypes[0] {
		return e.Instances, nil
	}
	return nil, &NotLoadedError{edge: "instances"}
}

// MetricsOrErr returns the Metrics value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) MetricsOrErr() ([]*Metric, error) {
	if e.loadedTypes[1] {
		return e.Metrics, nil
	}
	return nil, &NotLoadedError{edge: "metrics"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Task) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case task.FieldArgs:
			values[i] = new([]byte)
		case task.FieldActive, task.FieldDisplay:
			values[i] = new(sql.NullBool)
		case task.FieldID:
			values[i] = new(sql.NullInt64)
		case task.FieldType, task.FieldCode, task.FieldTitle, task.FieldDescription, task.FieldSchedule:
			values[i] = new(sql.NullString)
		case task.FieldCreateTime, task.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Task", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Task fields.
func (t *Task) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case task.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case task.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				t.CreateTime = value.Time
			}
		case task.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				t.UpdateTime = value.Time
			}
		case task.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				t.Type = task.Type(value.String)
			}
		case task.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				t.Code = value.String
			}
		case task.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				t.Title = value.String
			}
		case task.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		case task.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				t.Active = value.Bool
			}
		case task.FieldDisplay:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field display", values[i])
			} else if value.Valid {
				t.Display = value.Bool
			}
		case task.FieldSchedule:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field schedule", values[i])
			} else if value.Valid {
				t.Schedule = value.String
			}
		case task.FieldArgs:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field args", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.Args); err != nil {
					return fmt.Errorf("unmarshal field args: %w", err)
				}
			}
		}
	}
	return nil
}

// QueryInstances queries the "instances" edge of the Task entity.
func (t *Task) QueryInstances() *TaskInstanceQuery {
	return (&TaskClient{config: t.config}).QueryInstances(t)
}

// QueryMetrics queries the "metrics" edge of the Task entity.
func (t *Task) QueryMetrics() *MetricQuery {
	return (&TaskClient{config: t.config}).QueryMetrics(t)
}

// Update returns a builder for updating this Task.
// Note that you need to call Task.Unwrap() before calling this method if this Task
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Task) Update() *TaskUpdateOne {
	return (&TaskClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Task entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Task) Unwrap() *Task {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Task is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Task) String() string {
	var builder strings.Builder
	builder.WriteString("Task(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(t.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(t.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", t.Type))
	builder.WriteString(", code=")
	builder.WriteString(t.Code)
	builder.WriteString(", title=")
	builder.WriteString(t.Title)
	builder.WriteString(", description=")
	builder.WriteString(t.Description)
	builder.WriteString(", active=")
	builder.WriteString(fmt.Sprintf("%v", t.Active))
	builder.WriteString(", display=")
	builder.WriteString(fmt.Sprintf("%v", t.Display))
	builder.WriteString(", schedule=")
	builder.WriteString(t.Schedule)
	builder.WriteString(", args=")
	builder.WriteString(fmt.Sprintf("%v", t.Args))
	builder.WriteByte(')')
	return builder.String()
}

// Tasks is a parsable slice of Task.
type Tasks []*Task

func (t Tasks) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}