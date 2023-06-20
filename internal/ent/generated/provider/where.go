// Copyright Infratographer, Inc. and/or licensed to Infratographer, Inc. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.
//
// Code generated by entc, DO NOT EDIT.

package provider

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

// ID filters vertices based on their ID field.
func ID(id gidx.PrefixedID) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id gidx.PrefixedID) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id gidx.PrefixedID) predicate.Provider {
	return predicate.Provider(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...gidx.PrefixedID) predicate.Provider {
	return predicate.Provider(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...gidx.PrefixedID) predicate.Provider {
	return predicate.Provider(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id gidx.PrefixedID) predicate.Provider {
	return predicate.Provider(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id gidx.PrefixedID) predicate.Provider {
	return predicate.Provider(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id gidx.PrefixedID) predicate.Provider {
	return predicate.Provider(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id gidx.PrefixedID) predicate.Provider {
	return predicate.Provider(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldName, v))
}

// TenantID applies equality check predicate on the "tenant_id" field. It's identical to TenantIDEQ.
func TenantID(v gidx.PrefixedID) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldTenantID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Provider {
	return predicate.Provider(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Provider {
	return predicate.Provider(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Provider {
	return predicate.Provider(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Provider {
	return predicate.Provider(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Provider {
	return predicate.Provider(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Provider {
	return predicate.Provider(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Provider {
	return predicate.Provider(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Provider {
	return predicate.Provider(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Provider {
	return predicate.Provider(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Provider {
	return predicate.Provider(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Provider {
	return predicate.Provider(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Provider {
	return predicate.Provider(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Provider {
	return predicate.Provider(sql.FieldContainsFold(FieldName, v))
}

// TenantIDEQ applies the EQ predicate on the "tenant_id" field.
func TenantIDEQ(v gidx.PrefixedID) predicate.Provider {
	return predicate.Provider(sql.FieldEQ(FieldTenantID, v))
}

// TenantIDNEQ applies the NEQ predicate on the "tenant_id" field.
func TenantIDNEQ(v gidx.PrefixedID) predicate.Provider {
	return predicate.Provider(sql.FieldNEQ(FieldTenantID, v))
}

// TenantIDIn applies the In predicate on the "tenant_id" field.
func TenantIDIn(vs ...gidx.PrefixedID) predicate.Provider {
	return predicate.Provider(sql.FieldIn(FieldTenantID, vs...))
}

// TenantIDNotIn applies the NotIn predicate on the "tenant_id" field.
func TenantIDNotIn(vs ...gidx.PrefixedID) predicate.Provider {
	return predicate.Provider(sql.FieldNotIn(FieldTenantID, vs...))
}

// TenantIDGT applies the GT predicate on the "tenant_id" field.
func TenantIDGT(v gidx.PrefixedID) predicate.Provider {
	return predicate.Provider(sql.FieldGT(FieldTenantID, v))
}

// TenantIDGTE applies the GTE predicate on the "tenant_id" field.
func TenantIDGTE(v gidx.PrefixedID) predicate.Provider {
	return predicate.Provider(sql.FieldGTE(FieldTenantID, v))
}

// TenantIDLT applies the LT predicate on the "tenant_id" field.
func TenantIDLT(v gidx.PrefixedID) predicate.Provider {
	return predicate.Provider(sql.FieldLT(FieldTenantID, v))
}

// TenantIDLTE applies the LTE predicate on the "tenant_id" field.
func TenantIDLTE(v gidx.PrefixedID) predicate.Provider {
	return predicate.Provider(sql.FieldLTE(FieldTenantID, v))
}

// TenantIDContains applies the Contains predicate on the "tenant_id" field.
func TenantIDContains(v gidx.PrefixedID) predicate.Provider {
	vc := string(v)
	return predicate.Provider(sql.FieldContains(FieldTenantID, vc))
}

// TenantIDHasPrefix applies the HasPrefix predicate on the "tenant_id" field.
func TenantIDHasPrefix(v gidx.PrefixedID) predicate.Provider {
	vc := string(v)
	return predicate.Provider(sql.FieldHasPrefix(FieldTenantID, vc))
}

// TenantIDHasSuffix applies the HasSuffix predicate on the "tenant_id" field.
func TenantIDHasSuffix(v gidx.PrefixedID) predicate.Provider {
	vc := string(v)
	return predicate.Provider(sql.FieldHasSuffix(FieldTenantID, vc))
}

// TenantIDEqualFold applies the EqualFold predicate on the "tenant_id" field.
func TenantIDEqualFold(v gidx.PrefixedID) predicate.Provider {
	vc := string(v)
	return predicate.Provider(sql.FieldEqualFold(FieldTenantID, vc))
}

// TenantIDContainsFold applies the ContainsFold predicate on the "tenant_id" field.
func TenantIDContainsFold(v gidx.PrefixedID) predicate.Provider {
	vc := string(v)
	return predicate.Provider(sql.FieldContainsFold(FieldTenantID, vc))
}

// HasLoadBalancers applies the HasEdge predicate on the "load_balancers" edge.
func HasLoadBalancers() predicate.Provider {
	return predicate.Provider(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, LoadBalancersTable, LoadBalancersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLoadBalancersWith applies the HasEdge predicate on the "load_balancers" edge with a given conditions (other predicates).
func HasLoadBalancersWith(preds ...predicate.LoadBalancer) predicate.Provider {
	return predicate.Provider(func(s *sql.Selector) {
		step := newLoadBalancersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Provider) predicate.Provider {
	return predicate.Provider(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Provider) predicate.Provider {
	return predicate.Provider(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Provider) predicate.Provider {
	return predicate.Provider(func(s *sql.Selector) {
		p(s.Not())
	})
}
