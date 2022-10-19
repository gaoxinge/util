package type_util

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestGetPointer(t *testing.T) {
	v := []float32{1, 2, 3}
	vPtr := GetPointer(v)
	vSize := GetSize[float32]()
	assert.Equal(t, *(*float32)(vPtr), v[0])
	assert.Equal(t, *(*float32)(unsafe.Add(vPtr,  vSize)), v[1])
	assert.Equal(t, *(*float32)(unsafe.Add(vPtr,  2 * vSize)), v[2])
}
