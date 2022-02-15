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
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/metric"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/task"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/taskcategory"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/taskinstance"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/tasktag"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdateTime sets the "update_time" field.
func (tu *TaskUpdate) SetUpdateTime(t time.Time) *TaskUpdate {
	tu.mutation.SetUpdateTime(t)
	return tu
}

// SetTitle sets the "title" field.
func (tu *TaskUpdate) SetTitle(s string) *TaskUpdate {
	tu.mutation.SetTitle(s)
	return tu
}

// SetDescription sets the "description" field.
func (tu *TaskUpdate) SetDescription(s string) *TaskUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableDescription(s *string) *TaskUpdate {
	if s != nil {
		tu.SetDescription(*s)
	}
	return tu
}

// ClearDescription clears the value of the "description" field.
func (tu *TaskUpdate) ClearDescription() *TaskUpdate {
	tu.mutation.ClearDescription()
	return tu
}

// SetActive sets the "active" field.
func (tu *TaskUpdate) SetActive(b bool) *TaskUpdate {
	tu.mutation.SetActive(b)
	return tu
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableActive(b *bool) *TaskUpdate {
	if b != nil {
		tu.SetActive(*b)
	}
	return tu
}

// SetDisplay sets the "display" field.
func (tu *TaskUpdate) SetDisplay(b bool) *TaskUpdate {
	tu.mutation.SetDisplay(b)
	return tu
}

// SetNillableDisplay sets the "display" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableDisplay(b *bool) *TaskUpdate {
	if b != nil {
		tu.SetDisplay(*b)
	}
	return tu
}

// SetSchedule sets the "schedule" field.
func (tu *TaskUpdate) SetSchedule(s string) *TaskUpdate {
	tu.mutation.SetSchedule(s)
	return tu
}

// SetNillableSchedule sets the "schedule" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableSchedule(s *string) *TaskUpdate {
	if s != nil {
		tu.SetSchedule(*s)
	}
	return tu
}

// ClearSchedule clears the value of the "schedule" field.
func (tu *TaskUpdate) ClearSchedule() *TaskUpdate {
	tu.mutation.ClearSchedule()
	return tu
}

// SetArgs sets the "args" field.
func (tu *TaskUpdate) SetArgs(m map[string]interface{}) *TaskUpdate {
	tu.mutation.SetArgs(m)
	return tu
}

// ClearArgs clears the value of the "args" field.
func (tu *TaskUpdate) ClearArgs() *TaskUpdate {
	tu.mutation.ClearArgs()
	return tu
}

// AddInstanceIDs adds the "instances" edge to the TaskInstance entity by IDs.
func (tu *TaskUpdate) AddInstanceIDs(ids ...int) *TaskUpdate {
	tu.mutation.AddInstanceIDs(ids...)
	return tu
}

// AddInstances adds the "instances" edges to the TaskInstance entity.
func (tu *TaskUpdate) AddInstances(t ...*TaskInstance) *TaskUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddInstanceIDs(ids...)
}

// AddMetricIDs adds the "metrics" edge to the Metric entity by IDs.
func (tu *TaskUpdate) AddMetricIDs(ids ...int) *TaskUpdate {
	tu.mutation.AddMetricIDs(ids...)
	return tu
}

// AddMetrics adds the "metrics" edges to the Metric entity.
func (tu *TaskUpdate) AddMetrics(m ...*Metric) *TaskUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return tu.AddMetricIDs(ids...)
}

// SetCategoryID sets the "category" edge to the TaskCategory entity by ID.
func (tu *TaskUpdate) SetCategoryID(id int) *TaskUpdate {
	tu.mutation.SetCategoryID(id)
	return tu
}

// SetCategory sets the "category" edge to the TaskCategory entity.
func (tu *TaskUpdate) SetCategory(t *TaskCategory) *TaskUpdate {
	return tu.SetCategoryID(t.ID)
}

// AddTagIDs adds the "tags" edge to the TaskTag entity by IDs.
func (tu *TaskUpdate) AddTagIDs(ids ...int) *TaskUpdate {
	tu.mutation.AddTagIDs(ids...)
	return tu
}

// AddTags adds the "tags" edges to the TaskTag entity.
func (tu *TaskUpdate) AddTags(t ...*TaskTag) *TaskUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTagIDs(ids...)
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// ClearInstances clears all "instances" edges to the TaskInstance entity.
func (tu *TaskUpdate) ClearInstances() *TaskUpdate {
	tu.mutation.ClearInstances()
	return tu
}

