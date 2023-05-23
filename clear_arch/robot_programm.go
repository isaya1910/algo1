package main

import (
	"fmt"
	"strings"
)

/*Само задание я уже выполнил в ООП стиле настолько, насколько позваоляет Golang.
Из плюсов своего решения могу отметить что я сдела абстрактным CleaningState и Робота, то есть в будущем сама система
будет довольно гибкой и можно легко добавлять новые стейты. То же самое касается и робота,  есть возможность делать различные
вариации.

Из минус решения я назову саму main функцию, так как существенная часть логики программы, такая как логика принятия комманд
и отлавливание ошибок с парсинга не инкапсулированы в отдельные структуры.
*/

/*
	Рефлексия по поводу задания по написанию кода в ООП стиле:
		Я добавил новую структуру назвав ее RobotProgram и инкапсулировав в нее логику самой программы,
		парсинга входных данных. В предыдущем варианте, было также частично написано на ООП но сама программа
		была реализована в функции main что является признаком процедурного стиля.
		Также я изменил логику комманды Move на правильную, используя вектроное вычесление координат.
*/
/*
	Плюсы подхода:
		Операции и состояния инкапсулированы в сущность. Код лучше читается, писать unit тесты легче,
     	есть возможность замокать какие то зависимости.
	Минксы подхода:
		Мутабельность состояния сущности робота. State, его текуща точка в пространстве и тд. Пока отсутвует
		возможность работать в многопоточном режиме.
*/

type RobotProgram struct {
	Robot *SweeperRobot
}

func (r RobotProgram) runCommands(commands []string) {
	robot := *r.Robot
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

func main() {
	robot := &SweeperRobot{
		currentLocation: Point{},
		state:           Water{},
		angle:           0,
	}
	commands := []string{"move 100", "turn -90", "set soap", "move 120", "stop"}
	program := RobotProgram{
		Robot: robot,
	}

	program.runCommands(commands)

}
