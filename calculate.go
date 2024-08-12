package main

import "math"

func getDist(start point, end point) float64 {
	difX := end.x - start.x
	difY := end.y - start.y
	return math.Sqrt(difX*difX + difY*difY)
}
func getUpdatedCost(newLoadId int, d driver, i int, loads map[int]load) float64 {
	cost := d.milesDriven
	prev := depot
	if i > 0 {
		prev = loads[d.loads[i-1]].end
	}
	next := depot
	if i+1 < len(d.loads) {
		next = loads[d.loads[i+1]].start
	}
	prevLoadId := d.loads[i]
	cost -= (getDist(prev, loads[prevLoadId].start) + loads[prevLoadId].distance + getDist(loads[prevLoadId].end, next))
	cost += (getDist(prev, loads[newLoadId].start) + loads[newLoadId].distance + getDist(loads[newLoadId].end, next))
	return cost
}
func getUpdatedAverage(newLoadId int, d driver, i int, loads map[int]load) float64 {
	loadLen := float64(len(d.loads) + 1)
	cost := d.avgBetween * loadLen
	prev := depot
	if i > 0 {
		prev = loads[d.loads[i-1]].end
	}
	next := depot
	if i+1 < len(d.loads) {
		next = loads[d.loads[i+1]].start
	}
	prevLoadId := d.loads[i]
	cost -= (getDist(prev, loads[prevLoadId].start) + getDist(loads[prevLoadId].end, next))
	cost += (getDist(prev, loads[newLoadId].start) + getDist(loads[newLoadId].end, next))
	return cost / loadLen
}

func swapLoadsMinAvgHelper(first *driver, second *driver, loads map[int]load) {
	avgDist := (first.avgBetween + second.avgBetween) / 2
	for i := 0; i < len(first.loads); i++ {
		for j := 0; j < len(second.loads); j++ {
			firstDist := getUpdatedAverage(second.loads[j], *first, i, loads)
			secondDist := getUpdatedAverage(first.loads[i], *second, j, loads)
			nextAvg := (firstDist + secondDist) / 2
			firstDif := (firstDist - first.avgBetween) * float64(len(first.loads)+1)
			secondDif := (secondDist - second.avgBetween) * float64(len(second.loads)+1)
			newTimeFirst := firstDif + first.milesDriven + (loads[second.loads[j]].distance - loads[first.loads[i]].distance)
			newTimeSecond := secondDif + second.milesDriven + (loads[first.loads[i]].distance - loads[second.loads[j]].distance)
			if nextAvg < avgDist && newTimeFirst < maxTime && newTimeSecond < maxTime {
				first.milesDriven = newTimeFirst
				second.milesDriven = newTimeSecond
				first.avgBetween = firstDist
				second.avgBetween = secondDist
				avgDist = nextAvg
				swap(first.loads, i, second.loads, j)
			}
		}
	}
}

func swapLoadsMinAvgBetween(drivers []driver, loads map[int]load) {
	for i := 0; i < len(drivers); i++ {
		for j := i + 1; j < len(drivers); j++ {
			swapLoadsMinAvgHelper(&drivers[i], &drivers[j], loads)
		}
	}
}

func swapLoadsHelper(first *driver, second *driver, loads map[int]load) {
	avgDist := (first.milesDriven + second.milesDriven) / 2
	for i := 0; i < len(first.loads); i++ {
		for j := 0; j < len(second.loads); j++ {
			firstDist := getUpdatedCost(second.loads[j], *first, i, loads)
			secondDist := getUpdatedCost(first.loads[i], *second, j, loads)
			nextAvg := (firstDist + secondDist) / 2
			if nextAvg < avgDist && firstDist < maxTime && secondDist < maxTime {
				swap(first.loads, i, second.loads, j)
				first.milesDriven = firstDist
				second.milesDriven = secondDist
				avgDist = nextAvg
			}
		}
	}
}

func swapLoadsMinAvgTotal(drivers []driver, loads map[int]load) {
	for i := 0; i < len(drivers); i++ {
		for j := i + 1; j < len(drivers); j++ {
			swapLoadsHelper(&drivers[i], &drivers[j], loads)
		}
	}
}
func calculateTotalCost(drivers []driver) float64 {
	var totalCost float64 = 0
	for _, val := range drivers {
		totalCost += val.milesDriven
	}
	return totalCost + float64(len(drivers)*500)
}
