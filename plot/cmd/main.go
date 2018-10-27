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
	height   = 600
	width    = 800
	saveName = "plot_test.png"
	writer   = flag.String("w", "file", "where to draw the plot\neither one of 'browser', 'file'")
)

func main() {
	// parse command line flags
	flag.Parse()

	// create a plot
	canvas = plot.NewCanvas(width, height, &color.RGBA{255, 255, 255, 255})
	canvas.Rectangle(5, 5, width-5, height-5, 5, &color.RGBA{0, 0, 0, 255})
	if err := canvas.AddAxis([]int{0}, 3, "horizontal"); err != nil {
		log.Fatalf("error adding axis: %v\n", err)
	}
	if err := canvas.AddAxis([]int{0}, 3, "vertical"); err != nil {
		log.Fatalf("error adding axis: %v\n", err)
	}
	canvas.AddLabel(canvas.Width/2, int(float64(canvas.Height)*0.2),
		"hello gopher!", &color.RGBA{50, 50, 50, 255})

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
