package gocv

import "image"

type Object struct {
	Rect   image.Rectangle
	weight int
	tag    Tag
}
