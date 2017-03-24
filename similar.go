package gocv

type Similar struct {
	_sizeDifferenceMax float64
}

func (s Similar) Assest(o1 Object, o2 Object) bool {
	r1 := o1.Rect
	r2 := o2.Rect
	delta := int(s._sizeDifferenceMax * float64((min(r1.Dx(), r2.Dx()) + min(r1.Dy(), r2.Dy()))) * 0.5)

	return abs(r1.Min.X-r2.Min.X) <= delta &&
		abs(r1.Min.Y-r2.Min.Y) <= delta &&
		abs(r1.Max.X-r2.Max.X) <= delta &&
		abs(r1.Max.Y-r2.Max.Y) <= delta
}
