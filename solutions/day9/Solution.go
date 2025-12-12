package day9

import (
	"advent_of_code/utils"
	"sort"
	"strconv"
	"strings"
)

func solution1(original []Point) string {
	points := make([]Point, len(original))
	copy(points, original)

	sort.Slice(points, func(i, j int) bool {
		point1, point2 := points[i], points[j]
		if point1.X == point2.X {
			return point1.Y < point2.Y
		}

		return point1.X < point2.Y
	})

	maxSize := -1
	for i, point1 := range points {
		for j := i + 1; j < len(points)-1; j++ {
			point2 := points[j]
			maxSize = max(maxSize, rectangleArea(point1, point2))
		}
	}

	return strconv.Itoa(maxSize)
}

func solution2(points []Point) string {
	maxArea := -1
	for ix, a := range points {
		for jx := ix + 1; jx < len(points); jx++ {
			b, valid := points[jx], true
			for kx := 0; kx < len(points); kx++ {
				if (kx == ix) || (kx == jx) || kx-ix == -1 || kx-jx == -1 {
					continue
				}

				start, end := points[kx], points[0]
				if kx < len(points)-1 {
					end = points[kx+1]
				}

				s := Segment{start: start, end: end}
				if s.cuts(a, b) {
					valid = false
					break
				}
			}

			if valid {
				maxArea = max(maxArea, rectangleArea(a, b))
			}
		}
	}

	return strconv.Itoa(maxArea)
}

func Solution(day string) (string, string) {
	points := utils.GetInputPathLinesFormatted(day, func(s string) Point {
		coords := strings.Split(s, ",")
		X, _ := strconv.Atoi(coords[0])
		Y, _ := strconv.Atoi(coords[1])

		return Point{
			X: X,
			Y: Y,
		}
	})

	return solution1(points), solution2(points)
}
