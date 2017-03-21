package gocv

type Size Point

func (s Size) Div(scale float64) Size {
	//FIXME: check if round is used in C++
	return Size{x: Round(float64(s.x) / scale), y: Round(float64(s.y) / scale)}
}

func (s Size) Mul(scale float64) Size {
	return Size{x: Round(float64(s.x) * scale), y: Round(float64(s.y) * scale)}
}

func (s Size) Minus(size Size) Size {
	return Size{
		x: s.x - size.x,
		y: s.y - size.y,
	}
}

func (s Size) Equals(size Size) bool {
	return s.x == size.x && s.y == size.y
}
