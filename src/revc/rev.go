package main

import (
	"fmt"
	"goproj/point"
)

func main() {
	c := []int{1, 2, 3, 4, 5}
	rotateArray(c, 2)
	fmt.Println("--------------")
	fmt.Println(c)
	p := point.Point{1, 2}
	point.PointPrint(p)
}

func reversArray(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func rotateArray(s []int, n int) {
	reversArray(s[:n])
	reversArray(s[n:])
	reversArray(s[:])
}
