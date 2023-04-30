package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"go.infratographer.com/load-balancer-api/internal/ent/generated"
	graphapigen "go.infratographer.com/load-balancer-api/internal/graphapi/generated"
	"go.infratographer.com/x/idx"
)

func (r *entityResolver) FindLoadBalancerByID(ctx context.Context, id idx.PrefixedID) (*generated.LoadBalancer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *entityResolver) FindLocationByID(ctx context.Context, id idx.PrefixedID) (*graphapigen.Location, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *entityResolver) FindTenantByID(ctx context.Context, id idx.PrefixedID) (*graphapigen.Tenant, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns graphapigen.EntityResolver implementation.
func (r *Resolver) Entity() graphapigen.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
