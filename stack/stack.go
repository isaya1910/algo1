package stack

import "fmt"

type Stack[T any] struct {
	dataSlice []T
}

func (st *Stack[T]) Size() int {
	return len(st.dataSlice)
}

func (st *Stack[T]) Peek() (T, error) {
	var result T
	if st.Size() == 0 {
		return result, fmt.Errorf("stack is empty")
	}
	result = st.dataSlice[st.Size()-1]
	return result, nil
}

func (st *Stack[T]) Pop() (T, error) {
	var result T
	if st.Size() == 0 {
		return result, fmt.Errorf("stack is empty")
	}
	result = st.dataSlice[st.Size()-1]
	st.dataSlice = st.dataSlice[:st.Size()-1]
	return result, nil
}

func (st *Stack[T]) Push(itm T) {
	st.dataSlice = append(st.dataSlice, itm)
}
