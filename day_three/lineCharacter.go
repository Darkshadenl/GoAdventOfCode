package main

import "strconv"

type lineCharacter struct {
	character             string
	characterOnLeft       *lineCharacter
	characterOnRight      *lineCharacter
	firstPart             bool
	HasNeighbourCharacter bool
	neighbourCharacter    string
	y                     int
	x                     int
}

func getNumbersAround(p *lineCharacter, line []lineCharacter, checkLeft bool) []lineCharacter {
	var numberCharacters []lineCharacter

	for i := 1; i <= 2; i++ {
		var index int
		if checkLeft {
			index = p.x - i
		} else {
			index = p.x + i
		}

		if index < 0 || index >= len(line) {
			break
		}

		char := line[index]
		if isLineCharacterANumber(char) {
			char.HasNeighbourCharacter = true
			numberCharacters = append(numberCharacters, char)
		}
	}

	return numberCharacters
}

func isLineCharacterANumber(p lineCharacter) bool {
	if _, err := strconv.Atoi(p.character); err == nil {
		return true
	}
	return false
}
