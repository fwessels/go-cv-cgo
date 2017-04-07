package gocv

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

import (
	"image"
	"sort"
)

// ingroup contour
//SimdContourMetrics(const uint8_t * src, size_t srcStride, size_t width, size_t height, uint8_t * dst, size_t dstStride)
func ContourMetrics(src, dst View) {
	assert(SizeOpEquals(src.Size(), dst.Size()) && src.format == GRAY8 && dst.format == INT16)
	C.SimdContourMetrics((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height), (*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

//SimdContourMetricsMasked(const uint8_t * src, size_t srcStride, size_t width, size_t height, const uint8_t * mask, size_t maskStride, uint8_t indexMin, uint8_t * dst, size_t dstStride)
func ContourMetricsMasked(src, mask View, indexMin uint8, dst View) {
	assert(ViewCompatible(src, mask) && SizeOpEquals(src.Size(), dst.Size()) && src.format == GRAY8 && dst.format == INT16)
	C.SimdContourMetricsMasked((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height),
		(*C.uint8_t)(mask.data), C.size_t(mask.stride),
		C.uint8_t(indexMin),
		(*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

//SimdContourAnchors(const uint8_t * src, size_t srcStride, size_t width, size_t height, size_t step, int16_t threshold, uint8_t * dst, size_t dstStride)
func ContourAnchors(src View, step int, threshold int16, dst View) {
	C.SimdContourAnchors((*C.uint8_t)(src.data), C.size_t(src.stride), C.size_t(src.width), C.size_t(src.height),
		C.size_t(step), C.int16_t(threshold),
		(*C.uint8_t)(dst.data), C.size_t(dst.stride))
}

type Direction int16

const (
	Unknown Direction = iota - 1
	Up
	Down
	Right
	Left
)

type Contour struct {
	P []image.Point
}

type ContourDetector struct {
	_roi     image.Rectangle
	_m       View
	_a       View
	_e       View
	_anchors Anchors
}

func NewContourDetector() *ContourDetector {
	return &ContourDetector{}
}

func (c *ContourDetector) Init(size Size) {
	c._m.Recreate(size.X, size.Y, INT16)
	c._a.Recreate(size.X, size.Y, GRAY8)
	c._e.Recreate(size.X, size.Y, GRAY8)
}

func (c *ContourDetector) Detect(src View, contours []Contour) ([]Contour, bool) {
	mask := View{}
	indexMin := uint8(3)
	roi := image.Rectangle{}
	gradientThreshold := uint16(40)
	anchorThreshold := 0
	anchorScanInterval := 2
	minSegmentLength := 2

	if !ViewCompatible(src, c._a) {
		return nil, false
	}

	if mask.format != NONE && !ViewCompatible(mask, c._a) {
		return nil, false
	}

	if roi.Empty() {
		c._roi = image.Rect(0, 0, src.Size().X, src.Size().Y)
	} else {
		c._roi = roi
	}

	c._roi = RectOpIntersect(c._roi, image.Rect(0, 0, src.Size().X, src.Size().Y))

	c.ContourMetrics(src, mask, indexMin)

	if gradientThreshold < 0 {
		gradientThreshold = c.EstimateAdaptiveThreshold()
	}

	c.ContourAnchors(anchorThreshold, anchorScanInterval)

	contours = c.PerformSmartRouting(contours, minSegmentLength, gradientThreshold*2)

	return contours, true
}

func (c *ContourDetector) PerformSmartRouting(contours []Contour, minSegmentLength int, gradientThreshold uint16) []Contour {
	e := c._e.RegionRect(c._roi)
	frame := image.Rect(1, 1, e.width-1, e.height-1)
	Fill(e.RegionRect(frame), 0)
	FillFrame(e, frame, 255)

	for i := 0; i < len(c._anchors); i++ {
		anchor := c._anchors[i]
		if anchor.val > 0 {
			contour := Contour{}
			contour.P = make([]image.Point, 0, 200)
			contours, contour = c.SmartRoute(contours, contour, anchor.p.X, anchor.p.Y, minSegmentLength, gradientThreshold, Unknown)
			if len(contour.P) > minSegmentLength {
				contours = append(contours, contour)
			}
		}
	}

	return contours
}

func (c *ContourDetector) SmartRoute(contours []Contour, contour Contour, x, y, minSegmentLength int, gradientThreshold uint16, direction Direction) ([]Contour, Contour) {
	switch direction {
	case Unknown:
	case Left:
		for c.CheckMetricsForMagnitudeAndDirection(x, y, gradientThreshold, 1) {
			if c._e.PixGet(x, y) == 0 {
				c._e.PixSet(x, y, 255)
				if len(contour.P) != 0 && (abs(contour.P[len(contour.P)-1].X-x) > 1 || abs(contour.P[len(contour.P)-1].Y-y) > 1) {
					if len(contour.P) > minSegmentLength {
						contours = append(contours, contour)
					}
					contour.P = []image.Point{}
				}
				contour.P = append(contour.P, image.Point{X: x, Y: y})
			}
			if c.CheckMetricsForMagnitudeMaximum(x-1, y-1, x-1, y, x-1, y+1) {
				x--
				y--
			} else if c.CheckMetricsForMagnitudeMaximum(x-1, y+1, x-1, y, x-1, y-1) {
				x--
				y++
			} else {
				x--
			}
			if c._e.PixGet(x, y) != 0 {
				break
			}
		}
	case Right:
		for c.CheckMetricsForMagnitudeAndDirection(x, y, gradientThreshold, 1) {
			if c._e.PixGet(x, y) == 0 {
				c._e.PixSet(x, y, 255)
				if len(contour.P) != 0 && (abs(contour.P[len(contour.P)-1].X-x) > 1 || abs(contour.P[len(contour.P)-1].Y-y) > 1) {
					if len(contour.P) > minSegmentLength {
						contours = append(contours, contour)
					}
					contour.P = []image.Point{}
				}
				contour.P = append(contour.P, image.Point{X: x, Y: y})
			}
			if c.CheckMetricsForMagnitudeMaximum(x+1, y-1, x+1, y, x+1, y+1) {
				x++
				y--
			} else if c.CheckMetricsForMagnitudeMaximum(x+1, y+1, x+1, y, x+1, y-1) {
				x++
				y++
			} else {
				x++
			}
			if c._e.PixGet(x, y) != 0 {
				break
			}
		}
	case Up:
		for c.CheckMetricsForMagnitudeAndDirection(x, y, gradientThreshold, 0) {
			if c._e.PixGet(x, y) == 0 {
				c._e.PixSet(x, y, 255)
				if len(contour.P) != 0 && (abs(contour.P[len(contour.P)-1].X-x) > 1 || abs(contour.P[len(contour.P)-1].Y-y) > 1) {
					if len(contour.P) > minSegmentLength {
						contours = append(contours, contour)
					}
					contour.P = []image.Point{}
				}
				contour.P = append(contour.P, image.Point{X: x, Y: y})
			}
			if c.CheckMetricsForMagnitudeMaximum(x-1, y-1, x, y-1, x+1, y-1) {
				x--
				y--
			} else if c.CheckMetricsForMagnitudeMaximum(x+1, y-1, x, y-1, x-1, y-1) {
				x++
				y--
			} else {
				y--
			}
			if c._e.PixGet(x, y) != 0 {
				break
			}
		}
	case Down:
		for c.CheckMetricsForMagnitudeAndDirection(x, y, gradientThreshold, 0) {
			if c._e.PixGet(x, y) == 0 {
				c._e.PixSet(x, y, 255)
				if len(contour.P) != 0 && (abs(contour.P[len(contour.P)-1].X-x) > 1 || abs(contour.P[len(contour.P)-1].Y-y) > 1) {
					if len(contour.P) > minSegmentLength {
						contours = append(contours, contour)
					}
					contour.P = []image.Point{}
				}
				contour.P = append(contour.P, image.Point{X: x, Y: y})
			}
			if c.CheckMetricsForMagnitudeMaximum(x+1, y+1, x, y+1, x-1, y+1) {
				x++
				y++
			} else if c.CheckMetricsForMagnitudeMaximum(x-1, y+1, x, y+1, x+1, y+1) {
				x--
				y++
			} else {
				y++
			}
			if c._e.PixGet(x, y) != 0 {
				break
			}
		}
	}

	if c._e.PixGet(x, y) != 0 || c._m.PixGet16(x, y) < gradientThreshold {
		return contours, contour
	}

	d := c._m.PixGet16(x, y) & 1
	if d == 0 {
		contours, contour = c.SmartRoute(contours, contour, x, y, minSegmentLength, gradientThreshold, Up)
		contours, contour = c.SmartRoute(contours, contour, x, y, minSegmentLength, gradientThreshold, Down)
	} else if d == 1 {
		contours, contour = c.SmartRoute(contours, contour, x, y, minSegmentLength, gradientThreshold, Right)
		contours, contour = c.SmartRoute(contours, contour, x, y, minSegmentLength, gradientThreshold, Left)
	}

	return contours, contour
}

func (c *ContourDetector) CheckMetricsForMagnitudeAndDirection(x, y int, gradientThreshold, direction uint16) bool {
	m := c._m.PixGet16(x, y)
	return m >= gradientThreshold && (m&1) == direction
}

func (c *ContourDetector) CheckMetricsForMagnitudeMaximum(x0, y0, x1, y1, x2, y2 int) bool {
	m0 := c._m.PixGet16(x0, y0) | 1
	m1 := c._m.PixGet16(x1, y1) | 1
	m2 := c._m.PixGet16(x2, y2) | 1
	return m0 > m1 && m0 > m2
}

func (c *ContourDetector) ContourAnchors(anchorThreshold, anchorScanInterval int) {
	ContourAnchors(c._m.RegionRect(c._roi), anchorScanInterval, int16(anchorThreshold), c._a.RegionRect(c._roi))
	c._anchors = []Anchor{}
	for row := c._roi.Min.Y + 1; row < c._roi.Max.Y-1; row += anchorScanInterval {
		for col := c._roi.Min.X; col < c._roi.Max.X-1; col += anchorScanInterval {
			if c._a.PixGet(col, row) != 0 {
				c._anchors = append(c._anchors, Anchor{p: image.Point{X: col, Y: row}, val: c._m.PixGet16(col, row) / 2})
			}
		}
	}
	sort.Sort(c._anchors)
}

func (c *ContourDetector) EstimateAdaptiveThreshold() uint16 {
	roiSize := c._roi.Size()
	mSize := c._m.Size()
	if roiSize.X >= mSize.X || roiSize.Y >= mSize.Y {
		assert(true)
	}
	m := c._m.RegionRect(c._roi)
	size := m.Size()

	var value uint16
	var sum uint32
	var count int

	for i := 0; i < size.X; i++ {
		for j := 0; j < size.Y; j++ {
			value = m.PixGet16(i, j)
			if value != 0 {
				count++
				value = value >> 1
				sum += uint32(value)
			}
		}
	}
	meanThreshold := uint16(float64(sum) / float64(count))
	return meanThreshold
}

func (c *ContourDetector) ContourMetrics(src, mask View, indexMin uint8) {
	if mask.format == GRAY8 {
		ContourMetricsMasked(src.RegionRect(c._roi), mask.RegionRect(c._roi), indexMin, c._m.RegionRect(c._roi))
	} else {
		ContourMetrics(src.RegionRect(c._roi), c._m.RegionRect(c._roi))
	}
}
