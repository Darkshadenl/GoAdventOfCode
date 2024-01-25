package main

import "bufio"

type LineIterator struct {
	scanner *bufio.Scanner
}

func NewLineIterator(scanner *bufio.Scanner) *LineIterator {
	return &LineIterator{scanner}
}

func (li *LineIterator) Next() (string, bool) {
	if li.scanner.Scan() {
		return li.scanner.Text(), true
	}
	return "", false
}
