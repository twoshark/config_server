package fixture

//Type enumerates the supported fixtures
type Type string

const (
	//Basic Fixture
	Basic Type = "Basic" 
	//Spotlight Fixture
	Spotlight = "Spotlight"
	//Invalid Fixture
	Invalid = "Invalid"
)

//Index ...
type Index struct {
	Index int
	Type  Type
}

//NewFixtureIndex ...
func NewFixtureIndex(i int, t Type) Index {
	return Index{
		Index: i,
		Type:  t,
	}
}
