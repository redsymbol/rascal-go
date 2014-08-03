package main

type Point struct {
	X, Y int
}

func PointsAdjacent(x1, y1, x2, y2 int) bool {
	diff_x := x1 - x2
	diff_y := y1 - y2
	if diff_x < 0 {
		diff_x *= -1
	}
	if diff_y < 0 {
		diff_y *= -1
	}
	return diff_x <= 1 && diff_y <= 1
}
