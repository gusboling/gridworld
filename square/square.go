package square

import (
	"errors"
	"fmt"
)

const (
	northIndex = 0
	eastIndex  = 1
	southIndex = 2
	westIndex  = 3
)

type SquareInterface interface {
	MakeImpassable() error
	Occupy() error

	AttachNorth(foreignSquare *Square) error
	AttachEast(foreignSquare *Square) error
	AttachSouth(foreignSquare *Square) error
	AttachWest(foreignSquare *Square) error
}

// Square represents a single grid-square, connected to 4 neighboring squares
type Square struct {
	/*
		Absolute coordinates are not stored in Square because they're a property
		derived from the grid, not from the square itself. To prevent conflicting
		information, X/Y coordinates are generated/stored in the Grid struct and
		not duplicated here.
	*/

	// Pointers to neighboring squares, order: north,east,south,west
	neighbors []*Square

	// Properties
	isPassable bool
	isOccupied bool
}

// MakePassable returns an error if square is already impassable
func (s *Square) MakePassable() error {
	if s.isPassable {
		return errors.New("square is already passable")
	}
	s.isPassable = true
	return nil
}

// MakeOccupied returns an error if square is already impassable
func (s *Square) MakeOccupied() error {
	if s.isOccupied {
		return errors.New("square is already occupied")
	}
	s.isOccupied = true
	return nil
}

func (s *Square) AttachNorth(foreignSquare *Square) error {
	return s.Attach(foreignSquare, northIndex)
}

func (s *Square) AttachEast(foreignSquare *Square) error {
	return s.Attach(foreignSquare, eastIndex)
}

func (s *Square) AttachSouth(foreignSquare *Square) error {
	return s.Attach(foreignSquare, southIndex)
}

func (s *Square) AttachWest(foreignSquare *Square) error {
	return s.Attach(foreignSquare, westIndex)
}

func (s *Square) Attach(foreignSquare *Square, neighborIndex uint32) error {
	if foreignSquare == nil {
		return errors.New("cannot use nil pointer for attachment")
	}

	if neighborIndex < 0 || neighborIndex > 3 {
		return errors.New(fmt.Sprintf("cannot attach to neighbor at index %v", neighborIndex))
	}

	s.neighbors[neighborIndex] = foreignSquare
	return nil
}

func (s *Square) DetachNorth() error {
	return s.Detach(northIndex)
}

func (s *Square) DetachEast() error {
	return s.Detach(eastIndex)
}

func (s *Square) DetachSouth() error {
	return s.Detach(southIndex)
}

func (s *Square) DetachWest() error {
	return s.Detach(westIndex)
}

func (s *Square) Detach(neighborIndex uint32) error {

	if s.neighbors[neighborIndex] == nil {
		return errors.New(fmt.Sprintf("neighbor at index %v is already nil", neighborIndex))
	}

	if neighborIndex < 0 || neighborIndex > 3 {
		return errors.New(fmt.Sprintf("cannot detach from neighbor at index %v", neighborIndex))
	}

	s.neighbors[neighborIndex] = nil
	return nil
}
