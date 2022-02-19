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
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/item"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/metric"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/task"
)

// MetricUpdate is the builder for updating Metric entities.
type MetricUpdate struct {
	config
	hooks    []Hook
	mutation *MetricMutation
}

// Where appends a list predicates to the MetricUpdate builder.
func (mu *MetricUpdate) Where(ps ...predicate.Metric) *MetricUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUpdateTime sets the "update_time" field.
func (mu *MetricUpdate) SetUpdateTime(t time.Time) *MetricUpdate {
	mu.mutation.SetUpdateTime(t)
	return mu
}

// SetTitle sets the "title" field.
func (mu *MetricUpdate) SetTitle(s string) *MetricUpdate {
	mu.mutation.SetTitle(s)
	return mu
}

// SetDisplayTitle sets the "display_title" field.
func (mu *MetricUpdate) SetDisplayTitle(m map[string]string) *MetricUpdate {
	mu.mutation.SetDisplayTitle(m)
	return mu
}

// ClearDisplayTitle clears the value of the "display_title" field.
func (mu *MetricUpdate) ClearDisplayTitle() *MetricUpdate {
	mu.mutation.ClearDisplayTitle()
	return mu
}

// SetDescription sets the "description" field.
func (mu *MetricUpdate) SetDescription(s string) *MetricUpdate {
	mu.mutation.SetDescription(s)
	return mu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mu *MetricUpdate) SetNillableDescription(s *string) *MetricUpdate {
	if s != nil {
		mu.SetDescription(*s)
	}
	return mu
}

// ClearDescription clears the value of the "description" field.
func (mu *MetricUpdate) ClearDescription() *MetricUpdate {
	mu.mutation.ClearDescription()
	return mu
}

// SetMeta sets the "meta" field.
func (mu *MetricUpdate) SetMeta(m map[string]interface{}) *MetricUpdate {
	mu.mutation.SetMeta(m)
	return mu
}

// ClearMeta clears the value of the "meta" field.
func (mu *MetricUpdate) ClearMeta() *MetricUpdate {
	mu.mutation.ClearMeta()
	return mu
}

// AddItemIDs adds the "items" edge to the Item entity by IDs.
func (mu *MetricUpdate) AddItemIDs(ids ...int) *MetricUpdate {
	mu.mutation.AddItemIDs(ids...)
	return mu
}

// AddItems adds the "items" edges to the Item entity.
func (mu *MetricUpdate) AddItems(i ...*Item) *MetricUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return mu.AddItemIDs(ids...)
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (mu *MetricUpdate) SetTaskID(id int) *MetricUpdate {
	mu.mutation.SetTaskID(id)
	return mu
}

// SetTask sets the "task" edge to the Task entity.
func (mu *MetricUpdate) SetTask(t *Task) *MetricUpdate {
	return mu.SetTaskID(t.ID)
}

// Mutation returns the MetricMutation object of the builder.
func (mu *MetricUpdate) Mutation() *MetricMutation {
	return mu.mutation
}

// ClearItems clears all "items" edges to the Item entity.
func (mu *MetricUpdate) ClearItems() *MetricUpdate {
	mu.mutation.ClearItems()
	return mu
}

// RemoveItemIDs removes the "items" edge to Item entities by IDs.
func (mu *MetricUpdate) RemoveItemIDs(ids ...int) *MetricUpdate {
	mu.mutation.RemoveItemIDs(ids...)
	return mu
}

// RemoveItems removes "items" edges to Item entities.
func (mu *MetricUpdate) RemoveItems(i ...*Item) *MetricUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return mu.RemoveItemIDs(ids...)
}

// ClearTask clears the "task" edge to the Task entity.
func (mu *MetricUpdate) ClearTask() *MetricUpdate {
	mu.mutation.ClearTask()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MetricUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	mu.defaults()
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MetricMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MetricUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MetricUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MetricUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MetricUpdate) defaults() {
	if _, ok := mu.mutation.UpdateTime(); !ok {
		v := metric.UpdateDefaultUpdateTime()
		mu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MetricUpdate) check() error {
	if v, ok := mu.mutation.Title(); ok {
		if err := metric.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Metric.title": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Description(); ok {
		if err := metric.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Metric.description": %w`, err)}
		}
	}
	if _, ok := mu.mutation.TaskID(); mu.mutation.TaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Metric.task"`)
	}
	return nil
}

func (mu *MetricUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   metric.Table,
			Columns: metric.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metric.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metric.FieldUpdateTime,
		})
	}
	if value, ok := mu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: metric.FieldTitle,
		})
	}
	if value, ok := mu.mutation.DisplayTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: metric.FieldDisplayTitle,
		})
	}
	if mu.mutation.DisplayTitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: metric.FieldDisplayTitle,
		})
	}
	if value, ok := mu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: metric.FieldDescription,
		})
	}
	if mu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: metric.FieldDescription,
		})
	}
	if value, ok := mu.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: metric.FieldMeta,
		})
	}
	if mu.mutation.MetaCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: metric.FieldMeta,
		})
	}
	if mu.mutation.ItemsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedItemsIDs(); len(nodes) > 0 && !mu.mutation.ItemsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ItemsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.TaskCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.TaskIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{metric.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MetricUpdateOne is the builder for updating a single Metric entity.
type MetricUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MetricMutation
}

