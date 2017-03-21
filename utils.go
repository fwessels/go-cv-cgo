package gocv

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func i2b(i int) bool {
	if i == 0 {
		return true
	}
	return false
}

func assert(b bool) {
	if !b {
		panic("oups")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

// FIXME: Use accelerated version of Round
func Round(value float64) int {
	if value < 0 {
		return int(value - 0.5)
	}
	return int(value + 0.5)
}
