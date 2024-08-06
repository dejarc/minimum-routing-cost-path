package main

import "math"

type load struct {
	id       int
	startX   float64
	startY   float64
	endX     float64
	endY     float64
	distance float64
}

func getDist(sX float64, sY float64, eX float64, eY float64) float64 {
	difX := sX - eX
	difY := sY - eY
	return math.Sqrt(difX*difX + difY*difY)
}
func createLoad(id int, sX float64, sY float64, eX float64, eY float64) load {
	return load{id, sX, sY, eX, eY, getDist(sX, sY, eX, eY)}
}
