// Package schema ...
package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/idx"
)

// LoadBalancer holds the schema definition for the LoadBalancer entity.
type LoadBalancer struct {
	ent.Schema
}

// Mixin of the LoadBalancer.
func (LoadBalancer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		idx.PrimaryKeyMixin(idPrefixLoadBalancer),
		entx.ExternalEdge(
			entx.WithExternalEdgeType("Location"),
			entx.ExternalEdgeImmutable(true),
		),
		entx.ExternalEdge(
			entx.WithExternalEdgeType("Tenant"),
			entx.ExternalEdgeImmutable(true),
		),
		entx.TimestampsMixin{},
	}
}

// Fields of the LoadBalancer.
func (LoadBalancer) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.String("ip_address_id").
			GoType(idx.PrefixedID("")).
			Immutable(),
	}
}

// Edges of the LoadBalancer.
func (LoadBalancer) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.To("ports", Ports.Type),
		// edge.To("locations", Location.Type).Annotations(),
		// edge.To("pools", Pool.Type).Annotations(),
	}
}

// Annotations of the LoadBalancer.
func (LoadBalancer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
