package direction

import "fmt"

type Direction int32 //int32 because more convertable

const (
	// Enumerated type
	North   = Direction(0)
	East    = Direction(1)
	South   = Direction(2)
	West    = Direction(3)
	Unknown = Direction(4)
)

// Given a direction on a compass-rose, returns opposite direction
func OppositeDirection(d Direction) (Direction, error) {
	switch d {
	case North:
		return South, nil
	case East:
		return West, nil
	case South:
		return North, nil
	case West:
		return East, nil
	}

	return Unknown, fmt.Errorf("invalid direction given - opposite unknown")
}
