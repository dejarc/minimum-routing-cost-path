package main

type Point struct {
	X float64
	Y float64
}
type Load struct {
	Id       int
	Start    Point
	End      Point
	Distance float64
	Visited  bool
}

func CreateLoad(id int, sX float64, sY float64, eX float64, eY float64) Load {
	start := Point{sX, sY}
	end := Point{eX, eY}
	return Load{id, start, end, GetDist(start, end), false}
}
