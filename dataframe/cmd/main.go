package main

import (
	"fmt"
	"log"

	"github.com/DanielSchuette/bioinformatics/dataframe"
)

func main() {
	// create example data
	dataInt64 := []int64{1, 2, 3, 4}
	dataFloat64 := []float64{1.0, 2.0, 3.0, 4.0}
	dataString := []string{"a", "b", "c"}
	dataInvalid := []int{1, 2, 3, 4}

	// create new `Arrays' and do some computations
	arr1, err := dataframe.NewArray(dataInt64)
	if err != nil {
		log.Fatalf("error creating dataframe: %v\n", err)
	}
	arr1.Info()
	sum1, err := arr1.Sum()
	if err != nil {
		log.Fatalf("error computing sum: %v\n", err)
	}
	fmt.Printf("sum: %f\n", sum1)

	arr2, err := dataframe.NewArray(dataFloat64)
	if err != nil {
		log.Fatalf("error creating dataframe: %v\n", err)
	}
	arr2.Info()
	sum2, err := arr2.Sum()
	if err != nil {
		log.Fatalf("error computing sum: %v\n", err)
	}
	fmt.Printf("sum: %f\n", sum2)

	arr3, err := dataframe.NewArray(dataString)
	if err != nil {
		log.Fatalf("error creating dataframe: %v\n", err)
	}
	arr3.Info()
	sum3, err := arr3.Sum()
	if err != nil {
		log.Fatalf("error computing sum: %v\n", err)
	}
	fmt.Printf("sum: %f\n", sum3)

	arr4, err := dataframe.NewArray(dataInvalid)
	if err != nil {
		log.Fatalf("error creating dataframe: %v\n", err)
	}
	arr4.Info()
}
