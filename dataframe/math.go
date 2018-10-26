package dataframe

import "fmt"

// Info pretty-prints the properties of an `Array'
func (a *Array) Info() {
	s := fmt.Sprintf("Array Info:\nLength: %v\nType: %v\nData: %v",
		a.Len, a.Type, a.data)
	fmt.Printf("%s\n", s)
}

// Info pretty-prints the properties of a `DataFrame'
func (d *Dataframe) Info() {
	// TODO: implement
	fmt.Println("not yet implemented")
}
