package main

type point struct {
	x float64
	y float64
}
type load struct {
	id       int
	start    point
	end      point
	distance float64
	visited  bool
}

func createLoad(id int, sX float64, sY float64, eX float64, eY float64) load {
	start := point{sX, sY}
	end := point{eX, eY}
	return load{id, start, end, getDist(start, end), false}
}
