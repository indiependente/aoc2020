package part2

import "github.com/indiependente/aoc2020/day2/models"

func IsValid(p models.Password) bool {
	return (rune(p.Password[p.Policy.Low-1]) == p.Policy.Letter) !=
		(rune(p.Password[p.Policy.High-1]) == p.Policy.Letter)
}

func IsValidBulk(passwords []models.Password) int {
	valid := 0
	for _, p := range passwords {
		if IsValid(p) {
			valid++
		}
	}
	return valid
}
