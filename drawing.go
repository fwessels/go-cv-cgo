package gocv

import "image"

// ingroup drawing
//SimdAlphaBlending(const uint8_t * src, size_t srcStride, size_t width, size_t height, size_t channelCount, const uint8_t * alpha, size_t alphaStride, uint8_t * dst, size_t dstStride)

func DrawLine(canvas View, p1 image.Point, p2 image.Point, color uint8, width int) {

	assert(PixelSize(canvas.format) == 1)

	w := canvas.width - 1
	h := canvas.height - 1

	x1 := p1.X
	y1 := p1.Y
	x2 := p2.X
	y2 := p2.Y

	if x1 < 0 || y1 < 0 || x1 > w || y1 > h || x2 < 0 || y2 < 0 || x2 > w || y2 > h {
		if (x1 < 0 && x2 < 0) || (y1 < 0 && y2 < 0) || (x1 > w && x2 > w) || (y1 > h && y2 > h) {
			return
		}
		if y1 == y2 {
			x1 = min(max(x1, 0), w)
			x2 = min(max(x2, 0), w)
		} else if x1 == x2 {
			y1 = min(max(y1, 0), h)
			y2 = min(max(y2, 0), h)
		} else {
			x0 := (x1*y2 - y1*x2) / (y2 - y1)
			y0 := (y1*x2 - x1*y2) / (x2 - x1)
			xh := (x1*y2 - y1*x2 + h*(x2-x1)) / (y2 - y1)
			yw := (y1*x2 - x1*y2 + w*(y2-y1)) / (x2 - x1)
			if x1 < 0 {
				x1 = 0
				y1 = y0
			}
			if x2 < 0 {
				x2 = 0
				y2 = y0
			}
			if x1 > w {
				x1 = w
				y1 = yw
			}
			if x2 > w {
				x2 = w
				y2 = yw
			}

			if (y1 < 0 && y2 < 0) || (y1 > h && y2 > h) {
				return
			}

			if y1 < 0 {
				x1 = x0
				y1 = 0
			}
			if y2 < 0 {
				x2 = x0
				y2 = 0
			}
			if y1 > h {
				x1 = xh
				y1 = h
			}
			if y2 > h {
				x2 = xh
				y2 = h
			}
		}
	}

	inverse := abs(y2-y1) > abs(x2-x1)
	if inverse {
		x1, y1 = y1, x1
		x2, y2 = y2, x2
	}

	if x1 > x2 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}

	dx := float64(x2 - x1)
	dy := float64(abs(y2 - y1))

	err := float64(dx / 2.0)

	ystep := -1
	if y1 < y2 {
		ystep = 1
	}

	y0 := y1 - width/2

	for x := x1; x <= x2; x++ {
		for i := 0; i < width; i++ {
			y := y0 + i
			if y >= 0 {
				if inverse {
					if y < w {
						canvas.PixSet(y, x, color)
					}
				} else {
					if y < h {
						canvas.PixSet(x, y, color)
					}
				}
			}
		}

		//
		err -= dy
		if err < 0 {
			y0 += ystep
			err += dx
		}
	}
}
