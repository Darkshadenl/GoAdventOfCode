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
		if currentLine >= 3 {
			log.Printf("currentLine: %d", currentLine)
			doAssignment(q)
			q.Dequeue()
		}
		linedata := scanner.Text()
		q.Enqueue(linedata)
		currentLine++
	}
}

func doAssignment(lines Queue) {
	//validSigns := []string{"/", "+", "@", "%", "=", "$", "#", "-", "*", "&"}

	//lineOne, _ := lines.GetElement(0)
	//lineTwo, _ := lines.GetElement(1)
	lineThree, _ := lines.GetElement(2)
	runes := []rune(lineThree)

	for i, r := range runes {
		index := i
		character := string(r)
		fmt.Println(index, character)
	}

	//log.Println(lineOne)
	//log.Println(lineTwo)
	//log.Println(lineThree)
}
