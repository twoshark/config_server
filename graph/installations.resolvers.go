package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/twoshark/config_server/graph/model"
)

func (r *mutationResolver) CreateInstallation(ctx context.Context, input model.CreateInstallation) (*model.Installation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateInstallation(ctx context.Context, input model.UpdateInstallation) (*model.Installation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Installation(ctx context.Context, id string) (*model.Installation, error) {
	panic(fmt.Errorf("not implemented"))
}
