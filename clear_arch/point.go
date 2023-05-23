package main

import "errors"

type Point struct {
	x int32
	y int32
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
