package PointontheGraphics

import "fmt"

type CheckPointInShape interface {
	CheckPoint(x, y int) bool
}

var _ = CheckPointInShape(Round{})

func (round Round) CheckPoint(x, y int) bool {
	length := (x-round.X)*(x-round.X) + (y-round.Y)*(y-round.Y)
	return length <= round.Radius*round.Radius
}

type Round struct {
	X      int
	Y      int
	Radius int
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

func (coordinate *Coordinates) NewRound(X, Y, Radius, ID int) *Round {
	if X > coordinate.MAXX || X < 0 || Y > coordinate.MAXY || Y < 0 {
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

	for i := left / 1; i <= right/1; i++ {
		coordinate.XRounds[i] = append(coordinate.XRounds[i], round)
	}
	fmt.Printf("left/1 is %d, right/1 is %d\n", left/1, right/1+1)

	for j := down / 1; j <= up/1; j++ {
		coordinate.YRounds[j] = append(coordinate.YRounds[j], round)
	}
	fmt.Printf("down/1 is %d, up/1 is %d\n", down/1, up/1+1)

	coordinate.Rounds[round.ID] = &round
	return true
}

func (coordinate *Coordinates) FindRoundsonPoint(X, Y int) map[int]struct{} {
	XRounds := coordinate.XRounds[X-1]
	YRounds := coordinate.YRounds[Y-1]
	RoundsID := map[int]struct{}{}

	for _, v := range XRounds {
		RoundsID[v.ID] = struct{}{}
	}
	for _, v := range YRounds {
		RoundsID[v.ID] = struct{}{}
	}

	for roundid := range RoundsID {
		fmt.Printf("roundid is %v\n", roundid)
		if !coordinate.Rounds[roundid].CheckPoint(X, Y) {
			delete(RoundsID, roundid)
		}
	}
	return RoundsID
}
