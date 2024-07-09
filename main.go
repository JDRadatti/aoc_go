package main

import (
	"github.com/JDRadatti/aoc_go/inputs"
	"github.com/JDRadatti/aoc_go/internal/day01"
	"github.com/JDRadatti/aoc_go/internal/day02"
	"github.com/JDRadatti/aoc_go/internal/day03"
	"github.com/JDRadatti/aoc_go/internal/day04"
	"github.com/JDRadatti/aoc_go/internal/day05"
	"github.com/JDRadatti/aoc_go/internal/day06"
	"github.com/JDRadatti/aoc_go/internal/day07"
	"github.com/JDRadatti/aoc_go/internal/day08"
	"github.com/JDRadatti/aoc_go/internal/day09"
	"github.com/JDRadatti/aoc_go/pkg/utils"
	"log"
	"os"
	"path/filepath"
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
	case "day08":
		solutionA = day08.SolutionA
		solutionB = day08.SolutionB
	case "day09":
		solutionA = day09.SolutionA
		solutionB = day09.SolutionB
	}

	timeA := utils.TimeIt(day + " (A)")
	log.Printf("%s solution a: %v", day, solutionA(input))
	timeA()

	timeB := utils.TimeIt(day + " (B)")
	log.Printf("%s solution b: %v", day, solutionB(input))
	timeB()
}

func help() {
	os.Stdout.Write([]byte(`main.go - advent of code runner
USAGE: go run main.go dayXX [FLAG]

FLAGS: 
    -h Show this help message
    -g Generate new template for dayXX
`))
	os.Exit(0)
}

func generate(day string) {
	path := filepath.Join("internal", day)

	err := os.Mkdir(path, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	err = os.WriteFile(filepath.Join(path, day+".go"),
		[]byte("package "+day+`

// https://adventofcode.com/2023/day/
func SolutionA(input []byte) int {
        return 0
}

// https://adventofcode.com/2023/day/
func SolutionB(input []byte) int {
        return 0
}
`), 0660)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filepath.Join(path, day+"_test.go"),
		[]byte(`package `+day+`

import (
    "testing"
)

func Test`+`D`+day[1:]+`a(t *testing.T) {
	inputs := [][]byte{
		[]byte(`+"`"+"`"+`),
	}
	expected := [...]int{
		0,
	}
	for i := 0; i < len(inputs); i++ {
		result := SolutionA(inputs[i])

		if expected[i] != result {
			t.Errorf("Expected %d but got %d", expected[i], result)
		}
	}
}

func Test`+`D`+day[1:]+`b(t *testing.T) {
	inputs := [][]byte{
		[]byte(`+"`"+"`"+`),
	}
	expected := [...]int{
		0,
	}
	for i := 0; i < len(inputs); i++ {
		result := SolutionB(inputs[i])

		if expected[i] != result {
			t.Errorf("Expected %d but got %d", expected[i], result)
		}
	}
}`), 0660)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filepath.Join(filepath.Join("inputs", day+".txt")),
		[]byte(""), 0660)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	day := os.Args[1]
	if len(os.Args) == 3 && os.Args[2] == "-g" {
		generate(day)
	} else if len(os.Args) == 2 && os.Args[1] == "-h" {
		help()
	} else if len(os.Args) == 2 {
		run(day)
	} else {
		log.Printf("ERROR: invalid arguments.")
		help()
	}
}
