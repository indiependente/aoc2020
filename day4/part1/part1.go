package part1

import (
	"github.com/indiependente/aoc2020/day4/models"
)

func CountValid(passports []models.Passport) int {
	valid := 0
	for _, p := range passports {
		errs := p.ValidWithValidator(validator)
		if !errs.Nil() {
			if errs.HasOnly(models.ErrCountryIDInvalid) {
				valid++
				continue
			}
			continue
		}
		valid++
	}
	return valid
}

func validator(p models.Passport) models.InvalidPassportErr {
	ipe := models.InvalidPassportErr{
		Errs: []error{},
	}
	if p.BirthYear == "" {
		ipe.Errs = append(ipe.Errs, models.ErrBirthYearInvalid)
	}
	if p.IssueYear == "" {
		ipe.Errs = append(ipe.Errs, models.ErrIssueYearInvalid)
	}
	if p.ExpirationYear == "" {
		ipe.Errs = append(ipe.Errs, models.ErrExpirationYearInvalid)
	}
	if p.Height == "" {
		ipe.Errs = append(ipe.Errs, models.ErrHeightInvalid)
	}
	if p.HairColor == "" {
		ipe.Errs = append(ipe.Errs, models.ErrHairColorInvalid)
	}
	if p.EyeColor == "" {
		ipe.Errs = append(ipe.Errs, models.ErrEyeColorInvalid)
	}
	if p.PassportID == "" {
		ipe.Errs = append(ipe.Errs, models.ErrPassportIDInvalid)
	}
	if p.CountryID == "" {
		ipe.Errs = append(ipe.Errs, models.ErrCountryIDInvalid)
	}
	return ipe
}
