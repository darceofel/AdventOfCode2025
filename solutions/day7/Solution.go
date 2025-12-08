package day7

import (
	"advent_of_code/utils"
	"fmt"
	"strconv"
)

func solution1(lines []string) string {
	activeBeams, splitCount := make([]bool, len(lines[0])), 0
	for _, line := range lines {
		for ix, char := range line {
			if char == 'S' {
				activeBeams[ix] = true
			}

			if (char == '^') && activeBeams[ix] {
				activeBeams[ix] = false
				splitCount++
				if ix > 0 {
					activeBeams[ix-1] = true
				}
				if ix < len(lines[0])-1 {
					activeBeams[ix+1] = true
				}
			}
		}
	}

	return strconv.Itoa(splitCount)
}

func solution2(lines []string) string {
	activeBeams := make([]int, len(lines[0]))
	for _, line := range lines {
		for ix, char := range line {
			if char == 'S' {
				activeBeams[ix] = 1
			}

			if (char == '^') && (activeBeams[ix] > 0) {
				if ix > 0 {
					activeBeams[ix-1] += activeBeams[ix]
				}
				if ix < len(lines[0])-1 {
					activeBeams[ix+1] += activeBeams[ix]
				}
				activeBeams[ix] = 0
			}
		}
	}

	timelinesCount := 0
	for _, val := range activeBeams {
		fmt.Print(val)
		timelinesCount += val
	}
	return fmt.Sprintf("%d", timelinesCount)
}

func Solution(day string) (string, string) {
	lines := utils.GetInputPathLines(day)
	return solution1(lines), solution2(lines)
}
