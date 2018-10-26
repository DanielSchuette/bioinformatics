package dataframe

import "fmt"

// Package-level constants that determine certain
// behaviors. Most constants are not exposed to
// the user.
const (
	Version = "0.0.1"
)

type typeError struct {
	t interface{}
}

func (e *typeError) Error() string {
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
	// (float64), and "str" (string)
	Type string

	// data holds the actual data of the
	// `Array' and is a slice
	data interface{}
}

// NewArray returns a new Array. If `data' is
// empty, the resulting array will be empty as
// well.
func NewArray(data interface{}) (*Array, error) {
	a := &Array{}

	// determine type of the input data, then
	// get the actual data and size for `Array'
	switch v := data.(type) {
	case []int64:
		a.Type = "integer"
		d, ok := data.([]int64)
		if !ok {
			return nil, &typeError{v}
		}
		a.data = d
		a.Len = len(d)
	case []float64:
		a.Type = "float"
		d, ok := data.([]float64)
		if !ok {
			return nil, &typeError{v}
		}
		a.data = d
		a.Len = len(d)
	case []string:
		a.Type = "str"
		d, ok := data.([]string)
		if !ok {
			return nil, &typeError{v}
		}
		a.data = d
		a.Len = len(d)
	default:
		return nil, &typeError{v}
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
