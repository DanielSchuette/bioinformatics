package dataframe

import (
	"fmt"
)

// Info pretty-prints the properties of an `Array'.
func (a *Array) Info() {
	s := fmt.Sprintf("Array Info:\nLength: %v\nType: %v\nData: %v",
		a.Len, a.Type, a.data)
	fmt.Printf("%s\n", s)
}

// Info pretty-prints the properties of a `DataFrame'.
func (d *Dataframe) Info() {
	// TODO: implement
	fmt.Println("not yet implemented")
}

// Sum returns the sum of all elements of an integer
// or float `Array' or an error if the `Array' is of
// type string.
func (a *Array) Sum() (float64, error) {
	var sum float64 /* holds the sum */

	if a.Type == "string" {
		return sum, fmt.Errorf("math: `Array' %v", errString)
	}

	if a.Type == "integer" {
		for _, val := range a.data.([]int64) {
			sum += float64(val)
		}
	}

	if a.Type == "float" {
		for _, val := range a.data.([]float64) {
			sum += val
		}
	}

	return sum, nil
}
