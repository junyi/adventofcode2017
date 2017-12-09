package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const path = "input.txt"

func main() {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	re := regexp.MustCompile(`([a-z]+) (inc|dec) (-?\d+) if ([a-z]+) ((?:<|>|=|!)=?) (-?\d+)`)

	registers := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)

		reg := matches[1]
		op := matches[2]
		val, _ := strconv.Atoi(matches[3])
		cReg := matches[4]
		cOp := matches[5]
		cVal, _ := strconv.Atoi(matches[6])

		cRegVal := registers[cReg]
		cond := false
		switch cOp {
		case ">":
			cond = cRegVal > cVal
		case ">=":
			cond = cRegVal >= cVal
		case "<":
			cond = cRegVal < cVal
		case "<=":
			cond = cRegVal <= cVal
		case "==":
			cond = cRegVal == cVal
		case "!=":
			cond = cRegVal != cVal
		}

		if cond {
			switch op {
			case "inc":
				registers[reg] += val
			case "dec":
				registers[reg] -= val
			}
		}
	}

	maxVal := 0
	for _, v := range registers {
		if v > maxVal {
			maxVal = v
		}
	}

	fmt.Println("Result =", maxVal)
}
