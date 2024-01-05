package twentythree

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Day3 struct {
	data string
}

type number struct {
	start     int
	numString string
}

type symbol struct {
	position int
	char     byte
	numbers  []int
}

func (self *Day3) ParseData(data string) {
	self.data = data
}

func (self *Day3) processData() ([]*number, []*symbol) {
	rdr := strings.NewReader(self.data)
	scanner := bufio.NewScanner(rdr)

	summables := make([]*number, 0)
	symbols := make([]*symbol, 0)

	var lastNumbers, currentNumbers []*number
	var lastSymbolPositions, currentSymbolPositions []*symbol

	for scanner.Scan() {
		line := scanner.Text()
		charIndex := 0

		lastNumbers = currentNumbers
		lastSymbolPositions = currentSymbolPositions
		currentNumbers = make([]*number, 0, 10)
		currentSymbolPositions = make([]*symbol, 0, 10)

		for currentIndex := charIndex; currentIndex < len(line); currentIndex = charIndex {
			char := line[charIndex]
			if char == '.' {
				charIndex += 1
				continue
			}

			// parse the entire number if we have a digit.
			//fmt.Println("line substring: ", line[charIndex:])
			for charIndex < len(line) && unicode.IsDigit([]rune(line[charIndex:])[0]) {
				charIndex += 1
			}

			// we found digits. store it and check if we are next to a symbol.
			if charIndex > currentIndex {
				num := number{start: currentIndex, numString: line[currentIndex:charIndex]}
				currentNumbers = append(currentNumbers, &num)

				numberNumber, numberError := strconv.Atoi(num.numString)
				if numberError != nil {
					panic(fmt.Sprint("failed to convert number to string", numberError))
				}

				// check if the number is adjacent to a symbol so far.
				if len(currentSymbolPositions) > 0 && currentSymbolPositions[len(currentSymbolPositions)-1].position == currentIndex-1 {
					adjacentSymbol := currentSymbolPositions[len(currentSymbolPositions)-1]
					summables = append(summables, &num)

					adjacentSymbol.numbers = append(adjacentSymbol.numbers, numberNumber)
				}

				for symbolPosIndex := 0; symbolPosIndex < len(lastSymbolPositions) && lastSymbolPositions[symbolPosIndex].position < charIndex+1; symbolPosIndex += 1 {
					if lastSymbolPositions[symbolPosIndex].position >= currentIndex-1 && lastSymbolPositions[symbolPosIndex].position <= charIndex {
						adjacentSymbol := lastSymbolPositions[symbolPosIndex]
						summables = append(summables, &num)

						adjacentSymbol.numbers = append(adjacentSymbol.numbers, numberNumber)
					}
				}

				// continue so we don't evaluate the symbol logic.
				continue
			}

			currentSymbol := symbol{position: charIndex, char: char}
			currentSymbolPositions = append(currentSymbolPositions, &currentSymbol)
			symbols = append(symbols, &currentSymbol)
			charIndex += 1

			// check if the symbol is next to any numbers from the last row.
			if len(currentNumbers) > 0 {
				lastNumber := currentNumbers[len(currentNumbers)-1]
				if lastNumber.start-1 <= currentIndex && currentIndex <= lastNumber.start+len(lastNumber.numString) {
					summables = append(summables, currentNumbers[len(currentNumbers)-1])

					lastNumberNumber, lastNumberError := strconv.Atoi(lastNumber.numString)
					if lastNumberError != nil {
						panic("we failed to convert the last number to an int")
					}

					currentSymbol.numbers = append(currentSymbol.numbers, lastNumberNumber)
				}
			}

			for numberIndex := 0; numberIndex < len(lastNumbers) &&
				lastNumbers[numberIndex].start+len(lastNumbers[numberIndex].numString)+1 > numberIndex; numberIndex += 1 {
				if lastNumbers[numberIndex].start-1 <= currentIndex && lastNumbers[numberIndex].start+len(lastNumbers[numberIndex].numString) >= currentIndex {
					summables = append(summables, lastNumbers[numberIndex])

					lastNumberNumber, lastNumberError := strconv.Atoi(lastNumbers[numberIndex].numString)
					if lastNumberError != nil {
						panic("failed to convert a previous line number to an int")
					}

					currentSymbol.numbers = append(currentSymbol.numbers, lastNumberNumber)
				}
			}
		}
	}

	return summables, symbols
}

func (self *Day3) Part1() (int64, error) {
	summables, _ := self.processData()

	sum := 0
	for _, num := range summables {
		intValue, err := strconv.Atoi(num.numString)

		if err != nil {
			panic("error converting string to int")
		}

		sum += intValue
	}

	return int64(sum), nil
}

func (self *Day3) Part2() (int64, error) {
	_, symbols := self.processData()

	sum := 0
	for _, symbol := range symbols {
		symbolScalar := 1

		if symbol.char != '*' {
			continue
		}

		if len(symbol.numbers) < 2 {
			continue
		}

		for _, number := range symbol.numbers {
			symbolScalar *= number
		}

		sum += symbolScalar
	}

	return int64(sum), nil
}
