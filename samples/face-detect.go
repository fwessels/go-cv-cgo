package main

import (
	"log"

	gocv "github.com/fwessels/go-cv"
)

func main() {

	objects := []gocv.Object{}

	img := gocv.View{}

	// Load image
	if ok := img.Load("data/images/lena.pgm"); !ok {
		log.Fatal("Cannot load image")
	}

	// Initialize new detection engine
	detect := gocv.NewDetection()

	// Load face detection xml
	if ok := detect.Load("data/cascades/haar_face_0.xml"); !ok {
		log.Fatal("face detection xml not loaded")
	}

	// Init
	if ok := detect.Init(img.Size()); !ok {
		log.Fatal("cannot init detect with image size")
	}

	// Detect
	objects, ok := detect.Detect(img, objects)
	if !ok {
		log.Fatal("detection failed")
	}

	if len(objects) == 0 {
		log.Fatal("no objects found")
	}

	// Print found face coordinates
	log.Printf("%+v", objects[0].Rect)
}
