package part1

import (
	"math"

	"github.com/indiependente/aoc2020/day5/models"
)

func HighestSeatID(passes []*models.BoardingPass) int {
	maxSeatID := 0
	for _, p := range passes {
		p.Partition(BinarySpacePartitioner)
		if p.SeatID() > maxSeatID {
			maxSeatID = p.SeatID()
		}
	}
	return maxSeatID
}

func BinarySpacePartitioner(code string) (int, int) {
	r, c := bsearch(code[:7], 'F', 'B'), bsearch(code[7:], 'L', 'R')
	return r, c
}
func bsearch(code string, low, high rune) int {
	length := int(math.Pow(2, float64(len(code))))
	min, max := 0, length-1
	for _, c := range code {
		switch c {
		case low:
			max = min + (max-min)/2
		case high:
			min = max - ((max - min) / 2)
		}
	}
	return min
}
