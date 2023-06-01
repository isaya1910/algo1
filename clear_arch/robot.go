package main

import (
	"fmt"
	"math"
)

/*
	Интерефейс робота, в рамках текущего демо проекта его реализацией SweeperRobot.
	SweeperRobot в свою же очередь имеет зависомость RobotRepository.
*/
type Robot interface {
	Move(move Move)
	Turn(turn Turn)
	SetState(state CleaningState)
	Start()
	Stop()
}

// RobotRepository /* Реализовал stateless архитектуру через паттерн репозиторий
/*
	Убрал все состояния с реализации робота, теперь каждое состояние сохраняется и получается через слой
	абстракции репозитория. Разницы нет что под имплементацией, это может быть как SQL база данных, или key-value
	хранилища либо другие структуры данных.
*/
type RobotRepository interface {
	saveRobotLocation(point Point)
	getRobotLocation() Point
	saveNewState(state CleaningState)
	getState() CleaningState
	getAngle() Angle
	saveAngle(angle Angle)
}

type SweeperRobot struct {
	repo RobotRepository
}

func (s *SweeperRobot) Move(move Move) {
	angleRads := float64(s.repo.getAngle()) * (math.Pi / 180.0)
	repo := s.repo
	repo.saveRobotLocation(Point{
		x: repo.getRobotLocation().x + int32(float64(move.Movement.distance)*math.Cos(angleRads)),
		y: repo.getRobotLocation().y + int32(float64(move.Movement.distance)*math.Sin(angleRads)),
	})
	fmt.Println(repo.getRobotLocation())
}

func (s *SweeperRobot) SetState(state CleaningState) {
	s.repo.saveNewState(state)
	fmt.Println(s.repo.getState().Clean())
}

func (s *SweeperRobot) Start() {
	fmt.Sprintf("Started in location: ", s.repo.getRobotLocation())
	s.repo.getState().Clean()
}

func (s *SweeperRobot) Turn(turn Turn) {
	s.repo.saveAngle(turn.Angle)
	fmt.Println(s.repo.getAngle())
}

func (s *SweeperRobot) Stop() {
	fmt.Println("Work stopped")
}
