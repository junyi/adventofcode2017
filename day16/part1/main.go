package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var path = "input.txt"

type Program struct {
	bytes     []byte
	positions map[byte]int
	length    int
	offset    int
}

func newProgram(length int) Program {
	b := make([]byte, length)
	p := make(map[byte]int)
	for i := 0; i < length; i++ {
		b[i] = byte(i)
		p[b[i]] = i
	}
	return Program{
		bytes:     b,
		positions: p,
		length:    length,
		offset:    0,
	}
}

func (program *Program) Print() string {
	l := program.length
	tempBytes := make([]byte, l)
	for i := 0; i < l; i++ {
		tempBytes[i] = 'a' + program.bytes[(i-program.offset+l)%l]
	}
	return string(tempBytes)
}

func (program *Program) Swap(pos1 int, pos2 int) {
	temp := program.bytes[pos1]
	program.bytes[pos1] = program.bytes[pos2]
	program.bytes[pos2] = temp

	program.positions[program.bytes[pos1]] = pos1
	program.positions[program.bytes[pos2]] = pos2
}

func (program *Program) Move(input string) {
	switch input[0] {
	case 's':
		offset, _ := strconv.Atoi(input[1:])
		program.offset = (program.offset + offset) % program.length
	case 'x':
		l := program.length
		arr := strings.Split(input[1:], "/")
		pos1, _ := strconv.Atoi(arr[0])
		pos2, _ := strconv.Atoi(arr[1])
		epos1 := (pos1 - program.offset + l) % l
		epos2 := (pos2 - program.offset + l) % l
		program.Swap(epos1, epos2)
	case 'p':
		arr := strings.Split(input[1:], "/")
		pos1 := program.positions[arr[0][0]-'a']
		pos2 := program.positions[arr[1][0]-'a']
		program.Swap(pos1, pos2)
	}
}

func main() {
	b, _ := ioutil.ReadFile(path)
	str := string(b)
	arr := strings.Split(strings.Trim(str, "\n"), ",")

	program := newProgram(16)

	for _, move := range arr {
		program.Move(move)
	}

	fmt.Println("Result =", program.Print())
}
