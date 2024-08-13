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
var depot Point = Point{0, 0}

func GetFileLines(path string) []string {
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

func ConvertStringsToLoads(lines []string) map[int]Load {
	pattern := regexp.MustCompile(`(\d+) \((-?[0-9]+.[0-9]+),(-?[0-9]+.[0-9]+)\) \((-?[0-9]+.[0-9]+),(-?[0-9]+.[0-9]+)\)`)
	loads := make(map[int]Load)
	for _, val := range lines {
		next := pattern.FindStringSubmatch(val)
		id, _ := strconv.ParseInt(next[1], 10, 32)
		startX, _ := strconv.ParseFloat(next[2], 64)
		startY, _ := strconv.ParseFloat(next[3], 64)
		endingX, _ := strconv.ParseFloat(next[4], 64)
		endingY, _ := strconv.ParseFloat(next[5], 64)
		loads[int(id)] = CreateLoad(int(id), startX, startY, endingX, endingY)
	}
	return loads
}
func GetDistanceToHome(prevDistance float64, current Point, next Load) float64 {
	return prevDistance + GetDist(current, next.Start) + next.Distance + GetDist(next.End, depot)
}

func IsValid(prevDistance float64, current Point, next Load) bool {
	return !next.Visited && GetDistanceToHome(prevDistance, current, next) < maxTime
}

func LoadsToString(drivers []Driver) []string {
	loadStrings := make([]string, 0)
	for _, val := range drivers {
		var str strings.Builder
		str.WriteString("[")
		for i := 0; i < len(val.Loads)-1; i++ {
			str.WriteString(fmt.Sprintf("%d, ", val.Loads[i]))
		}
		str.WriteString(fmt.Sprintf("%d", val.Loads[len(val.Loads)-1]))
		str.WriteString("]")
		loadStrings = append(loadStrings, str.String())
	}
	return loadStrings
}

func FindOptimalLoads(loads map[int]Load) []Driver {
	totalLoads := len(loads)
	var drivers []Driver
	curDriver := CreateDriver()
	loadsDelivered := 0
	for loadsDelivered < totalLoads {
		var origin Point
		if curDriver.MilesDriven == 0 {
			origin = depot
		} else {
			prevId := curDriver.Loads[len(curDriver.Loads)-1]
			origin = loads[prevId].End
		}
		minMiles := math.MaxFloat64
		minId := 0
		for index, val := range loads {
			if IsValid(curDriver.MilesDriven, origin, val) && GetDist(origin, val.Start) < minMiles { // get minimum distance to next stop
				minMiles = GetDist(origin, val.Start)
				minId = index
			}
		}
		if minMiles != math.MaxFloat64 { // add load to current driver
			loadsDelivered++
			curDriver.MilesDriven += (minMiles + loads[minId].Distance)
			l := loads[minId]
			l.Visited = true
			loads[minId] = l
			curDriver.Loads = append(curDriver.Loads, loads[minId].Id)
		} else {
			curDriver.MilesDriven += GetDist(origin, depot)
			drivers = append(drivers, curDriver)
			curDriver = CreateDriver()
		}
	}
	lastId := curDriver.Loads[len(curDriver.Loads)-1]
	curDriver.MilesDriven += GetDist(loads[lastId].End, depot)
	drivers = append(drivers, curDriver)
	return drivers
}
func PrintLoadStrings(loadStrings []string) {
	for _, loadStr := range loadStrings {
		fmt.Println(loadStr)
	}
}
func main() {
	path := os.Args[1:][0]
	lines := GetFileLines(path)
	loads := ConvertStringsToLoads(lines)
	drivers := FindOptimalLoads(loads)
	loadStrings := LoadsToString(drivers)
	PrintLoadStrings(loadStrings)
}
