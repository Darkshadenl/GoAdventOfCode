package main

import (
	"adventofcode/utils"
	"bufio"
	"io"
	"strings"
)

type Subset struct {
	red   int
	green int
	blue  int
	power int
}

type Game struct {
	id       string
	possible bool
	subsets  []Subset
	power    int
}

func objectifyInput(r io.Reader) []Game {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	games := make([]Game, 0)

	for scanner.Scan() {
		code := scanner.Text()

		if strings.Contains(code, ":") {
			game := Game{}
			game.possible = false
			splitString := strings.Split(code, ":")
			splitSubsets := strings.Split(splitString[1], ";")
			gameSplit := strings.Split(splitString[0], " ")
			gameId := gameSplit[1]
			game.id = gameId

			for _, subset := range splitSubsets {
				subsetje := Subset{}
				subsetSplit := strings.Split(subset, ",")
				for _, value := range subsetSplit {
					stripped := strings.TrimSpace(value)
					index := strings.Index(stripped, " ")
					color := stripped[index+1:]
					_, colorNr := utils.StringToInt(stripped[:index])

					switch color {
					case "red":
						subsetje.red = colorNr
					case "green":
						subsetje.green = colorNr
					case "blue":
						subsetje.blue = colorNr
					}
				}
				game.subsets = append(game.subsets, subsetje)
			}
			games = append(games, game)
		}
	}
	return games
}
