package gocv

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"
import (
	"fmt"
	"image"
	"image/draw"
	"io/ioutil"
	"os"
	"unsafe"

	"github.com/lazywei/go-opencv/opencv"
)

type Format uint8

// http://ninghang.blogspot.com/2012/11/list-of-mat-type-in-opencv.html
const (
	CV_8UC1  = 0
	CV_8UC2  = 8
	CV_8UC3  = 16
	CV_8UC4  = 24
	CV_16SC1 = 3
	CV_32SC1 = 4
	CV_32FC1 = 5
	CV_64FC1 = 6
)

const (
	NONE Format = iota
	GRAY8
	UV16
	BGR24
	BGRA32
	INT16
	INT32
	INT64
	FLOAT
	DOUBLE
	BAYERGRBG
	BAYERGBRG
	BAYERRGGB
	BAYERBGGR
	HSV24
	HSL24
)

func PixelSize(f Format) int {
	switch f {
	case NONE:
		return 0
	case GRAY8:
		return 1
	case UV16:
		return 2
	case BGR24:
		return 3
	case BGRA32:
		return 4
	case INT16:
		return 2
	case INT32:
		return 4
	case INT64:
		return 8
	case FLOAT:
		return 4
	case DOUBLE:
		return 8
	case BAYERGRBG:
		return 1
	case BAYERGBRG:
		return 1
	case BAYERRGGB:
		return 1
	case BAYERBGGR:
		return 1
	case HSV24:
		return 3
	case HSL24:
		return 3
	default:
		return 0
	}
}

func ChannelCount(f Format) int {
	switch f {
	case NONE:
		return 0
	case GRAY8:
		return 1
	case UV16:
		return 2
	case BGR24:
		return 3
	case BGRA32:
		return 4
	case INT16:
		return 1
	case INT32:
		return 1
	case INT64:
		return 1
	case FLOAT:
		return 1
	case DOUBLE:
		return 1
	case BAYERGRBG:
		return 1
	case BAYERGBRG:
		return 1
	case BAYERRGGB:
		return 1
	case BAYERBGGR:
		return 1
	case HSV24:
		return 3
	case HSL24:
		return 3
	default:
		return 0
	}
}

type View struct {
	width, height int
	format        Format
	stride        int
	owner         bool
	data          unsafe.Pointer
}

func (v *View) Size() Size {
	return Size{x: v.width, y: v.height}
}

func (v *View) ChannelCount() int {
	return ChannelCount(v.format)
}

func (v View) RegionPos(size Size, position Position) View {
	switch position {
	case TopLeft:
		return v.RegionRect(Rect{0, 0, size.x, size.y})
	case TopCenter:
		return v.RegionRect(Rect{(v.width - size.x) / 2, 0, (v.width + size.x) / 2, size.y})
	case TopRight:
		return v.RegionRect(Rect{v.width - size.x, 0, v.width, size.y})
	case MiddleLeft:
		return v.RegionRect(Rect{0, (v.height - size.y) / 2, size.x, (v.height + size.y) / 2})
	case MiddleCenter:
		return v.RegionRect(Rect{(v.width - size.x) / 2, (v.height - size.y) / 2, (v.width + size.x) / 2, (v.height + size.y) / 2})
	case MiddleRight:
		return v.RegionRect(Rect{v.width - size.x, (v.height - size.y) / 2, v.width, (v.height + size.y) / 2})
	case BottomLeft:
		return v.RegionRect(Rect{0, v.height - size.y, size.x, v.height})
	case BottomCenter:
		return v.RegionRect(Rect{(v.width - size.x) / 2, v.height - size.y, (v.width + size.x) / 2, v.height})
	case BottomRight:
		return v.RegionRect(Rect{v.width - size.x, v.height - size.y, v.width, v.height})
	default:
		assert(false)
	}
	return View{}
}

func (v View) RegionRect(rect Rect) View {
	if v.data != nil && rect.right >= rect.left && rect.bottom >= rect.top {
		left := min(max(rect.left, 0), v.width)
		top := min(max(rect.top, 0), v.height)
		right := min(max(rect.right, 0), v.width)
		bottom := min(max(rect.bottom, 0), v.height)

		return View{
			width:  right - left,
			height: bottom - top,
			stride: v.stride,
			format: v.format,
			// data:   unsafe.Pointer(&(*(*[]uint8)(v.data))[top*v.stride+left*PixelSize(v.format)]),
			data: unsafe.Pointer(uintptr(v.data) + uintptr(top*v.stride) + uintptr(left*PixelSize(v.format))),
		}
	} else {
		return View{}
	}

}

