// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/item"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/metric"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/task"
)

// MetricCreate is the builder for creating a Metric entity.
type MetricCreate struct {
	config
	mutation *MetricMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (mc *MetricCreate) SetCreateTime(t time.Time) *MetricCreate {
	mc.mutation.SetCreateTime(t)
	return mc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (mc *MetricCreate) SetNillableCreateTime(t *time.Time) *MetricCreate {
	if t != nil {
		mc.SetCreateTime(*t)
	}
	return mc
}

// SetUpdateTime sets the "update_time" field.
func (mc *MetricCreate) SetUpdateTime(t time.Time) *MetricCreate {
	mc.mutation.SetUpdateTime(t)
	return mc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (mc *MetricCreate) SetNillableUpdateTime(t *time.Time) *MetricCreate {
	if t != nil {
		mc.SetUpdateTime(*t)
	}
	return mc
}

// SetTitle sets the "title" field.
func (mc *MetricCreate) SetTitle(s string) *MetricCreate {
	mc.mutation.SetTitle(s)
	return mc
}

// SetDescription sets the "description" field.
func (mc *MetricCreate) SetDescription(s string) *MetricCreate {
	mc.mutation.SetDescription(s)
	return mc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mc *MetricCreate) SetNillableDescription(s *string) *MetricCreate {
	if s != nil {
		mc.SetDescription(*s)
	}
	return mc
}

// SetMeta sets the "meta" field.
func (mc *MetricCreate) SetMeta(m map[string]interface{}) *MetricCreate {
	mc.mutation.SetMeta(m)
	return mc
}

// AddItemIDs adds the "items" edge to the Item entity by IDs.
func (mc *MetricCreate) AddItemIDs(ids ...int) *MetricCreate {
	mc.mutation.AddItemIDs(ids...)
	return mc
}

// AddItems adds the "items" edges to the Item entity.
func (mc *MetricCreate) AddItems(i ...*Item) *MetricCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return mc.AddItemIDs(ids...)
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (mc *MetricCreate) SetTaskID(id int) *MetricCreate {
	mc.mutation.SetTaskID(id)
	return mc
}

// SetTask sets the "task" edge to the Task entity.
func (mc *MetricCreate) SetTask(t *Task) *MetricCreate {
	return mc.SetTaskID(t.ID)
}

// Mutation returns the MetricMutation object of the builder.
func (mc *MetricCreate) Mutation() *MetricMutation {
	return mc.mutation
}

// Save creates the Metric in the database.
func (mc *MetricCreate) Save(ctx context.Context) (*Metric, error) {
	var (
		err  error
		node *Metric
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MetricMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MetricCreate) SaveX(ctx context.Context) *Metric {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MetricCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MetricCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MetricCreate) defaults() {
	if _, ok := mc.mutation.CreateTime(); !ok {
		v := metric.DefaultCreateTime()
		mc.mutation.SetCreateTime(v)
	}
	if _, ok := mc.mutation.UpdateTime(); !ok {
		v := metric.DefaultUpdateTime()
		mc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MetricCreate) check() error {
	if _, ok := mc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Metric.create_time"`)}
	}
	if _, ok := mc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Metric.update_time"`)}
	}
	if _, ok := mc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Metric.title"`)}
	}
	if v, ok := mc.mutation.Title(); ok {
		if err := metric.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Metric.title": %w`, err)}
		}
	}
	if v, ok := mc.mutation.Description(); ok {
		if err := metric.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Metric.description": %w`, err)}
		}
	}
	if _, ok := mc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task", err: errors.New(`ent: missing required edge "Metric.task"`)}
	}
	return nil
}

func (mc *MetricCreate) sqlSave(ctx context.Context) (*Metric, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mc *MetricCreate) createSpec() (*Metric, *sqlgraph.CreateSpec) {
	var (
		_node = &Metric{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: metric.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metric.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metric.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := mc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metric.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := mc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: metric.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := mc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: metric.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := mc.mutation.Meta(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: metric.FieldMeta,
		})
		_node.Meta = value
	}
	if nodes := mc.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metric.ItemsTable,
			Columns: []string{metric.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   metric.TaskTable,
			Columns: []string{metric.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.task_metrics = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MetricCreateBulk is the builder for creating many Metric entities in bulk.
type MetricCreateBulk struct {
	config
	builders []*MetricCreate
}

// Save creates the Metric entities in the database.
func (mcb *MetricCreateBulk) Save(ctx context.Context) ([]*Metric, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Metric, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MetricMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MetricCreateBulk) SaveX(ctx context.Context) []*Metric {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MetricCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MetricCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}