// SetUpdateTime sets the "update_time" field.
func (muo *MetricUpdateOne) SetUpdateTime(t time.Time) *MetricUpdateOne {
	muo.mutation.SetUpdateTime(t)
	return muo
}

// SetTitle sets the "title" field.
func (muo *MetricUpdateOne) SetTitle(s string) *MetricUpdateOne {
	muo.mutation.SetTitle(s)
	return muo
}

// SetDisplayTitle sets the "display_title" field.
func (muo *MetricUpdateOne) SetDisplayTitle(m map[string]string) *MetricUpdateOne {
	muo.mutation.SetDisplayTitle(m)
	return muo
}

// ClearDisplayTitle clears the value of the "display_title" field.
func (muo *MetricUpdateOne) ClearDisplayTitle() *MetricUpdateOne {
	muo.mutation.ClearDisplayTitle()
	return muo
}

// SetDescription sets the "description" field.
func (muo *MetricUpdateOne) SetDescription(s string) *MetricUpdateOne {
	muo.mutation.SetDescription(s)
	return muo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (muo *MetricUpdateOne) SetNillableDescription(s *string) *MetricUpdateOne {
	if s != nil {
		muo.SetDescription(*s)
	}
	return muo
}

// ClearDescription clears the value of the "description" field.
func (muo *MetricUpdateOne) ClearDescription() *MetricUpdateOne {
	muo.mutation.ClearDescription()
	return muo
}

// SetMeta sets the "meta" field.
func (muo *MetricUpdateOne) SetMeta(m map[string]interface{}) *MetricUpdateOne {
	muo.mutation.SetMeta(m)
	return muo
}

// ClearMeta clears the value of the "meta" field.
func (muo *MetricUpdateOne) ClearMeta() *MetricUpdateOne {
	muo.mutation.ClearMeta()
	return muo
}

// AddItemIDs adds the "items" edge to the Item entity by IDs.
func (muo *MetricUpdateOne) AddItemIDs(ids ...int) *MetricUpdateOne {
	muo.mutation.AddItemIDs(ids...)
	return muo
}

// AddItems adds the "items" edges to the Item entity.
func (muo *MetricUpdateOne) AddItems(i ...*Item) *MetricUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return muo.AddItemIDs(ids...)
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (muo *MetricUpdateOne) SetTaskID(id int) *MetricUpdateOne {
	muo.mutation.SetTaskID(id)
	return muo
}

// SetTask sets the "task" edge to the Task entity.
func (muo *MetricUpdateOne) SetTask(t *Task) *MetricUpdateOne {
	return muo.SetTaskID(t.ID)
}

// Mutation returns the MetricMutation object of the builder.
func (muo *MetricUpdateOne) Mutation() *MetricMutation {
	return muo.mutation
}

// ClearItems clears all "items" edges to the Item entity.
func (muo *MetricUpdateOne) ClearItems() *MetricUpdateOne {
	muo.mutation.ClearItems()
	return muo
}

// RemoveItemIDs removes the "items" edge to Item entities by IDs.
func (muo *MetricUpdateOne) RemoveItemIDs(ids ...int) *MetricUpdateOne {
	muo.mutation.RemoveItemIDs(ids...)
	return muo
}

// RemoveItems removes "items" edges to Item entities.
func (muo *MetricUpdateOne) RemoveItems(i ...*Item) *MetricUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return muo.RemoveItemIDs(ids...)
}

// ClearTask clears the "task" edge to the Task entity.
func (muo *MetricUpdateOne) ClearTask() *MetricUpdateOne {
	muo.mutation.ClearTask()
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MetricUpdateOne) Select(field string, fields ...string) *MetricUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Metric entity.
func (muo *MetricUpdateOne) Save(ctx context.Context) (*Metric, error) {
	var (
		err  error
		node *Metric
	)
	muo.defaults()
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MetricMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MetricUpdateOne) SaveX(ctx context.Context) *Metric {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MetricUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MetricUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MetricUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdateTime(); !ok {
		v := metric.UpdateDefaultUpdateTime()
		muo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MetricUpdateOne) check() error {
	if v, ok := muo.mutation.Title(); ok {
		if err := metric.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Metric.title": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Description(); ok {
		if err := metric.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Metric.description": %w`, err)}
		}
	}
	if _, ok := muo.mutation.TaskID(); muo.mutation.TaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Metric.task"`)
	}
	return nil
}

func (muo *MetricUpdateOne) sqlSave(ctx context.Context) (_node *Metric, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   metric.Table,
			Columns: metric.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: metric.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Metric.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, metric.FieldID)
		for _, f := range fields {
			if !metric.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != metric.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: metric.FieldUpdateTime,
		})
	}
	if value, ok := muo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: metric.FieldTitle,
		})
	}
	if value, ok := muo.mutation.DisplayTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: metric.FieldDisplayTitle,
		})
	}
	if muo.mutation.DisplayTitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: metric.FieldDisplayTitle,
		})
	}
	if value, ok := muo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: metric.FieldDescription,
		})
	}
	if muo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: metric.FieldDescription,
		})
	}
	if value, ok := muo.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: metric.FieldMeta,
		})
	}
	if muo.mutation.MetaCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: metric.FieldMeta,
		})
	}
	if muo.mutation.ItemsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedItemsIDs(); len(nodes) > 0 && !muo.mutation.ItemsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ItemsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.TaskCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.TaskIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Metric{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{metric.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
