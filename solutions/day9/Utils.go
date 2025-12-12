package day9

import (
	"advent_of_code/utils"
)

var neighborOffsets = [8][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

type Point struct {
	X int
	Y int
}

type Segment struct {
	start Point
	end   Point
}

func rectangleArea(a, b Point) int {
	return (utils.Abs(a.X-b.X) + 1) * (utils.Abs(a.Y-b.Y) + 1)
}

func segmentsCross(a, b, c, d Point) bool {
	// Orientation of ordered triplet (p, q, r)
	orientation := func(p, q, r Point) int {
		val := (q.Y-p.Y)*(r.X-q.X) - (q.X-p.X)*(r.Y-q.Y)
		if val == 0 {
			return 0
		}
		if val > 0 {
			return 1
		}
		return 2
	}

	// Check if r is between p and q (inclusive)
	onSegment := func(p, q, r Point) bool {
		return r.X >= min(p.X, q.X) && r.X <= max(p.X, q.X) &&
			r.Y >= min(p.Y, q.Y) && r.Y <= max(p.Y, q.Y)
	}

	o1 := orientation(a, b, c)
	o2 := orientation(a, b, d)
	o3 := orientation(c, d, a)
	o4 := orientation(c, d, b)

	// Case 1: proper intersection
	if o1 != o2 && o3 != o4 {
		return true
	}

	// Case 2: touching at endpoints only (allowed)
	// c lies on AB and is exactly A or B
	if o1 == 0 && onSegment(a, b, c) && (c == a || c == b) {
		return true
	}

	// d lies on AB and is exactly A or B
	if o2 == 0 && onSegment(a, b, d) && (d == a || d == b) {
		return true
	}

	// a lies on CD and is exactly C or D
	if o3 == 0 && onSegment(c, d, a) && (a == c || a == d) {
		return true
	}

	// b lies on CD and is exactly C or D
	if o4 == 0 && onSegment(c, d, b) && (b == c || b == d) {
		return true
	}

	// All other cases: no intersection
	return false
}

func (s Segment) cuts(a, b Point) bool {
	maxX, minX := max(a.X, b.X), min(a.X, b.X)
	maxY, minY := max(a.Y, b.Y), min(a.Y, b.Y)

	topRight, topLeft := Point{X: maxX, Y: maxY}, Point{X: minX, Y: maxY}
	bottomRight, bottomLeft := Point{X: maxX, Y: minY}, Point{X: minX, Y: minY}

	return segmentsCross(s.start, s.end, topRight, topLeft) || segmentsCross(s.start, s.end, topRight, bottomRight) || segmentsCross(s.start, s.end, bottomLeft, topLeft) || segmentsCross(s.start, s.end, bottomRight, bottomLeft)
}