// RemoveInstanceIDs removes the "instances" edge to TaskInstance entities by IDs.
func (tu *TaskUpdate) RemoveInstanceIDs(ids ...int) *TaskUpdate {
	tu.mutation.RemoveInstanceIDs(ids...)
	return tu
}

// RemoveInstances removes "instances" edges to TaskInstance entities.
func (tu *TaskUpdate) RemoveInstances(t ...*TaskInstance) *TaskUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveInstanceIDs(ids...)
}

// ClearMetrics clears all "metrics" edges to the Metric entity.
func (tu *TaskUpdate) ClearMetrics() *TaskUpdate {
	tu.mutation.ClearMetrics()
	return tu
}

// RemoveMetricIDs removes the "metrics" edge to Metric entities by IDs.
func (tu *TaskUpdate) RemoveMetricIDs(ids ...int) *TaskUpdate {
	tu.mutation.RemoveMetricIDs(ids...)
	return tu
}

// RemoveMetrics removes "metrics" edges to Metric entities.
func (tu *TaskUpdate) RemoveMetrics(m ...*Metric) *TaskUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return tu.RemoveMetricIDs(ids...)
}

// ClearCategory clears the "category" edge to the TaskCategory entity.
func (tu *TaskUpdate) ClearCategory() *TaskUpdate {
	tu.mutation.ClearCategory()
	return tu
}

// ClearTags clears all "tags" edges to the TaskTag entity.
func (tu *TaskUpdate) ClearTags() *TaskUpdate {
	tu.mutation.ClearTags()
	return tu
}

// RemoveTagIDs removes the "tags" edge to TaskTag entities by IDs.
func (tu *TaskUpdate) RemoveTagIDs(ids ...int) *TaskUpdate {
	tu.mutation.RemoveTagIDs(ids...)
	return tu
}

// RemoveTags removes "tags" edges to TaskTag entities.
func (tu *TaskUpdate) RemoveTags(t ...*TaskTag) *TaskUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tu.defaults()
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TaskUpdate) defaults() {
	if _, ok := tu.mutation.UpdateTime(); !ok {
		v := task.UpdateDefaultUpdateTime()
		tu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TaskUpdate) check() error {
	if v, ok := tu.mutation.Title(); ok {
		if err := task.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Task.title": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Description(); ok {
		if err := task.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Task.description": %w`, err)}
		}
	}
	if _, ok := tu.mutation.CategoryID(); tu.mutation.CategoryCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Task.category"`)
	}
	return nil
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   task.Table,
			Columns: task.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldUpdateTime,
		})
	}
	if value, ok := tu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldTitle,
		})
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldDescription,
		})
	}
	if tu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldDescription,
		})
	}
	if value, ok := tu.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: task.FieldActive,
		})
	}
	if value, ok := tu.mutation.Display(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: task.FieldDisplay,
		})
	}
	if value, ok := tu.mutation.Schedule(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldSchedule,
		})
	}
	if tu.mutation.ScheduleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldSchedule,
		})
	}
	if value, ok := tu.mutation.Args(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: task.FieldArgs,
		})
	}
	if tu.mutation.ArgsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: task.FieldArgs,
		})
	}
	if tu.mutation.InstancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.InstancesTable,
			Columns: []string{task.InstancesColumn},
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
	if nodes := tu.mutation.RemovedInstancesIDs(); len(nodes) > 0 && !tu.mutation.InstancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.InstancesTable,
			Columns: []string{task.InstancesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.InstancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.InstancesTable,
			Columns: []string{task.InstancesColumn},
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
	if tu.mutation.MetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.MetricsTable,
			Columns: []string{task.MetricsColumn},
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
	if nodes := tu.mutation.RemovedMetricsIDs(); len(nodes) > 0 && !tu.mutation.MetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.MetricsTable,
			Columns: []string{task.MetricsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.MetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.MetricsTable,
			Columns: []string{task.MetricsColumn},
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
	if tu.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.CategoryTable,
			Columns: []string{task.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: taskcategory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.CategoryTable,
			Columns: []string{task.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: taskcategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   task.TagsTable,
			Columns: task.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasktag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTagsIDs(); len(nodes) > 0 && !tu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   task.TagsTable,
			Columns: task.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasktag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   task.TagsTable,
			Columns: task.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasktag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskMutation
}

// SetUpdateTime sets the "update_time" field.
func (tuo *TaskUpdateOne) SetUpdateTime(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetUpdateTime(t)
	return tuo
}

// SetTitle sets the "title" field.
func (tuo *TaskUpdateOne) SetTitle(s string) *TaskUpdateOne {
	tuo.mutation.SetTitle(s)
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TaskUpdateOne) SetDescription(s string) *TaskUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableDescription(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetDescription(*s)
	}
	return tuo
}

// ClearDescription clears the value of the "description" field.
func (tuo *TaskUpdateOne) ClearDescription() *TaskUpdateOne {
	tuo.mutation.ClearDescription()
	return tuo
}

// SetActive sets the "active" field.
func (tuo *TaskUpdateOne) SetActive(b bool) *TaskUpdateOne {
	tuo.mutation.SetActive(b)
	return tuo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableActive(b *bool) *TaskUpdateOne {
	if b != nil {
		tuo.SetActive(*b)
	}
	return tuo
}

// SetDisplay sets the "display" field.
func (tuo *TaskUpdateOne) SetDisplay(b bool) *TaskUpdateOne {
	tuo.mutation.SetDisplay(b)
	return tuo
}

// SetNillableDisplay sets the "display" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableDisplay(b *bool) *TaskUpdateOne {
	if b != nil {
		tuo.SetDisplay(*b)
	}
	return tuo
}

// SetSchedule sets the "schedule" field.
func (tuo *TaskUpdateOne) SetSchedule(s string) *TaskUpdateOne {
	tuo.mutation.SetSchedule(s)
	return tuo
}

// SetNillableSchedule sets the "schedule" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableSchedule(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetSchedule(*s)
	}
	return tuo
}

