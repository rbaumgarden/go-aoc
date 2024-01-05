package twentythree

import (
	"bufio"
	"strconv"
	"strings"
)

type Day2 struct {
	data string
}

var cubeRestrictions = map[string]int{"red": 12, "green": 13, "blue": 14}

func (self *Day2) ParseData(data string) {
	self.data = data
}

func getGameData(gameText string) (int, map[string]int) {
	colorMaxCounts := map[string]int{"red": 0, "green": 0, "blue": 0}
	separatorIndex := strings.Index(gameText, ":")
	gameId, _ := strconv.Atoi(gameText[5:separatorIndex])
	// skip the semicolon and following space.
	separatorIndex += 2

	tokenChars := make([]rune, 0, 5)

	for separatorIndex < len(gameText) {
		// get box quantity.
		for _, char := range gameText[separatorIndex:] {
			separatorIndex += 1

			if char == ' ' {
				break
			}
			tokenChars = append(tokenChars, char)
		}
		quantity, _ := strconv.Atoi(string(tokenChars))

		// clear it to parse the box color.
		tokenChars = tokenChars[:0]

	colorloop:
		for _, char := range gameText[separatorIndex:] {
			separatorIndex += 1
			switch char {
			case ',':
				separatorIndex += 1 // for the following space
				break colorloop
			case ';':
				separatorIndex += 1 // for the following space
				break colorloop
			case ' ':
				panic("i am not managing the line-parsing correctly")
			default:
				tokenChars = append(tokenChars, char)
			}
		}
		color := string(tokenChars)
		tokenChars = tokenChars[:0]

		if colorMaxCounts[color] < quantity {
			colorMaxCounts[color] = quantity
		}
	}

	return gameId, colorMaxCounts
}

func getIdOfPossibleGame(gameText string) (int, bool) {

	gameId, maxGameColors := getGameData(gameText)
	for key, value := range cubeRestrictions {
		if value < maxGameColors[key] {
			return gameId, false
		}
	}
	//if cubeRestrictions[color] < quantity {
	//	return gameId, false
	//}

	return gameId, true
}

func (self *Day2) Part1() (int64, error) {
	validSum := 0
	stringReader := strings.NewReader(self.data)
	scanner := bufio.NewScanner(stringReader)

	for scanner.Scan() {
		gameId, valid := getIdOfPossibleGame(scanner.Text())

		if valid {
			validSum += gameId
		}
	}

	return int64(validSum), nil
}

func (self *Day2) Part2() (int64, error) {
	stringReader := strings.NewReader(self.data)
	scanner := bufio.NewScanner(stringReader)
	sumPowers := 0

	for scanner.Scan() {
		_, maxGameColors := getGameData(scanner.Text())
		gamePower := 1
		for _, count := range maxGameColors {
			gamePower *= count
		}

		sumPowers += gamePower
	}
    
    return int64(sumPowers), nil
}
