package store

import (
	"sync"

	"github.com/twoshark/config_server/fixture"
	"github.com/twoshark/config_server/graph/model"
)

var (
	//instance is the Store Singleton
	instance *Store
	once     sync.Once
)

//Store ...
type Store struct {
	Fixtures      FixtureStore
	Installations InstallationStore
}

//GetStore ....
func GetStore() *Store {
	once.Do(func() { // <-- atomic, does not allow repeating
		instance = new(Store)
		instance.init()
	})
	return instance
}

//Init initializes the store arrays
func (s *Store) init() {
	s.Fixtures.Basics = make([]model.Basic, 0, 24)
	s.Fixtures.Spotlights = make([]model.Spotlight, 0, 24)
}

//FixtureStore ...
type FixtureStore struct {
	Basics     []model.Basic
	Spotlights []model.Spotlight
}

//FindByID ...
func (f FixtureStore) FindByID(id string) fixture.Index {
	for i, basic := range f.Basics {
		if basic.ID == id {
			return fixture.NewFixtureIndex(i, fixture.Basic)
		}
	}
	for i, spotlight := range f.Spotlights {
		if spotlight.ID == id {
			return fixture.NewFixtureIndex(i, fixture.Spotlight)
		}
	}
	return fixture.NewFixtureIndex(-1, fixture.Invalid)
}

//GetFixture ...
func (f FixtureStore) GetFixture(fixIndex fixture.Index) model.Fixture {
	switch fixIndex.Type {
	case fixture.Spotlight:
		return f.Spotlights[fixIndex.Index]
	case fixture.Basic:
		return f.Basics[fixIndex.Index]
	default:
		return nil
	}
}

//InstallationStore ...
type InstallationStore []model.Installation

//FindByID ...
func (is InstallationStore) FindByID(id string) int {
	for i, install := range is {
		if install.ID == id {
			return i
		}
	}
	return -1
}
