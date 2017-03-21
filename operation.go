package gocv

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

// ingroup operation
//SimdOperationBinary16i(const uint8_t * a, size_t aStride, const uint8_t * b, size_t bStride, size_t width, size_t height, uint8_t * dst, size_t dstStride, SimdOperationBinary16iType type)
//SimdVectorProduct(const uint8_t * vertical, const uint8_t * horizontal, uint8_t * dst, size_t stride, size_t width, size_t height)

//SimdOperationBinary8u(const uint8_t * a, size_t aStride, const uint8_t * b, size_t bStride, size_t width, size_t height, size_t channelCount, uint8_t * dst, size_t dstStride, SimdOperationBinary8uType type)
func OperationBinary8u(a, b, dst View, operationBinary8u SimdOperationBinary8uType) {
	//FIXME asset
	C.SimdOperationBinary8u((*C.uint8_t)(a.data), C.size_t(a.stride), (*C.uint8_t)(b.data), C.size_t(b.stride), C.size_t(a.width), C.size_t(a.height), C.size_t(a.ChannelCount()), (*C.uint8_t)(dst.data), C.size_t(dst.stride), C.SimdOperationBinary8uType(operationBinary8u))
}
