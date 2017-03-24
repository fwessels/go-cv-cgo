package gocv

import "image"

type Size image.Point

func SizeOpEquals(s1 Size, s2 Size) bool {
	return s1.X == s2.X && s1.Y == s2.Y
}

func SizeOpDiv(s Size, scale float64) Size {
	//FIXME: check if round is used in C++
	return Size{X: Round(float64(s.X) / scale), Y: Round(float64(s.Y) / scale)}
}

func SizeOpMul(s Size, scale float64) Size {
	return Size{X: Round(float64(s.X) * scale), Y: Round(float64(s.Y) * scale)}
}

func SizeOpMinus(s1 Size, s2 Size) Size {
	return Size{
		X: s1.X - s2.X,
		Y: s1.Y - s2.Y,
	}
}
