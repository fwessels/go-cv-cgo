package gocv

type Rect struct {
	left   int
	top    int
	right  int
	bottom int
}

func (r Rect) Height() int {
	// FIXME: weird, this should not be negative
	return r.bottom - r.top
}

func (r Rect) Width() int {
	// FIXME: weird, this should not be negative
	return r.right - r.left
}

func (r Rect) Area() int {
	return r.Width() * r.Height()
}

func (r Rect) Empty() bool {
	return r.Area() == 0
}

func (r Rect) Div(s float64) Rect {
	//FIXME: is Round really used
	return Rect{
		left:   Round(float64(r.left) / s),
		top:    Round(float64(r.top) / s),
		right:  Round(float64(r.right) / s),
		bottom: Round(float64(r.bottom) / s),
	}
}

func (r Rect) Mul(s float64) Rect {
	//FIXME: is Round really used
	return Rect{
		left:   Round(float64(r.left) * s),
		top:    Round(float64(r.top) * s),
		right:  Round(float64(r.right) * s),
		bottom: Round(float64(r.bottom) * s),
	}
}

func (r Rect) Add(rect Rect) Rect {
	// FIXME: Use Convert()
	return Rect{
		left:   r.left + rect.left,
		top:    r.top + rect.top,
		right:  r.right + rect.right,
		bottom: r.bottom + rect.bottom,
	}
}

func (r Rect) Copy() Rect {
	return Rect{
		left:   r.left,
		top:    r.top,
		right:  r.right,
		bottom: r.bottom,
	}
}

func (r *Rect) Assign(rect Rect) {
	r.left = rect.left
	r.top = rect.top
	r.right = rect.right
	r.bottom = r.bottom
}

// FIXME: test needed
func (r Rect) Or(rect Rect) Rect {
	if r.Empty() {
		return r.Copy()
	}
	if rect.Empty() {
		return rect.Copy()
	}
	_r := rect.Copy()
	new := Rect{}
	new.left = min(r.left, _r.left)
	new.top = min(r.top, _r.top)
	new.right = max(r.right, _r.right)
	new.bottom = max(r.bottom, _r.bottom)
	return new
}

func (r Rect) And(rect Rect) Rect {
	if r.Empty() {
		return r.Copy()
	}
	if rect.Empty() {
		return rect.Copy()
	}
	new := Rect{}
	_r := rect.Copy()
	if r.left < _r.left {
		new.left = min(_r.left, r.right)
	}
	if r.top < _r.top {
		new.top = min(_r.top, r.bottom)
	}
	if r.right > _r.right {
		new.right = max(_r.right, r.left)
	}
	if r.bottom > _r.bottom {
		new.bottom = max(_r.bottom, r.top)
	}
	return new
}

func (r Rect) Intersection(rect Rect) Rect {
	left := max(r.left, rect.left)
	top := max(r.top, rect.top)
	right := max(left, min(r.right, rect.right))
	bottom := max(top, min(r.bottom, rect.bottom))
	return Rect{left: left, top: top, right: right, bottom: bottom}
}

func (r Rect) Shifted(shift Size) Rect {
	return Rect{
		left:   r.left + shift.x,
		top:    r.top + shift.y,
		right:  r.right + shift.x,
		bottom: r.bottom + shift.y,
	}
}

func Size2Rect(s Size) Rect {
	return Rect{
		left:   0,
		top:    0,
		right:  s.x,
		bottom: s.y,
	}
}
