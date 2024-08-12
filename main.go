package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var maxTime float64 = 720
var depot point = point{0, 0}

func getFileLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("error with file%v", err)
		os.Exit(1)
	}
	reader := bufio.NewReader(file)
	var nextLine []byte
	var lines []string
	reader.ReadLine() // eliminate header line
	for {
		nextLine, _, err = reader.ReadLine()
		if err != nil {
			break
		}
		lines = append(lines, string(nextLine))
	}
	return lines
}

func convertStringsToLoads(lines []string) map[int]load {
	pattern := regexp.MustCompile(`(\d+) \((-?[0-9]+.[0-9]+),(-?[0-9]+.[0-9]+)\) \((-?[0-9]+.[0-9]+),(-?[0-9]+.[0-9]+)\)`)
	loads := make(map[int]load)
	for _, val := range lines {
		next := pattern.FindStringSubmatch(val)
		id, _ := strconv.ParseInt(next[1], 10, 32)
		startX, _ := strconv.ParseFloat(next[2], 64)
		startY, _ := strconv.ParseFloat(next[3], 64)
		endingX, _ := strconv.ParseFloat(next[4], 64)
		endingY, _ := strconv.ParseFloat(next[5], 64)
		loads[int(id)] = createLoad(int(id), startX, startY, endingX, endingY)
	}
	return loads
}
func getDistanceToHome(prevDistance float64, current point, next load) float64 {
	return prevDistance + getDist(current, next.start) + next.distance + getDist(next.end, depot)
}

func isValid(prevDistance float64, current point, next load) bool {
	return !next.visited && getDistanceToHome(prevDistance, current, next) < maxTime
}

func printLoads(drivers []driver) {
	for _, val := range drivers {
		var str strings.Builder
		str.WriteString("[")
		for i := 0; i < len(val.loads)-1; i++ {
			str.WriteString(fmt.Sprintf("%d, ", val.loads[i]))
		}
		str.WriteString(fmt.Sprintf("%d", val.loads[len(val.loads)-1]))
		str.WriteString("]")
		fmt.Println(str.String())
	}
}

func findOptimalLoads(loads map[int]load) []driver {
	totalLoads := len(loads)
	var drivers []driver
	curDriver := createDriver()
	loadsDelivered := 0
	for loadsDelivered < totalLoads {
		var origin point
		if curDriver.milesDriven == 0 {
			origin = depot
		} else {
			prevId := curDriver.loads[len(curDriver.loads)-1]
			origin = loads[prevId].end
		}
		minMiles := math.MaxFloat64
		minId := 0
		for index, val := range loads {
			if isValid(curDriver.milesDriven, origin, val) && getDist(origin, val.start) < minMiles { // get minimum distance to next stop
				minMiles = getDist(origin, val.start)
				minId = index
			}
		}
		if minMiles != math.MaxFloat64 { // add load to current driver
			loadsDelivered++
			curDriver.milesDriven += (minMiles + loads[minId].distance)
			l := loads[minId]
			l.visited = true
			x := float64(len(curDriver.loads))
			curDriver.avgBetween = (curDriver.avgBetween*x + minMiles) / (x + 1)
			loads[minId] = l
			curDriver.loads = append(curDriver.loads, loads[minId].id)
		} else {
			distHome := getDist(origin, depot)
			curDriver.milesDriven += distHome
			x := float64(len(curDriver.loads))
			curDriver.avgBetween = (curDriver.avgBetween*x + distHome) / (x + 1)
			drivers = append(drivers, curDriver)
			swapLoadsMinAvgBetween(drivers, loads)
			curDriver = createDriver()
		}
	}
	lastId := curDriver.loads[len(curDriver.loads)-1]
	homeDist := getDist(loads[lastId].end, depot)
	curDriver.milesDriven += homeDist
	x := float64(len(curDriver.loads))
	curDriver.avgBetween = (curDriver.avgBetween*x + homeDist) / (x + 1)
	drivers = append(drivers, curDriver)
	swapLoadsMinAvgTotal(drivers, loads)
	return drivers
}
func swap(first []int, firstIndex int, second []int, secondIndex int) {
	temp := first[firstIndex]
	first[firstIndex] = second[secondIndex]
	second[secondIndex] = temp
}

func main() {
	path := os.Args[1:][0]
	lines := getFileLines(path)
	loads := convertStringsToLoads(lines)
	drivers := findOptimalLoads(loads)
	cost := calculateTotalCost(drivers)
	_ = cost
	printLoads(drivers)
}
