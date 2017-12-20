package main

import (
	"bufio"
	"fmt"
	"os"
)

var path = "input.txt"

type Grid struct {
	startX     int
	maxX, maxY int
	grid       [][]byte
}

func NewGrid(scanner *bufio.Scanner) Grid {
	grid := make([][]byte, 0, 100)

	var i, j, maxX, maxY, startX int
	for scanner.Scan() {
		text := scanner.Text()
		maxX = len(text)
		if i == 0 {
			for j = 0; j < maxX; j++ {
				if text[j] == '|' {
					startX = j
					break
				}
			}
		}
		grid = append(grid, []byte(text))
		i++
	}
	maxY = i

	return Grid{
		startX,
		maxX,
		maxY,
		grid,
	}
}

func (g *Grid) Traverse() int {
	var x, y, count int
	letters := make([]byte, 0)
	directions := map[byte][]int{
		'u': []int{0, -1},
		'd': []int{0, 1},
		'l': []int{-1, 0},
		'r': []int{1, 0},
	}
	traversed := make(map[int]bool)

	// Move down first
	dir := byte('d')
	x = g.startX

	getChar := func(x int, y int) byte {
		i := x*g.maxX + y
		if _, present := traversed[i]; present {
			return 0
		}

		if x < 0 || x >= g.maxX {
			return 0
		}

		if y < 0 || y >= g.maxY {
			return 0
		}

		return g.grid[y][x]
	}

	switchDirection := func(dir byte) byte {
		if dir == 'u' || dir == 'd' {
			left := getChar(x-1, y)
			if left == 0 || left == ' ' {
				return 'r'
			}
			return 'l'
		}

		up := getChar(x, y-1)
		if up == 0 || up == ' ' {
			return 'd'
		}
		return 'u'
	}

	for {
		char := getChar(x, y)
		if char == ' ' {
			break
		}

		i := x*g.maxX + y
		traversed[i] = true

		if char != 0 {
			switch {
			case char == '+':
				dir = switchDirection(dir)
			case char >= 'A' && char <= 'Z':
				letters = append(letters, char)
			}
		}

		x += directions[dir][0]
		y += directions[dir][1]
		count++
	}

	return count
}

func main() {
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	grid := NewGrid(scanner)
	result := grid.Traverse()

	fmt.Println("Result =", result)
}
