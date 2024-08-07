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

func convertStringsToLoads(lines []string) []load {
	pattern := regexp.MustCompile(`(\d+) \((-?[0-9]+.[0-9]+),(-?[0-9]+.[0-9]+)\) \((-?[0-9]+.[0-9]+),(-?[0-9]+.[0-9]+)\)`)
	loads := make([]load, len(lines)+1)
	for index, val := range lines {
		next := pattern.FindStringSubmatch(val)
		_ = next
		id, _ := strconv.ParseInt(next[1], 10, 32)
		startX, _ := strconv.ParseFloat(next[2], 64)
		startY, _ := strconv.ParseFloat(next[3], 64)
		endingX, _ := strconv.ParseFloat(next[4], 64)
		endingY, _ := strconv.ParseFloat(next[5], 64)
		loads[index+1] = createLoad(int(id), startX, startY, endingX, endingY)
	}
	return loads
}

func isValid(d driver, current point, next load, visited []bool) bool {
	if visited[next.id] { // has been visited before
		return false
	}
	newMiles := d.milesDriven + getDist(current, next.start) + next.distance + getDist(next.end, depot) // (new start - prior end) + new load distance + distance from end to home
	return newMiles < maxTime
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

func findOptimalLoads(loads []load) []driver {
	visited := make([]bool, len(loads))
	totalLoads := len(loads) - 1
	visited[0] = true
	var drivers []driver
	curDriver := createDriver()
	loadsDelivered := 0
	for loadsDelivered < totalLoads {
		var origin point
		if curDriver.milesDriven == 0 {
			origin = depot
		} else {
			x := curDriver.loads[len(curDriver.loads)-1]
			origin = loads[x].end
		}
		minMiles := math.MaxFloat64
		minLoad := 0
		for index, val := range loads {
			if isValid(curDriver, origin, val, visited) && getDist(origin, val.start)+val.distance < minMiles { // find shortest load that also allows driver to get home under constraint 
				minMiles = getDist(origin, val.start) + val.distance
				minLoad = index
			}
		}
		if minLoad != 0 { // add load to current driver
			loadsDelivered++
			curDriver.milesDriven += minMiles
			visited[minLoad] = true
			curDriver.loads = append(curDriver.loads, minLoad)
		} else { // 
			curDriver.milesDriven += getDist(origin, depot)
			drivers = append(drivers, curDriver)
			curDriver = createDriver()
		}
	}
	index := curDriver.loads[len(curDriver.loads)-1]
	curDriver.milesDriven += getDist(loads[index].end, depot)
	drivers = append(drivers, curDriver)
	return drivers
}

func main() {
	// calculateAverageCost()
	path := os.Args[1:][0]
	lines := getFileLines(path)
	_ = lines
	loads := convertStringsToLoads(lines)
	drivers := findOptimalLoads(loads)
	printLoads(drivers)
}
