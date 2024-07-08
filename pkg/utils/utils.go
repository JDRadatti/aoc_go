package utils

import (
	"log"
	"time"
)

// Least common divisor of elements in elems
func LCM(elems []int) int {
	for i := 1; i < len(elems); i++ {
		elems[i] = elems[i] * elems[i-1] / gcd(elems[i], elems[i-1])
	}
	return elems[len(elems)-1]
}

// Greatest common divisor of elements in elems
func GCD(elems []int) int {
	for i := 1; i < len(elems); i++ {
		// Euclids Algorithm
		elems[i] = gcd(elems[i], elems[i-1])
	}
	return elems[len(elems)-1]
}

func gcd(a, b int) int {
	if a <= b {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
    return a
}

// Timeit measures the duration of a function call
// call with defer TimeIt("func name")()
func TimeIt(name string) func() {
	start := time.Now()
	return func() {
		log.Printf("function %s executed in %v", name, time.Since(start))
	}
}

// Min of two integers
func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

// Max of two integers
func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

// BAtoI converts a Byte Array to Int
func BAtoI(bytes []byte) int {
	value := 0
	for i := range bytes {
		value = value*10 + int(bytes[i]) - '0'
	}
	return value
}
