// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/dimension"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/item"
	"github.com/DanielTitkov/dashboars/internal/repository/entgo/ent/predicate"
)

// DimensionUpdate is the builder for updating Dimension entities.
type DimensionUpdate struct {
	config
	hooks    []Hook
	mutation *DimensionMutation
}

// Where appends a list predicates to the DimensionUpdate builder.
func (du *DimensionUpdate) Where(ps ...predicate.Dimension) *DimensionUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetTitle sets the "title" field.
func (du *DimensionUpdate) SetTitle(s string) *DimensionUpdate {
	du.mutation.SetTitle(s)
	return du
}

// SetValue sets the "value" field.
func (du *DimensionUpdate) SetValue(s string) *DimensionUpdate {
	du.mutation.SetValue(s)
	return du
}

// SetMeta sets the "meta" field.
func (du *DimensionUpdate) SetMeta(m map[string]interface{}) *DimensionUpdate {
	du.mutation.SetMeta(m)
	return du
}

// ClearMeta clears the value of the "meta" field.
func (du *DimensionUpdate) ClearMeta() *DimensionUpdate {
	du.mutation.ClearMeta()
	return du
}

// AddItemIDs adds the "item" edge to the Item entity by IDs.
func (du *DimensionUpdate) AddItemIDs(ids ...int) *DimensionUpdate {
	du.mutation.AddItemIDs(ids...)
	return du
}

// AddItem adds the "item" edges to the Item entity.
func (du *DimensionUpdate) AddItem(i ...*Item) *DimensionUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return du.AddItemIDs(ids...)
}

// Mutation returns the DimensionMutation object of the builder.
func (du *DimensionUpdate) Mutation() *DimensionMutation {
	return du.mutation
}

// ClearItem clears all "item" edges to the Item entity.
func (du *DimensionUpdate) ClearItem() *DimensionUpdate {
	du.mutation.ClearItem()
	return du
}

// RemoveItemIDs removes the "item" edge to Item entities by IDs.
func (du *DimensionUpdate) RemoveItemIDs(ids ...int) *DimensionUpdate {
	du.mutation.RemoveItemIDs(ids...)
	return du
}

// RemoveItem removes "item" edges to Item entities.
func (du *DimensionUpdate) RemoveItem(i ...*Item) *DimensionUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return du.RemoveItemIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DimensionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		if err = du.check(); err != nil {
			return 0, err
		}
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DimensionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = du.check(); err != nil {
				return 0, err
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DimensionUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DimensionUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DimensionUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DimensionUpdate) check() error {
	if v, ok := du.mutation.Title(); ok {
		if err := dimension.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Dimension.title": %w`, err)}
		}
	}
	if v, ok := du.mutation.Value(); ok {
		if err := dimension.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "Dimension.value": %w`, err)}
		}
	}
	return nil
}

func (du *DimensionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dimension.Table,
			Columns: dimension.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dimension.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dimension.FieldTitle,
		})
	}
	if value, ok := du.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dimension.FieldValue,
		})
	}
	if value, ok := du.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dimension.FieldMeta,
		})
	}
	if du.mutation.MetaCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: dimension.FieldMeta,
		})
	}
	if du.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dimension.ItemTable,
			Columns: dimension.ItemPrimaryKey,
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
	if nodes := du.mutation.RemovedItemIDs(); len(nodes) > 0 && !du.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dimension.ItemTable,
			Columns: dimension.ItemPrimaryKey,
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
	if nodes := du.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dimension.ItemTable,
			Columns: dimension.ItemPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dimension.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DimensionUpdateOne is the builder for updating a single Dimension entity.
type DimensionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DimensionMutation
}

// SetTitle sets the "title" field.
func (duo *DimensionUpdateOne) SetTitle(s string) *DimensionUpdateOne {
	duo.mutation.SetTitle(s)
	return duo
}

// SetValue sets the "value" field.
func (duo *DimensionUpdateOne) SetValue(s string) *DimensionUpdateOne {
	duo.mutation.SetValue(s)
	return duo
}

// SetMeta sets the "meta" field.
func (duo *DimensionUpdateOne) SetMeta(m map[string]interface{}) *DimensionUpdateOne {
	duo.mutation.SetMeta(m)
	return duo
}

// ClearMeta clears the value of the "meta" field.
func (duo *DimensionUpdateOne) ClearMeta() *DimensionUpdateOne {
	duo.mutation.ClearMeta()
	return duo
}

// AddItemIDs adds the "item" edge to the Item entity by IDs.
func (duo *DimensionUpdateOne) AddItemIDs(ids ...int) *DimensionUpdateOne {
	duo.mutation.AddItemIDs(ids...)
	return duo
}

// AddItem adds the "item" edges to the Item entity.
func (duo *DimensionUpdateOne) AddItem(i ...*Item) *DimensionUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return duo.AddItemIDs(ids...)
}

// Mutation returns the DimensionMutation object of the builder.
func (duo *DimensionUpdateOne) Mutation() *DimensionMutation {
	return duo.mutation
}

// ClearItem clears all "item" edges to the Item entity.
func (duo *DimensionUpdateOne) ClearItem() *DimensionUpdateOne {
	duo.mutation.ClearItem()
	return duo
}

// RemoveItemIDs removes the "item" edge to Item entities by IDs.
func (duo *DimensionUpdateOne) RemoveItemIDs(ids ...int) *DimensionUpdateOne {
	duo.mutation.RemoveItemIDs(ids...)
	return duo
}

// RemoveItem removes "item" edges to Item entities.
func (duo *DimensionUpdateOne) RemoveItem(i ...*Item) *DimensionUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return duo.RemoveItemIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DimensionUpdateOne) Select(field string, fields ...string) *DimensionUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Dimension entity.
func (duo *DimensionUpdateOne) Save(ctx context.Context) (*Dimension, error) {
	var (
		err  error
		node *Dimension
	)
	if len(duo.hooks) == 0 {
		if err = duo.check(); err != nil {
			return nil, err
		}
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DimensionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = duo.check(); err != nil {
				return nil, err
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DimensionUpdateOne) SaveX(ctx context.Context) *Dimension {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DimensionUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DimensionUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DimensionUpdateOne) check() error {
	if v, ok := duo.mutation.Title(); ok {
		if err := dimension.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Dimension.title": %w`, err)}
		}
	}
	if v, ok := duo.mutation.Value(); ok {
		if err := dimension.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "Dimension.value": %w`, err)}
		}
	}
	return nil
}

func (duo *DimensionUpdateOne) sqlSave(ctx context.Context) (_node *Dimension, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dimension.Table,
			Columns: dimension.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dimension.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Dimension.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dimension.FieldID)
		for _, f := range fields {
			if !dimension.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dimension.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dimension.FieldTitle,
		})
	}
	if value, ok := duo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dimension.FieldValue,
		})
	}
	if value, ok := duo.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dimension.FieldMeta,
		})
	}
	if duo.mutation.MetaCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: dimension.FieldMeta,
		})
	}
	if duo.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dimension.ItemTable,
			Columns: dimension.ItemPrimaryKey,
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
	if nodes := duo.mutation.RemovedItemIDs(); len(nodes) > 0 && !duo.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dimension.ItemTable,
			Columns: dimension.ItemPrimaryKey,
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
	if nodes := duo.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dimension.ItemTable,
			Columns: dimension.ItemPrimaryKey,
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
	_node = &Dimension{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dimension.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}