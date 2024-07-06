package day05

import (
	"bytes"
	"errors"
	"github.com/JDRadatti/aoc_go/pkg/utils"
	"math"
)

type ParsedAlmanac struct {
	seeds []int
	maps  []CategoryMap
}

func (almanac *ParsedAlmanac) ClosestLocation() int {
	var closest int = math.MaxInt
	for _, seed := range almanac.seeds {
		seedLocation := almanac.SeedLocation(seed)
		closest = utils.Min(closest, seedLocation)
	}
	return closest
}

func (almanac *ParsedAlmanac) ClosestLocationWithSeedRanges() int {
	var closest int = math.MaxInt
	for i := 1; i < len(almanac.seeds)-1; i += 2 {
		start := almanac.seeds[i-1]
		end := start + almanac.seeds[i]
		for j := start; j < end; j++ {
			seedLocation := almanac.SeedLocation(j)
			closest = utils.Min(closest, seedLocation)
		}
	}
	return closest

}

func (almanac *ParsedAlmanac) SeedLocation(seedNumber int) int {
	dest := seedNumber
	for _, m := range almanac.maps {
		dest = m.GetDestFromMap(dest)
	}
	return dest
}

func NewParsedAlmanac(almanac []byte) (ParsedAlmanac, error) {
	almanacLines := bytes.Split(almanac, []byte("\n"))
	parsed := ParsedAlmanac{}

	// Parse seeds
	seeds, err := ParseSeeds(almanacLines[0])
	if err != nil {
		panic(err)
	}
	parsed.seeds = seeds

	// Parse each map (separated by blank line)
	categoryMap := CategoryMap{}
	for i := 1; i < len(almanacLines); i++ {
		if len(almanacLines[i]) == 0 || i == len(almanacLines)-1 {
			if len(categoryMap.ranges) > 0 {
				parsed.maps = append(parsed.maps, categoryMap)
			}
			categoryMap = CategoryMap{}
			i++ // Skip next line
			continue
		}
		categoryMap.AddRange(almanacLines[i])
	}
	return parsed, nil
}

func ParseSeeds(seedLine []byte) ([]int, error) {
	seeds2d := bytes.Split(seedLine, []byte(" "))
	if bytes.Compare(seeds2d[0], []byte("seeds:")) != 0 {
		return nil, nil
	}
	var seeds1d []int
	for s := 1; s < len(seeds2d); s++ {
		seeds1d = append(seeds1d, utils.BAtoI(seeds2d[s]))
	}
	return seeds1d, nil
}

type CategoryMap struct {
	ranges []CategoryRange
}

func (c *CategoryMap) GetDestFromMap(source int) int {
	var dest int
	for _, r := range c.ranges {
		dest = r.GetDestFromRange(source)
		if dest >= 0 {
	        return dest
		}
	}
	return source
}

func (c *CategoryMap) AddRange(rangeBytes []byte) {
	r, err := NewCategoryRange(rangeBytes)
	if err != nil {
		panic(err)
	}
	c.ranges = append(c.ranges, r)
}

type CategoryRange struct {
	SourceStart      int
	DestinationStart int
	Size             int
}

// GetDestFromRange returns -1 when source: destination mapping not found
func (c *CategoryRange) GetDestFromRange(source int) int {
	if c.InRange(source) {
		offset := source - c.SourceStart
		return c.DestinationStart + offset
	}
	return -1
}

func (c *CategoryRange) InRange(source int) bool {
	return source >= c.SourceStart && source < c.SourceStart+c.Size
}

func NewCategoryRange(line []byte) (CategoryRange, error) {
	lineSplit := bytes.Split(line, []byte(" "))
	if len(lineSplit) != 3 {
		return CategoryRange{}, errors.New("invalid argument")
	}
	return CategoryRange{
		SourceStart:      utils.BAtoI(lineSplit[1]),
		DestinationStart: utils.BAtoI(lineSplit[0]),
		Size:             utils.BAtoI(lineSplit[2]),
	}, nil
}



// https://adventofcode.com/2023/day/5
func SolutionA(input []byte) int {
	parsedAlmanac, err := NewParsedAlmanac(input)
	if err != nil {
		panic(err)
	}
	return parsedAlmanac.ClosestLocation()
}
