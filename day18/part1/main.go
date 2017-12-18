package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var path = "input.txt"

type Program struct {
	registers    map[byte]int
	instructions [][]interface{}
	counter      int
	LastPlayed   int
}

func (p *Program) Get(input string) (result int) {
	if len(input) == 1 && input[0] >= 'a' && input[0] <= 'z' {
		result = p.registers[input[0]]
	} else {
		result, _ = strconv.Atoi(input)
	}
	return
}

/*
snd X plays a sound with a frequency equal to the value of X.
set X Y sets register X to the value of Y.
add X Y increases register X by the value of Y.
mul X Y sets register X to the result of multiplying the value contained in register X by the value of Y.
mod X Y sets register X to the remainder of dividing the value contained in register X by the value of Y (that is, it sets X to the result of X modulo Y).
rcv X recovers the frequency of the last sound played, but only when the value of X is not zero. (If it is zero, the command does nothing.)
jgz X Y jumps with an offset of the value of Y, but only if the value of X is greater than zero. (An offset of 2 skips the next instruction, an offset of -1 jumps to the previous instruction, and so on.)
*/
func (p *Program) Process() bool {
	if p.counter < 0 || p.counter >= len(p.instructions) {
		return false
	}

	instruction := p.instructions[p.counter]
	cmd := instruction[0].(string)
	x := instruction[1]
	y := instruction[2]

	switch {
	case cmd == "snd":
		reg := x.(byte)
		if val, ok := p.registers[reg]; ok && val > 0 {
			p.LastPlayed = val
		}
	case cmd == "rcv":
		reg := x.(byte)
		if val, ok := p.registers[reg]; ok && val > 0 {
			return false
		}
	case cmd == "set" || cmd == "add" || cmd == "mul" || cmd == "mod":
		reg := x.(byte)
		vy := p.Get(y.(string))
		switch cmd {
		case "set":
			p.registers[reg] = vy
		case "add":
			p.registers[reg] += vy
		case "mul":
			p.registers[reg] *= vy
		case "mod":
			p.registers[reg] %= vy
		}
	case cmd == "jgz":
		vx := p.Get(x.(string))
		vy := p.Get(y.(string))
		if vx > 0 {
			p.counter += vy
			return true
		}
	}
	p.counter++

	return true
}

func ParseInput(input string) []interface{} {
	result := make([]interface{}, 3)
	arr := strings.Split(input, " ")
	result[0] = arr[0]

	switch result[0] {
	case "snd":
		result[1] = arr[1][0]
	case "rcv":
		result[1] = arr[1][0]
	case "set":
		fallthrough
	case "add":
		fallthrough
	case "mul":
		fallthrough
	case "mod":
		result[1] = arr[1][0]
		result[2] = arr[2]
	case "jgz":
		result[1] = arr[1]
		result[2] = arr[2]
	}

	return result
}

func NewProgram(scanner *bufio.Scanner) Program {
	registers := make(map[byte]int)
	instructions := make([][]interface{}, 0)

	for scanner.Scan() {
		text := scanner.Text()
		instructions = append(instructions, ParseInput(text))
	}

	return Program{
		registers:    registers,
		instructions: instructions,
		counter:      0,
	}
}

func main() {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	program := NewProgram(scanner)

	for program.Process() {
	}

	fmt.Println("Result =", program.LastPlayed)
}
