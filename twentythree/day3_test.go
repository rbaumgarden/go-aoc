package twentythree_test

import (
	"aoc/twentythree"
	"testing"
)

func TestDay3Part1(t *testing.T) {
	day := twentythree.Day3{}

	day.ParseData(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..` )
	result, error := day.Part1()

	if result != 4361 || error != nil {
		t.Fatalf("Day3.Part1() = %d, %s for value %d", result, error,  4361)
	}
}

func TestDay3Part2(t *testing.T) {
	day := twentythree.Day3{}

	day.ParseData(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`)

	result, error := day.Part2()

	if result != 467835 || error != nil {
	    t.Fatalf("Day3.Part2() = %d, %s for value %d", result, error, 467835)
	}
}
