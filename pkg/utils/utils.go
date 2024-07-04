package utils

import (
    "time"
    "log"
)

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
