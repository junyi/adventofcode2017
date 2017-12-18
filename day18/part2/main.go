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
	name         string
	registers    map[byte]int
	instructions [][]interface{}
	Blocked      bool
	buffer       []int
	counter      int
	sendCounter  int
	peer         *Program
}

func (p *Program) Send(val int) {
	p.sendCounter++
	p.peer.buffer = append(p.peer.buffer, val)
}

func (p *Program) Receive() (int, error) {
	p.Blocked = true
	if len(p.buffer) > 0 {
		p.Blocked = false
		val := p.buffer[0]
		p.buffer = p.buffer[1:]
		return val, nil
	}
	return 0, fmt.Errorf("")
}

func (p *Program) SetPeer(program *Program) {
	p.peer = program
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
		val := p.Get(x.(string))
		p.Send(val)
	case cmd == "rcv":
		reg := x.(byte)
		val, err := p.Receive()
		if err != nil {
			return true
		}
		p.registers[reg] = val
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
		result[1] = arr[1]
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

func (p *Program) Clone(name string, id int) Program {
	registers := make(map[byte]int)
	registers['p'] = id
	instructions := make([][]interface{}, len(p.instructions))
	for i, v := range p.instructions {
		instructions[i] = make([]interface{}, 3)
		copy(instructions[i], v)
	}
	return Program{
		name:         name,
		registers:    registers,
		instructions: instructions,
		counter:      0,
		buffer:       make([]int, 0),
	}
}

func NewProgram(name string, id int, scanner *bufio.Scanner) Program {
	registers := make(map[byte]int)
	registers['p'] = id
	instructions := make([][]interface{}, 0)

	for scanner.Scan() {
		text := scanner.Text()
		instructions = append(instructions, ParseInput(text))
	}

	return Program{
		name:         name,
		registers:    registers,
		instructions: instructions,
		counter:      0,
		buffer:       make([]int, 0),
	}
}

func main() {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	p0 := NewProgram("p0", 0, scanner)
	p1 := p0.Clone("p1", 1)

	p0.SetPeer(&p1)
	p1.SetPeer(&p0)

	for {
		v1 := p0.Process()
		v2 := p1.Process()
		if !v1 || !v2 {
			break
		}
		if p0.Blocked && p1.Blocked {
			break
		}
	}

	fmt.Println("Result =", p1.sendCounter)
}
