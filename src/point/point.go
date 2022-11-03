package point

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func PointPrint(p Point) {
	fmt.Printf("this is the Point with the points of %d, %d ", p.X, p.Y)
}
