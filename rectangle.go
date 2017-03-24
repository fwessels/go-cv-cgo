package gocv

import "image"

func RectOpDiv(r image.Rectangle, s float64) image.Rectangle {
	//FIXME: is Round really used
	return image.Rect(
		Round(float64(r.Min.X)/s),
		Round(float64(r.Min.Y)/s),
		Round(float64(r.Max.X)/s),
		Round(float64(r.Max.Y)/s),
	)
}

func RectOpMul(r image.Rectangle, s float64) image.Rectangle {
	//FIXME: is Round really used
	return image.Rect(
		Round(float64(r.Min.X)*s),
		Round(float64(r.Min.Y)*s),
		Round(float64(r.Max.X)*s),
		Round(float64(r.Max.Y)*s),
	)
}

func RectOpAdd(r1 image.Rectangle, r2 image.Rectangle) image.Rectangle {
	// FIXME: Use Convert()
	return image.Rect(
		r1.Min.X+r2.Min.X,
		r1.Min.Y+r2.Min.Y,
		r1.Max.X+r2.Max.X,
		r1.Max.Y+r2.Max.Y,
	)
}

func RectCopy(r image.Rectangle) image.Rectangle {
	return image.Rect(r.Min.X, r.Min.Y, r.Max.X, r.Max.Y)
}

func RectAssign(r1 image.Rectangle, r2 image.Rectangle) {
	r1.Min.X = r2.Min.X
	r1.Min.Y = r2.Min.Y
	r1.Max.X = r2.Max.X
	r1.Max.Y = r2.Max.Y
}

// FIXME: test needed
func RectOpOr(r1 image.Rectangle, r2 image.Rectangle) image.Rectangle {
	if r1.Empty() {
		return RectCopy(r1)
	}
	if r2.Empty() {
		return RectCopy(r2)
	}
	return image.Rect(
		min(r1.Min.X, r2.Min.X),
		min(r1.Min.Y, r2.Min.Y),
		max(r1.Max.X, r2.Max.X),
		max(r1.Max.Y, r2.Max.Y),
	)
}

func RectOpAnd(r1 image.Rectangle, r2 image.Rectangle) image.Rectangle {
	if r1.Empty() {
		return RectCopy(r1)
	}
	if r2.Empty() {
		return RectCopy(r2)
	}

	var left, top, right, bottom int

	if r1.Min.X < r2.Min.X {
		left = min(r2.Min.X, r1.Max.X)
	}
	if r1.Min.Y < r2.Min.Y {
		top = min(r2.Min.Y, r1.Max.Y)
	}
	if r1.Max.X > r2.Max.X {
		right = max(r2.Max.X, r1.Min.X)
	}
	if r1.Max.Y > r2.Max.Y {
		bottom = max(r2.Max.Y, r1.Min.Y)
	}
	return image.Rect(left, top, right, bottom)
}

func RectOpIntersect(r1 image.Rectangle, r2 image.Rectangle) image.Rectangle {
	return r1.Intersect(r2)
}

func RectOpShift(r1 image.Rectangle, shift Size) image.Rectangle {
	return r1.Add(image.Point(shift))
}

func Size2Rect(s Size) image.Rectangle {
	return image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{s.X, s.Y},
	}
}
