package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Queue []string

func (q *Queue) Enqueue(value string) {
	*q = append(*q, value)
}

func (q *Queue) Dequeue() (string, bool) {
	if len(*q) == 0 {
		return "", false
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, true
}

func (q *Queue) GetElement(index int) (string, bool) {
	if index < 0 || index >= len(*q) {
		return "", false
	}
	return (*q)[index], true
}

func main() {
	data, err := os.Open("./day_three/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)
	lines := -1

	for scanner.Scan() {
		lines++
	}

	data.Seek(0, 0)
	scanner = bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)

	currentLine := 0
	q := Queue{}
	for scanner.Scan() {
		if currentLine >= 3 && currentLine < lines {
			log.Printf("currentLine: %d", currentLine)
			doAssignment(q, false)
			q.Dequeue()
		} else {
			doAssignment(q, true)
			q.Dequeue()
		}
		linedata := scanner.Text()
		q.Enqueue(linedata)
		currentLine++
	}
}

func doAssignment(lines Queue, lastBatch bool) {
	validSigns := map[string]bool{"/": true, "+": true, "@": true, "%": true, "=": true, "$": true, "#": true, "-": true, "*": true, "&": true}

	lineZero, _ := lines.GetElement(0)
	lineOne, _ := lines.GetElement(1)
	lineTwo, _ := lines.GetElement(2)

	runes := []rune(lineOne)

	for i, r := range runes {
		index := i
		character := string(r)
		if validSigns[character] {
			println("found valid sign")
			fmt.Println(index, character)
			// go through all the neighbours, and check if they contain a number.

			//getNeighbours(index, [lineZero, lineOne, lineTwo])

		}
	}

	//log.Println(lineZero)
	//log.Println(lineOne)
	//log.Println(lineTwo)
}
