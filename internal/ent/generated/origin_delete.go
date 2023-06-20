// Copyright Infratographer, Inc. and/or licensed to Infratographer, Inc. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/origin"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/predicate"
)

// OriginDelete is the builder for deleting a Origin entity.
type OriginDelete struct {
	config
	hooks    []Hook
	mutation *OriginMutation
}

// Where appends a list predicates to the OriginDelete builder.
func (od *OriginDelete) Where(ps ...predicate.Origin) *OriginDelete {
	od.mutation.Where(ps...)
	return od
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (od *OriginDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, od.sqlExec, od.mutation, od.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (od *OriginDelete) ExecX(ctx context.Context) int {
	n, err := od.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (od *OriginDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(origin.Table, sqlgraph.NewFieldSpec(origin.FieldID, field.TypeString))
	if ps := od.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, od.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	od.mutation.done = true
	return affected, err
}

// OriginDeleteOne is the builder for deleting a single Origin entity.
type OriginDeleteOne struct {
	od *OriginDelete
}

// Where appends a list predicates to the OriginDelete builder.
func (odo *OriginDeleteOne) Where(ps ...predicate.Origin) *OriginDeleteOne {
	odo.od.mutation.Where(ps...)
	return odo
}

// Exec executes the deletion query.
func (odo *OriginDeleteOne) Exec(ctx context.Context) error {
	n, err := odo.od.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{origin.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (odo *OriginDeleteOne) ExecX(ctx context.Context) {
	if err := odo.Exec(ctx); err != nil {
		panic(err)
	}
}
