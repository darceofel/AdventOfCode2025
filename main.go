package main

import (
	"advent_of_code/solutions/day1"
	"advent_of_code/solutions/day2"
	"advent_of_code/solutions/day3"
	"advent_of_code/solutions/day4"
	"advent_of_code/solutions/day5"
	"advent_of_code/solutions/day6"
	"advent_of_code/solutions/day7"
	"advent_of_code/solutions/day8"
	"advent_of_code/solutions/day9"
	"fmt"
	"os"
)

var solutions = map[string]func(string) (string, string){
	"9": day9.Solution,
	"8": day8.Solution,
	"7": day7.Solution,
	"6": day6.Solution,
	"5": day5.Solution,
	"4": day4.Solution,
	"3": day3.Solution,
	"2": day2.Solution,
	"1": day1.Solution,
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please provide a day number as an argument.")
		return
	}

	day := args[0]
	solutionFunc, ok := solutions[day]
	if !ok {
		fmt.Println("Solution for day", day, "not found.")
		return
	}

	inputDay := day
	if len(args) >= 2 && args[1] == "test" {
		inputDay = "-1"
	}

	result1, result2 := solutionFunc(inputDay)

	fmt.Printf("Day %s results\n", day)
	fmt.Println("- Problem 1:", result1)
	fmt.Println("- Problem 2:", result2)
}
