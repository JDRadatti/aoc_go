package main

import (
	"github.com/JDRadatti/aoc_go/inputs"
	"github.com/JDRadatti/aoc_go/pkg/utils"
	"github.com/JDRadatti/aoc_go/internal/day01"
	"github.com/JDRadatti/aoc_go/internal/day02"
	"github.com/JDRadatti/aoc_go/internal/day03"
	"github.com/JDRadatti/aoc_go/internal/day04"
	"github.com/JDRadatti/aoc_go/internal/day05"
	"github.com/JDRadatti/aoc_go/internal/day06"
	"github.com/JDRadatti/aoc_go/internal/day07"
	"log"
	"os"
)

func run(day string) {
	input, err := inputs.ReadInput(day)
	if err != nil {
		log.Fatalf("failed to read input for %s \n%v", day, err)
	}

    var solutionA (func([]byte) int)
    var solutionB (func([]byte) int)

	switch day {
	default:
		log.Fatal("unexpected day")
	case "day01":
		solutionA = day01.SolutionA
        solutionB = day01.SolutionB
	case "day02":
		solutionA = day02.SolutionA
		solutionB = day02.SolutionB
    case "day03":
		solutionA = day03.SolutionA
		solutionB = day03.SolutionB
    case "day04":
		solutionA = day04.SolutionA
		solutionB = day04.SolutionB
    case "day05":
		solutionA = day05.SolutionA
		solutionB = day05.SolutionB
    case "day06":
		solutionA = day06.SolutionA
		solutionB = day06.SolutionB
    case "day07":
		solutionA = day07.SolutionA
		solutionB = day07.SolutionB
	}

    timeA := utils.TimeIt(day + " (A)")
    log.Printf("%s solution a: %v", day, solutionA(input))
    timeA()

    timeB := utils.TimeIt(day + " (B)")
	log.Printf("%s solution b: %v", day, solutionB(input))
    timeB() 
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
