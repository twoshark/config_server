package uuid

import (
	"crypto/rand"
	"fmt"
	"log"

	"github.com/twoshark/config_server/fixture"
	"github.com/twoshark/config_server/store"
)

//IDType ...
type IDType string

const (
	//Installation ...
	Installation IDType = "Installation"
	//Fixture ...
	Fixture = "Fixture"
)

//Generate a Unique ID for the Store
func Generate(idType IDType) string {
	//Create Array
	b := make([]byte, 16)
	//Fill with Random Numbers
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	//Format UUID
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	if Exists(uuid, idType) {
		uuid = Generate(idType)
	}

	return uuid
}

//Exists retruns whether or not this uuid is already in use
func Exists(id string, idType IDType) bool {
	s := store.GetStore()
	var found bool
	switch idType {
	case Installation:
		found = s.Installations.FindByID(id) > 0
	case Fixture:
		fixtureIndex := s.Fixtures.FindByID(id)
		found = fixtureIndex.Type == fixture.Invalid
	}

	return found
}
