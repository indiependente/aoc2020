package part2

import (
	"sort"

	"github.com/indiependente/aoc2020/day5/models"
)

func MySeatID(passes []*models.BoardingPass) int {
	var mySeatID int
	ids := []int{}
	for _, p := range passes {
		ids = append(ids, p.SeatID())
	}
	sort.Ints(ids)
	for i := 0; i < len(ids)-1; i++ {
		if ids[i]+2 == ids[i+1] {
			mySeatID = ids[i] + 1
		}
	}
	return mySeatID
}
