package main

import "io/ioutil"
import "strings"

import "fmt"

var path = "input.txt"

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

func main() {
	file, _ := ioutil.ReadFile(path)
	input := strings.Trim(string(file), "\n")
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
	hashNums := make([]byte, 32)
	for i := 0; i < 16; i++ {
		num := list[i*16]
		for j := 1; j < 16; j++ {
			num ^= list[i*16+j]
		}
		hex := fmt.Sprintf("%02x", num)
		hashNums[2*i] = hex[0]
		hashNums[2*i+1] = hex[1]
	}

	fmt.Println("Result =", string(hashNums))
}
