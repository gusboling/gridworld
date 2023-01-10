package space

import (
	"testing"

	"github.com/grid_simulator/model/direction"
	"github.com/stretchr/testify/assert"
)

func Test_NewSquare(t *testing.T) {
	newStruct := NewSpace()

	assert.Len(t, newStruct.neighbors, 4)
	assert.Empty(t, newStruct.neighbors)
	assert.False(t, newStruct.isOccupied)
	assert.False(t, newStruct.isPassable)
}

func Test_MakePassable(t *testing.T) {

	testSquare := new(Space)

	errMakePassable := testSquare.MakePassable()
	assert.Nil(t, errMakePassable)

	errMakePassable = testSquare.MakePassable()
	assert.NotNil(t, errMakePassable)
}

func Test_MakeOccupied(t *testing.T) {

	testSquare := new(Space)

	errMakeOccupied := testSquare.MakeOccupied()
	assert.Nil(t, errMakeOccupied)

	errMakeOccupied = testSquare.MakeOccupied()
	assert.NotNil(t, errMakeOccupied)
}

func Test_Attach(t *testing.T) {
	// TODO(gus): implement

	alpha := NewSpace()
	assert.NotNil(t, alpha)
	beta := NewSpace()
	assert.NotNil(t, beta)

	errAttachAlpha := alpha.AttachEast(beta)
	assert.Nil(t, errAttachAlpha)
	assert.Equal(t, beta, alpha.neighbors[direction.East])
	assert.Equal(t, alpha, beta.neighbors[direction.West])
}

func Test_Detach(t *testing.T) {

	alpha := NewSpace()
	assert.NotNil(t, alpha)
	beta := NewSpace()
	assert.NotNil(t, beta)

	errAttachAlpha := alpha.AttachEast(beta)
	assert.Nil(t, errAttachAlpha)
	assert.Equal(t, beta, alpha.neighbors[direction.East])
	assert.Equal(t, alpha, beta.neighbors[direction.West])

	errDetachAlpha := alpha.Detach(direction.East)
	assert.Nil(t, errDetachAlpha)
	assert.Nil(t, alpha.neighbors[direction.East])
	assert.Nil(t, beta.neighbors[direction.West])
}
