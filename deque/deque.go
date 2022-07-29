package deque

import (
	"fmt"
)

type Deque[T any] struct {
	dataSlice []T
}

func (d *Deque[T]) Size() int {
	return len(d.dataSlice)
}

func (d *Deque[T]) AddFront(itm T) {
	d.dataSlice = append([]T{itm}, d.dataSlice...)
}

func (d *Deque[T]) AddTail(itm T) {
	d.dataSlice = append(d.dataSlice, itm)
}

func (d *Deque[T]) RemoveFront() (T, error) {
	var result T
	if len(d.dataSlice) == 0 {
		return result, fmt.Errorf("deque is empty")
	}

	result = d.dataSlice[0]
	d.dataSlice = d.dataSlice[1:]
	return result, nil
}

func (d *Deque[T]) RemoveTail() (T, error) {
	var result T
	if len(d.dataSlice) == 0 {
		return result, fmt.Errorf("deque is empty")
	}

	result = d.dataSlice[d.Size()-1]

	d.dataSlice = d.dataSlice[:d.Size()-1]

	return result, nil
}
