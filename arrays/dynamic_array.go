package arrays

import (
	"fmt"
)

type DynamicArray[T comparable] struct {
	arr      []T
	size     int
	capacity int
}

func NewDynamicArray[T comparable](initialCapacity int) (*DynamicArray[T], error) {
	if initialCapacity < 0 {
		return nil, fmt.Errorf("initialCapacity cannot be negative: %d.", initialCapacity)
	}
	return &DynamicArray[T]{
		arr:      make([]T, initialCapacity),
		size:     0,
		capacity: initialCapacity,
	}, nil
}

func doubleSize[T comparable](da *DynamicArray[T]) {
	if da.size == da.capacity {
		newCap := da.capacity * 2
		if newCap == 0 {
			newCap = 1
		}
		newArr := make([]T, newCap)
		copy(newArr, da.arr[:da.size])
		da.arr = newArr
		da.capacity = newCap
	}
}

func halveSize[T comparable](da *DynamicArray[T]) {
	if da.capacity > 1 && da.size <= da.capacity/4 {
		newCap := da.capacity / 2
		newArr := make([]T, newCap)
		copy(newArr, da.arr[:da.size])
		da.arr = newArr
		da.capacity = newCap
	}
}

func (da *DynamicArray[T]) Insert(value T) {
	if da.size >= da.capacity {
		doubleSize(da)
	}
	da.arr[da.size] = value
	da.size++
}

func (da *DynamicArray[T]) Find(target T) int {
	for i := range da.size {
		if da.arr[i] == target {
			return i
		}
	}
	return -1
}

func (da *DynamicArray[T]) Delete(target T) error {
	index := da.Find(target)
	if index == -1 {
		return fmt.Errorf("Unable to delete element %v: the entry is not in the array.", target)
	}
	for i := index; i < da.size-1; i++ {
		da.arr[i] = da.arr[i+1]
	}
	da.size--
	if da.capacity > 1 && da.size <= da.capacity/4 {
		halveSize(da)
	}
	return nil
}
