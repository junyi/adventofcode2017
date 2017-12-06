package main

import (
	"io/ioutil"
	"log"
)

func main() {
	input, err := ioutil.ReadFile("test_p2.txt")
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	count := 0
	// Remove the extra newline char at EOF
	input = input[:len(input)-1]
	length := len(input)
	offset := length / 2
	for i := offset; i < length+offset; i++ {
		char := input[(i % length)]
		prevChar := input[(i-offset+length)%length]
		if prevChar == char {
			count += int(prevChar) - '0'
		}
		prevChar = char
	}

	println("Result =", count)
}
