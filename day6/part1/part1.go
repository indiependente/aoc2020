package part1

import "github.com/indiependente/aoc2020/day6/models"

func SumQuestionCount(groups []models.Group) int {
	sum := 0
	for _, g := range groups {
		sum += len(g.DistinctQuestions())
	}
	return sum
}
