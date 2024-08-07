package main

func calculateTotalCost(drivers []driver) float64 {
	var totalCost float64 = 0
	for _, val := range drivers {
		totalCost += val.milesDriven
	}
	return totalCost + float64(len(drivers)*500)
}