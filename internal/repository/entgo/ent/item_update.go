// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/dimension"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/item"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/metric"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/taskinstance"
)

// ItemUpdate is the builder for updating Item entities.
type ItemUpdate struct {
	config
	hooks    []Hook
	mutation *ItemMutation
}

// Where appends a list predicates to the ItemUpdate builder.
func (iu *ItemUpdate) Where(ps ...predicate.Item) *ItemUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetUpdateTime sets the "update_time" field.
func (iu *ItemUpdate) SetUpdateTime(t time.Time) *ItemUpdate {
	iu.mutation.SetUpdateTime(t)
	return iu
}

// SetValue sets the "value" field.
func (iu *ItemUpdate) SetValue(f float64) *ItemUpdate {
	iu.mutation.ResetValue()
	iu.mutation.SetValue(f)
	return iu
}

// AddValue adds f to the "value" field.
func (iu *ItemUpdate) AddValue(f float64) *ItemUpdate {
	iu.mutation.AddValue(f)
	return iu
}

// SetTimestamp sets the "timestamp" field.
func (iu *ItemUpdate) SetTimestamp(t time.Time) *ItemUpdate {
	iu.mutation.SetTimestamp(t)
	return iu
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (iu *ItemUpdate) SetNillableTimestamp(t *time.Time) *ItemUpdate {
	if t != nil {
		iu.SetTimestamp(*t)
	}
	return iu
}

// SetMeta sets the "meta" field.
func (iu *ItemUpdate) SetMeta(m map[string]interface{}) *ItemUpdate {
	iu.mutation.SetMeta(m)
	return iu
}

// ClearMeta clears the value of the "meta" field.
func (iu *ItemUpdate) ClearMeta() *ItemUpdate {
	iu.mutation.ClearMeta()
	return iu
}

// AddDimensionIDs adds the "dimensions" edge to the Dimension entity by IDs.
func (iu *ItemUpdate) AddDimensionIDs(ids ...int) *ItemUpdate {
	iu.mutation.AddDimensionIDs(ids...)
	return iu
}

// AddDimensions adds the "dimensions" edges to the Dimension entity.
func (iu *ItemUpdate) AddDimensions(d ...*Dimension) *ItemUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return iu.AddDimensionIDs(ids...)
}

// SetTaskInstanceID sets the "task_instance" edge to the TaskInstance entity by ID.
func (iu *ItemUpdate) SetTaskInstanceID(id int) *ItemUpdate {
	iu.mutation.SetTaskInstanceID(id)
	return iu
}

// SetTaskInstance sets the "task_instance" edge to the TaskInstance entity.
func (iu *ItemUpdate) SetTaskInstance(t *TaskInstance) *ItemUpdate {
	return iu.SetTaskInstanceID(t.ID)
}

// SetMetricID sets the "metric" edge to the Metric entity by ID.
func (iu *ItemUpdate) SetMetricID(id int) *ItemUpdate {
	iu.mutation.SetMetricID(id)
	return iu
}

// SetMetric sets the "metric" edge to the Metric entity.
func (iu *ItemUpdate) SetMetric(m *Metric) *ItemUpdate {
	return iu.SetMetricID(m.ID)
}

// Mutation returns the ItemMutation object of the builder.
func (iu *ItemUpdate) Mutation() *ItemMutation {
	return iu.mutation
}

// ClearDimensions clears all "dimensions" edges to the Dimension entity.
func (iu *ItemUpdate) ClearDimensions() *ItemUpdate {
	iu.mutation.ClearDimensions()
	return iu
}

// RemoveDimensionIDs removes the "dimensions" edge to Dimension entities by IDs.
func (iu *ItemUpdate) RemoveDimensionIDs(ids ...int) *ItemUpdate {
	iu.mutation.RemoveDimensionIDs(ids...)
	return iu
}

// RemoveDimensions removes "dimensions" edges to Dimension entities.
func (iu *ItemUpdate) RemoveDimensions(d ...*Dimension) *ItemUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return iu.RemoveDimensionIDs(ids...)
}

// ClearTaskInstance clears the "task_instance" edge to the TaskInstance entity.
func (iu *ItemUpdate) ClearTaskInstance() *ItemUpdate {
	iu.mutation.ClearTaskInstance()
	return iu
}

