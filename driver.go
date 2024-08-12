package main

type driver struct {
	loads       []int
	milesDriven float64
	avgBetween	float64
}

func createDriver() driver {
	return driver{make([]int, 0), 0, 0}
}
