package atomic

import (
	"sync/atomic"
)

// Uint is atomic uint implementation.
type Uint[T ~uint8 | ~uint16] struct {
	v uint32
}

func NewUint[T ~uint8 | ~uint16]() *T {
	return new(T)
}

func (i *Uint[T]) Add(delta T) (new T) {
	return T(atomic.AddUint32(&(i.v), uint32(delta)))
}

func (i *Uint[T]) Load() (val T) {
	return T(atomic.LoadUint32(&(i.v)))
}

func (i *Uint[T]) Store(val T) {
	atomic.StoreUint32(&(i.v), uint32(val))
}

func (i *Uint[T]) Swap(new T) (old T) {
	return T(atomic.SwapUint32(&(i.v), uint32(new)))
}

func (i *Uint[T]) CompareAndSwap(old, new T) (swapped bool) {
	return atomic.CompareAndSwapUint32(&(i.v), uint32(old), uint32(new))
}

type Uint8 Uint[uint8]
type Uint16 Uint[uint16]
