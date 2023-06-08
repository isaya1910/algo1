package main

import (
	"fmt"
	"math"
)

/*
	Интерефейс робота, в рамках текущего демо проекта представлен его реализацией SweeperRobot.
	SweeperRobot в свою же очередь имеет зависомость RobotRepository который также является интерфейсом.
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

/*
	Заменил interface RobotRepository на фунциональную иньекцию зависимостей
*/
/*
	Не смог придумать заменить это одной
*/
type SaveRobotLocationFunc func(point Point)
type GetRobotLocationFunc func() Point
type SaveNewStateFunc func(state CleaningState)
type GetStateFunc func() CleaningState
type GetAngleFunc func() Angle
type SaveAngleFunc func(angle Angle)

type SweeperRobot struct {
	saveRobotLocation SaveRobotLocationFunc
	getRobotLocation  GetRobotLocationFunc
	saveNewState      SaveNewStateFunc
	getState          GetStateFunc
	getAngle          GetAngleFunc
	saveAngle         SaveAngleFunc
}

func (s *SweeperRobot) Move(move Move) {
	angleRads := float64(s.getAngle()) * (math.Pi / 180.0)
	repo := s
	repo.saveRobotLocation(Point{
		x: repo.getRobotLocation().x + int32(float64(move.Movement.distance)*math.Cos(angleRads)),
		y: repo.getRobotLocation().y + int32(float64(move.Movement.distance)*math.Sin(angleRads)),
	})
	fmt.Println(repo.getRobotLocation())
}

func (s *SweeperRobot) SetState(state CleaningState) {
	s.saveNewState(state)
	fmt.Println(s.getState().Clean())
}

func (s *SweeperRobot) Start() {
	fmt.Sprintf("Started in location: ", s.getRobotLocation())
	s.getState().Clean()
}

func (s *SweeperRobot) Turn(turn Turn) {
	s.saveAngle(turn.Angle)
	fmt.Println(s.getAngle())
}

func (s *SweeperRobot) Stop() {
	fmt.Println("Work stopped")
}
