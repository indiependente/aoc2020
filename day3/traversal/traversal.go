package traversal

import "github.com/indiependente/aoc2020/day3/models"

func Traversal(m models.Map, step models.Step) (trees, open int) {
	rows, _ := m.Size()
	j := step.Right
	for i := step.Down; i < rows; i += step.Down {
		if m.At(i, j).IsTree() {
			trees++
		} else {
			open++
		}
		j += step.Right
	}
	return trees, open
}
