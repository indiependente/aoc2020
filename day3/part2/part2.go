package part2

import (
	"github.com/indiependente/aoc2020/day3/models"
	"github.com/indiependente/aoc2020/day3/traversal"
)

func Trees(m models.Map, steps ...models.Step) int {
	mulTrees := 1
	for _, step := range steps {
		trees, _ := traversal.Traversal(m, step)
		mulTrees *= trees
	}
	return mulTrees
}
