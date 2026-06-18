package personell

import (
	"catchupgame/maps"
	_ "github.com/stretchr/testify"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"testing"
)

type TestStruct struct {
	pers Person
	mapa maps.Map
	way  string
}

func TestPerson_Move(t *testing.T) { //write correct test

	//var test_template map[TestStruct]bool
	//
	//s1:= TestStruct{
	//	{{{XCoord: 1, YCoord: 2}}, {}, ""
	//}
	//test_template[]
	//
	assert.Equal(t, true, true)
}
