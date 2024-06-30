package day04

import (
	"bytes"
    "log"
	"math"
)

// https://adventofcode.com/2023/day/4
func SolutionA(input []byte) int {
	cards := bytes.Split(input, []byte("\n"))
	sum := 0.0
	for c := range cards {
        if len(cards[c]) == 0 {
            continue
        }
		card := Card{}
		card.processCard(cards[c])
		if card.MatchingCount > 0 {
			sum += math.Pow(float64(2), float64(card.MatchingCount-1))
		}
	}
	return int(sum)
}

// https://adventofcode.com/2023/day/4
func SolutionB(input []byte) int {
	lines := bytes.Split(input, []byte("\n"))
	cards := []Card{}
	for c := range lines {
        if len(lines[c]) == 0 {
            continue
        }
		card := Card{Copies: 1}
		card.processCard(lines[c])
		cards = append(cards, card)
	}

	for c := range cards {
		card := cards[c]
		for i := range card.MatchingCount { 
			cards[c + i + 1].Copies += card.Copies // Guaranteed to be in bounds
		}
	}

	totalCopies := 0
	for c := range cards {
        log.Println(cards[c].Id, cards[c].Copies)
		totalCopies += cards[c].Copies
	}

	return totalCopies
}

type Card struct {
	Id            int
	MatchingCount int
	Copies        int
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
				c.MatchingCount++
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
