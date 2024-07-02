package utils

// Min of two integers
func Min(x int, y int) int {
    if x < y {
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
