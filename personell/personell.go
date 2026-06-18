package personell

import (
	"catchupgame/controls"
	"catchupgame/maps"
)

func init() {

}

type Character interface {
	Move()
}
type Person struct {
	pos maps.Position
}

type Enemy struct {
	Person
}

func (e *Enemy) FindShortWay() {

}

func (p *Person) Move(way string, mapa maps.Map) bool {

	wayCode, wayExists := controls.GetWays(way)
	var expectedPos maps.Position

	switch wayCode {
	case 1:
		expectedPos.XCoord = p.pos.XCoord - 1
	case 2:
		expectedPos.YCoord = p.pos.YCoord + 1
	case 3:
		expectedPos.YCoord = p.pos.YCoord - 1
	case 4:
		expectedPos.XCoord = p.pos.XCoord + 1

	}

	if mapa.PositionCorrect(expectedPos) { //destination available on map and way exists

		return wayExists
	}

	wayExists = false
	return wayExists

}
