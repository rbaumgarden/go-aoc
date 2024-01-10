
package twentythree_test

import (
	"aoc/twentythree"
	"testing"
)

func TestDay4Part1(t *testing.T) {
	day := twentythree.Day4{}

	day.ParseData(`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11` )
	result, error := day.Part1()

	if result != 13 || error != nil {
		t.Fatalf("Day4.Part1() = %d, %s for value %d", result, error,  13)
	}
}

func TestDay4Part2(t *testing.T) {
	day := twentythree.Day4{}

	day.ParseData(`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`)

	result, error := day.Part2()

	if result != 30 || error != nil {
	    t.Fatalf("Day4.Part2() = %d, %s for value %d", result, error, 30)
	}
}
