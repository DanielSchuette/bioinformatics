package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"net/http"

	"github.com/DanielSchuette/bioinformatics/plot"
)

const port = ":8080"

var (
	canvas   *plot.Canvas
	saveName = "plot_test.png"
	writer   = flag.String("w", "file", "where to draw the plot\neither one of 'browser', 'file'")
)

func main() {
	// parse command line flags
	flag.Parse()

	// create a plot
	canvas = plot.NewCanvas(400, 600, &color.RGBA{255, 255, 255, 255})
	canvas.Rectangle(10, 10, 390, 590, 8, &color.RGBA{0, 0, 0, 255})
	canvas.AddLabel(50, 50, "hello gopher!", &color.RGBA{50, 50, 50, 255})

	switch *writer {
	case "browser":
		// handle image request at '/'
		http.HandleFunc("/", drawHandler)

		// listen and serve on port 8080
		fmt.Printf("visit 127.0.0.1%s in your browser...\n", port)
		log.Fatal(http.ListenAndServe(port, nil))
	case "file":
		canvas.SaveToFile(saveName)
		fmt.Printf("plot: saved plot as %s\n", saveName)
	default:
		log.Fatalf("plot: don't know option %s\n", *writer)
	}
}

// drawHandler opens a canvas/plot in a browser window
func drawHandler(w http.ResponseWriter, r *http.Request) {
	canvas.EncodePNG(w)
}
