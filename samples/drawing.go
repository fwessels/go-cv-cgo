package main

import (
	"image"
	"log"

	_ "image/jpeg"

	gocv "github.com/fwessels/go-cv"
)

func main() {

	view := gocv.View{}
	if ok := view.LoadPGM("data/images/lena.pgm"); !ok {
		log.Fatal("Unable to load PGM image")
	}

	gocv.DrawRectangle(view, image.Rect(10, 10, 100, 100), 255, 1)

	view.Save("/tmp/result.pgm")
}
