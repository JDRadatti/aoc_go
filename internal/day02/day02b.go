package day02

type Cubes map[byte]int

func (c *Cubes) powerSet() int {
    result := 1
    for _, value := range *c {
        if value != 0 {
            result *= value
        }
    }
    return result
}

func (c *Cubes) zero() {
    for key := range *c {
        (*c)[key] = 0
    }
}

func SolutionB(input []byte) int {
	cubes := Cubes{
		'r': 0,
		'g': 0,
		'b': 0,
	}
    currCount := 0 
    sum := 0

	for i := 0; i < len(input); i++ {

		if input[i] == ':' {
            currCount = 0
		} else if input[i] >= 48 && input[i] <= 57 { // input[i] is int
            currCount = currCount * 10 + int(input[i]) - 48
		} else if input[i] == 10 { // New line
            sum += cubes.powerSet()
            cubes.zero()
            currCount = 0
        } else if currCount > 0 {
            if cube, ok := cubes[input[i]]; ok {
                cubes[input[i]] = max(cube, currCount)
                currCount = 0
            }
		}
    }
	return sum 
}
