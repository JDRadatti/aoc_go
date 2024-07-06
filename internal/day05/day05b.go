package day05

//import ()
// https://adventofcode.com/2023/day/5
func SolutionB(input []byte) int {
	parsedAlmanac, err := NewParsedAlmanac(input)
	if err != nil {
		panic(err)
	}
	return parsedAlmanac.ClosestLocationWithSeedRanges()
}
