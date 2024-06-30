package day04

import (
	"bytes"
	"math"
)

// https://adventofcode.com/2023/day/4
func SolutionA(input []byte) int {
	cards := bytes.Split(input, []byte("\n"))
	sum := 0.0
	for c := range cards {
		card := Card{}
		card.processCard(cards[c])
		if card.matchingCount > 0 {
			sum += math.Pow(float64(2), float64(card.matchingCount-1))
		}
	}
	return int(sum)
}

// https://adventofcode.com/2023/day/4
func SolutionB(input []byte) int {
	cards := bytes.Split(input, []byte("\n"))
	for c := range cards {
		card := Card{}
		card.processCard(cards[c])

	}
	return 0
}



type Card struct {
	Id            int
	matchingCount int
}

func (c *Card) processCard(card []byte) {
	tokens := bytes.Split(card, []byte(" "))
	if len(tokens) <= 1 {
		return
	}
    t := 1 // Skip "Game"
	for len(tokens[t]) == 0 {
        t++
	} // Skip extra spaces

	c.Id = BAtoI(tokens[t][:len(tokens[t])-1])
    t++

	numbersBefore := map[int]struct{}{}
	for t < len(tokens) {
		if bytes.Compare(tokens[t], []byte("|")) == 0 {
			t++
			break
		}

        if n := BAtoI(tokens[t]); n != 0 { // ignore whitespace
              numbersBefore[n] = struct{}{}
        }
		t++
	}

	for t < len(tokens) {
		currNumber := BAtoI(tokens[t])
		if _, ok := numbersBefore[currNumber]; ok {
			if currNumber != 0 {
				c.matchingCount++
			}
		}
		t++
	}
}

// BAtoI converts a Byte Array to Int
func BAtoI(bytes []byte) int {
	value := 0
	for i := range bytes {
		value = value*10 + int(bytes[i]) - '0'
	}
	return value
}


