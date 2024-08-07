package main

import "math"

func getDist(start point, end point) float64 {
	difX := end.x - start.x
	difY := end.y - start.y
	return math.Sqrt(difX*difX + difY*difY)
}

func calculateTotalCost(drivers []driver) float64 {
	var totalCost float64 = 0
	for _, val := range drivers {
		totalCost += val.milesDriven
	}
	return totalCost + float64(len(drivers)*500)
}
