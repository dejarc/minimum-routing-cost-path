package main

import (
	"reflect"
	"testing"
)

func TestConvertStringsToLoads(t *testing.T) {
	cases := map[int]load{
		1: {1, point{-50.1, 80.0}, point{90.1, 12.2}, 155.7333618721435, false},
		2: {2, point{-24.5, -19.2}, point{98.5, 0}, 124.48951763100378, false},
		3: {3, point{0.3, 8.9}, point{40.9, 55}, 61.42939035998974, false},
		4: {4, point{5.3, -61.1}, point{77.8, -5.4}, 91.42614505708966, false},
	}
	sampleLines := []string{
		"1 (-50.1,80.0) (90.1,12.2)",
		"2 (-24.5,-19.2) (98.5,1,8)",
		"3 (0.3,8.9) (40.9,55.0)",
		"4 (5.3,-61.1) (77.8,-5.4)",
	}
	loads := convertStringsToLoads(sampleLines)
	for id, val := range loads {
		x := cases[id]
		if !reflect.DeepEqual(x, val) {
			t.Fatalf("\nexpected value %v\nreceived value %v", cases[id], val)
		}
	}
}

func TestGetDistanceToHome(t *testing.T) {
	prevDis := 0
	cur := point{x: 0, y: 0}
	nextLoad:= load{ 1,  point {x: -50.1, y: 80}, point {x: 90.1, y: 12.2},  155.7333618721435, false}
	expected := 341.0484306841512
	received := getDistanceToHome(float64(prevDis), cur, nextLoad)
	if received != expected {
		t.Fatalf("\nexpected distance %f received distance %v", expected, received)
	}
}

func TestIsValidWithVisitedLoad(t *testing.T) {
	testLoad := load{1, point{-50.1, 80.0}, point{90.1, 12.2}, 155.7333618721435, true}
	if isValid(0, point{0,0}, testLoad) {
		t.Fatalf("load has been previously visited %v", testLoad)
	}
}

func TestIsValidWithDistanceViolation(t *testing.T) {
	testLoad := load{1, point{-50.1, 80.0}, point{90.1, 12.2}, 155.7333618721435, false}
	prevDistance := 500.000
	totalDistance := 837.458181561743
	if isValid(prevDistance, point{-1.99,2.99}, testLoad) {
		t.Fatalf("load has a distance of %f greater than the maximum distance of %f", totalDistance, maxTime)
	}
}