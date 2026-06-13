package maps

const (
	minSize = 3
	maxSize = 7
)

type Position struct {
	XCoord uint8
	YCoord uint8
}

func PositionCorrect(position Position) bool {

	return false
}

type MapGen struct {
	seed int
}

type Map struct {
	Size      int
	Territory map[Position]bool
	Type      string //square rect
}

func (mg *MapGen) GenerateMap() Map {

	return Map{}
}
