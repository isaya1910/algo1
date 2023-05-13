package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/*Само задание я уже выполнил в процедурном стиле также используя golang структуры и интефейсы.
Из плюсов своего решения могу отметить что я сдела абстрактным CleaningState и Робота, то есть в будущем сама система
будет довольно гибкой и можно легко добавлять новые стейты. То же самое касается и робота,  есть возможность делать различные
вариации.

Из минус решения я назову саму main функцию, так как существенная часть логики программы, такая как логика принятия комманд
и отлавливание ошибок с парсинга не инкапсулированы в отдельные структуры.
*/

type Point struct {
	x int32
	y int32
}

type Movement struct {
	x int32
	y int32
}

type Angle int32

func ParseTurn(angleText string) (*Turn, error) {
	angleValue, err := strconv.Atoi(angleText)
	if err != nil {
		return nil, err
	}
	return &Turn{
		Angle(int32(angleValue)),
	}, nil
}

type Move struct {
	Movement Movement
}

func ParseMove(pointsText string) (*Move, error) {
	movementText := strings.Split(pointsText, ",")
	if len(movementText) != 2 {
		return nil, errors.New("parse error")
	}

	xValue, err := strconv.Atoi(movementText[0])
	if err != nil {
		return nil, err
	}

	yValue, err := strconv.Atoi(movementText[1])
	if err != nil {
		return nil, err
	}

	return &Move{
		Movement: Movement{
			x: int32(xValue),
			y: int32(yValue),
		},
	}, nil
}

type Turn struct {
	Angle Angle
}

type CleaningState interface {
	Clean() string
}

func ParseCleaningState(stateText string) (CleaningState, error) {
	switch stateText {
	case "water":
		return Water{}, nil
	case "soap":
		return Soap{}, nil
	case "brush":
		return Brush{}, nil
	default:
		return nil, errors.New("parse error")
	}
}

type Water struct{}

func (Water) Clean() string {
	return "Clean with water"
}

type Soap struct{}

func (Soap) Clean() string {
	return "Clean with soap"
}

type Brush struct{}

func (Brush) Clean() string {
	return "Clean with brush"
}

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
	s.currentLocation = Point{
		x: s.currentLocation.x + move.Movement.x,
		y: s.currentLocation.y + move.Movement.y,
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

func main() {
	robot := SweeperRobot{
		currentLocation: Point{},
		state:           Water{},
		angle:           0,
	}

	commands := []string{"move 100,100", "turn -90", "set soap", "move 50,-50", "stop"}

	for _, command := range commands {
		parsedCommand := strings.Split(command, " ")
		if len(parsedCommand) == 0 {
			fmt.Print("invalid command")
			break
		}

		if parsedCommand[0] != "start" && parsedCommand[0] != "stop" && len(parsedCommand) != 2 {
			fmt.Println("invalid command")
			break
		}

		switch parsedCommand[0] {
		case "move":
			move, err := ParseMove(parsedCommand[1])
			if err != nil {
				fmt.Println(err)
				break
			}
			robot.Move(*move)
		case "set":
			state, err := ParseCleaningState(parsedCommand[1])
			if err != nil {
				fmt.Println(err)
				break
			}
			robot.SetState(state)
		case "turn":
			turn, err := ParseTurn(parsedCommand[1])
			if err != nil {
				fmt.Println(err)
				break
			}
			robot.Turn(*turn)

		case "start":
			robot.Start()

		case "stop":
			robot.Stop()
		}

	}

}
