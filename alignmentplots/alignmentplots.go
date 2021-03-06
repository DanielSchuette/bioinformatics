package alignmentplots

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

// Package-level constants that determine certain
// behaviors of e.g. the plotting functionality.
const (
	// returns the current version of `alignmentplots'
	Version = "0.0.1"

	// constants that are used internally
	// they are not exported
	majorMatchIdentifier = "0"
	minorMatchIdentifier = "x"
	noMatchIdentifier    = "."
	rowLabelDelimiter    = "||"
	colLabelDelimiter    = "="
)

// Alignment holds two protein sequences and optionally
// their alignment matrix.
type Alignment struct {
	// sequence A will be plotted in rows
	SeqA string

	// sequence B will be plotted in columns
	SeqB string

	// the alignment matrix indicates sequence similarities
	AlignmentMatrix [][]bool
}

// String ensures that `Alignment' implements the
// `Stringer' interface for pretty printing.
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
// `Alignment' struct.
// It must be called before an alignment plot can be
// created.
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
// of two protein sequences.
func (a *Alignment) Plot(title string) error {
	// check input validity
	if (len(a.AlignmentMatrix) == 0) ||
		(len(a.AlignmentMatrix[0]) == 0) {
		return errors.New("alignmentplots error: compute an alignment matrix first")
	}
	// print a header for the plot
	fmt.Printf("alignment plot, v%s\n%s\n", Version, title)

	// print column labels (i.e. sequence B)
	fmt.Printf("      ") /* some padding */
	for _, val := range a.SeqB {
		fmt.Printf("%s ", string(val))
	}
	fmt.Println() /* new line */

	// print a delimiter between column labels and data
	fmt.Printf("   \\\\") /* some padding */
	for i, l := 0, len(a.SeqB); i < l; i++ {
		fmt.Printf("%s%s", colLabelDelimiter, colLabelDelimiter)
	}
	fmt.Println() /* new line */

	// iterate over rows and columns and print them to
	// stdout (row labels are equal to sequence A)
	for i, lenA := 0, len(a.SeqA); i < lenA; i++ {
		for j, lenB := 0, len(a.SeqB); j < lenB; j++ {
			// for the first element in every row,
			// also print the row label and a delimiter
			if j == 0 {
				fmt.Printf(" %s %s ", string(a.SeqA[i]), rowLabelDelimiter)
			}

			// if a certain element is a match,
			// decide what to do
			if a.AlignmentMatrix[i][j] {
				// determine whether the match is on a main
				// diagonal while considering edge cases
				// print appropriate ascii characters
				switch {
				case (i == 0) && (j == 0):
					printWithDelimiter(majorMatchIdentifier, color.BgRed)
				case (i == 0) || (j == 0):
					if a.AlignmentMatrix[i+1][j+1] {
						printWithDelimiter(majorMatchIdentifier, color.BgRed)
					} else {
						printWithDelimiter(minorMatchIdentifier, color.BgBlue)
					}
				case (i == (lenA - 1)) && (j == (lenB - 1)):
					printWithDelimiter(majorMatchIdentifier, color.BgRed)
				case (i == (lenA - 1)) || (j == (lenB - 1)):
					if a.AlignmentMatrix[i-1][j-1] {
						printWithDelimiter(majorMatchIdentifier, color.BgRed)
					} else {
						printWithDelimiter(minorMatchIdentifier, color.BgBlue)
					}
				case (a.AlignmentMatrix[i-1][j-1]) ||
					(a.AlignmentMatrix[i+1][j+1]):
					printWithDelimiter(majorMatchIdentifier, color.BgRed)
				default:
					printWithDelimiter(minorMatchIdentifier, color.BgBlue)
				}
			} else {
				fmt.Printf("%v ", noMatchIdentifier)
			}
		}
		fmt.Println()
	}
	return nil
}

// a helper function for colorful printing of certain delimiters
func printWithDelimiter(delim string, col color.Attribute) {
	color.Set(col, color.Bold)
	fmt.Printf("%v", delim)
	color.Unset()
	fmt.Printf(" ")
}
