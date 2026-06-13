package personell

import (
	"catchupgame/controls"
	"catchupgame/maps"
	"errors"
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

func (p *Person) Move(way string) error {

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

	if maps.PositionCorrect(expectedPos) { //destination availible on map and way exists

		return nil
	}
	wrongTurn := errors.New("invalid turn")
	return wrongTurn

}
