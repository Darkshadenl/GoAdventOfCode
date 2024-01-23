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

func uint8ToInt(s uint8) (bool, int) {
	if i, err := strconv.Atoi(string(s)); err == nil {
		//log.Println("uint8 is number")
		return true, i
	}
	//log.Println("uint8 is not number")
	return false, -1
}