// ClearMetric clears the "metric" edge to the Metric entity.
func (iu *ItemUpdate) ClearMetric() *ItemUpdate {
	iu.mutation.ClearMetric()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ItemUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	iu.defaults()
	if len(iu.hooks) == 0 {
		if err = iu.check(); err != nil {
			return 0, err
		}
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iu.check(); err != nil {
				return 0, err
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			if iu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ItemUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ItemUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ItemUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *ItemUpdate) defaults() {
	if _, ok := iu.mutation.UpdateTime(); !ok {
		v := item.UpdateDefaultUpdateTime()
		iu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *ItemUpdate) check() error {
	if _, ok := iu.mutation.TaskInstanceID(); iu.mutation.TaskInstanceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Item.task_instance"`)
	}
	if _, ok := iu.mutation.MetricID(); iu.mutation.MetricCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Item.metric"`)
	}
	return nil
}

func (iu *ItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   item.Table,
			Columns: item.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: item.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldUpdateTime,
		})
	}
	if value, ok := iu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: item.FieldValue,
		})
	}
	if value, ok := iu.mutation.AddedValue(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: item.FieldValue,
		})
	}
	if value, ok := iu.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldTimestamp,
		})
	}
	if value, ok := iu.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: item.FieldMeta,
		})
	}
	if iu.mutation.MetaCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: item.FieldMeta,
		})
	}
	if iu.mutation.DimensionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   item.DimensionsTable,
			Columns: item.DimensionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dimension.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedDimensionsIDs(); len(nodes) > 0 && !iu.mutation.DimensionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   item.DimensionsTable,
			Columns: item.DimensionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dimension.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.DimensionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   item.DimensionsTable,
			Columns: item.DimensionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dimension.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.TaskInstanceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.TaskInstanceTable,
			Columns: []string{item.TaskInstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: taskinstance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.TaskInstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.TaskInstanceTable,
			Columns: []string{item.TaskInstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: taskinstance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.MetricCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.MetricTable,
			Columns: []string{item.MetricColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metric.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.MetricIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.MetricTable,
			Columns: []string{item.MetricColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metric.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ItemUpdateOne is the builder for updating a single Item entity.
type ItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ItemMutation
}

// SetUpdateTime sets the "update_time" field.
func (iuo *ItemUpdateOne) SetUpdateTime(t time.Time) *ItemUpdateOne {
	iuo.mutation.SetUpdateTime(t)
	return iuo
}

// SetValue sets the "value" field.
func (iuo *ItemUpdateOne) SetValue(f float64) *ItemUpdateOne {
	iuo.mutation.ResetValue()
	iuo.mutation.SetValue(f)
	return iuo
}

// AddValue adds f to the "value" field.
func (iuo *ItemUpdateOne) AddValue(f float64) *ItemUpdateOne {
	iuo.mutation.AddValue(f)
	return iuo
}

// SetTimestamp sets the "timestamp" field.
func (iuo *ItemUpdateOne) SetTimestamp(t time.Time) *ItemUpdateOne {
	iuo.mutation.SetTimestamp(t)
	return iuo
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (iuo *ItemUpdateOne) SetNillableTimestamp(t *time.Time) *ItemUpdateOne {
	if t != nil {
		iuo.SetTimestamp(*t)
	}
	return iuo
}

// SetMeta sets the "meta" field.
func (iuo *ItemUpdateOne) SetMeta(m map[string]interface{}) *ItemUpdateOne {
	iuo.mutation.SetMeta(m)
	return iuo
}

// ClearMeta clears the value of the "meta" field.
func (iuo *ItemUpdateOne) ClearMeta() *ItemUpdateOne {
	iuo.mutation.ClearMeta()
	return iuo
}

// AddDimensionIDs adds the "dimensions" edge to the Dimension entity by IDs.
func (iuo *ItemUpdateOne) AddDimensionIDs(ids ...int) *ItemUpdateOne {
	iuo.mutation.AddDimensionIDs(ids...)
	return iuo
}

// AddDimensions adds the "dimensions" edges to the Dimension entity.
func (iuo *ItemUpdateOne) AddDimensions(d ...*Dimension) *ItemUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return iuo.AddDimensionIDs(ids...)
}

// SetTaskInstanceID sets the "task_instance" edge to the TaskInstance entity by ID.
func (iuo *ItemUpdateOne) SetTaskInstanceID(id int) *ItemUpdateOne {
	iuo.mutation.SetTaskInstanceID(id)
	return iuo
}

// SetTaskInstance sets the "task_instance" edge to the TaskInstance entity.
func (iuo *ItemUpdateOne) SetTaskInstance(t *TaskInstance) *ItemUpdateOne {
	return iuo.SetTaskInstanceID(t.ID)
}

// SetMetricID sets the "metric" edge to the Metric entity by ID.
func (iuo *ItemUpdateOne) SetMetricID(id int) *ItemUpdateOne {
	iuo.mutation.SetMetricID(id)
	return iuo
}

// SetMetric sets the "metric" edge to the Metric entity.
func (iuo *ItemUpdateOne) SetMetric(m *Metric) *ItemUpdateOne {
	return iuo.SetMetricID(m.ID)
}

// Mutation returns the ItemMutation object of the builder.
func (iuo *ItemUpdateOne) Mutation() *ItemMutation {
	return iuo.mutation
}

// ClearDimensions clears all "dimensions" edges to the Dimension entity.
func (iuo *ItemUpdateOne) ClearDimensions() *ItemUpdateOne {
	iuo.mutation.ClearDimensions()
	return iuo
}

// RemoveDimensionIDs removes the "dimensions" edge to Dimension entities by IDs.
func (iuo *ItemUpdateOne) RemoveDimensionIDs(ids ...int) *ItemUpdateOne {
	iuo.mutation.RemoveDimensionIDs(ids...)
	return iuo
}

// RemoveDimensions removes "dimensions" edges to Dimension entities.
func (iuo *ItemUpdateOne) RemoveDimensions(d ...*Dimension) *ItemUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return iuo.RemoveDimensionIDs(ids...)
}

// ClearTaskInstance clears the "task_instance" edge to the TaskInstance entity.
func (iuo *ItemUpdateOne) ClearTaskInstance() *ItemUpdateOne {
	iuo.mutation.ClearTaskInstance()
	return iuo
}

// ClearMetric clears the "metric" edge to the Metric entity.
func (iuo *ItemUpdateOne) ClearMetric() *ItemUpdateOne {
	iuo.mutation.ClearMetric()
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ItemUpdateOne) Select(field string, fields ...string) *ItemUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Item entity.
func (iuo *ItemUpdateOne) Save(ctx context.Context) (*Item, error) {
	var (
		err  error
		node *Item
	)
	iuo.defaults()
	if len(iuo.hooks) == 0 {
		if err = iuo.check(); err != nil {
			return nil, err
		}
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iuo.check(); err != nil {
				return nil, err
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			if iuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ItemUpdateOne) SaveX(ctx context.Context) *Item {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ItemUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ItemUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *ItemUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdateTime(); !ok {
		v := item.UpdateDefaultUpdateTime()
		iuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *ItemUpdateOne) check() error {
	if _, ok := iuo.mutation.TaskInstanceID(); iuo.mutation.TaskInstanceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Item.task_instance"`)
	}
	if _, ok := iuo.mutation.MetricID(); iuo.mutation.MetricCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Item.metric"`)
	}
	return nil
}

func (iuo *ItemUpdateOne) sqlSave(ctx context.Context) (_node *Item, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   item.Table,
			Columns: item.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: item.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Item.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, item.FieldID)
		for _, f := range fields {
			if !item.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != item.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldUpdateTime,
		})
	}
	if value, ok := iuo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: item.FieldValue,
		})
	}
	if value, ok := iuo.mutation.AddedValue(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: item.FieldValue,
		})
	}
	if value, ok := iuo.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldTimestamp,
		})
	}
	if value, ok := iuo.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: item.FieldMeta,
		})
	}
	if iuo.mutation.MetaCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: item.FieldMeta,
		})
	}
	if iuo.mutation.DimensionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   item.DimensionsTable,
			Columns: item.DimensionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dimension.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedDimensionsIDs(); len(nodes) > 0 && !iuo.mutation.DimensionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   item.DimensionsTable,
			Columns: item.DimensionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dimension.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.DimensionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   item.DimensionsTable,
			Columns: item.DimensionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dimension.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.TaskInstanceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.TaskInstanceTable,
			Columns: []string{item.TaskInstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: taskinstance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.TaskInstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.TaskInstanceTable,
			Columns: []string{item.TaskInstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: taskinstance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.MetricCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.MetricTable,
			Columns: []string{item.MetricColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metric.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.MetricIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.MetricTable,
			Columns: []string{item.MetricColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metric.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Item{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