// ClearSchedule clears the value of the "schedule" field.
func (tuo *TaskUpdateOne) ClearSchedule() *TaskUpdateOne {
	tuo.mutation.ClearSchedule()
	return tuo
}

// SetArgs sets the "args" field.
func (tuo *TaskUpdateOne) SetArgs(m map[string]interface{}) *TaskUpdateOne {
	tuo.mutation.SetArgs(m)
	return tuo
}

// ClearArgs clears the value of the "args" field.
func (tuo *TaskUpdateOne) ClearArgs() *TaskUpdateOne {
	tuo.mutation.ClearArgs()
	return tuo
}

// AddInstanceIDs adds the "instances" edge to the TaskInstance entity by IDs.
func (tuo *TaskUpdateOne) AddInstanceIDs(ids ...int) *TaskUpdateOne {
	tuo.mutation.AddInstanceIDs(ids...)
	return tuo
}

// AddInstances adds the "instances" edges to the TaskInstance entity.
func (tuo *TaskUpdateOne) AddInstances(t ...*TaskInstance) *TaskUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddInstanceIDs(ids...)
}

// AddMetricIDs adds the "metrics" edge to the Metric entity by IDs.
func (tuo *TaskUpdateOne) AddMetricIDs(ids ...int) *TaskUpdateOne {
	tuo.mutation.AddMetricIDs(ids...)
	return tuo
}

// AddMetrics adds the "metrics" edges to the Metric entity.
func (tuo *TaskUpdateOne) AddMetrics(m ...*Metric) *TaskUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return tuo.AddMetricIDs(ids...)
}

// SetCategoryID sets the "category" edge to the TaskCategory entity by ID.
func (tuo *TaskUpdateOne) SetCategoryID(id int) *TaskUpdateOne {
	tuo.mutation.SetCategoryID(id)
	return tuo
}

// SetCategory sets the "category" edge to the TaskCategory entity.
func (tuo *TaskUpdateOne) SetCategory(t *TaskCategory) *TaskUpdateOne {
	return tuo.SetCategoryID(t.ID)
}

// AddTagIDs adds the "tags" edge to the TaskTag entity by IDs.
func (tuo *TaskUpdateOne) AddTagIDs(ids ...int) *TaskUpdateOne {
	tuo.mutation.AddTagIDs(ids...)
	return tuo
}

// AddTags adds the "tags" edges to the TaskTag entity.
func (tuo *TaskUpdateOne) AddTags(t ...*TaskTag) *TaskUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTagIDs(ids...)
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// ClearInstances clears all "instances" edges to the TaskInstance entity.
func (tuo *TaskUpdateOne) ClearInstances() *TaskUpdateOne {
	tuo.mutation.ClearInstances()
	return tuo
}

