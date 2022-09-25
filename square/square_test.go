package square

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MakePassable(t *testing.T) {

	testSquare := new(Square)

	errMakePassable := testSquare.MakePassable()
	assert.Nil(t, errMakePassable)

	errMakePassable = testSquare.MakePassable()
	assert.NotNil(t, errMakePassable)
}

func Test_MakeOccupied(t *testing.T) {

	testSquare := new(Square)

	errMakeOccupied := testSquare.MakeOccupied()
	assert.Nil(t, errMakeOccupied)

	errMakeOccupied = testSquare.MakeOccupied()
	assert.NotNil(t, errMakeOccupied)
}

func Test_Attach(t *testing.T) {
	// TODO(gus): implement
}

func Test_Detach(t *testing.T) {
	// TODO(gus): implement
}
