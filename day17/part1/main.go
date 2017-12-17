package main

import "fmt"

var steps = 348

type Buffer struct {
	position int
	bytes    []int
	steps    int
	count    int
}

func NewBuffer(steps int) Buffer {
	return Buffer{
		position: 0,
		bytes:    []int{0},
		steps:    steps,
		count:    0,
	}
}

func (buffer Buffer) Print() int {
	return buffer.bytes[buffer.position+1]
}
func (buffer *Buffer) Move() {
	position := (buffer.position + buffer.steps) % len(buffer.bytes)
	newPosition := position + 1
	buffer.bytes = append(buffer.bytes, 0)
	copy(buffer.bytes[position+2:], buffer.bytes[position+1:])
	buffer.bytes[newPosition] = buffer.count + 1
	buffer.count++
	buffer.position = newPosition
}

func main() {
	buffer := NewBuffer(steps)
	for i := 0; i < 2017; i++ {
		buffer.Move()
	}
	fmt.Println("Result =", buffer.Print())
}
