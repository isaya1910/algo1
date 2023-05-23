package main

import (
	"fmt"
	"math"
)

type Robot interface {
	Move(move Move)
	Turn(turn Turn)
	SetState(state CleaningState)
	Start()
	Stop()
}

type SweeperRobot struct {
	currentLocation Point
	state           CleaningState
	angle           Angle
}

func (s *SweeperRobot) Move(move Move) {
	angleRads := float64(s.angle) * (math.Pi / 180.0)
	s.currentLocation = Point{
		x: s.currentLocation.x + int32(float64(move.Movement.distance)*math.Cos(angleRads)),
		y: s.currentLocation.y + int32(float64(move.Movement.distance)*math.Sin(angleRads)),
	}
	fmt.Println(s.currentLocation)
}

func (s *SweeperRobot) SetState(state CleaningState) {
	s.state = state
	fmt.Println(s.state.Clean())
}

func (s *SweeperRobot) Start() {
	fmt.Sprintf("Started in location: ", s.currentLocation)
	s.state.Clean()
}

func (s *SweeperRobot) Turn(turn Turn) {
	s.angle += turn.Angle
	fmt.Println(s.angle)
}

func (s *SweeperRobot) Stop() {
	fmt.Println("Work stopped")
}
