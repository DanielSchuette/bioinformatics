# Alignment Plots Package
Alignment plots visualize the relationship between nucleotide or amino acid sequences.

## Example

The following example demonstrates how to print an alignment plot of two sequences to `stdout`.

```go
package main

import (
	"fmt"
	"log"

	"github.com/DanielSchuette/bioinformatics/alignmentplots"
)

func main() {
	// create an `Alignment' struct
	align := &alignmentplots.Alignment{
		SeqA: "ASDFJJKSKALSKSKDJSKASKD",
		SeqB: "ASDFJHHJKSKALSKSGKDJSKAHSKD",
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
```

## License

The code in this repository is MIT licensed.
