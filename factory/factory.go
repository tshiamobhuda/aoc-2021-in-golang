package factory

import (
	"fmt"

	"github.com/tshiamobhuda/aoc-2021-in-go/day"
)

type Factory interface {
	GetAocDay(day uint) (AocDay, error)
}

type AocDay interface {
	Part1(input string) string
	Part2(input string) string
}

type AocFactory struct{}

func (a AocFactory) GetAocDay(d uint) (AocDay, error) {
	// @TODO Refactor this - maybe use reflection/type switches
	switch d {
	case 1:
		return &day.Day1{}, nil
	// case 2:
	// 	return &day.Day2{}, nil
	// case 3:
	// 	return &day.Day3{}, nil
	case 4:
		return &day.Day4{}, nil
	// case 5:
	// 	return &day.Day5{}, nil
	// case 6:
	// 	return &day.Day6{}, nil
	// case 7:
	// 	return &day.Day7{}, nil
	// case 8:
	// 	return &day.Day8{}, nil
	// case 9:
	// 	return &day.Day9{}, nil
	// case 10:
	// 	return &day.Day10{}, nil
	// case 11:
	// 	return &day.Day11{}, nil
	// case 12:
	// 	return &day.Day12{}, nil
	// case 13:
	// 	return &day.Day13{}, nil
	// case 14:
	// 	return &day.Day14{}, nil
	// case 15:
	// 	return &day.Day15{}, nil
	// case 16:
	// 	return &day.Day16{}, nil
	// case 17:
	// 	return &day.Day17{}, nil
	// case 18:
	// 	return &day.Day18{}, nil
	// case 19:
	// 	return &day.Day19{}, nil
	// case 20:
	// 	return &day.Day20{}, nil
	// case 21:
	// 	return &day.Day21{}, nil
	// case 22:
	// 	return &day.Day22{}, nil
	// case 23:
	// 	return &day.Day23{}, nil
	// case 24:
	// 	return &day.Day24{}, nil
	// case 25:
	// 	return &day.Day25{}, nil

	default:
		return nil, fmt.Errorf("invalid day provided")
	}
}