// Recreate
func (v *View) Recreate(w, h int, f Format) {

	if v.owner && v.data != nil {
		Free(v.data)
		v.data = nil
		v.owner = false
	}
	v.width = w
	v.height = h
	v.format = f
	v.stride = Align(v.width*PixelSize(v.format), Alignment())
	v.data = Allocate(v.height*v.stride, Alignment())
}

// Load
func (v *View) Load(path string) bool {

	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	var imgtype string
	var w, h, d int

	for i := 0; i < 3; i++ {
		err = nil
		switch i {
		case 0:
			_, err = fmt.Fscanf(file, "%s\n", &imgtype)
		case 1:
			_, err = fmt.Fscanf(file, "%d %d\n", &w, &h)
		case 2:
			_, err = fmt.Fscanf(file, "%d\n", &d)
		}
		if err != nil {
			return false
		}
	}

	if imgtype != "P5" {
		return false
	}

	if d != 255 {
		return false
	}

	v.Recreate(w, h, GRAY8)

	if data, err := ioutil.ReadAll(file); err == nil {
		v.data = C.CBytes(data)
	} else {
		return false
	}

	return true
}

func (v *View) Save(path string) bool {
	if v.format != GRAY8 {
		return false
	}
	f, err := os.Create(path)
	if err != nil {
		return false
	}
	defer f.Close()

	header := fmt.Sprintf("P5\n%d %d\n%d\n", v.width, v.height, 255)

	f.WriteString(header)

	data := C.GoBytes(v.data, C.int(v.width*v.height))

	f.Write(data)

	return true
}

func (v *View) CopyFrom(img *image.RGBA) error {

	for y := 0; y < img.Bounds().Size().Y; y++ {
		start := y * img.Bounds().Size().X * 4
		psrcstart := unsafe.Pointer(uintptr(v.data) + uintptr(start))
		for x := 0; x < img.Bounds().Size().X; x++ {
			psrcdata := (*C.uint)(unsafe.Pointer(uintptr(psrcstart) + uintptr(x*4)))
			*psrcdata = C.uint(uint32(img.Pix[start+x*4+0])<<16 +
				uint32(img.Pix[start+x*4+1])<<8 +
				uint32(img.Pix[start+x*4+2])<<0 +
				uint32(img.Pix[start+x*4+3])<<24)
		}
	}

	return nil
}

func (v *View) CopyTo(img *image.RGBA) error {

	for y := 0; y < img.Bounds().Size().Y; y++ {
		start := y * img.Bounds().Size().X * 4
		psrcstart := unsafe.Pointer(uintptr(v.data) + uintptr(start))
		for x := 0; x < img.Bounds().Size().X; x++ {
			psrcdata := (*C.uint)(unsafe.Pointer(uintptr(psrcstart) + uintptr(x*4)))
			img.Pix[start+x*4+0] = uint8((*psrcdata >> 16) & 0xff)
			img.Pix[start+x*4+1] = uint8((*psrcdata >> 8) & 0xff)
			img.Pix[start+x*4+2] = uint8((*psrcdata >> 0) & 0xff)
			img.Pix[start+x*4+3] = uint8((*psrcdata >> 24) & 0xff)
		}
	}

	return nil
}

// AsRGBA returns an RGBA copy of the supplied image.
func AsRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	img := image.NewRGBA(bounds)
	draw.Draw(img, bounds, src, bounds.Min, draw.Src)
	return img
}

func OcvTo(cvtype int) Format {
	switch cvtype {
	case CV_8UC1:
		return GRAY8
	case CV_8UC2:
		return UV16
	case CV_8UC3:
		return BGR24
	case CV_8UC4:
		return BGRA32
	case CV_16SC1:
		return INT16
	case CV_32SC1:
		return INT32
	case CV_32FC1:
		return FLOAT
	case CV_64FC1:
		return DOUBLE
	}
	return NONE
}

func MatToView(mat *opencv.Mat) View {
	matData := mat.GetData()
	return View{
		width:  mat.Cols(),
		height: mat.Rows(),
		stride: mat.Step(),
		format: OcvTo(mat.Type()),
		data:   unsafe.Pointer(&matData),
	}
}
