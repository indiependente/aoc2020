package part1

import "github.com/indiependente/aoc2020/day2/models"

func IsValid(p models.Password) bool {
	occurrences := 0
	for _, c := range p.Password {
		if c == p.Policy.Letter {
			occurrences++
		}
	}
	return p.Policy.Low <= occurrences && p.Policy.High >= occurrences
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
