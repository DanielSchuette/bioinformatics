package main

import (
	"github.com/DanielSchuette/bioinformatics/alignmentplots"
)

func main() {
	// create an `Alignment' struct
	align := &alignmentplots.Alignment{
		SeqA: "ASDFJJKSKALSKSKDJSKASKD",
		SeqB: "ASDFJHHJKSKALSKSGKDJSKAHSKD",
	}

	// align the two input sequences `SeqA' and `SeqB'
	align.Align()
	// fmt.Println(align)

	// plot the alignment
	align.Plot("My First Alignment Plot")
}
