package dynArra

import (
	"fmt"
)

type DynArray[T any] struct {
	count    int
	capacity int
	array    []T
}

func (da *DynArray[T]) Init() {
	da.count = 0
	da.MakeArray(16)
}

func (da *DynArray[T]) MakeArray(sz int) {
	var arr = make([]T, sz)
	//  копируем содержимое array в arr ...
	copy(arr, da.array)
	da.capacity = sz
	da.array = arr //
}

func (da *DynArray[T]) Insert(itm T, index int) error {
	if da.count == index {
		da.Append(itm)
		return nil
	}
	if index >= da.count || index < 0 {
		return fmt.Errorf("bad index '%d'", index)
	}

	if da.count == da.capacity {
		da.MakeArray(2 * da.capacity)
	}

	for i := da.count - 1; i > index; i-- {
		da.array[i] = da.array[i-1]
	}
	da.array[index] = itm
	da.count++
	return nil
}

func (da *DynArray[T]) Remove(index int) error {
	if index >= len(da.array) || index < 0 {
		return fmt.Errorf("bad index '%d'", index)
	}
	for i := index; i < da.count-1; i++ {
		da.array[i] = da.array[i+1]
	}
	da.count--

	capacityPercent := float32(da.count) / float32(da.capacity)
	newCapacity := (da.capacity * 2) / 3

	if capacityPercent < 0.5 && newCapacity >= 16 {
		da.capacity = newCapacity
		return nil
	}
	if capacityPercent < 0.5 && newCapacity < 16 {
		da.capacity = 16
		return nil
	}
	return nil
}

func (da *DynArray[T]) Append(itm T) {
	if da.count == da.capacity {
		da.MakeArray(2 * da.capacity)
	}
	da.array[da.count] = itm
	da.count++
}

func (da *DynArray[T]) GetItem(index int) (T, error) {
	var result T
	if index >= len(da.array) || index < 0 {
		return result, fmt.Errorf("bad index '%d'", index)
	}
	result = da.array[index]
	return result, nil
}
