package atomic

import (
	"sync/atomic"
)

// Atomic int16 implementation.
type Int16 struct {
	v int32
}

func NewInt16() *Int16 {
	return new(Int16)
}

func (i *Int16) Add(delta int16) (new int16) {
	return int16(atomic.AddInt32(&(i.v), int32(delta)))
}

func (i *Int16) Load() (val int16) {
	return int16(atomic.LoadInt32(&(i.v)))
}

func (i *Int16) Store(val int16) {
	atomic.StoreInt32(&(i.v), int32(val))
}

func (i *Int16) Swap(new int16) (old int16) {
	return int16(atomic.SwapInt32(&(i.v), int32(new)))
}

func (i *Int16) CompareAndSwap(old, new int16) (swapped bool) {
	return atomic.CompareAndSwapInt32(&(i.v), int32(old), int32(new))
}
