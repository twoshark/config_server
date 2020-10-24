package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/twoshark/config_server/graph/generated"
	"github.com/twoshark/config_server/graph/model"
)

func (r *mutationResolver) CreateBasic(ctx context.Context, input model.NewFixture) (*model.Basic, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateBasic(ctx context.Context, input model.UpdateBasic) (*model.Basic, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateSpotlight(ctx context.Context, input model.NewFixture) (*model.Spotlight, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSpotlight(ctx context.Context, input model.UpdateSpotlight) (*model.Spotlight, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Fixture(ctx context.Context, id string) (model.Fixture, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