// RemoveInstanceIDs removes the "instances" edge to TaskInstance entities by IDs.
func (tuo *TaskUpdateOne) RemoveInstanceIDs(ids ...int) *TaskUpdateOne {
	tuo.mutation.RemoveInstanceIDs(ids...)
	return tuo
}

// RemoveInstances removes "instances" edges to TaskInstance entities.
func (tuo *TaskUpdateOne) RemoveInstances(t ...*TaskInstance) *TaskUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveInstanceIDs(ids...)
}

// ClearMetrics clears all "metrics" edges to the Metric entity.
func (tuo *TaskUpdateOne) ClearMetrics() *TaskUpdateOne {
	tuo.mutation.ClearMetrics()
	return tuo
}

// RemoveMetricIDs removes the "metrics" edge to Metric entities by IDs.
func (tuo *TaskUpdateOne) RemoveMetricIDs(ids ...int) *TaskUpdateOne {
	tuo.mutation.RemoveMetricIDs(ids...)
	return tuo
}

// RemoveMetrics removes "metrics" edges to Metric entities.
func (tuo *TaskUpdateOne) RemoveMetrics(m ...*Metric) *TaskUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return tuo.RemoveMetricIDs(ids...)
}

// ClearCategory clears the "category" edge to the TaskCategory entity.
func (tuo *TaskUpdateOne) ClearCategory() *TaskUpdateOne {
	tuo.mutation.ClearCategory()
	return tuo
}

// ClearTags clears all "tags" edges to the TaskTag entity.
func (tuo *TaskUpdateOne) ClearTags() *TaskUpdateOne {
	tuo.mutation.ClearTags()
	return tuo
}

// RemoveTagIDs removes the "tags" edge to TaskTag entities by IDs.
func (tuo *TaskUpdateOne) RemoveTagIDs(ids ...int) *TaskUpdateOne {
	tuo.mutation.RemoveTagIDs(ids...)
	return tuo
}

// RemoveTags removes "tags" edges to TaskTag entities.
func (tuo *TaskUpdateOne) RemoveTags(t ...*TaskTag) *TaskUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTagIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	var (
		err  error
		node *Task
	)
	tuo.defaults()
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TaskUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdateTime(); !ok {
		v := task.UpdateDefaultUpdateTime()
		tuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TaskUpdateOne) check() error {
	if v, ok := tuo.mutation.Title(); ok {
		if err := task.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Task.title": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Description(); ok {
		if err := task.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Task.description": %w`, err)}
		}
	}
	if _, ok := tuo.mutation.CategoryID(); tuo.mutation.CategoryCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Task.category"`)
	}
	return nil
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   task.Table,
			Columns: task.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Task.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for _, f := range fields {
			if !task.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != task.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldUpdateTime,
		})
	}
	if value, ok := tuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldTitle,
		})
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldDescription,
		})
	}
	if tuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldDescription,
		})
	}
	if value, ok := tuo.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: task.FieldActive,
		})
	}
	if value, ok := tuo.mutation.Display(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: task.FieldDisplay,
		})
	}
	if value, ok := tuo.mutation.Schedule(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldSchedule,
		})
	}
	if tuo.mutation.ScheduleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldSchedule,
		})
	}
	if value, ok := tuo.mutation.Args(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: task.FieldArgs,
		})
	}
	if tuo.mutation.ArgsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: task.FieldArgs,
		})
	}
	if tuo.mutation.InstancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.InstancesTable,
			Columns: []string{task.InstancesColumn},
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
	if nodes := tuo.mutation.RemovedInstancesIDs(); len(nodes) > 0 && !tuo.mutation.InstancesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.InstancesTable,
			Columns: []string{task.InstancesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.InstancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.InstancesTable,
			Columns: []string{task.InstancesColumn},
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
	if tuo.mutation.MetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.MetricsTable,
			Columns: []string{task.MetricsColumn},
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
	if nodes := tuo.mutation.RemovedMetricsIDs(); len(nodes) > 0 && !tuo.mutation.MetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.MetricsTable,
			Columns: []string{task.MetricsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.MetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.MetricsTable,
			Columns: []string{task.MetricsColumn},
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
	if tuo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.CategoryTable,
			Columns: []string{task.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: taskcategory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.CategoryTable,
			Columns: []string{task.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: taskcategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   task.TagsTable,
			Columns: task.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasktag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !tuo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   task.TagsTable,
			Columns: task.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasktag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   task.TagsTable,
			Columns: task.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasktag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Task{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
