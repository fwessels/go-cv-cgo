package gocv

import "image"

type Anchor struct {
	p   image.Point
	val uint16
}

type Anchors []Anchor

func (a Anchors) Len() int      { return len(a) }
func (a Anchors) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Anchors) Less(i, j int) bool {
	if a[i].val != a[j].val {
		return a[i].val > a[j].val
	}
	if a[i].p.X != a[j].p.X {
		return a[i].p.X > a[j].p.X
	}
	return a[i].p.Y > a[j].p.Y
}
