# Package dataframe

[![GoDoc](https://godoc.org/github.com/DanielSchuette/bioinformatics/dataframe?status.svg)](https://godoc.org/github.com/DanielSchuette/bioinformatics/dataframe)

## Overview

This package implements a dataframe that is inspired by Python's `pandas.DataFrame`. A dataframe is a two-dimensional matrix-like data structure in which every column has a single data type and all columns are of equal length. In a dataframe, column names are (which must be unique) are used to identify columns that might be used for further computations.

The core structures `Array` (1-D) and `Dataframe` (2-D) along with the most basic methods are defined in `dataframe.go`. Computations that can be done on those two data structures are implemented in `math.go`.

## Contributions

Contributions are highly welcome!

## License

The code in this repository is MIT licensed.
