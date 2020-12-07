package part2

import (
	"github.com/indiependente/aoc2020/day6/models"
)

func SumOnlyCommonQuestionCount(groups []models.Group) int {
	sum := 0
	for _, g := range groups {
		sum += len(g.OnlyCommon())
	}
	return sum
}
