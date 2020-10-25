package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/twoshark/config_server/fixture"
	"github.com/twoshark/config_server/graph/generated"
	"github.com/twoshark/config_server/graph/model"
	"github.com/twoshark/config_server/store"
	"github.com/twoshark/config_server/uuid"
)

func (r *mutationResolver) CreateFixture(ctx context.Context, input model.NewFixture) (model.Fixture, error) {
	s := store.GetStore()

	var t fixture.Type
	err := json.Unmarshal([]byte(*input.Type), &t)
	if err != nil {
		return nil, err
	}

	switch t {
	case fixture.Basic:
		newBasic := model.Basic{
			ID:            uuid.Generate(uuid.Fixture),
			Name:          input.Name,
			Description:   input.Description,
			Installations: input.Installations,
			Opacity:       0,
			Animation:     0,
			Option:        0,
			Speed:         0,
			Strobe:        0,
		}
		s.Fixtures.Basics = append(s.Fixtures.Basics, newBasic)

		return &newBasic, nil

	case fixture.Spotlight:
		opacity := 0
		newSpot := model.Spotlight{
			ID:            uuid.Generate(uuid.Fixture),
			Name:          input.Name,
			Description:   input.Description,
			Installations: input.Installations,
			Opacity:       &opacity,
		}
		s.Fixtures.Spotlights = append(s.Fixtures.Spotlights, newSpot)

		return &newSpot, nil

	default:
		return nil, errors.New("Invalid Type")
	}
}

func (r *mutationResolver) UpdateBasic(ctx context.Context, input model.UpdateBasic) (*model.Basic, error) {
	s := store.GetStore()

	fx := s.Fixtures.FindByID(input.ID)
	if fx.Type == fixture.Invalid {
		return nil, errors.New("This Fixture Cannot Be Found")
	}
	updateString(input.Name, s.Fixtures.Basics[fx.Index].Name)
	updateString(input.Description, s.Fixtures.Basics[fx.Index].Description)
	updateStringArray(&input.Installations, &s.Fixtures.Basics[fx.Index].Installations)
	updateInt(input.Opacity, &s.Fixtures.Basics[fx.Index].Opacity)
	updateInt(input.Animation, &s.Fixtures.Basics[fx.Index].Animation)
	updateInt(input.Option, &s.Fixtures.Basics[fx.Index].Option)
	updateInt(input.Strobe, &s.Fixtures.Basics[fx.Index].Strobe)
	updateInt(input.Speed, &s.Fixtures.Basics[fx.Index].Speed)

	return &s.Fixtures.Basics[fx.Index], nil
}

func (r *mutationResolver) UpdateSpotlight(ctx context.Context, input model.UpdateSpotlight) (*model.Spotlight, error) {
	s := store.GetStore()

	fx := s.Fixtures.FindByID(input.ID)
	if fx.Type == fixture.Invalid {
		return nil, errors.New("This Fixture Cannot Be Found")
	}
	updateString(input.Name, s.Fixtures.Basics[fx.Index].Name)
	updateString(input.Description, s.Fixtures.Basics[fx.Index].Description)
	updateStringArray(&input.Installations, &s.Fixtures.Basics[fx.Index].Installations)
	updateInt(input.Opacity, &s.Fixtures.Basics[fx.Index].Opacity)

	return &s.Fixtures.Spotlights[fx.Index], nil
}

func (r *queryResolver) Fixture(ctx context.Context, id string) (model.Fixture, error) {
	s := store.GetStore()

	fx := s.Fixtures.FindByID(id)
	if fx.Type == fixture.Invalid {
		return nil, errors.New("This Fixture Cannot Be Found")
	}
	return s.Fixtures.GetFixture(fx), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
