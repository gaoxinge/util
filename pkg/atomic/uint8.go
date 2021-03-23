package atomic

import (
	"sync/atomic"
)

// Atomic uint8 implementation.
type Uint8 struct {
	v uint32
}

func NewUint8() *Uint8 {
	return new(Uint8)
}

func (i *Uint8) Add(delta uint8) (new uint8) {
	return uint8(atomic.AddUint32(&(i.v), uint32(delta)))
}

func (i *Uint8) Load() (val uint8) {
	return uint8(atomic.LoadUint32(&(i.v)))
}

func (i *Uint8) Store(val uint8) {
	atomic.StoreUint32(&(i.v), uint32(val))
}

func (i *Uint8) Swap(new uint8) (old uint8) {
	return uint8(atomic.SwapUint32(&(i.v), uint32(new)))
}

func (i *Uint8) CompareAndSwap(old, new uint8) (swapped bool) {
	return atomic.CompareAndSwapUint32(&(i.v), uint32(old), uint32(new))
}
