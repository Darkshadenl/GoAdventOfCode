package main

var lineStackG = make(lineCharacterQueue, 0)

func getCharacter(Y, X int) (*lineCharacter, bool) {
	if Y >= 0 && Y < len(lineStackG) &&
		X >= 0 && X < len(lineStackG[Y]) {
		if !lineStackG[Y][X].hasFirstPart && !lineStackG[Y][X].firstPart {
			return &lineStackG[Y][X], true
		}
	}
	return nil, false
}

func getRightNeighbour(character *lineCharacter) (*lineCharacter, bool) {
	x := character.x + 1
	y := character.y
	return getCharacter(y, x)
}

func getLeftNeighbour(character *lineCharacter) (*lineCharacter, bool) {
	x := character.x - 1
	y := character.y
	return getCharacter(y, x)
}

func getTopNeighbour(character *lineCharacter) (*lineCharacter, bool) {
	x := character.x
	y := character.y - 1
	return getCharacter(y, x)
}

func getBottomNeighbour(character *lineCharacter) (*lineCharacter, bool) {
	x := character.x
	y := character.y + 1
	return getCharacter(y, x)
}

func getTopLeftNeighbour(character *lineCharacter) (*lineCharacter, bool) {
	x := character.x - 1
	y := character.y - 1
	return getCharacter(y, x)
}

func getTopRightNeighbour(character *lineCharacter) (*lineCharacter, bool) {
	x := character.x + 1
	y := character.y - 1
	return getCharacter(y, x)
}

func getBottomLeftNeighbour(character *lineCharacter) (*lineCharacter, bool) {
	x := character.x - 1
	y := character.y + 1
	return getCharacter(y, x)
}

func getBottomRightNeighbour(character *lineCharacter) (*lineCharacter, bool) {
	x := character.x + 1
	y := character.y + 1
	return getCharacter(y, x)
}

func getNeighbourCoordinates(lineQueue lineCharacterQueue, character *lineCharacter) []*lineCharacter {
	lineStackG = lineQueue
	coordinates := make([]*lineCharacter, 0)
	functions := []func(*lineCharacter) (*lineCharacter, bool){
		getTopLeftNeighbour,
		getLeftNeighbour,
		getBottomLeftNeighbour,
		getTopNeighbour,
		getBottomNeighbour,
		getTopRightNeighbour,
		getRightNeighbour,
		getBottomRightNeighbour,
	}
	for _, f := range functions {
		neighbour, ok := f(character)
		if ok {
			coordinates = append(coordinates, neighbour)
		}
	}

	return coordinates
}
