package gocv

import (
	"image"
	"unsafe"
)

type DetectPtr func(hid unsafe.Pointer, left, top, right, bottom int, mask, dst View)

type FuncPtr func(src *float32, size int, dst *float32)

type Hid struct {
	handle Handle
	data   *Data
	detect DetectPtr
}

func Parallel(begin, end int, function, threadNumber, blockStepMin int) {
}

func (h *Hid) Detect(mask View, rect image.Rectangle, dst View, threadNumber int, throughColumn bool) {
	s := SizeOpMinus(dst.Size(), h.data.size)
	m := mask.RegionPos(s, MiddleCenter)

	// r := RectOpIntersect(RectOpShift(rect, h.data.size.Div(-2)), Size2Rect(s))
	r := RectOpIntersect(RectOpShift(rect, SizeOpDiv(h.data.size, -2)), Size2Rect(s))

	Fill(dst, 0)
	DetectionPrepare(unsafe.Pointer(h.handle))

	h.detect(unsafe.Pointer(h.handle), r.Min.X, r.Min.Y, r.Max.X, r.Max.Y, m, dst)

	//FIXME: support parallel detection
}
