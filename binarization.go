package gocv

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

// ingroup binarization
//SimdBinarization(const uint8_t * src, size_t srcStride, size_t width, size_t height, uint8_t value, uint8_t positive, uint8_t negative, uint8_t * dst, size_t dstStride, SimdCompareType compareType)
//SimdAveragingBinarization(const uint8_t * src, size_t srcStride, size_t width, size_t height, uint8_t value, size_t neighborhood, uint8_t threshold, uint8_t positive, uint8_t negative, uint8_t * dst, size_t dstStride, SimdCompareType compareType)

func Binarization(src View, value uint8, positive uint8, negative uint8, dst View, compareType SimdCompareType) {
	C.SimdBinarization(
		(*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height),
		C.uint8_t(value),
		C.uint8_t(positive),
		C.uint8_t(negative),
		(*C.uint8_t)(dst.data), C.size_t(dst.stride),
		C.SimdCompareType(compareType),
	)
}
