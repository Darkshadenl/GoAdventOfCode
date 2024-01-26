package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./day_three/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	lines := -1

	for scanner.Scan() {
		lines++
	}

	file.Seek(0, 0)

	scanner = bufio.NewScanner(file)
	iterator := NewLineIterator(scanner)
	doAssignment(iterator)
}

func identifyLineCharacters(line string, y int) []lineCharacter {
	characters := make([]lineCharacter, 0)
	for x, r := range line {
		character := string(r)
		if character == "\n" {
			continue
		}
		lc := lineCharacter{
			character: character,
			y:         y,
			x:         x,
		}
		characters = append(characters, lc)
	}
	return characters
}

var iteration = 0

func doAssignment(iterator *LineIterator) {
	queueQ := lineCharacterQueue{}
	saveList := make([]string, 0)

	for {
		line, ok := iterator.Next()
		if !ok {
			break
		}
		if iteration > 2 {
			queueQ.EnqueueList(identifyLineCharacters(line, 2))
		} else {
			queueQ.EnqueueList(identifyLineCharacters(line, iteration))
		}
		if iteration > 1 {
			// identify the numbers with characters as a neighbour and link them
			res := identificationOfPartNumbers(queueQ)
			// put all the valid numbers in a list of row 0.
			saveList = append(saveList, getPartNumbers(res)...)
			// once done, dequeue the first line
			_, _ = queueQ.Dequeue()
			// fix Y coordinates of all remaining lines
			queueQ = fixYCoordinates(queueQ)
		}
		iteration++
	}

	// print the list of numbers
	for i := range saveList {
		log.Printf("%s ", saveList[i])
	}

	log.Println()
}

func fixYCoordinates(lineQueue lineCharacterQueue) lineCharacterQueue {
	for i := 0; i < len(lineQueue); i++ {
		for j := 0; j < len(lineQueue[i]); j++ {
			lineQueue[i][j].y = lineQueue[i][j].y - 1
		}
	}
	return lineQueue
}

func getPartNumbers(lineQueue lineCharacterQueue) []string {
	list := make([]string, 0)
	needed, ok := lineQueue.GetElement(0)
	if ok {
		for i := 0; i < len(needed); i++ {
			if !needed[i].firstPart {
				continue
			}
			part := needed[i]
			first := part.character
			second := ""
			third := ""
			if part.characterOnRight != nil {
				second = part.characterOnRight.character
				if part.characterOnRight.characterOnRight != nil {
					third = part.characterOnRight.characterOnRight.character
				}
			}
			final := first + second + third
			list = append(list, final)
		}
	}
	return list
}

func identificationOfPartNumbers(lineQueue lineCharacterQueue) lineCharacterQueue {
	validSigns := map[string]bool{"/": true, "+": true, "@": true, "%": true, "=": true, "$": true, "#": true, "-": true, "*": true, "&": true}

	for _, l := range lineQueue[1] {
		if validSigns[l.character] {
			characters := getNeighbourCoordinates(lineQueue, &l)

			for i := range characters {
				c := characters[i]
				if _, err := strconv.Atoi((*c).character); err == nil {
					c.HasNeighbourCharacter = true
					c.neighbourCharacter = l.character
					numbersOnLeft := getNumbersAround(c, lineQueue[c.y], true)
					numbersOnLeft = reverse(numbersOnLeft)
					numbersOnLeft = append(numbersOnLeft, *c)
					//numbersOnLeft = linkCharacters(numbersOnLeft, true)

					//if len(numbersOnLeft) == 0 {
					//	c.firstPart = true
					//} else {
					//	if c.characterOnLeft == nil {
					//		c.characterOnLeft = &numbersOnLeft[len(numbersOnLeft)-1]
					//	}
					//}

					numbersOnRight := getNumbersAround(c, lineQueue[c.y], false)
					numbersOnRight = linkCharacters(numbersOnRight, false)

					if len(numbersOnRight) > 0 {
						if c.characterOnRight == nil {
							c.characterOnRight = &numbersOnRight[0]
						}
						if numbersOnRight[0].characterOnLeft == nil {
							numbersOnRight[0].characterOnLeft = c
						}
					}
				}
			}
		}
	}
	return lineQueue
}

func linkCharacters(characters []lineCharacter, firstOnLeft bool) []lineCharacter {
	if len(characters) == 0 {
		return characters
	}
	if firstOnLeft {
		characters[0].firstPart = true
	} else {
		characters[len(characters)-1].firstPart = true
	}

	for i := range characters {
		if i > 0 {
			if characters[i].characterOnLeft == nil {
				characters[i].characterOnLeft = &characters[i-1]
			}
		}
		if i < len(characters)-1 {
			if characters[i].characterOnRight == nil {
				characters[i].characterOnRight = &characters[i+1]
			}
		}
	}

	return characters
}
