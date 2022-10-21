package atomic

import (
	"sync/atomic"
)

// Int is atomic int implementation.
type Int[T ~int8 | ~int16] struct {
	v int32
}

func NewInt[T ~int8 | ~int16]() *T {
	return new(T)
}

func (i *Int[T]) Add(delta T) (new T) {
	return T(atomic.AddInt32(&(i.v), int32(delta)))
}

func (i *Int[T]) Load() (val T) {
	return T(atomic.LoadInt32(&(i.v)))
}

func (i *Int[T]) Store(val T) {
	atomic.StoreInt32(&(i.v), int32(val))
}

func (i *Int[T]) Swap(new T) (old T) {
	return T(atomic.SwapInt32(&(i.v), int32(new)))
}

func (i *Int[T]) CompareAndSwap(old, new T) (swapped bool) {
	return atomic.CompareAndSwapInt32(&(i.v), int32(old), int32(new))
}

type Int8 Int[int8]
type Int16 Int[int16]
