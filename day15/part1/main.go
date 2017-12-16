package main

import "fmt"

var (
	startA = 699
	startB = 124
)

func generateA(prevValue uint64) uint64 {
	return (prevValue * 16807) % 2147483647
}

func generateB(prevValue uint64) uint64 {
	return (prevValue * 48271) % 2147483647
}

func compare(valA uint64, valB uint64) bool {
	sixteen := uint64(1<<16) - 1
	return (valA & sixteen) == (valB & sixteen)
}

func main() {
	valA := uint64(startA)
	valB := uint64(startB)

	count := 0
	for i := 0; i < 40*1000*1000; i++ {
		valA = generateA(valA)
		valB = generateB(valB)
		if compare(valA, valB) {
			count++
		}
	}
	fmt.Println("Result =", count)
}
