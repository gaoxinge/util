package atomic

import (
	"sync/atomic"
)

// Atomic int8 implementation.
type Int8 struct {
	v int32
}

func NewInt8() *Int8 {
	return new(Int8)
}

func (i *Int8) Add(delta int8) (new int8) {
	return int8(atomic.AddInt32(&(i.v), int32(delta)))
}

func (i *Int8) Load() (val int8) {
	return int8(atomic.LoadInt32(&(i.v)))
}

func (i *Int8) Store(val int8) {
	atomic.StoreInt32(&(i.v), int32(val))
}

func (i *Int8) Swap(new int8) (old int8) {
	return int8(atomic.SwapInt32(&(i.v), int32(new)))
}

func (i *Int8) CompareAndSwap(old, new int8) (swapped bool) {
	return atomic.CompareAndSwapInt32(&(i.v), int32(old), int32(new))
}
