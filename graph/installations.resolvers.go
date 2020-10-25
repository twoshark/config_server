package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/twoshark/config_server/graph/model"
	"github.com/twoshark/config_server/store"
	"github.com/twoshark/config_server/uuid"
)

func (r *mutationResolver) CreateInstallation(ctx context.Context, input model.CreateInstallation) (*model.Installation, error) {
	s := store.GetStore()

	found := s.Installations.FindByID(input.ID) > 0
	if found {
		return nil, errors.New("This ID Already Exists")
	}

	newInstallation := model.Installation{
		ID:   uuid.Generate(uuid.Installation),
		Name: input.Name,
	}
	updateString(input.Description, newInstallation.Description)

	s.Installations = append(s.Installations, newInstallation)
	return &newInstallation, nil
}

func (r *mutationResolver) UpdateInstallation(ctx context.Context, input model.UpdateInstallation) (*model.Installation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Installation(ctx context.Context, id string) (*model.Installation, error) {
	panic(fmt.Errorf("not implemented"))
}
