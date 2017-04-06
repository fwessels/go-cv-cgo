package main

import (
	"log"

	_ "image/jpeg"

	gocv "github.com/fwessels/go-cv"
)

func main() {

	contours := []gocv.Contour{}
	view := gocv.View{}

	// go-cv loads image
	if ok := view.LoadPGM("data/images/lena.pgm"); !ok {
		log.Fatal("Cannot load image")
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

	// Print found face coordinates
	log.Printf("%+v", contours[0])

}
