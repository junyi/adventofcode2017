package main

import "io/ioutil"
import "strings"
import "strconv"
import "fmt"

var path = "input.txt"
var listLength = 256

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
	input := strings.Split(string(file), ",")
	inputLengths := make([]int, len(input))

	for i, v := range input {
		inputLengths[i], _ = strconv.Atoi(strings.Trim(v, "\n"))
	}

	list := make([]int, listLength)
	for i := 0; i < listLength; i++ {
		list[i] = i
	}

	end := 0
	curPos := 0
	skipSize := 0
	for _, length := range inputLengths {
		if length > listLength {
			continue
		}

		end = (curPos + length - 1) % listLength
		if length >= 1 {
			reverse(list, curPos, end)
		}

		curPos = (curPos + length + skipSize) % listLength
		skipSize++
	}

	fmt.Println("Result =", list[0]*list[1])
}
