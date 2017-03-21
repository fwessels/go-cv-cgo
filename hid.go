package gocv

import "unsafe"

type DetectPtr func(hid unsafe.Pointer, left, top, right, bottom int, mask, dst View)

type FuncPtr func(src *float32, size int, dst *float32)

type Hid struct {
	handle Handle
	data   *Data
	detect DetectPtr
}

func Parallel(begin, end int, function, threadNumber, blockStepMin int) {
}

func (h *Hid) Detect(mask View, rect Rect, dst View, threadNumber int, throughColumn bool) {
	s := dst.Size().Minus(h.data.size)
	m := mask.RegionPos(s, MiddleCenter)

	r := rect.Shifted(h.data.size.Div(-2)).Intersection(Size2Rect(s))
	Fill(dst, 0)
	DetectionPrepare(unsafe.Pointer(h.handle))

	h.detect(unsafe.Pointer(h.handle), r.left, r.top, r.right, r.bottom, m, dst)

	//FIXME: support parallel detection
}
