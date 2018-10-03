package main

import (
	"fmt"
	"log"

	"github.com/DanielSchuette/bioinformatics/alignmentplots"
)

func main() {
	// create an `Alignment' struct
	align := &alignmentplots.Alignment{
		SeqA: "MEEPQSDPSVEPPLSQET",   /* human TP53 AA 1-18 */
		SeqB: "EPQSDPSVEPPLSQETFSDL", /* human TP53 AA 3-22 */
	}

	// plot the resulting struct
	fmt.Println(align)

	// align the two input sequences `SeqA' and `SeqB'
	align.Align()

	// plot the alignment
	err := align.Plot("My First Alignment Plot")
	if err != nil {
		log.Fatalf("error while plotting: %v\n", err)
	}
}
