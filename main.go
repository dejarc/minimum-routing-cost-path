package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parseFileFromPath(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("error with file%v", err)
		os.Exit(1)
	}
	reader := bufio.NewReader(file)
	var nextLine []byte
	var lines []string
	for err == nil {
		nextLine, _, err = reader.ReadLine()
		lines = append(lines, string(nextLine))
	}
	return lines[1 : len(lines)-1]
}
func parseLinesToLoads(lines []string) []load {
	pattern := regexp.MustCompile(`(\d+) \((-?[0-9]+.[0-9]+),(-?[0-9]+.[0-9]+)\) \((-?[0-9]+.[0-9]+),(-?[0-9]+.[0-9]+)\)`)
	loads := make([]load, len(lines))
	for index, val := range lines {
		next := pattern.FindStringSubmatch(val)
		_ = next
		x, _ := strconv.ParseInt(next[1], 10, 32)
		x2, _ := strconv.ParseFloat(next[2], 64)
		x4, _ := strconv.ParseFloat(next[3], 64)
		x6, _ := strconv.ParseFloat(next[4], 64)
		x8, _ := strconv.ParseFloat(next[5], 64)
		loads[index] = createLoad(int(x), x2, x4, x6, x8)
	}
	return loads
}
func main() {

	path := os.Args[1:][0]
	lines := parseFileFromPath(path)
	_ = lines
	loads := parseLinesToLoads(lines)
	_ = loads
	fmt.Printf(path)
}
