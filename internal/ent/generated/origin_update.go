// Copyright Infratographer, Inc. and/or licensed to Infratographer, Inc. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/origin"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/predicate"
)

// OriginUpdate is the builder for updating Origin entities.
type OriginUpdate struct {
	config
	hooks    []Hook
	mutation *OriginMutation
}

// Where appends a list predicates to the OriginUpdate builder.
func (ou *OriginUpdate) Where(ps ...predicate.Origin) *OriginUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetName sets the "name" field.
func (ou *OriginUpdate) SetName(s string) *OriginUpdate {
	ou.mutation.SetName(s)
	return ou
}

// SetTarget sets the "target" field.
func (ou *OriginUpdate) SetTarget(s string) *OriginUpdate {
	ou.mutation.SetTarget(s)
	return ou
}

// SetPortNumber sets the "port_number" field.
func (ou *OriginUpdate) SetPortNumber(i int) *OriginUpdate {
	ou.mutation.ResetPortNumber()
	ou.mutation.SetPortNumber(i)
	return ou
}

// AddPortNumber adds i to the "port_number" field.
func (ou *OriginUpdate) AddPortNumber(i int) *OriginUpdate {
	ou.mutation.AddPortNumber(i)
	return ou
}

// SetActive sets the "active" field.
func (ou *OriginUpdate) SetActive(b bool) *OriginUpdate {
	ou.mutation.SetActive(b)
	return ou
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (ou *OriginUpdate) SetNillableActive(b *bool) *OriginUpdate {
	if b != nil {
		ou.SetActive(*b)
	}
	return ou
}

// Mutation returns the OriginMutation object of the builder.
func (ou *OriginUpdate) Mutation() *OriginMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OriginUpdate) Save(ctx context.Context) (int, error) {
	ou.defaults()
	return withHooks(ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OriginUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OriginUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OriginUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *OriginUpdate) defaults() {
	if _, ok := ou.mutation.UpdatedAt(); !ok {
		v := origin.UpdateDefaultUpdatedAt()
		ou.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OriginUpdate) check() error {
	if v, ok := ou.mutation.Name(); ok {
		if err := origin.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Origin.name": %w`, err)}
		}
	}
	if v, ok := ou.mutation.Target(); ok {
		if err := origin.TargetValidator(v); err != nil {
			return &ValidationError{Name: "target", err: fmt.Errorf(`generated: validator failed for field "Origin.target": %w`, err)}
		}
	}
	if v, ok := ou.mutation.PortNumber(); ok {
		if err := origin.PortNumberValidator(v); err != nil {
			return &ValidationError{Name: "port_number", err: fmt.Errorf(`generated: validator failed for field "Origin.port_number": %w`, err)}
		}
	}
	if _, ok := ou.mutation.PoolID(); ou.mutation.PoolCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Origin.pool"`)
	}
	return nil
}

func (ou *OriginUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(origin.Table, origin.Columns, sqlgraph.NewFieldSpec(origin.FieldID, field.TypeString))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.UpdatedAt(); ok {
		_spec.SetField(origin.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ou.mutation.Name(); ok {
		_spec.SetField(origin.FieldName, field.TypeString, value)
	}
	if value, ok := ou.mutation.Target(); ok {
		_spec.SetField(origin.FieldTarget, field.TypeString, value)
	}
	if value, ok := ou.mutation.PortNumber(); ok {
		_spec.SetField(origin.FieldPortNumber, field.TypeInt, value)
	}
	if value, ok := ou.mutation.AddedPortNumber(); ok {
		_spec.AddField(origin.FieldPortNumber, field.TypeInt, value)
	}
	if value, ok := ou.mutation.Active(); ok {
		_spec.SetField(origin.FieldActive, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{origin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OriginUpdateOne is the builder for updating a single Origin entity.
type OriginUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OriginMutation
}

// SetName sets the "name" field.
func (ouo *OriginUpdateOne) SetName(s string) *OriginUpdateOne {
	ouo.mutation.SetName(s)
	return ouo
}

// SetTarget sets the "target" field.
func (ouo *OriginUpdateOne) SetTarget(s string) *OriginUpdateOne {
	ouo.mutation.SetTarget(s)
	return ouo
}

// SetPortNumber sets the "port_number" field.
func (ouo *OriginUpdateOne) SetPortNumber(i int) *OriginUpdateOne {
	ouo.mutation.ResetPortNumber()
	ouo.mutation.SetPortNumber(i)
	return ouo
}

// AddPortNumber adds i to the "port_number" field.
func (ouo *OriginUpdateOne) AddPortNumber(i int) *OriginUpdateOne {
	ouo.mutation.AddPortNumber(i)
	return ouo
}

// SetActive sets the "active" field.
func (ouo *OriginUpdateOne) SetActive(b bool) *OriginUpdateOne {
	ouo.mutation.SetActive(b)
	return ouo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (ouo *OriginUpdateOne) SetNillableActive(b *bool) *OriginUpdateOne {
	if b != nil {
		ouo.SetActive(*b)
	}
	return ouo
}

// Mutation returns the OriginMutation object of the builder.
func (ouo *OriginUpdateOne) Mutation() *OriginMutation {
	return ouo.mutation
}

// Where appends a list predicates to the OriginUpdate builder.
func (ouo *OriginUpdateOne) Where(ps ...predicate.Origin) *OriginUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OriginUpdateOne) Select(field string, fields ...string) *OriginUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Origin entity.
func (ouo *OriginUpdateOne) Save(ctx context.Context) (*Origin, error) {
	ouo.defaults()
	return withHooks(ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OriginUpdateOne) SaveX(ctx context.Context) *Origin {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OriginUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OriginUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *OriginUpdateOne) defaults() {
	if _, ok := ouo.mutation.UpdatedAt(); !ok {
		v := origin.UpdateDefaultUpdatedAt()
		ouo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OriginUpdateOne) check() error {
	if v, ok := ouo.mutation.Name(); ok {
		if err := origin.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Origin.name": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.Target(); ok {
		if err := origin.TargetValidator(v); err != nil {
			return &ValidationError{Name: "target", err: fmt.Errorf(`generated: validator failed for field "Origin.target": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.PortNumber(); ok {
		if err := origin.PortNumberValidator(v); err != nil {
			return &ValidationError{Name: "port_number", err: fmt.Errorf(`generated: validator failed for field "Origin.port_number": %w`, err)}
		}
	}
	if _, ok := ouo.mutation.PoolID(); ouo.mutation.PoolCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Origin.pool"`)
	}
	return nil
}

func (ouo *OriginUpdateOne) sqlSave(ctx context.Context) (_node *Origin, err error) {
	if err := ouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(origin.Table, origin.Columns, sqlgraph.NewFieldSpec(origin.FieldID, field.TypeString))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Origin.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, origin.FieldID)
		for _, f := range fields {
			if !origin.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != origin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.UpdatedAt(); ok {
		_spec.SetField(origin.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ouo.mutation.Name(); ok {
		_spec.SetField(origin.FieldName, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Target(); ok {
		_spec.SetField(origin.FieldTarget, field.TypeString, value)
	}
	if value, ok := ouo.mutation.PortNumber(); ok {
		_spec.SetField(origin.FieldPortNumber, field.TypeInt, value)
	}
	if value, ok := ouo.mutation.AddedPortNumber(); ok {
		_spec.AddField(origin.FieldPortNumber, field.TypeInt, value)
	}
	if value, ok := ouo.mutation.Active(); ok {
		_spec.SetField(origin.FieldActive, field.TypeBool, value)
	}
	_node = &Origin{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{origin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
