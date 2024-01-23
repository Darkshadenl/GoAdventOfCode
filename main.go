package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var writtenNumbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}
var firstLetterWrittenNumbers = []string{"o", "t", "f", "s", "e", "n"}
var lastLetterWrittenNumbers = []string{"e", "o", "r", "x", "n", "t"}

var combinedWrittenNumbers = map[string]int{
	"twone":     21,
	"oneight":   18,
	"eightwo":   82,
	"eighthree": 83,
	"threeight": 38,
	"fiveight":  58,
	"sevenine":  79,
	"nineight":  98,
}

func main() {
	data, err := os.Open("calibrationcodes.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	answer := findCalibrationValues(data)
	fmt.Println(answer)

}

func findCalibrationValues(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	calibrationValues := make([]int, 0)

	for scanner.Scan() {
		code := scanner.Text()
		codeOnlyNrs := replaceTextualNrs(code)
		log.Println(code)
		log.Println(codeOnlyNrs)

		firstNumber, position := scanForNumberv2(codeOnlyNrs, 0)
		lastNumber, _ := scanForNumberv2(codeOnlyNrs, position)
		log.Printf("%d, %d", firstNumber, lastNumber)
		combinedStr := fmt.Sprintf("%d%d", firstNumber, lastNumber)

		if combined, err := strconv.Atoi(combinedStr); err == nil {
			log.Printf("added %d", combined)
			calibrationValues = append(
				calibrationValues,
				combined,
			)
		} else {
			//log.Fatal("failed to add new calibration value")
		}
	}

	number := 0
	log.Println(len(calibrationValues))
	for _, v := range calibrationValues {
		number += v
	}
	return number
}

func scanForNumberv2(code string, skipPositions int) (int, int) {
	sliced := code

	if skipPositions > 0 {
		sliced = reverseString(sliced)
	}

	for position := 0; position < len(sliced); position++ {
		isNumber, number := uint8ToInt(sliced[position])
		if isNumber {
			return number, position + 1
		}
	}
	return -1, -1
}

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func scanForNumber(code string, reverse bool, position int) (int, int) {
	if reverse {
		for i := len(code) - 1; i > -1; i-- {
			isNumber, s := uint8ToInt(code[i])
			if isNumber {
				return s, i
			}
		}
	} else {
		for i := 0; i < len(code); i++ {
			isNumber, s := uint8ToInt(code[i])
			if isNumber {
				return s, i
			}
		}
	}
	return -1, -1
}

func replaceTextualNrs(code string) string {
	returnVal := code
	codeLength := len(code)

	for word, combinedWrittenNumber := range combinedWrittenNumbers {
		if strings.Contains(code, word) {
			returnVal = strings.Replace(returnVal, word, strconv.Itoa(combinedWrittenNumber), -1)
		}
	}

	for i := 3; i < codeLength+1; i++ {
		partial := code[:i]

		for word := range writtenNumbers {
			if strings.Contains(partial, word) {
				returnVal = strings.Replace(returnVal, word, strconv.Itoa(writtenNumbers[word]), -1)
			}
		}
	}

	return returnVal
}

func removeNonMatchingItems(codes *[]string, letter string, position int) {
	for i, v := range *codes {
		l := string(v[position])
		if l == letter {
			continue
		}
		// remove item at current index from codes
		if len(*codes) == 1 {
			*codes = []string{}
		} else {
			*codes = append((*codes)[:i], (*codes)[i+1:]...)
		}
	}
}

func getWordsStartingWithLetter(letter string) []string {
	words := make([]string, 0)
	for k := range writtenNumbers {
		if strings.HasPrefix(k, letter) {
			words = append(words, k)
		}
	}
	return words
}

func arrayContainsLetter(arr []string, letter string) bool {
	for _, v := range arr {
		if strings.Contains(v, letter) {
			return true
		}
	}
	return false
}

func uint8ToInt(s uint8) (bool, int) {
	if i, err := strconv.Atoi(string(s)); err == nil {
		//log.Println("uint8 is number")
		return true, i
	}
	//log.Println("uint8 is not number")
	return false, -1
}

func stringToInt(s string) (bool, int) {
	if i, err := strconv.Atoi(s); err == nil {
		//log.Println("string is number")
		return true, i
	}
	//log.Println("string is not number")
	return false, -1
}
