package twentythree_test

import (
	"aoc/twentythree"
	"testing"
)

func TestPart1(t *testing.T) {
	day := twentythree.Day1{}

	day.ParseData(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)
	result, error := day.Part1()

	if result != 142 || error != nil {
		t.Fatalf("Day1.Part1() = %d, %s for value %d", result, error, 142)
	}
}

func TestPart2(t *testing.T) {
	day := twentythree.Day1{}

	day.ParseData(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`)

	result, error := day.Part2()

	if result != 281 || error != nil {
	    t.Fatalf("Day1.Part2() = %d, %s for value %d", result, error, 281)
	}

}
