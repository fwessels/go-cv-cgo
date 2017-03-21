package gocv

type Similar struct {
	_sizeDifferenceMax float64
}

func (s Similar) Assest(o1 Object, o2 Object) bool {
	r1 := o1.Rect
	r2 := o2.Rect
	delta := int(s._sizeDifferenceMax * float64((min(r1.Width(), r2.Width()) + min(r1.Height(), r2.Height()))) * 0.5)

	return abs(r1.left-r2.left) <= delta &&
		abs(r1.top-r2.top) <= delta &&
		abs(r1.right-r2.right) <= delta &&
		abs(r1.bottom-r2.bottom) <= delta
}
