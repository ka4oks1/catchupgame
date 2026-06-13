package maps

type MapGen struct {
	seed int
}

type Map struct {
	Size int
	Type string //square rect
}

func (mg *MapGen) GenerateMap() Map {

	return Map{}
}
