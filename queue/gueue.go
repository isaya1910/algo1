package queue

import (
	"fmt"
)

type Queue[T any] struct {
	dataSlice []T
}

func (q *Queue[T]) Size() int {
	return len(q.dataSlice)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var result T
	if q.Size() == 0 {
		return result, fmt.Errorf("queue is empty")
	}
	result = q.dataSlice[q.Size()-1]
	q.dataSlice = q.dataSlice[:q.Size()-1]
	return result, nil
}

func (q *Queue[T]) Enqueue(itm T) {
	q.dataSlice = append([]T{itm}, q.dataSlice...)
}
