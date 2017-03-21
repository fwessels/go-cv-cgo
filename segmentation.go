package gocv

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"
import "unsafe"

// ingroup segmentation
//SimdSegmentationChangeIndex(uint8_t * mask, size_t stride, size_t width, size_t height, uint8_t oldIndex, uint8_t newIndex)
//SimdSegmentationFillSingleHoles(uint8_t * mask, size_t stride, size_t width, size_t height, uint8_t index)
//SimdSegmentationPropagate2x2(const uint8_t * parent, size_t parentStride, size_t width, size_t height, uint8_t * child, size_t childStride, const uint8_t * difference, size_t differenceStride, uint8_t currentIndex, uint8_t invalidIndex, uint8_t emptyIndex, uint8_t differenceThreshold)

//SimdSegmentationShrinkRegion(const uint8_t * mask, size_t stride, size_t width, size_t height, uint8_t index, ptrdiff_t * left, ptrdiff_t * top, ptrdiff_t * right, ptrdiff_t * bottom)
func SegmentationShrinkRegion(mask View, index uint8, rect Rect) {
	C.SimdSegmentationShrinkRegion(
		(*C.uint8_t)(mask.data), C.size_t(mask.stride), C.size_t(mask.width), C.size_t(mask.height),
		C.uint8_t(index),
		(*C.ptrdiff_t)(unsafe.Pointer(&rect.left)),
		(*C.ptrdiff_t)(unsafe.Pointer(&rect.top)),
		(*C.ptrdiff_t)(unsafe.Pointer(&rect.right)),
		(*C.ptrdiff_t)(unsafe.Pointer(&rect.bottom)),
	)
}
