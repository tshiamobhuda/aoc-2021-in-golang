package main

import (
	"flag"
	"fmt"
	"os"

	f "github.com/tshiamobhuda/aoc-2021-in-go/factory"
)

func main() {
	dy := flag.Uint("day", 0, "day of the solution to run")
	part := flag.Uint("part", 0, "part of the solution to run. Either 1 or 2")
	testOnly := flag.Bool("t", false, "only run test for the specified day")
	flag.Parse()

	if *dy < 1 {
		fmt.Println("Please specify which day's solution you'd like to run using the -day flag")
		os.Exit(2)
	}

	if *part < 1 || *part > 2 {
		fmt.Println("Please specify a valid part of the puzzle. Use either 1 or 2")
	}

	puzzleInput, err := getPuzzle(*dy, *testOnly)

	if err != nil {
		fmt.Printf("Could not find PuzzleInput for day %d part %d | Error:\n%v\n", *dy, *part, err.Error())
		os.Exit(2)
	}

	var factory f.AocFactory

	solution, _ := factory.GetAocDay(*dy)

	fmt.Printf("\nAdvent Of Code - Day: [%d] part [%d] \n\r", *dy, *part)

	if *part == 1 {
		fmt.Printf("Part %d: %v\n", *part, solution.Part1(puzzleInput))
	}

	if *part == 2 {
		fmt.Printf("Part %d: %v\n", *part, solution.Part2(puzzleInput))
	}
}

func getPuzzle(dy uint, test bool) (string, error) {
	var name string

	if test {
		name = fmt.Sprintf("day%d_test.txt", dy)
	} else {
		name = fmt.Sprintf("day%d.txt", dy)
	}

	data, err := os.ReadFile("./puzzle/" + name)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
