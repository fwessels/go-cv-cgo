package main

import (
	"image"
	"log"
	"os"

	_ "image/jpeg"

	gocv "github.com/fwessels/go-cv"
)

func main() {

	contours := []gocv.Contour{}
	view := gocv.View{}

	// Load jpeg image
	file, err := os.Open("data/images/lena.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode file using go's image API
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// go-cv loads image
	if err := view.LoadImage(img, gocv.GRAY8); err != nil {
		log.Fatal("Cannot load image, ", err)
	}

	// Initialize new contouer detection
	detect := gocv.NewContourDetector()

	// Init
	detect.Init(view.Size())

	// Detect contours
	contours, ok := detect.Detect(view, contours)

	if !ok {
		log.Fatal("Contour detection failed.")
	}

	if len(contours) == 0 {
		log.Fatal("no contours found")
	} else {
		log.Printf("contours found = %d\n", len(contours))
	}

	for i := 0; i < len(contours); i++ {
		for j := 1; j < len(contours[i].P); j++ {
			gocv.DrawLine(view, contours[i].P[j-1], contours[i].P[j], 255, 1)
		}
	}

	view.Save("/tmp/result.pgm")
}
