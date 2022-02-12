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
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/taskinstance"
)

// ItemCreate is the builder for creating a Item entity.
type ItemCreate struct {
	config
	mutation *ItemMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (ic *ItemCreate) SetCreateTime(t time.Time) *ItemCreate {
	ic.mutation.SetCreateTime(t)
	return ic
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ic *ItemCreate) SetNillableCreateTime(t *time.Time) *ItemCreate {
	if t != nil {
		ic.SetCreateTime(*t)
	}
	return ic
}

// SetUpdateTime sets the "update_time" field.
func (ic *ItemCreate) SetUpdateTime(t time.Time) *ItemCreate {
	ic.mutation.SetUpdateTime(t)
	return ic
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ic *ItemCreate) SetNillableUpdateTime(t *time.Time) *ItemCreate {
	if t != nil {
		ic.SetUpdateTime(*t)
	}
	return ic
}

// SetValue sets the "value" field.
func (ic *ItemCreate) SetValue(f float64) *ItemCreate {
	ic.mutation.SetValue(f)
	return ic
}

// SetTimestamp sets the "timestamp" field.
func (ic *ItemCreate) SetTimestamp(t time.Time) *ItemCreate {
	ic.mutation.SetTimestamp(t)
	return ic
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (ic *ItemCreate) SetNillableTimestamp(t *time.Time) *ItemCreate {
	if t != nil {
		ic.SetTimestamp(*t)
	}
	return ic
}

// SetMeta sets the "meta" field.
func (ic *ItemCreate) SetMeta(m map[string]interface{}) *ItemCreate {
	ic.mutation.SetMeta(m)
	return ic
}

// SetTaskInstanceID sets the "task_instance" edge to the TaskInstance entity by ID.
func (ic *ItemCreate) SetTaskInstanceID(id int) *ItemCreate {
	ic.mutation.SetTaskInstanceID(id)
	return ic
}

// SetTaskInstance sets the "task_instance" edge to the TaskInstance entity.
func (ic *ItemCreate) SetTaskInstance(t *TaskInstance) *ItemCreate {
	return ic.SetTaskInstanceID(t.ID)
}

// Mutation returns the ItemMutation object of the builder.
func (ic *ItemCreate) Mutation() *ItemMutation {
	return ic.mutation
}

// Save creates the Item in the database.
func (ic *ItemCreate) Save(ctx context.Context) (*Item, error) {
	var (
		err  error
		node *Item
	)
	ic.defaults()
	if len(ic.hooks) == 0 {
		if err = ic.check(); err != nil {
			return nil, err
		}
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ic.check(); err != nil {
				return nil, err
			}
			ic.mutation = mutation
			if node, err = ic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			if ic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ItemCreate) SaveX(ctx context.Context) *Item {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *ItemCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *ItemCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *ItemCreate) defaults() {
	if _, ok := ic.mutation.CreateTime(); !ok {
		v := item.DefaultCreateTime()
		ic.mutation.SetCreateTime(v)
	}
	if _, ok := ic.mutation.UpdateTime(); !ok {
		v := item.DefaultUpdateTime()
		ic.mutation.SetUpdateTime(v)
	}
	if _, ok := ic.mutation.Timestamp(); !ok {
		v := item.DefaultTimestamp()
		ic.mutation.SetTimestamp(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *ItemCreate) check() error {
	if _, ok := ic.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Item.create_time"`)}
	}
	if _, ok := ic.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Item.update_time"`)}
	}
	if _, ok := ic.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "Item.value"`)}
	}
	if _, ok := ic.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "timestamp", err: errors.New(`ent: missing required field "Item.timestamp"`)}
	}
	if _, ok := ic.mutation.TaskInstanceID(); !ok {
		return &ValidationError{Name: "task_instance", err: errors.New(`ent: missing required edge "Item.task_instance"`)}
	}
	return nil
}

func (ic *ItemCreate) sqlSave(ctx context.Context) (*Item, error) {
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ic *ItemCreate) createSpec() (*Item, *sqlgraph.CreateSpec) {
	var (
		_node = &Item{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: item.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: item.FieldID,
			},
		}
	)
	if value, ok := ic.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ic.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := ic.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: item.FieldValue,
		})
		_node.Value = value
	}
	if value, ok := ic.mutation.Timestamp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldTimestamp,
		})
		_node.Timestamp = value
	}
	if value, ok := ic.mutation.Meta(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: item.FieldMeta,
		})
		_node.Meta = value
	}
	if nodes := ic.mutation.TaskInstanceIDs(); len(nodes) > 0 {
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
		_node.task_instance_items = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ItemCreateBulk is the builder for creating many Item entities in bulk.
type ItemCreateBulk struct {
	config
	builders []*ItemCreate
}

// Save creates the Item entities in the database.
func (icb *ItemCreateBulk) Save(ctx context.Context) ([]*Item, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Item, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ItemMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *ItemCreateBulk) SaveX(ctx context.Context) []*Item {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *ItemCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *ItemCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
