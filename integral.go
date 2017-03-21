package gocv

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

// ingroup integral
func Integral4(src, sum, sqsum, tilted View) {
	C.SimdIntegral((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), (*C.uint8_t)(sum.data), C.size_t(sum.stride), (*C.uint8_t)(sqsum.data), C.size_t(sqsum.stride), (*C.uint8_t)(tilted.data), C.size_t(tilted.stride), C.SimdPixelFormatType(sum.format), C.SimdPixelFormatType(sqsum.format))
}

func Integral3(src, sum, sqsum View) {
	assert(src.width+1 == sum.width && src.height+1 == sum.height && sum.Size().Equals(sqsum.Size()))
	assert(src.format == GRAY8 && sum.format == INT32 && (sqsum.format == INT32 || sqsum.format == DOUBLE))
	C.SimdIntegral((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), (*C.uint8_t)(sum.data), C.size_t(sum.stride), (*C.uint8_t)(sqsum.data), C.size_t(sqsum.stride), nil, 0, C.SimdPixelFormatType(sum.format), C.SimdPixelFormatType(sqsum.format))
}

func Integral2(src, sum View) {
	assert(src.width+1 == sum.width && src.height+1 == sum.height)
	assert(src.format == GRAY8 && sum.format == INT32)
	C.SimdIntegral((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), (*C.uint8_t)(sum.data), C.size_t(sum.stride), nil, 0, nil, 0, C.SimdPixelFormatType(sum.format), C.SimdPixelFormatType(SimdPixelFormatNone))
}
