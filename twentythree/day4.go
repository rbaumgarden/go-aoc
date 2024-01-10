package twentythree

import (
	"bufio"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Day4 struct {
	data string
}

func (self *Day4) ParseData(s string) {
	self.data = s
}

func parseNumber(s string) (number int, remaining string) {
	// skip any spaces used for formatting
	for s[0] == ' ' {
		s = s[1:]
	}

	spaceIndex := strings.IndexAny(s, " ")
	// last item won't have a following space to terminate
	if spaceIndex == -1 {
		spaceIndex = len(s)
	}
	number, parseErr := strconv.Atoi(s[0:spaceIndex])

	if parseErr != nil {
		panic("failed to parse int:\n\n" + parseErr.Error())
	}

	if spaceIndex < len(s) {
		spaceIndex += 1
	}
	return number, s[spaceIndex:]
}

func processCard(s string, numbers chan<- int) (winning_numbers []int) {
	winning_numbers = make([]int, 10)

	//skip the trash.
	colonIndex := strings.IndexRune(s, ':')
	if colonIndex < 0 {
		panic("failed to find header separator")
	}
	s = s[colonIndex+2:] // also sip following space character

	// parse winning numbers.
	for s[0] != '|' {
		number, remaining := parseNumber(s)
		s = remaining

		winning_numbers = append(winning_numbers, number)
	}
	s = s[2:]

	// spawn goroutine to parse card numbers.
	go func() {
		for len(s) > 0 {
			number, remaining := parseNumber(s)
			s = remaining

			numbers <- number
		}

		close(numbers)
	}()

	return winning_numbers
}

func (self *Day4) Part1() (int64, error) {
	stringReader := strings.NewReader(self.data)
	lineScanner := bufio.NewScanner(stringReader)

	sumScores := 0
	for lineScanned := lineScanner.Scan(); lineScanned == true; lineScanned = lineScanner.Scan() {
		line := lineScanner.Text()

		fmt.Println(line)

		lottoNumCh := make(chan int)
		winningNums := processCard(line, lottoNumCh)
		slices.Sort(winningNums)

		winCount := 0
		for lottoNum := range lottoNumCh {
			if slices.Contains(winningNums, lottoNum) {
				if _, found := slices.BinarySearch(winningNums, lottoNum); found {
					winCount += 1
				}
			}
		}

		score := int(math.Pow(2, float64(winCount-1)))
		sumScores += score
		fmt.Println("wincount: ", winCount, "score: ", score, "sumScores: ", sumScores)
	}

	return int64(sumScores), nil
}

func (self *Day4) Part2() (int64, error) {
	cardCounts := make([]int, 0, 10)
	stringReader := strings.NewReader(self.data)
	lineScanner := bufio.NewScanner(stringReader)

	sumScores := 0
	for lineScanned := lineScanner.Scan(); lineScanned == true; lineScanned = lineScanner.Scan() {
		line := lineScanner.Text()

		if len(cardCounts) == 0 {
			cardCounts = append(cardCounts, 1)
		}

		fmt.Println(line)

		lottoNumCh := make(chan int)
		winningNums := processCard(line, lottoNumCh)
		slices.Sort(winningNums)

		winCount := 0
		for lottoNum := range lottoNumCh {
			if slices.Contains(winningNums, lottoNum) {
				if _, found := slices.BinarySearch(winningNums, lottoNum); found {
					winCount += 1
					if len(cardCounts) < winCount+1 {
						cardCounts = append(cardCounts, 1)
					}
					// we win for each copy of the current card, so add that sum.
					cardCounts[winCount] += cardCounts[0]
				}
			}
		}

		// multiply the number of wins by the number of copies we have of this card.
		winCount = cardCounts[0]
		sumScores += winCount
		fmt.Println("current card count: ", cardCounts[0], "wincount: ", winCount, "sumScores: ", sumScores)

		// discard the first element; we are dont processing this card.
		cardCounts = cardCounts[1:]
	}

	return int64(sumScores), nil
}
