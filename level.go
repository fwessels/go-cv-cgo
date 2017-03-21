package gocv

type Level struct {
	hids  []Hid
	scale float64

	src  View
	roi  View
	mask View

	rect Rect

	sum    View
	sqsum  View
	tilted View

	dst View

	throughColumn bool
	needSqsum     bool
	needTilted    bool
}
