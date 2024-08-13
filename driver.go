package main

type Driver struct {
	Loads       []int
	MilesDriven float64
}

func CreateDriver() Driver {
	return Driver{make([]int, 0), 0}
}
