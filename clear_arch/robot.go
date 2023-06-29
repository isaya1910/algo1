package main

import (
	"math"
)

// SweeperRobotState Монада состояний так как я это понял
type SweeperRobotState struct {
	location Point
	angle    Angle
	state    CleaningState
}

// Move теперь возвращает новое состояние робота и результат вычислений
func (s *SweeperRobotState) Move(move Move) (SweeperRobotState, Point) {
	angleRads := float64(s.angle) * (math.Pi / 180.0)
	newLocation := Point{
		x: s.location.x + int32(float64(move.Movement.distance)*math.Cos(angleRads)),
		y: s.location.y + int32(float64(move.Movement.distance)*math.Sin(angleRads)),
	}
	return SweeperRobotState{
		newLocation,
		s.angle,
		s.state,
	}, newLocation
}

// SetState возвращает новое состояние робота с новым CleaningState
func (s *SweeperRobotState) SetState(newState CleaningState) SweeperRobotState {
	return SweeperRobotState{
		location: s.location,
		angle:    s.angle,
		state:    newState,
	}
}

// Turn возвращает новое состояние робота с новым углом
func (s *SweeperRobotState) Turn(turn Turn) SweeperRobotState {
	return SweeperRobotState{
		location: s.location,
		angle:    turn.Angle,
		state:    s.state,
	}
}
