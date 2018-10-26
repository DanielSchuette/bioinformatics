package dataframe

import (
	"errors"
	"fmt"
)

// Package-level constants that determine certain
// behaviors. Most constants are not exposed to
// the user.
const (
	Version = "0.0.1"
)

// an `errString' is returned whenever a float
// or integer but not string is expected by a
// method or function
var errString = errors.New("must be float or integer, not string")

// an `errType' is returned whenever a type
// is used in the wrong context (e.g. a type
// other than int64, float64, string to
// initialize an `Array')
type errType struct {
	t interface{}
}

func (e *errType) Error() string {
	return fmt.Sprintf("dataframe: don't know type %T", e.t)
}

// An Array holds 1-D data of a single type
// for either computations or to add them to
// a `Dataframe'. An `Array' has convenience
// methods for adding and deleting items and
// simple computations if the `Array' is of
// a numeric type.
type Array struct {
	// `Len' returns the `Array' length
	Len int

	// `Type' returns the `Array' type
	// at this point, an `Array' can be
	// of type "integer" (int64), "float"
	// (float64), and "string" (string)
	Type string

	// data holds the actual data of the
	// `Array' and is a slice
	data interface{}
}

// NewArray returns a new Array. `data' must be
// a slice of either int64, float64, or string,
// otherwise an error is returned to the caller.
func NewArray(data interface{}) (*Array, error) {
	a := &Array{}

	// determine type of the input data, then
	// get the actual data and size for `Array'
	switch v := data.(type) {
	case []int64:
		a.Type = "integer"
		d, ok := data.([]int64)
		if !ok {
			return nil, &errType{v}
		}
		a.data = d
		a.Len = len(d)
	case []float64:
		a.Type = "float"
		d, ok := data.([]float64)
		if !ok {
			return nil, &errType{v}
		}
		a.data = d
		a.Len = len(d)
	case []string:
		a.Type = "string"
		d, ok := data.([]string)
		if !ok {
			return nil, &errType{v}
		}
		a.data = d
		a.Len = len(d)
	default:
		return nil, &errType{v}
	}

	return a, nil
}

// A Dataframe holds 2-D data in a column-based fashion.
// A single column can only hold data of one type, but
// different columns can hold different data types.
// Every column must be an `Array' of the correct type
// and can be added or deleted.
// Columns are organized in a map and thus they are
// uniquely identified by their names.
type Dataframe struct {
	// Dataframe dimensions: (rows, columns)
	Dims [2]int

	// `columns' holds the actual data
	// every element must be of `Array' type
	columns map[string]Array
}

// NewDataframe returns a new `Dataframe'. A slice
// of `Array' is expected as an input, along with a
// slice of column names of the same length. Further,
// all elements in `arr' must be of the same length.
// Otherwise, an error is returned to the caller.
func NewDataframe(arr []*Array, names []string) (*Dataframe, error) {
	if len(arr) != len(names) {
		return nil, nil
		// TODO
	}
	return nil, nil
	// TODO
}
