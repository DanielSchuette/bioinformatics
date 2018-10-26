package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", drawHandler)

	// listen and serve on port 8080
	fmt.Printf("visit 127.0.0.1%s in your browser...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func rect(x1, y1, x2, y2, thickness int, img *image.RGBA) {
	col := color.RGBA{255, 255, 255, 255}

	for t := 0; t < thickness; t++ {
		// draw horizontal lines
		for x := x1; x <= x2; x++ {
			img.Set(x, y1+t, col)
			img.Set(x, y2-t, col)
		}
		// draw vertical lines
		for y := y1; y <= y2; y++ {
			img.Set(x1+t, y, col)
			img.Set(x2-t, y, col)
		}
	}
}

// handler to test
func drawHandler(w http.ResponseWriter, r *http.Request) {
	img := image.NewRGBA(image.Rect(0, 0, 1200, 1800))
	rect(5, 5, 1195, 1795, 2, img)
	png.Encode(w, img)
}
