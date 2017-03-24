package gocv

import "image"

type Level struct {
	hids  []Hid
	scale float64

	src  View
	roi  View
	mask View

	rect image.Rectangle

	sum    View
	sqsum  View
	tilted View

	dst View

	throughColumn bool
	needSqsum     bool
	needTilted    bool
}
