package main

import "math"

func GetDist(start Point, end Point) float64 {
	difX := end.X - start.X
	difY := end.Y - start.Y
	return math.Sqrt(difX*difX + difY*difY)
}

func CalculateTotalCost(drivers []Driver) float64 {
	var totalCost float64 = 0
	for _, val := range drivers {
		totalCost += val.MilesDriven
	}
	return totalCost + float64(len(drivers)*500)
}
