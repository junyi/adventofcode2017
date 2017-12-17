package main

import "fmt"

var steps = 348

type Buffer struct {
	position int
	steps    int
	count    int
}

func NewBuffer(steps int) Buffer {
	return Buffer{
		position: 0,
		steps:    steps,
		count:    1,
	}
}

func (buffer Buffer) Position() int {
	return buffer.position
}

func (buffer *Buffer) SimpleMove() {
	buffer.position = (buffer.position+buffer.steps)%buffer.count + 1
	buffer.count++
}

func main() {
	buffer := NewBuffer(steps)
	lastNum := 0
	for i := 0; i < 50000000; i++ {
		buffer.SimpleMove()
		if buffer.Position() == 1 {
			lastNum = i + 1
		}
	}
	fmt.Println("Result =", lastNum)
}
