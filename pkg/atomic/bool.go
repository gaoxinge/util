package atomic

import (
	"sync/atomic"
)

// Atomic bool implementation.
// Reference:
//   - https://gist.github.com/mahan/6256149
//   - https://github.com/tevino/abool
type Bool struct {
	v int32
}

func NewBool() *Bool {
	return new(Bool)
}

func (b *Bool) Negative() (new bool) {
	var (
		v       int32
		swapped bool
	)
	for !swapped {
		v = b.v
		swapped = atomic.CompareAndSwapInt32(&(b.v), v, 1 - v)
	}
	return (1 - v) == 1
}

func (b *Bool) Load() (val bool) {
	return atomic.LoadInt32(&(b.v)) == 1
}

func (b *Bool) Store(val bool) {
	var v int32
	if val {
		v = 1
	}
	atomic.StoreInt32(&(b.v), v)
}

func (b *Bool) Swap(new bool) (old bool) {
	var v int32
	if new {
		v = 1
	}
	return atomic.SwapInt32(&(b.v), v) == 1
}

func (b *Bool) CompareAndSwap(old, new bool) (swapped bool) {
	var o, n int32
	if old {
		o = 1
	}
	if new {
		n = 1
	}
	return atomic.CompareAndSwapInt32(&(b.v), o, n)
}
