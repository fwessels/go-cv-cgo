package gocv

import "C"

// type ptrdiff_t int64

type Point struct {
	x int
	y int
}

func NewPoint(w int, h int) Point {
	return Point{x: w, y: h}
}
