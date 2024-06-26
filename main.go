package main

import (
	"github.com/JDRadatti/aoc_go/inputs"
	"github.com/JDRadatti/aoc_go/internal/day01"
	"github.com/JDRadatti/aoc_go/internal/day02"
	"github.com/JDRadatti/aoc_go/internal/day03"
	"github.com/JDRadatti/aoc_go/internal/day04"
	"log"
	"os"
)

func run(day string) {
	input, err := inputs.ReadInput(day)
	if err != nil {
		log.Fatalf("failed to read input for %s \n%v", day, err)
	}

    solutionA, solutionB := 0, 0 

	switch day {
	default:
		log.Fatal("unexpected day")
	case "day01":
		solutionA = day01.SolutionA(input)
        solutionB = day01.SolutionB(input)
	case "day02":
		solutionA = day02.SolutionA(input)
		solutionB = day02.SolutionB(input)
    case "day03":
		solutionA = day03.SolutionA(input)
		solutionB = day03.SolutionB(input)
    case "day04":
		solutionA = day04.SolutionA(input)
		solutionB = day04.SolutionB(input)
	}
    log.Printf("%s solution a: %v", day, solutionA)
	log.Printf("%s solution b: %v", day, solutionB)

}

func help() {
	os.Stdout.Write([]byte("\nmain.go - advent of code runner\n" +
		"USAGE: go run main.go dayXX\n"))
	os.Exit(0)
}

func main() {
	if len(os.Args) != 2 {
		log.Printf("ERROR: invalid arguments.")
		help()
	}
	dirname := os.Args[1]
	run(dirname)
}
