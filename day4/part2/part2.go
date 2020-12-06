package part2

import (
	"regexp"

	"github.com/indiependente/aoc2020/day4/models"
)

func CountValid(passports []models.Passport) int {
	valid := 0
	for _, p := range passports {
		errs := p.ValidWithValidator(validator)
		if !errs.Nil() {
			continue
		}
		valid++
	}
	return valid
}

func validator(p models.Passport) models.InvalidPassportErr {
	ipe := models.InvalidPassportErr{}
	if !validateBirthdayYear(p.BirthYear) {
		ipe.Add(models.ErrBirthYearInvalid)
	}
	if !validateIssueYear(p.IssueYear) {
		ipe.Add(models.ErrIssueYearInvalid)
	}
	if !validateExpirationYear(p.ExpirationYear) {
		ipe.Add(models.ErrExpirationYearInvalid)
	}
	if !validateHeight(p.Height) {
		ipe.Add(models.ErrHeightInvalid)
	}
	if !validateHairColor(p.HairColor) {
		ipe.Add(models.ErrHairColorInvalid)
	}
	if !validateEyeColor(p.EyeColor) {
		ipe.Add(models.ErrEyeColorInvalid)
	}
	if !validatePassportID(p.PassportID) {
		ipe.Add(models.ErrPassportIDInvalid)
	}
	return ipe
}

func validateBirthdayYear(byr string) bool {
	return len(byr) == 4 && byr >= "1920" && byr <= "2002"
}

func validateIssueYear(iyr string) bool {
	return len(iyr) == 4 && iyr >= "2010" && iyr <= "2020"
}

func validateExpirationYear(eyr string) bool {
	return len(eyr) == 4 && eyr >= "2020" && eyr <= "2030"
}

func validateHeight(hgt string) bool {
	if len(hgt) == 0 {
		return false
	}
	switch hgt[len(hgt)-2:] {
	case "cm":
		return hgt[:len(hgt)-2] >= "150" && hgt[:len(hgt)-2] <= "193"
	case "in":
		return hgt[:len(hgt)-2] >= "59" && hgt[:len(hgt)-2] <= "76"
	default:
		return false
	}
}

func validateHairColor(hcl string) bool {
	return len(hcl) == 7 && regexp.MustCompile(`^#[a-zA-Z0-9]*$`).MatchString(hcl)
}

func validateEyeColor(ecl string) bool {
	return ecl == "amb" ||
		ecl == "blu" ||
		ecl == "brn" ||
		ecl == "gry" ||
		ecl == "grn" ||
		ecl == "hzl" ||
		ecl == "oth"
}

func validatePassportID(pid string) bool {
	return len(pid) == 9
}
