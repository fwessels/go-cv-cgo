package gocv

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"
import (
	"image"
	"math"
	"unsafe"
)

// ingroup object_detection

func DetectionLoadA(path string) unsafe.Pointer {

	return C.SimdDetectionLoadA(C.CString(path))
}

func DetectionInfo(data unsafe.Pointer, width *C.size_t, height *C.size_t, flags *C.SimdDetectionInfoFlags) {
	C.SimdDetectionInfo(data, width, height, flags)
}

func DetectionInit(data unsafe.Pointer, sum, sqsum, tilted View, throughColumn, int16 int) unsafe.Pointer {

	return C.SimdDetectionInit(data, (*C.uint8_t)(sum.data), C.size_t(sum.stride), C.size_t(sum.width), C.size_t(sum.height), (*C.uint8_t)(sqsum.data), C.size_t(sqsum.stride), (*C.uint8_t)(tilted.data), C.size_t(tilted.stride), C.int(throughColumn), C.int(int16))
}

func DetectionPrepare(hid unsafe.Pointer) {

	C.SimdDetectionPrepare(hid)
}

func DetectionHaarDetect32fp(hid unsafe.Pointer, left, top, right, bottom int, mask, dst View) {

	C.SimdDetectionHaarDetect32fp(hid, (*C.uint8_t)(mask.data), C.size_t(mask.stride), C.ptrdiff_t(left), C.ptrdiff_t(top), C.ptrdiff_t(right), C.ptrdiff_t(bottom), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func DetectionHaarDetect32fi(hid unsafe.Pointer, left, top, right, bottom int, mask, dst View) {

	C.SimdDetectionHaarDetect32fi(hid, (*C.uint8_t)(mask.data), C.size_t(mask.stride), C.ptrdiff_t(left), C.ptrdiff_t(top), C.ptrdiff_t(right), C.ptrdiff_t(bottom), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func DetectionLbpDetect32fp(hid unsafe.Pointer, left, top, right, bottom int, mask, dst View) {

	C.SimdDetectionLbpDetect32fp(hid, (*C.uint8_t)(mask.data), C.size_t(mask.stride), C.ptrdiff_t(left), C.ptrdiff_t(top), C.ptrdiff_t(right), C.ptrdiff_t(bottom), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func DetectionLbpDetect32fi(hid unsafe.Pointer, left, top, right, bottom int, mask, dst View) {

	C.SimdDetectionLbpDetect32fi(hid, (*C.uint8_t)(mask.data), C.size_t(mask.stride), C.ptrdiff_t(left), C.ptrdiff_t(top), C.ptrdiff_t(right), C.ptrdiff_t(bottom), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func DetectionLbpDetect16ip(hid unsafe.Pointer, left, top, right, bottom int, mask, dst View) {

	C.SimdDetectionLbpDetect16ip(hid, (*C.uint8_t)(mask.data), C.size_t(mask.stride), C.ptrdiff_t(left), C.ptrdiff_t(top), C.ptrdiff_t(right), C.ptrdiff_t(bottom), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func DetectionLbpDetect16ii(hid unsafe.Pointer, left, top, right, bottom int, mask, dst View) {

	C.SimdDetectionLbpDetect16ii(hid, (*C.uint8_t)(mask.data), C.size_t(mask.stride), C.ptrdiff_t(left), C.ptrdiff_t(top), C.ptrdiff_t(right), C.ptrdiff_t(bottom), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

func DetectionFree(data unsafe.Pointer) {

	C.SimdDetectionFree(data)
}

type Tag int
type Handle unsafe.Pointer

type Detection struct {
	_data              []Data
	_imageSize         Size
	_threadNumber      int
	_needNormalization bool
	_levels            []Level
}

const UNDEFINED_OBJECT_TAG = Tag(-1)

func NewDetection() *Detection {
	return &Detection{
		_data:   []Data{},
		_levels: []Level{},
	}
}

func (d *Detection) Load(path string) bool {

	handle := DetectionLoadA(path)
	//
	if handle != nil {
		data := Data{}
		data.handle = Handle(handle)
		data.tag = UNDEFINED_OBJECT_TAG
		DetectionInfo(handle,
			(*C.size_t)(unsafe.Pointer(&data.size.X)),
			(*C.size_t)(unsafe.Pointer(&data.size.Y)),
			&data.flags)
		d._data = append(d._data, data)

	}
	return handle != nil
}

func (d *Detection) Init(s Size) bool {
	if len(d._data) == 0 {
		return false
	}
	d._imageSize = s
	d._threadNumber = 1
	return d.InitLevels(1.1, Size{0, 0}, Size{math.MaxInt32, math.MaxInt32}, View{})
}

func (d *Detection) InitLevels(scaleFactor float64, sizeMin Size, sizeMax Size, roi View) bool {
	d._needNormalization = false
	d._levels = make([]Level, 0, 100)
	scale := float64(1.0)
	for {
		inserts := make([]bool, len(d._data))
		exit := true
		insert := false
		for i := 0; i < len(d._data); i++ {
			windowSize := SizeOpMul(d._data[i].size, scale)
			if windowSize.X <= sizeMax.X && windowSize.Y <= sizeMax.Y && windowSize.X <= d._imageSize.X && windowSize.Y <= d._imageSize.Y {
				if windowSize.X >= sizeMin.X && windowSize.Y >= sizeMin.Y {
					insert = true
					inserts[i] = true
				}
				exit = false
			}
		}
		if exit {
			break
		}
		if insert {
			level := Level{}
			level.scale = scale
			level.throughColumn = scale <= 2.0
			scaledSize := SizeOpDiv(d._imageSize, scale)

			level.src.Recreate(scaledSize.X, scaledSize.Y, GRAY8)
			level.roi.Recreate(scaledSize.X, scaledSize.Y, GRAY8)
			level.mask.Recreate(scaledSize.X, scaledSize.Y, GRAY8)

			level.sum.Recreate(scaledSize.X+1, scaledSize.Y+1, INT32)
			level.sqsum.Recreate(scaledSize.X+1, scaledSize.Y+1, INT32)
			level.tilted.Recreate(scaledSize.X+1, scaledSize.Y+1, INT32)

			level.dst.Recreate(scaledSize.X, scaledSize.Y, GRAY8)

			level.needSqsum = false
			level.needTilted = false

			for i := 0; i < len(d._data); i++ {
				if !inserts[i] {
					continue
				}
				handle := DetectionInit(unsafe.Pointer(d._data[i].handle), level.sum, level.sqsum, level.tilted,
					b2i(level.throughColumn), b2i(d._data[i].Int16()))
				if handle != nil {
					hid := Hid{}
					hid.handle = Handle(handle)
					hid.data = &d._data[i]
					if d._data[i].Haar() {
						if level.throughColumn {
							hid.detect = DetectionHaarDetect32fi
						} else {
							hid.detect = DetectionHaarDetect32fp
						}
					} else {
						if d._data[i].Int16() {
							if level.throughColumn {
								hid.detect = DetectionLbpDetect16ii
							} else {
								hid.detect = DetectionLbpDetect16ip
							}
						} else {
							if level.throughColumn {
								hid.detect = DetectionLbpDetect32fi
							} else {
								hid.detect = DetectionLbpDetect32fp
							}
						}
					}
					level.hids = append(level.hids, hid)
				} else { // handle
					return false
				}

				level.needSqsum = level.needSqsum || d._data[i].Haar()
				level.needTilted = level.needTilted || d._data[i].Tilted()
				d._needNormalization = d._needNormalization || d._data[i].Haar()
			}
			level.rect = Size2Rect(level.roi.Size())
			if roi.format == NONE {
				Fill(level.roi, 255)
			} else {
				ResizeBilinear(roi, level.roi)
				Binarization(level.roi, 0, 255, 0, level.roi, SimdCompareGreater)
				SegmentationShrinkRegion(level.roi, 255, level.rect)
			}

			d._levels = append(d._levels, level)
		}
		scale *= scaleFactor
	}

	return !(len(d._levels) == 0)
}

func (d *Detection) FillMotionMask(rects []image.Rectangle, level Level, rect image.Rectangle) {
	Fill(level.mask, 0)
	for i := 0; i < len(rects); i++ {
		r := RectOpDiv(rects[i], level.scale)
		RectAssign(rect, RectOpOr(rect, r))
		Fill(level.mask.RegionRect(r), 0xff)
	}
	RectAssign(rect, RectOpAnd(rect, level.rect))
	OperationBinary8u(level.mask, level.roi, level.mask, SimdOperationBinary8uAnd)
}

func (d *Detection) FillLevels(src View) {
	gray := View{}
	if src.format != GRAY8 {
		gray.Recreate(src.Size().X, src.Size().Y, GRAY8)
		Convert(src, gray)
		src = gray
	}

	ResizeBilinear(src, d._levels[0].src)

	if d._needNormalization {
		NormalizeHistogram(d._levels[0].src, d._levels[0].src)
	}
	d.EstimateIntegral(d._levels[0])
	for i := 1; i < len(d._levels); i++ {
		ResizeBilinear(d._levels[0].src, d._levels[i].src)
		d.EstimateIntegral(d._levels[i])
	}
}

func (d *Detection) EstimateIntegral(level Level) {
	if level.needSqsum {
		if level.needTilted {
			Integral4(level.src, level.sum, level.sqsum, level.tilted)
		} else {
			Integral3(level.src, level.sum, level.sqsum)
		}
	} else {
		Integral2(level.src, level.sum)
	}
}

func addr2uint8(addr uintptr, index int) uint8 {
	newAddr := addr + uintptr(index)
	return *(*uint8)(unsafe.Pointer(newAddr))
}

func (d *Detection) AddObjects(objs []Object, dst View, rect image.Rectangle, size Size, scale float64, step int, tag Tag) []Object {
	s := SizeOpMinus(dst.Size(), size)
	r := RectOpIntersect(RectOpShift(rect, SizeOpDiv(size, -2)), Size2Rect(s))
	for row := r.Min.Y; row < r.Max.Y; row += step {
		// mask := (*(*[]uint8)(dst.data))[row*dst.stride : r.right]
		mask := uintptr(dst.data) + uintptr(row*dst.stride)
		for col := r.Min.X; col < r.Max.X; col += step {
			if addr2uint8(mask, col) != 0 {
				objs = append(objs,
					Object{
						Rect:   RectOpMul(image.Rect(col, row, col+size.X, row+size.Y), scale),
						weight: 1,
						tag:    tag,
					},
				)
			}
		}
	}
	return objs
}

func (d *Detection) Partition(vec []Object, labels []int, sizeDifferenceMax float64) (int, []int) {
	similar := Similar{sizeDifferenceMax}
	i := len(vec)
	j := i
	N := i
	const PARENT = 0
	const RANK = 1

	nodes := make([][2]int, N)

	for i = 0; i < N; i++ {
		nodes[i][PARENT] = -1
		nodes[i][RANK] = 0
	}

	for i = 0; i < N; i++ {
		root := i
		for nodes[root][PARENT] >= 0 {
			root = nodes[root][PARENT]
		}
		for j = 0; j < N; j++ {
			if i == j || !similar.Assest(vec[i], vec[j]) {
				continue
			}
			root2 := j
			for nodes[root2][PARENT] >= 0 {
				root2 = nodes[root2][PARENT]
			}
			if root2 != root {
				rank := nodes[root][RANK]
				rank2 := nodes[root2][RANK]
				if rank > rank2 {
					nodes[root2][PARENT] = root
				} else {
					nodes[root][PARENT] = root2
					if rank == rank2 {
						nodes[root2][RANK]++
					}
					root = root2
				}
				assert(nodes[root][PARENT] < 0)
				k := j
				parent := nodes[k][PARENT]
				for ; parent >= 0; parent = nodes[k][PARENT] {
					nodes[k][PARENT] = root
					k = parent
				}

				k = i
				parent = nodes[k][PARENT]
				for ; parent >= 0; parent = nodes[k][PARENT] {
					nodes[k][PARENT] = root
					k = parent
				}
			}
		}
	}

	old := labels
	labels = make([]int, len(old)+N)
	copy(labels, old)

	nclasses := 0

	for i = 0; i < N; i++ {
		root := i
		for nodes[root][PARENT] >= 0 {
			root = nodes[root][PARENT]
		}
		if nodes[root][RANK] >= 0 {
			nodes[root][RANK] = ^nclasses
			nclasses++
		}
		labels[i] = ^nodes[root][RANK]
	}

	return nclasses, labels
}

func (d *Detection) GroupObjects(dst []Object, src []Object, groupSizeMin int, sizeDifferenceMax float64) []Object {

	if groupSizeMin == 0 || len(src) < groupSizeMin {
		return nil
	}

	nclasses, labels := d.Partition(src, []int{}, sizeDifferenceMax)

	buffer := make([]Object, nclasses)

	for i := 0; i < len(labels); i++ {
		cls := labels[i]
		buffer[cls].Rect = RectOpAdd(buffer[cls].Rect, src[i].Rect)
		buffer[cls].weight++
		buffer[cls].tag = src[i].tag
	}
	for i := 0; i < len(buffer); i++ {
		buffer[i].Rect = RectOpDiv(buffer[i].Rect, float64(buffer[i].weight))
	}
	for i := 0; i < len(buffer); i++ {
		r1 := buffer[i].Rect
		n1 := buffer[i].weight
		if n1 < int(groupSizeMin) {
			continue
		}
		j := 0
		for j = 0; j < len(buffer); j++ {
			n2 := buffer[j].weight
			if j == i || n2 < int(groupSizeMin) {
				continue
			}

			r2 := buffer[j].Rect

			dx := Round(float64(r2.Dx()) * sizeDifferenceMax)
			dy := Round(float64(r2.Dy()) * sizeDifferenceMax)

			if i != j && (n2 > max(3, n1) || n1 < 3) &&
				r1.Min.X >= r2.Min.X-dx &&
				r1.Min.Y >= r2.Min.Y-dy &&
				r1.Max.X <= r2.Max.X+dx &&
				r1.Max.Y <= r2.Max.Y+dy {
				break
			}
		}
		if j == len(buffer) {
			dst = append(dst, buffer[i])
		}
	}

	return dst
}

func (d *Detection) Detect(src View, objects []Object) ([]Object, bool) {
	groupSizeMin := 3
	sizeDifferenceMax := 0.2
	motionMask := false
	motionsRegions := []image.Rectangle{}

	if len(d._levels) == 0 || !SizeOpEquals(src.Size(), d._imageSize) {
		return nil, false
	}

	d.FillLevels(src)

	candidates := make(map[Tag][]Object)
	for i := 0; i < len(d._levels); i++ {
		level := d._levels[i]
		mask := level.roi
		rect := level.rect
		if motionMask {
			d.FillMotionMask(motionsRegions, level, rect)
			mask = level.mask
		}
		if rect.Empty() {
			continue
		}
		for j := 0; j < len(level.hids); j++ {
			hid := level.hids[j]
			hid.Detect(mask, rect, level.dst, d._threadNumber, level.throughColumn)
			step := 1
			if level.throughColumn {
				step += 1
			}
			candidates[hid.data.tag] = d.AddObjects(candidates[hid.data.tag], level.dst, rect, hid.data.size, level.scale, step, hid.data.tag)
		}
	}

	for _, v := range candidates {
		objects = d.GroupObjects(objects, v, groupSizeMin, sizeDifferenceMax)
	}

	return objects, true
}
