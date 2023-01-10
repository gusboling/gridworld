package direction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DirectionValues(t *testing.T) {
	assert.EqualValues(t, int32(0), North)
	assert.EqualValues(t, int32(1), East)
	assert.EqualValues(t, int32(2), South)
	assert.EqualValues(t, int32(3), West)
}

func Test_GetOppositeDirection(t *testing.T) {

	oppositeNorthValue, errGetOppositeNorth := OppositeDirection(North)
	assert.Nil(t, errGetOppositeNorth)
	assert.Equal(t, South, oppositeNorthValue)

	oppositeEastValue, errGetOppositeEast := OppositeDirection(East)
	assert.Nil(t, errGetOppositeEast)
	assert.Equal(t, West, oppositeEastValue)

	oppositeSouthValue, errGetOppositeSouth := OppositeDirection(South)
	assert.Nil(t, errGetOppositeSouth)
	assert.Equal(t, North, oppositeSouthValue)

	oppositeSouthWest, errGetOppositeWest := OppositeDirection(West)
	assert.Nil(t, errGetOppositeWest)
	assert.Equal(t, East, oppositeSouthWest)
}
