package main

import (
	"io/ioutil"
	"log"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading file", err)
	}

	count := 0
	// Remove the extra newline char at EOF
	input = input[:len(input)-1]
	length := len(input)
	var prevChar byte
	for i := 0; i <= length; i++ {
		char := input[(i % length)]
		if prevChar == char {
			count += int(prevChar) - '0'
		}
		prevChar = char
	}

	println("Result =", count)
}
