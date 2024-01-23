package main

import (
	"adventofcode/utils"
	"log"
	"os"
)

func main() {
	data, err := os.Open("./day_two/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	games := objectifyInput(data)
	determined := determinePossibleGames(games)
	log.Println(sumPossibles(determined))
	log.Println(sumPower(determined))
}

func sumPower(determined []Game) any {
	sum := 0
	for _, game := range determined {
		sum += game.power
	}
	return sum
}

func sumPossibles(games []Game) int {
	sum := 0
	for _, game := range games {
		if game.possible {
			_, n := utils.StringToInt(game.id)
			sum += n
		}
	}
	return sum
}

func determinePossibleGames(games []Game) []Game {
	for i := range games {
		ps := checkSubsets(games[i].subsets)
		games[i].power = determinePower(games[i].subsets)
		games[i].possible = ps
	}
	return games
}

func determinePower(subsets []Subset) int {
	power := 0
	minGreen := -1
	minRed := -1
	minBlue := -1
	for _, subset := range subsets {
		// find minimum number of each color
		if minGreen == -1 || minRed == -1 || minBlue == -1 {
			minGreen = subset.green
			minRed = subset.red
			minBlue = subset.blue
		} else {
			if subset.green > minGreen {
				minGreen = subset.green
			}
			if subset.red > minRed {
				minRed = subset.red
			}
			if subset.blue > minBlue {
				minBlue = subset.blue
			}
		}
	}
	// calculate power
	power = minGreen * minRed * minBlue
	return power
}

func checkSubsets(subsets []Subset) bool {
	truths := make([]bool, len(subsets))

	for i, subset := range subsets {
		if checkSubset(subset) {
			truths[i] = true
		}
	}
	for _, truth := range truths {
		if !truth {
			return false
		}
	}
	return true
}

func checkSubset(subset Subset) bool {
	red := 12
	green := 13
	blue := 14
	if subset.red > red || subset.green > green || subset.blue > blue {
		return false
	}

	return true
}
