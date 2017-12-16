package main

import (
	"fmt"
)

func reverse(list []int, start int, end int) {
	listLength := len(list)
	var length int
	if end >= start {
		length = end - start + 1
	} else {
		length = (listLength + end) - start + 1
	}
	for o := 0; o < int(length/2); o++ {
		i := (start + o) % listLength
		j := (length - 1 - o + start + listLength) % listLength
		temp := list[i]
		list[i] = list[j]
		list[j] = temp
	}
}

func countOnes(input string) []uint16 {
	inputLengths := make([]byte, len(input))
	listLength := 256

	for i, v := range input {
		inputLengths[i] = byte(v)
	}

	inputLengths = append(inputLengths, []byte{17, 31, 73, 47, 23}...)

	list := make([]int, listLength)
	for i := 0; i < listLength; i++ {
		list[i] = i
	}

	// Sparse hash
	end := 0
	curPos := 0
	skipSize := 0
	round := 64
	for i := 0; i < round; i++ {
		for _, length := range inputLengths {
			l := int(length)
			if l > listLength {
				continue
			}

			end = (curPos + l - 1) % listLength
			if length >= 1 {
				reverse(list, curPos, end)
			}

			curPos = (curPos + l + skipSize) % listLength
			skipSize++
		}
	}

	// Dense hash
	arr := make([]uint16, 32)
	for i := 0; i < 16; i++ {
		num := list[i*16]
		for j := 1; j < 16; j++ {
			num ^= list[i*16+j]
		}
		arr[2*i] = uint16(num / 16)
		arr[2*i+1] = uint16(num % 16)
	}

	return arr
}

func formatCoord(x int, y int) int {
	return 128*x + y
}

func getValue(grid [][]uint16, x int, y int) uint16 {
	innerX := uint16(x / 4)
	offsetX := 3 - uint(x%4)
	val := grid[y][innerX]

	v := (val >> offsetX) & 1
	return v
}

func getNeighbours(grid [][]uint16, x int, y int) []int {
	val := getValue(grid, x, y)

	n := make([]int, 0)
	if val == 0 {
		return n
	}

	if x > 0 {
		if v := getValue(grid, x-1, y); v == 1 {
			n = append(n, formatCoord(x-1, y))
		}
	}
	if x < 127 {
		if v := getValue(grid, x+1, y); v == 1 {
			n = append(n, formatCoord(x+1, y))
		}
	}

	if y > 0 {
		if v := getValue(grid, x, y-1); v == 1 {
			n = append(n, formatCoord(x, y-1))
		}
	}
	if y < 127 {
		if v := getValue(grid, x, y+1); v == 1 {
			n = append(n, formatCoord(x, y+1))
		}
	}

	return n
}

func traverse(grid [][]uint16, visited map[int]bool, start int) {
	queue := make([]int, 1)
	visited[start] = true
	queue[0] = start

	for {
		if len(queue) == 0 {
			break
		}

		coord := queue[0]
		queue = queue[1:]
		x := coord / 128
		y := coord % 128

		neighbours := getNeighbours(grid, x, y)
		for _, neighbour := range neighbours {
			if _, ok := visited[neighbour]; !ok {
				queue = append(queue, neighbour)
				visited[neighbour] = true
			}
		}

	}
}

var input = "hwlqcszp"

func main() {
	grid := make([][]uint16, 128)
	for i := 0; i < 128; i++ {
		str := fmt.Sprintf("%s-%d", input, i)
		arr := countOnes(str)
		grid[i] = arr
	}

	regions := 0
	visited := make(map[int]bool)
	for y := 0; y < 128; y++ {
		for x := 0; x < 128; x++ {
			coord := formatCoord(x, y)
			if _, ok := visited[coord]; !ok {
				if val := getValue(grid, x, y); val == 1 {
					traverse(grid, visited, coord)
					regions++
				}
			}
		}
	}
	fmt.Println("Result =", regions)
}
