package gocv

import "image"

type Anchor struct {
	p   image.Point
	val uint16
}

/* func AnchorOpCompare(a, b Anchor) bool {
	return a.val > b.val
} */

// AxisSorter sorts planets by axis.
type Anchors []Anchor

func (a Anchors) Len() int           { return len(a) }
func (a Anchors) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Anchors) Less(i, j int) bool { return a[i].val < a[j].val }
