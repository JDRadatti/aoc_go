package day07

import (
	"bytes"
	"github.com/JDRadatti/aoc_go/pkg/utils"
	"sort"
)

// https://adventofcode.com/2023/day/7
func SolutionA(input []byte) int {
	hands := ParseHands(bytes.Split(input, []byte("\n")))
	sort.Sort(hands)
	sum := 0
	for i, hand := range hands {
		sum += (i + 1) * hand.Bid
	}
	return sum
}

// https://adventofcode.com/2023/day/7
func SolutionB(input []byte) int {
	return 0
}

type Label int

var LabelStrength map[byte]Label = map[byte]Label{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
	'1': 0,
}

type Hand struct {
	Bid    int
	Labels []Label
	Score  int
}

func (h Hand) Compare(other Hand) int {
	if h.Score > other.Score {
		return 1
	} else if h.Score < other.Score {
		return -1
	}

	for k := range h.Labels {
		if h.Labels[k] > other.Labels[k] {
			return 1
		} else if h.Labels[k] < other.Labels[k] {
			return -1
		}
	}
	return 0
}

type Hands []Hand

func ParseHands(handsInput [][]byte) Hands {

	hands := Hands{}
	for _, hand := range handsInput {
		split := bytes.Split(hand, []byte(" "))
		if len(split) != 2 {
			break
		}

		labels := make([]Label, 5)
		counts := map[Label]int{}
		for i, label := range split[0] {
			labels[i] = LabelStrength[label]
			counts[labels[i]]++
		}

		var handScore int
		for _, count := range counts {
			handScore += (count * count)
		}

		hands = append(hands, Hand{
			Bid:    utils.BAtoI(split[1]),
			Labels: labels,
			Score:  handScore,
		})
	}
	return hands
}

func (a Hands) Len() int      { return len(a) }
func (a Hands) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Hands) Less(i, j int) bool {
	if len(a[i].Labels) == 0 {
		return true
	} else if len(a[j].Labels) == 0 {
		return false
	}
	return a[i].Compare(a[j]) < 0
}
