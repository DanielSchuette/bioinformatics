package alignmentplots

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

// package-level constants that determine certain
// behaviors of e.g. the plotting functionality
const (
	majorMatchIdentifier = "0"
	minorMatchIdentifier = "x"
	noMatchIdentifier    = "."
)

// Alignment holds two protein sequences and optionally
// their alignment matrix
type Alignment struct {
	// sequence A will be plotted in rows
	SeqA string

	// sequence B will be plotted in columns
	SeqB string

	// the alignment matrix indicates sequence similarities
	AlignmentMatrix [][]bool
}

// String ensures that `Alignment' implements the
// `Stringer' interface for pretty printing
func (a *Alignment) String() string {
	var matrix string
	if a.AlignmentMatrix != nil {
		for i, lenA := 0, len(a.SeqA); i < lenA; i++ {
			// add a new row to the `matrix' string
			matrix += fmt.Sprintf("%v", a.AlignmentMatrix[i][:])

			// do not add an end-of-line character after the last
			// row
			if i == (lenA - 1) {
				break
			}

			// append a new line character to the evolving string
			matrix += "\n"
		}
	} else {
		matrix = ""
	}
	return fmt.Sprintf("seq A: %s\nseq B: %s\nmatrix:\n%v",
		a.SeqA, a.SeqB, matrix)
}

// Align populates the `AlignmentMatrix' field of an
// `Alignment' struct
// It must be called before an alignment plot can be
// created
func (a *Alignment) Align() {
	// initialize an empty 2-D array with correct dims
	lenA := len(a.SeqA) /* number of rows */
	lenB := len(a.SeqB) /* number of columns */

	// create the first dimension of the matrix (rows)
	a.AlignmentMatrix = make([][]bool, lenA)

	// for every row, populate with a new slice
	for col := 0; col < lenA; col++ {
		arr := make([]bool, lenB)
		a.AlignmentMatrix[col] = arr
	}

	// i iterates over the rows of the alignment
	// matrix (i.e. sequence A)
	for i := 0; i < lenA; i++ {
		// j iterates over the columns of the alignment
		// matrix (i.e. sequence B)
		for j := 0; j < lenB; j++ {
			if a.SeqA[i] == a.SeqB[j] {
				a.AlignmentMatrix[i][j] = true
				continue
			}
			a.AlignmentMatrix[i][j] = false
		}
	}
}

// Plot creates a dot plot that visualizes the alignment
// of two protein sequences
func (a *Alignment) Plot(title string) error {
	// check input validity
	if (len(a.AlignmentMatrix) == 0) ||
		(len(a.AlignmentMatrix[0]) == 0) {
		return errors.New("alignmentplots error: compute an alignment matrix first")
	}
	// print a header for the plot
	fmt.Printf("alignment plot, v0.0.1\n%s\n", title)

	// print column labels (i.e. sequence B)
	fmt.Printf("  ")
	for _, val := range a.SeqB {
		fmt.Printf("%s ", string(val))
	}
	fmt.Println()

	// iterate over rows and columns and print them to
	// stdout (row labels are equal to sequence A)
	for i, lenA := 0, len(a.SeqA); i < lenA; i++ {
		for j, lenB := 0, len(a.SeqB); j < lenB; j++ {
			// for the first element in every row,
			// also print the row label
			if j == 0 {
				fmt.Printf("%s ", string(a.SeqA[i]))
			}

			// if a certain element is a match,
			// decide what to do
			if a.AlignmentMatrix[i][j] {
				// determine whether the match is on a main
				// diagonal while considering edge cases
				// print appropriate ascii characters
				switch {
				case (i == 0) && (j == 0):
					printMajor()
				case (i == 0) || (j == 0):
					if a.AlignmentMatrix[i+1][j+1] {
						printMajor()
					} else {
						printMinor()
					}
				case (i == (lenA - 1)) && (j == (lenB - 1)):
					printMajor()
				case (i == (lenA - 1)) || (j == (lenB - 1)):
					if a.AlignmentMatrix[i-1][j-1] {
						printMajor()
					} else {
						printMinor()
					}
				case (a.AlignmentMatrix[i-1][j-1]) ||
					(a.AlignmentMatrix[i+1][j+1]):
					printMajor()
				default:
					printMinor()
				}
			} else {
				fmt.Printf("%v ", noMatchIdentifier)
			}
		}
		fmt.Println()
	}
	return nil
}

func printMajor() {
	color.Set(color.BgRed, color.Bold)
	fmt.Printf("%v", majorMatchIdentifier)
	color.Unset()
	fmt.Printf(" ")
}

func printMinor() {
	color.Set(color.BgBlue)
	fmt.Printf("%v", minorMatchIdentifier)
	color.Unset()
	fmt.Printf(" ")
}