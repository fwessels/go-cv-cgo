package main

import (
	"image"
	"log"
	"os"

	_ "image/jpeg"

	gocv "github.com/fwessels/go-cv"
)

func main() {

	objects := []gocv.Object{}
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

	// Initialize new detection engine
	detect := gocv.NewDetection()

	// Load face detection xml
	if ok := detect.Load("data/cascades/haar_face_0.xml"); !ok {
		log.Fatal("face detection xml not loaded")
	}

	// Init
	if ok := detect.Init(view.Size()); !ok {
		log.Fatal("cannot init detect with image size")
	}

	// Detect faces
	objects, ok := detect.Detect(view, objects)
	if !ok {
		log.Fatal("detection failed")
	}

	if len(objects) == 0 {
		log.Fatal("no objects found")
	}

	// Print found face coordinates
	log.Printf("%+v", objects[0].Rect)

}
