package main

import (
	"fmt"
	"image/color"
	"log"
	"net/http"

	"github.com/DanielSchuette/bioinformatics/plot"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", drawHandler)

	// listen and serve on port 8080
	fmt.Printf("visit 127.0.0.1%s in your browser...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// drawHandler opens a canvas/plot in a browser window
func drawHandler(w http.ResponseWriter, r *http.Request) {
	canvas := plot.NewCanvas(400, 600, &color.RGBA{255, 255, 255, 255})
	canvas.Rectangle(10, 10, 390, 590, 8, &color.RGBA{0, 0, 0, 255})
	canvas.AddLabel(50, 50, "hello gopher!", &color.RGBA{50, 50, 50, 255})
	canvas.EncodePNG(w)
}
