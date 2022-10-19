package type_util

import (
	"reflect"
	"unsafe"
)

func GetPointer[T any](s []T) unsafe.Pointer {
	// return unsafe.Pointer(&s[0])  // len(s) > 0
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	return unsafe.Pointer(sh.Data)
}

func GetSize[T Numeric]() uintptr {
	return unsafe.Sizeof(T(0))
}
