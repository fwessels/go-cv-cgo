package gocv

// #cgo pkg-config: simd
// #include "stdlib.h"
// #include "Simd/SimdLib.h"
// #cgo LDFLAGS: -lstdc++
import "C"

const (
	/*! A HAAR cascade classifier type. */
	SimdDetectionInfoFeatureHaar = 0
	/*! A LBP cascade classifier type. */
	SimdDetectionInfoFeatureLbp = 1
	/*! A mask to select cascade classifier type. */
	SimdDetectionInfoFeatureMask = 3
	/*! A flag which defines existence of tilted features in the HAAR cascade. */
	SimdDetectionInfoHasTilted = 4
	/*! A flag which defines possibility to use 16-bit integers for calculation. */
	SimdDetectionInfoCanInt16 = 8
)

type Data struct {
	handle Handle
	size   Size
	flags  C.SimdDetectionInfoFlags
	tag    Tag
}

func (d *Data) Int16() bool {
	return d.flags&SimdDetectionInfoCanInt16 != 0
}

func (d *Data) Haar() bool {
	return d.flags&SimdDetectionInfoFeatureMask == SimdDetectionInfoFeatureHaar
}

func (d *Data) Tilted() bool {
	return d.flags&SimdDetectionInfoHasTilted != 0
}
