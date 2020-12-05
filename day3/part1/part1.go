package part1

import (
	"github.com/indiependente/aoc2020/day3/models"
	"github.com/indiependente/aoc2020/day3/traversal"
)

func Trees(m models.Map) int {
	trees, _ := traversal.Traversal(m, models.Step{Right: 3, Down: 1})
	return trees
}
