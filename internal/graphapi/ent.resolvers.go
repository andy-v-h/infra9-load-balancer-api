package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"go.infratographer.com/load-balancer-api/internal/ent/generated"
	graphapigen "go.infratographer.com/load-balancer-api/internal/graphapi/generated"
	"go.infratographer.com/x/idx"
)

func (r *loadBalancerResolver) IPAddressID(ctx context.Context, obj *generated.LoadBalancer) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *loadBalancerResolver) Location(ctx context.Context, obj *generated.LoadBalancer) (*graphapigen.Location, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *loadBalancerResolver) Tenant(ctx context.Context, obj *generated.LoadBalancer) (*graphapigen.Tenant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) LoadBalancers(ctx context.Context, after *entgql.Cursor[idx.PrefixedID], first *int, before *entgql.Cursor[idx.PrefixedID], last *int, orderBy *generated.LoadBalancerOrder, where *generated.LoadBalancerWhereInput) (*generated.LoadBalancerConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *createLoadBalancerInputResolver) IPAddressID(ctx context.Context, obj *generated.CreateLoadBalancerInput, data string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *loadBalancerWhereInputResolver) IPAddressID(ctx context.Context, obj *generated.LoadBalancerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *loadBalancerWhereInputResolver) IPAddressIDNeq(ctx context.Context, obj *generated.LoadBalancerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *loadBalancerWhereInputResolver) IPAddressIDIn(ctx context.Context, obj *generated.LoadBalancerWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *loadBalancerWhereInputResolver) IPAddressIDNotIn(ctx context.Context, obj *generated.LoadBalancerWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *loadBalancerWhereInputResolver) IPAddressIDGt(ctx context.Context, obj *generated.LoadBalancerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *loadBalancerWhereInputResolver) IPAddressIDGte(ctx context.Context, obj *generated.LoadBalancerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *loadBalancerWhereInputResolver) IPAddressIDLt(ctx context.Context, obj *generated.LoadBalancerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *loadBalancerWhereInputResolver) IPAddressIDLte(ctx context.Context, obj *generated.LoadBalancerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *loadBalancerWhereInputResolver) IPAddressIDContains(ctx context.Context, obj *generated.LoadBalancerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *loadBalancerWhereInputResolver) IPAddressIDHasPrefix(ctx context.Context, obj *generated.LoadBalancerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *loadBalancerWhereInputResolver) IPAddressIDHasSuffix(ctx context.Context, obj *generated.LoadBalancerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *loadBalancerWhereInputResolver) IPAddressIDEqualFold(ctx context.Context, obj *generated.LoadBalancerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *loadBalancerWhereInputResolver) IPAddressIDContainsFold(ctx context.Context, obj *generated.LoadBalancerWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented"))
}

// LoadBalancer returns graphapigen.LoadBalancerResolver implementation.
func (r *Resolver) LoadBalancer() graphapigen.LoadBalancerResolver { return &loadBalancerResolver{r} }

// Query returns graphapigen.QueryResolver implementation.
func (r *Resolver) Query() graphapigen.QueryResolver { return &queryResolver{r} }

// CreateLoadBalancerInput returns graphapigen.CreateLoadBalancerInputResolver implementation.
func (r *Resolver) CreateLoadBalancerInput() graphapigen.CreateLoadBalancerInputResolver {
	return &createLoadBalancerInputResolver{r}
}

// LoadBalancerWhereInput returns graphapigen.LoadBalancerWhereInputResolver implementation.
func (r *Resolver) LoadBalancerWhereInput() graphapigen.LoadBalancerWhereInputResolver {
	return &loadBalancerWhereInputResolver{r}
}

type loadBalancerResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type createLoadBalancerInputResolver struct{ *Resolver }
type loadBalancerWhereInputResolver struct{ *Resolver }
