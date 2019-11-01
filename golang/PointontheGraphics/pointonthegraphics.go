package PointontheGraphics

import (
	"math"
)

type CheckPointInShape interface {
	CheckPoint(x, y float64) bool
}

var _ = CheckPointInShape(Round{})

func (round Round) CheckPoint(x, y float64) bool {
	length := (x-round.X)*(x-round.X) + (y-round.Y)*(y-round.Y)
	return length <= round.Radius*round.Radius
}

type Round struct {
	X      float64
	Y      float64
	Radius float64
	ID     int
}

type Coordinates struct {
	MAXX    int
	MAXY    int
	XRounds map[int][]Round
	YRounds map[int][]Round
	Rounds  map[int]*Round
}

func NewCoordinates(X, Y int) *Coordinates {
	return &Coordinates{
		MAXX:    X,
		MAXY:    Y,
		XRounds: map[int][]Round{},
		YRounds: map[int][]Round{},
		Rounds:  map[int]*Round{},
	}
}

func (coordinate *Coordinates) NewRound(X float64, Y float64, Radius float64, ID int) *Round {
	if X > float64(coordinate.MAXX) || X < 0 || Y > float64(coordinate.MAXY) || Y < 0 {
		return nil
	}
	return &Round{
		X:      X,
		Y:      Y,
		Radius: Radius,
		ID:     ID,
	}
}

func (coordinate *Coordinates) AddRound(round Round) bool {

	left := round.X - round.Radius
	right := round.X + round.Radius

	down := round.Y - round.Radius
	up := round.Y + round.Radius

	for i := math.Floor(left); i <= math.Floor(right); i++ {
		coordinate.XRounds[int(i)] = append(coordinate.XRounds[int(i)], round)
	}

	for j := math.Floor(down); j <= math.Floor(up); j++ {
		coordinate.YRounds[int(j)] = append(coordinate.YRounds[int(j)], round)
	}

	coordinate.Rounds[round.ID] = &round
	return true
}

func (coordinate *Coordinates) FindRoundsonPoint(X, Y float64) map[int]struct{} {
	XRounds := coordinate.XRounds[int(math.Floor(X))]
	YRounds := coordinate.YRounds[int(math.Floor(Y))]
	RoundsID := map[int]struct{}{}

	for _, v := range XRounds {
		RoundsID[v.ID] = struct{}{}
	}
	for _, v := range YRounds {
		RoundsID[v.ID] = struct{}{}
	}

	for roundid := range RoundsID {
		if !coordinate.Rounds[roundid].CheckPoint(X, Y) {
			delete(RoundsID, roundid)
		}
	}
	return RoundsID
}
