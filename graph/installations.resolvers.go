package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/twoshark/config_server/graph/model"
	"github.com/twoshark/config_server/store"
	"github.com/twoshark/config_server/uuid"
)

func (r *mutationResolver) CreateInstallation(ctx context.Context, input model.CreateInstallation) (*model.Installation, error) {
	s := store.GetStore()
	newInstallation := model.Installation{
		ID:   uuid.Generate(uuid.Installation),
		Name: input.Name,
	}
	newInstallation.Description = updateString(input.Description, newInstallation.Description)

	s.Installations = append(s.Installations, newInstallation)
	return &newInstallation, nil
}

func (r *mutationResolver) UpdateInstallation(ctx context.Context, input model.UpdateInstallation) (*model.Installation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Installations(ctx context.Context) ([]*model.Installation, error) {
	s := store.GetStore()

	ptrs := make([]*model.Installation, len(s.Installations))
	for i := range s.Installations {
		ptrs[i] = &s.Installations[i]
	}
	return ptrs, nil
}

func (r *queryResolver) Installation(ctx context.Context, id string) (*model.Installation, error) {
	panic(fmt.Errorf("not implemented"))
}
