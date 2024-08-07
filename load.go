package main

import "math"

type point struct {
	x float64
	y float64
}
type load struct {
	id       int
	start    point
	end      point
	distance float64
}

func getDist(start point, end point) float64 {
	difX := end.x - start.x
	difY := end.y - start.y
	return math.Sqrt(difX*difX + difY*difY)
}

func createLoad(id int, sX float64, sY float64, eX float64, eY float64) load {
	l := load{id, point{sX, sY}, point{eX, eY}, 0}
	l.distance = getDist(l.start, l.end)
	return l
}
