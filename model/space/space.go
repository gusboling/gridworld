package space

import (
	"errors"
	"fmt"

	"github.com/grid_simulator/model/direction"
)

const (
	max_neighbors = 4
)

type SpaceInterface interface {
	MakeImpassable() error
	Occupy() error

	AttachNorth(foreignSquare *Space) error
	AttachEast(foreignSquare *Space) error
	AttachSouth(foreignSquare *Space) error
	AttachWest(foreignSquare *Space) error
}

// Space represents a single grid-space, connected to 4 neighboring spaces
type Space struct {
	/*
		Absolute coordinates are not stored in Square because they're a property
		derived from the grid, not from the space itself. To prevent conflicting
		information, X/Y coordinates are generated/stored in the Grid struct and
		not duplicated here.
	*/

	// Pointers to neighboring spaces, order: north,east,south,west
	neighbors [max_neighbors]*Space

	// Properties
	isPassable bool
	isOccupied bool
}

func NewSpace() *Space {
	newObj := new(Space)
	return newObj
}

// MakePassable returns an error if space is already passable
func (s *Space) MakePassable() error {
	if s.isPassable {
		return errors.New("space is already passable")
	}
	s.isPassable = true
	return nil
}

// MakePassable returns an error if space is already impassable
func (s *Space) MakeImpassable() error {
	if !s.isPassable {
		return errors.New("space is already impassable")
	}
	s.isPassable = false
	return nil
}

// MakeOccupied returns an error if space is already occupied
func (s *Space) MakeOccupied() error {
	if s.isOccupied {
		return errors.New("space is already occupied")
	}
	s.isOccupied = true
	return nil
}

// MakeOccupied returns an error if space is already unoccupied
func (s *Space) MakeUnoccupied() error {
	if !s.isOccupied {
		return errors.New("space is already unoccupied")
	}
	s.isOccupied = false
	return nil
}

// Attach adds foreignSquare as a neighbor to s, and s as a neighbor to foreignSquare
func (s *Space) Attach(foreignSquare *Space, neighboringDirection direction.Direction) error {
	if foreignSquare == nil {
		return errors.New("cannot use nil pointer for attachment")
	}

	if neighboringDirection > 3 {
		return fmt.Errorf("cannot attach to neighbor at index %v", neighboringDirection)
	}

	s.neighbors[neighboringDirection] = foreignSquare

	direction, errOpposite := direction.OppositeDirection(neighboringDirection)
	if errOpposite == nil {
		foreignSquare.neighbors[direction] = s
		return nil
	}

	return errOpposite
}

// Detach removes s as a neighbor from a given neighboring space, and removes
// the corresponding neighboring space a as a neighbor from s
func (s *Space) Detach(neighboringDirection direction.Direction) error {

	if s.neighbors[neighboringDirection] == nil {
		return fmt.Errorf("neighbor at index %v is already nil", neighboringDirection)
	}

	invalidDirection := neighboringDirection > 3
	if invalidDirection {
		return fmt.Errorf("cannot detach from neighbor at index %v", neighboringDirection)
	}

	// ordering here is important, first detach the neighbor's reference to this
	// space, then detach this space's reference to neighbor2

	direction, errOpposite := direction.OppositeDirection(neighboringDirection)
	if errOpposite == nil {
		s.neighbors[neighboringDirection].neighbors[direction] = nil
		s.neighbors[neighboringDirection] = nil
		return nil
	}

	return errOpposite
}

// Convenience Methods

func (s *Space) AttachNorth(foreignSquare *Space) error {
	return s.Attach(foreignSquare, direction.North)
}

func (s *Space) AttachEast(foreignSquare *Space) error {
	return s.Attach(foreignSquare, direction.East)
}

func (s *Space) AttachSouth(foreignSquare *Space) error {
	return s.Attach(foreignSquare, direction.South)
}

func (s *Space) AttachWest(foreignSquare *Space) error {
	return s.Attach(foreignSquare, direction.West)
}

func (s *Space) DetachNorth() error {
	return s.Detach(direction.North)
}

func (s *Space) DetachEast() error {
	return s.Detach(direction.East)
}

func (s *Space) DetachSouth() error {
	return s.Detach(direction.South)
}

func (s *Space) DetachWest() error {
	return s.Detach(direction.West)
}
