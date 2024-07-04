package rangeset

import (
	"github.com/JDRadatti/aoc_go/pkg/utils"
	"log"
	"slices"
	"sort"
)

type Range struct {
	Start int
	End   int
}

func NewRange(start, end int) Range     { return Range{start, end} }
func (r *Range) Contains(elem int) bool { return elem >= r.Start && elem <= r.End }
func (r *Range) UpdateStart(elem int)   { r.Start = elem }
func (r *Range) UpdateEnd(elem int)     { r.End = elem }
func (r *Range) Compare(elem int) int {
	if elem < r.Start {
		return -1
	} else if elem > r.End {
		return 1
	}
	return 0
}

type Ranges []Range

func (r Ranges) Len() int           { return len(r) }
func (r Ranges) Less(i, j int) bool { return r[i].Start < r[j].Start }
func (r Ranges) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

type RangeSet struct {
	Ranges Ranges //Increasing order, each subarray is disjoint
}

func (self *RangeSet) Insert(r Range, i int) {
	self.Ranges = slices.Insert(self.Ranges, i, r)
}

func (self *RangeSet) Delete(i, j int) {
	self.Ranges = slices.Delete(self.Ranges, i, j)
}

func NewRangeSet() RangeSet { return RangeSet{} }

// Add to RangeSet.
func (self *RangeSet) Add(start int, end int) {
	if start > end {
		end, start = start, end
	} else if self.size() == 0 {
		// add new range
		self.Ranges = append(self.Ranges, NewRange(start, end))
		return
	}

	startIndex := self.Search(start)
	endIndex := self.Search(end)
	startRange := &self.Ranges[startIndex]
	endRange := &self.Ranges[endIndex]

	// Start a {b} c
	// End x {y} z
	// ax {} or {} cz == Insert new Range ({} are the same)
	// a {b} c   {}...{}   x {y} z     3^2=9 possibilites
	// cx -- insert cx and remove intermediates
	// ax, bx -- update startRange and remove all between startRange and endRange
	// cy, cz -- update endRange and remove all between startRange and endRange
	// ay, az, by, bz -- update startRange and remove all from start to endRange(included)

	if startIndex == endIndex {
		if end < startRange.Start { // ax
			self.Insert(NewRange(start, end), startIndex)
		} else if start > startRange.End { // cz
			self.Insert(NewRange(start, end), startIndex+1) // what happens at last index?
		} else {
			startRange.UpdateStart(utils.Min(start, startRange.Start))
			startRange.UpdateEnd(utils.Max(end, startRange.End))
		}
	} else {
		if start > startRange.End && end < endRange.Start { // cx
            self.Delete(startIndex+1, endIndex)
            self.Insert(NewRange(start, end), startIndex+1)
		} else if end < endRange.Start { // ax or bx
			startRange.UpdateStart(utils.Min(start, startRange.Start))
			startRange.UpdateEnd(utils.Max(end, startRange.End))
			self.Delete(startIndex+1, endIndex)
		} else if start > startRange.End { // cy or cz
			endRange.UpdateStart(utils.Min(start, endRange.Start))
			endRange.UpdateEnd(utils.Max(end, endRange.End))
			self.Delete(startIndex+1, endIndex)
		} else { // ay, az, by, bz
			startRange.UpdateStart(utils.Min(start, startRange.Start))
			startRange.UpdateEnd(utils.Max(end, endRange.End))
			self.Delete(startIndex+1, endIndex+1)
		}
	}
}
func (self *RangeSet) Union(other RangeSet) {
}

func (self *RangeSet) Intersection(other RangeSet) {
}

func (self *RangeSet) Difference(other RangeSet) {
}

func (self *RangeSet) Contains() {
}

func (self *RangeSet) Empty() bool {
	return self.size() == 0
}

// Size of Ranges, NOT of all the elements in the set
func (self *RangeSet) size() int {
	return len(self.Ranges)
}

// Search performs a binary search
// Returns index of ELEM, even if in gutter
// gutter = between two ranges
// In my case, I don't need to differentiate between range and gutter
//
// Example of return values: [-1 {0}, -2 {1}, -3 {2} -4]
// {x} = range and corresponding index x
// -x = gutter
// If in gutter -x, returns index of range to the right
func (self *RangeSet) Search(elem int) int {
	if self.size() == 0 {
		return 0
	}

	low, high := 0, self.size()-1
	mid := high / 2
	for low <= high {
		if low == high {
			mid = high
		} else {
            mid = low + ((high - low + 1) / 2)
        }
		if self.Ranges[mid].Contains(elem) {
			return mid
		} else if self.Ranges[mid].Compare(elem) < 0 {
			high = mid-1
		} else {
			low = mid+1
		} 	
    }
	return utils.Min(self.size()-1, low)
	// Returns -1 * (low + 1) for where to insert(gutter)
}
