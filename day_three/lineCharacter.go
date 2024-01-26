package main

import "strconv"

type lineCharacter struct {
	character             string
	characterOnLeft       *lineCharacter
	characterOnRight      *lineCharacter
	firstPart             bool
	hasFirstPart          bool
	HasNeighbourCharacter bool
	neighbourCharacter    string
	y                     int
	x                     int
}

func getNumbersAround(p *lineCharacter, line []lineCharacter, checkLeft bool) []*lineCharacter {
	var numberCharacters []*lineCharacter
	validSigns := map[string]bool{"/": true, "+": true, "@": true, "%": true, "=": true, "$": true, "#": true, "-": true, "*": true, "&": true, ".": true}

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

		char := &line[index] // Wijzig dit naar een pointer
		if validSigns[char.character] {
			break
		}
		if isLineCharacterANumber(*char) {
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

func reverse(slice []lineCharacter) []lineCharacter {
	if len(slice) == 1 {
		return slice
	}
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func reversePointers(slice []*lineCharacter) []*lineCharacter {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
