package maps

const (
	minSize = 3
	maxSize = 7
)

type Position struct {
	XCoord uint8
	YCoord uint8
}

type MapGen struct {
	seed int
}

type Map struct {
	Size      int
	Territory map[Position]bool
	Type      string //square rect
}

// false here is good like 0 is empty place where character may stand and go to
func (m *Map) PositionCorrect(position Position) bool {
	if m.Territory[position] == false {
		return true
	}
	return false
}
func (mg *MapGen) GenerateMap() Map {

	return Map{}
}
