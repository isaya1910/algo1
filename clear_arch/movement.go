package main

import "strconv"

type Movement struct {
	distance int32
}

type Move struct {
	Movement Movement
}

func ParseMove(distanceText string) (*Move, error) {
	distance, err := strconv.Atoi(distanceText)
	if err != nil {
		return nil, err
	}

	return &Move{
		Movement: Movement{
			distance: int32(distance),
		},
	}, nil
}
