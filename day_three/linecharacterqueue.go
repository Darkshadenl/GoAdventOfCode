package main

type lineCharacterQueue [][]lineCharacter

func (q *lineCharacterQueue) EnqueueList(values []lineCharacter) {
	*q = append(*q, values)
}

func (q *lineCharacterQueue) Dequeue() ([]lineCharacter, bool) {
	if len(*q) == 0 {
		return []lineCharacter{}, false
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, true
}

func (q *lineCharacterQueue) GetElement(index int) ([]lineCharacter, bool) {
	if index < 0 || index >= len(*q) {
		return []lineCharacter{}, false
	}
	return (*q)[index], true
}
