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
		//log.Println(code)
		codeOnlyNrs := replaceTextualNrs(code)
		log.Printf("code: %s, codeonlynrs: %s\n", code, codeOnlyNrs)

		firstNumber := scanForNumber(codeOnlyNrs, false)
		lastNumber := scanForNumber(codeOnlyNrs, true)
		combinedStr := fmt.Sprintf("%d%d", firstNumber, lastNumber)

		if combined, err := strconv.Atoi(combinedStr); err == nil {
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

func scanForNumber(code string, reverse bool) int {
	if reverse {
		for i := len(code) - 1; i > -1; i-- {
			isNumber, s := uint8ToInt(code[i])
			if isNumber {
				return s
			}
		}
	} else {
		for i := 0; i < len(code); i++ {
			isNumber, s := uint8ToInt(code[i])
			if isNumber {
				return s
			}
		}
	}
	return -1
}

func replaceTextualNrs(code string) string {
	returnVal := code
	skip := 0

	for i := 0; i < len(code); i++ {
		if skip > 0 {
			skip--
			continue
		}
		letter := string(code[i])
		// check if letter is in firstLetterWrittenNumbers

		foundStartLetter := arrayContainsLetter(firstLetterWrittenNumbers, letter)
		if !foundStartLetter {
			continue
		}

		// if found start letter, retrieve words that start with this letter from writtenNumbers
		matchingWords := getWordsStartingWithLetter(letter)
		if len(matchingWords) == 0 {
			continue
		}

		nextLetterIndex := i
		checkIndex := 0
		for len(matchingWords) > 1 {
			// now check next letter. Narrow down the list of words. If no words remain, continue
			nextLetterIndex++
			checkIndex++
			lenCode := len(code)
			if nextLetterIndex < lenCode-1 {
				nextLetter := string(code[nextLetterIndex])
				removeNonMatchingItems(&matchingWords, nextLetter, checkIndex)
			} else {
				break
			}
		}
		if len(matchingWords) == 1 {
			// found a match. Replace the word with the number
			mWord := matchingWords[0]
			mWordIndex := 0
			breaked := false

			for j := i; j < i+len(mWord)-1; j++ {
				if j < len(code) {
					nextLetter := string(code[j])
					mWordLetter := string(mWord[mWordIndex])
					mWordIndex++
					if nextLetter != mWordLetter {
						breaked = true
						break
					}
				} else {
					// Handle the situation when j is out of range
					breaked = true
					break
				}
			}

			if !breaked {
				matchingValue := strconv.Itoa(writtenNumbers[matchingWords[0]])
				replacedString := strings.Replace(returnVal, mWord, matchingValue, 1)
				if replacedString != returnVal {
					returnVal = replacedString
					skip = len(mWord) - 1
				}

			}
		} else {
			// no match found. Continue
			//log.Println("no match found")
			continue
		}
	}
	//log.Println("returnval", returnVal)
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
