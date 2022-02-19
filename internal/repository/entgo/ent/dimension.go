// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/dimension"
)

// Dimension is the model entity for the Dimension schema.
type Dimension struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// DisplayTitle holds the value of the "display_title" field.
	DisplayTitle map[string]string `json:"display_title,omitempty"`
	// DisplayValue holds the value of the "display_value" field.
	DisplayValue map[string]string `json:"display_value,omitempty"`
	// Meta holds the value of the "meta" field.
	Meta map[string]interface{} `json:"meta,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DimensionQuery when eager-loading is set.
	Edges DimensionEdges `json:"edges"`
}

// DimensionEdges holds the relations/edges for other nodes in the graph.
type DimensionEdges struct {
	// Item holds the value of the item edge.
	Item []*Item `json:"item,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ItemOrErr returns the Item value or an error if the edge
// was not loaded in eager-loading.
func (e DimensionEdges) ItemOrErr() ([]*Item, error) {
	if e.loadedTypes[0] {
		return e.Item, nil
	}
	return nil, &NotLoadedError{edge: "item"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Dimension) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case dimension.FieldDisplayTitle, dimension.FieldDisplayValue, dimension.FieldMeta:
			values[i] = new([]byte)
		case dimension.FieldID:
			values[i] = new(sql.NullInt64)
		case dimension.FieldTitle, dimension.FieldValue:
			values[i] = new(sql.NullString)
		case dimension.FieldCreateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Dimension", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Dimension fields.
func (d *Dimension) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dimension.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case dimension.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				d.CreateTime = value.Time
			}
		case dimension.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				d.Title = value.String
			}
		case dimension.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				d.Value = value.String
			}
		case dimension.FieldDisplayTitle:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field display_title", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &d.DisplayTitle); err != nil {
					return fmt.Errorf("unmarshal field display_title: %w", err)
				}
			}
		case dimension.FieldDisplayValue:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field display_value", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &d.DisplayValue); err != nil {
					return fmt.Errorf("unmarshal field display_value: %w", err)
				}
			}
		case dimension.FieldMeta:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field meta", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &d.Meta); err != nil {
					return fmt.Errorf("unmarshal field meta: %w", err)
				}
			}
		}
	}
	return nil
}

// QueryItem queries the "item" edge of the Dimension entity.
func (d *Dimension) QueryItem() *ItemQuery {
	return (&DimensionClient{config: d.config}).QueryItem(d)
}

// Update returns a builder for updating this Dimension.
// Note that you need to call Dimension.Unwrap() before calling this method if this Dimension
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Dimension) Update() *DimensionUpdateOne {
	return (&DimensionClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Dimension entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Dimension) Unwrap() *Dimension {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Dimension is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Dimension) String() string {
	var builder strings.Builder
	builder.WriteString("Dimension(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(d.CreateTime.Format(time.ANSIC))
	builder.WriteString(", title=")
	builder.WriteString(d.Title)
	builder.WriteString(", value=")
	builder.WriteString(d.Value)
	builder.WriteString(", display_title=")
	builder.WriteString(fmt.Sprintf("%v", d.DisplayTitle))
	builder.WriteString(", display_value=")
	builder.WriteString(fmt.Sprintf("%v", d.DisplayValue))
	builder.WriteString(", meta=")
	builder.WriteString(fmt.Sprintf("%v", d.Meta))
	builder.WriteByte(')')
	return builder.String()
}

// Dimensions is a parsable slice of Dimension.
type Dimensions []*Dimension

func (d Dimensions) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
