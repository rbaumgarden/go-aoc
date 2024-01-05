package twentythree

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Day1 struct {
	data string
}

var stringDigitMap = [10][]byte{[]byte("zero"), []byte("one"), []byte("two"), []byte("three"),
	[]byte("four"), []byte("five"), []byte("six"), []byte("seven"), []byte("eight"), []byte("nine"),
}

func parseNextDigit(chars []byte) (int, int) {
	count := 0
	for charIndex, char := range chars {
		count += 1
		if unicode.IsDigit(rune(char)) {
			return int(char - '0'), count
		}

		digit := -1
		for digitIndex, digitBytes := range stringDigitMap {
            if bytes.HasPrefix(chars[charIndex:], digitBytes) {
				digit = digitIndex
			}
		}
            
		if digit >= 0 {
			return digit, count
		}
	}

	return 0, 0
}

func (self *Day1) ParseData(data string) {
	self.data = data
}

func (self *Day1) Part1() (int64, error) {
	reader := bufio.NewReader(strings.NewReader(self.data))

	sum := 0
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		digits := make([]int, 0, 10)
		line := scanner.Text()
		lineBuf := line[0:]
		for digit, parsed := parseNextDigit([]byte(lineBuf)); parsed > 0; digit, parsed = parseNextDigit([]byte(lineBuf)) {
			lineBuf = lineBuf[parsed:]
			digits = append(digits, digit)
		}
		firstDigit := strconv.Itoa(digits[0])
		lastDigit := strconv.Itoa(digits[len(digits)-1])
		number, err := strconv.Atoi(firstDigit + lastDigit)

		if err != nil {
			panic("failed to convert a digit")
		}

		sum += number
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error: ", err)
	}

	return int64(sum), nil
}

func (self *Day1) Part2() (int64, error) {
	return self.Part1()
}
