package main

import (
	"aoc/twentythree"
	"fmt"
	"os"
)

func readInput(inputFile string) (string, error) {
	inputText, err := os.ReadFile(inputFile)

	if err != nil {
		panic(fmt.Sprintln("failed to read input file {}", err))
	}

	return string(inputText), nil
}

func main() {
	fmt.Println("executing")

    inputText, err := readInput(`E:\.dev\golang\aoc\input\20231204_part1_prod.txt`)

    if err != nil {
        panic("cannot proceed without input")
    }

	day := new(twentythree.Day4)
	day.ParseData(inputText)
	result, _ := day.Part1()
    fmt.Println("part 1: ", result)

    result, _ = day.Part2()
    fmt.Println("part 2: ", result)

}
