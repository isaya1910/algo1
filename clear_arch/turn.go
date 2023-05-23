package main

import "strconv"

type Turn struct {
	Angle Angle
}

func ParseTurn(angleText string) (*Turn, error) {
	angleValue, err := strconv.Atoi(angleText)
	if err != nil {
		return nil, err
	}
	return &Turn{
		Angle(int32(angleValue)),
	}, nil
}
