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
var startWrittenNumbers = []string{"o", "t", "f", "s", "e", "n"}
var endWrittenNumbers = []string{"e", "o", "r", "x", "n", "t"}

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
		log.Println(code)

		firstNumber := scanForNumber(code, false)
		lastNumber := scanForNumber(code, true)
		combinedStr := fmt.Sprintf("%d%d", firstNumber, lastNumber)
		combined, err := strconv.Atoi(combinedStr)
		if err != nil {
			log.Fatal(err)
		}
		calibrationValues = append(
			calibrationValues,
			combined,
		)
	}

	number := 0
	for k, v := range calibrationValues {
		log.Printf("k: %d, v: %d\n", k, v)
		number += v
	}
	return number
}

func scanForNumber(code string, reverse bool) int {
	// use writtenNumbers. Replace all occurences of writtenNumbers with numbers

	for k, v := range writtenNumbers {
		code = strings.ReplaceAll(code, k, strconv.Itoa(v))
	}

	if reverse {
		for i := len(code) - 1; i > -1; i-- {
			isNumber, s := isNumber(code[i])
			if isNumber {
				return s
			}
		}
	} else {
		for i := 0; i < len(code); i++ {
			isNumber, s := isNumber(code[i])
			if isNumber {
				return s
			}
		}
	}
	return -1
}

//func isWrittenNumber(code string, reverse bool) bool {
//	for i := 0; i < len(code); i++ {
//		letter := code[i]
//		if reverse {
//			// check if letter is in endWrittenNumbers
//			for _, v := range endWrittenNumbers {
//				if v == string(letter) {
//					// take all the writtenNumbers that end with this letter
//					for _, w := range writtenNumbers {
//						wordLength := len(w)
//
//					}
//				}
//			}
//		} else {
//			// check if letter is in startWrittenNumbers
//			for _, v := range startWrittenNumbers {
//				if v == string(letter) {
//					// take all the writtenNumbers that start with this letter
//					// also take
//				}
//			}
//		}
//
//	}
//
//	return false
//}

func isNumber(s uint8) (bool, int) {
	i, err := strconv.Atoi(string(s))
	if err != nil {
		// is nr
		return false, i
	} else {
		return true, i
	}
}
