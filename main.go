package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

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

func isNumber(s uint8) (bool, int) {
	i, err := strconv.Atoi(string(s))
	if err != nil {
		// is nr
		return false, i
	} else {
		return true, i
	}
}
