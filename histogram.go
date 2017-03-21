package gocv

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

import "unsafe"

// ingroup histogram
//SimdAbsSecondDerivativeHistogram(const uint8_t * src, size_t width, size_t height, size_t stride, size_t step, size_t indent, uint32_t * histogram)
//SimdHistogram(const uint8_t * src, size_t width, size_t height, size_t stride, uint32_t * histogram)
//SimdHistogramMasked(const uint8_t * src, size_t srcStride, size_t width, size_t height, const uint8_t * mask, size_t maskStride, uint8_t index, uint32_t * histogram)
//SimdHistogramConditional(const uint8_t * src, size_t srcStride, size_t width, size_t height, const uint8_t * mask, size_t maskStride, uint8_t value, SimdCompareType compareType, uint32_t * histogram)

const HISTOGRAM_SIZE = 256 // = UCHAR_MAX + 1

// SimdHistogram(const uint8_t * src, size_t width, size_t height, size_t stride, uint32_t * histogram)
func Histogram(src View, histogram []uint32) {
	C.SimdHistogram((*C.uint8_t)(src.data), C.size_t(src.width), C.size_t(src.height), C.size_t(src.stride), (*C.uint32_t)(unsafe.Pointer(&histogram)))
}

//SimdNormalizeHistogram(const uint8_t * src, size_t srcStride, size_t width, size_t height, uint8_t * dst, size_t dstStride)
func NormalizeHistogram(src, dst View) {
	C.SimdNormalizeHistogram((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}
