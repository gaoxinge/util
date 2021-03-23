package atomic

import (
	"sync/atomic"
)

// Atomic uint16 implementation.
type Uint16 struct {
	v uint32
}

func NewUint16() *Uint16 {
	return new(Uint16)
}

func (i *Uint16) Add(delta uint16) (new uint16) {
	return uint16(atomic.AddUint32(&(i.v), uint32(delta)))
}

func (i *Uint16) Load() (val uint16) {
	return uint16(atomic.LoadUint32(&(i.v)))
}

func (i *Uint16) Store(val uint16) {
	atomic.StoreUint32(&(i.v), uint32(val))
}

func (i *Uint16) Swap(new uint16) (old uint16) {
	return uint16(atomic.SwapUint32(&(i.v), uint32(new)))
}

func (i *Uint16) CompareAndSwap(old, new uint16) (swapped bool) {
	return atomic.CompareAndSwapUint32(&(i.v), uint32(old), uint32(new))
}
