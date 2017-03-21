package gocv

type SimdCompareType int

const (
	/*! equal to: a == b */
	SimdCompareEqual SimdCompareType = iota
	/*! equal to: a != b */
	SimdCompareNotEqual
	/*! equal to: a > b */
	SimdCompareGreater
	/*! equal to: a >= b */
	SimdCompareGreaterOrEqual
	/*! equal to: a < b */
	SimdCompareLesser
	/*! equal to: a <= b */
	SimdCompareLesserOrEqual
)

type SimdPixelFormatType int

const (
	/*! An undefined pixel format. */
	SimdPixelFormatNone = iota
	/*! A 8-bit gray pixel format. */
	SimdPixelFormatGray8
	/*! A 16-bit (2 8-bit channels) pixel format (UV plane of NV12 pixel format). */
	SimdPixelFormatUv16
	/*! A 24-bit (3 8-bit channels) BGR (Blue, Green, Red) pixel format. */
	SimdPixelFormatBgr24
	/*! A 32-bit (4 8-bit channels) BGRA (Blue, Green, Red, Alpha) pixel format. */
	SimdPixelFormatBgra32
	/*! A single channel 16-bit integer pixel format. */
	SimdPixelFormatInt16
	/*! A single channel 32-bit integer pixel format. */
	SimdPixelFormatInt32
	/*! A single channel 64-bit integer pixel format. */
	SimdPixelFormatInt64
	/*! A single channel 32-bit float point pixel format. */
	SimdPixelFormatFloat
	/*! A single channel 64-bit float point pixel format. */
	SimdPixelFormatDouble
	/*! A 8-bit Bayer pixel format (GRBG). */
	SimdPixelFormatBayerGrbg
	/*! A 8-bit Bayer pixel format (GBRG). */
	SimdPixelFormatBayerGbrg
	/*! A 8-bit Bayer pixel format (RGGB). */
	SimdPixelFormatBayerRggb
	/*! A 8-bit Bayer pixel format (BGGR). */
	SimdPixelFormatBayerBggr
	/*! A 24-bit (3 8-bit channels) HSV (Hue, Saturation, Value) pixel format. */
	SimdPixelFormatHsv24
	/*! A 24-bit (3 8-bit channels) HSL (Hue, Saturation, Lightness) pixel format. */
	SimdPixelFormatHsl24
)

type SimdOperationBinary8uType int

const (
	/*! Computes the average value for every channel of every point of two images. \n Average(a, b) = (a + b + 1)/2. */
	SimdOperationBinary8uAverage SimdOperationBinary8uType = iota
	/*! Computes the bitwise AND between two images. */
	SimdOperationBinary8uAnd
	/*! Computes the bitwise OR between two images. */
	SimdOperationBinary8uOr
	/*! Computes maximal value for every channel of every point of two images. */
	SimdOperationBinary8uMaximum
	/*! Computes minimal value for every channel of every point of two images. */
	SimdOperationBinary8uMinimum
	/*!Subtracts unsigned 8-bit integer b from unsigned 8-bit integer a and saturates (for every channel of every point of the images). */
	SimdOperationBinary8uSaturatedSubtraction
	/*!Adds unsigned 8-bit integer b from unsigned 8-bit integer a and saturates (for every channel of every point of the images). */
	SimdOperationBinary8uSaturatedAddition
)

type Position int

const (
	TopLeft      Position = iota /*!< A position in the top-left corner. */
	TopCenter                    /*!< A position at the top center. */
	TopRight                     /*!< A position in the top-right corner. */
	MiddleLeft                   /*!< A position of the left in the middle. */
	MiddleCenter                 /*!< A central position. */
	MiddleRight                  /*!< A position of the right in the middle. */
	BottomLeft                   /*!< A position in the bottom-left corner. */
	BottomCenter                 /*!< A position at the bottom center. */
	BottomRight                  /*!< A position in the bottom-right corner. */
)
