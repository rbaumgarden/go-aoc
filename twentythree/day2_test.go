package twentythree_test

import (
	"aoc/twentythree"
	"testing"
)

func TestDay2Part1(t *testing.T) {
	day := twentythree.Day2{}

	day.ParseData(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`)
	result, error := day.Part1()

	if result != 8 || error != nil {
		t.Fatalf("Day2.Part1() = %d, %s for value %d", result, error, 8)
	}
}

func TestDay2Part2(t *testing.T) {
	day := twentythree.Day2{}

	day.ParseData(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`)

	result, error := day.Part2()

	if result != 2286 || error != nil {
	    t.Fatalf("Day2.Part2() = %d, %s for value %d", result, error, 2286)
	}

}
