package day01

// SolutionA sums (first * 10 + last) of each line in input
// where first is the first digit and last is the last digit.
// Example:
// input = "1jvsl2/nk3four5/n"
// SolutionA(input) = 12 + 35 = 47
func SolutionA(input []byte) int {
	values := 0
	first, last := -1, -1

	for i := 0; i < len(input); i++ {
		if input[i] >= 48 && input[i] <= 57 {
			if first == -1 {
				first = int(input[i]) - 48
			} else {
				last = int(input[i]) - 48
			}
		} else if input[i] == 10 {
			if last == -1 {
				last = first
			}
			values += first*10 + last
			first, last = -1, -1
		}
	}
	return values
}

// SolutionB is the same as SolutionA with the caveat that
// spelled numbers, like "nine" also count as digits
// Example:
// input = "1jvsl2/nk3four/n"
// SolutionA(input) = 12 + 34 = 46
func SolutionB(input []byte) int {
	charToWord := map[byte][][]byte{
		'o': {[]byte("one")},
		't': {[]byte("two"), []byte("three")},
		'f': {[]byte("four"), []byte("five")},
		's': {[]byte("six"), []byte("seven")},
		'e': {[]byte("eight")},
		'n': {[]byte("nine")},
	}
	wordToValues := map[byte][]int{
		'o': {1},
		't': {2, 3},
		'f': {4, 5},
		's': {6, 7},
		'e': {8},
		'n': {9},
	}

	values := 0
	first, last := -1, -1

	for i := 0; i < len(input); i++ {
		cSlice := charToWord[input[i]]
		if cSlice != nil {
			for j := 0; j < len(cSlice); j++ {
				if i+len(cSlice[j]) > len(input) {
					break
				}
				window := input[i : i+len(cSlice[j])]
				eq := 0
				for k := 0; k < len(cSlice[j]); k++ {
					if cSlice[j][k] != window[k] {
						break
					}
					eq += 1
				}
				if eq == len(cSlice[j]) {
					if first == -1 {
						first = wordToValues[input[i]][j]
					} else {
						last = wordToValues[input[i]][j]
					}
					if charToWord[window[len(window)-1]] != nil {
						i -= 1 // edge case of oneight. last should be 8
					}
					i += len(cSlice[j]) - 1
				}
			}
		} else if input[i] >= 48 && input[i] <= 57 {
			if first == -1 {
				first = int(input[i]) - 48
			} else {
				last = int(input[i]) - 48
			}
		} else if input[i] == 10 {
			if last == -1 {
				last = first
			}
			values += first*10 + last
			first, last = -1, -1
		}
	}
	return values
}
