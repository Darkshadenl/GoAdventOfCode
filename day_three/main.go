package main

import (
	"bufio"
	"io"
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
	data, err := os.Open("./day_three/small_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)
	lines := -1
	currentLine := 0

	for scanner.Scan() {
		lines++
	}

	data.Seek(0, 0)
	scanner = bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)

	q := Queue{}
	for scanner.Scan() {
		if currentLine >= 3 {
			log.Printf("currentLine: %d", currentLine)
			q.Dequeue()
		}
		linedata := scanner.Text()
		q.Enqueue(linedata)
		//data = append(data, linedata)
		currentLine++
	}
}

func doAssignment(r io.Reader) {

}
