package main

type CleaningState interface {
	Clean() string
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
