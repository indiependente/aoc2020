package models

import (
	"errors"
	"strings"
)

const (
	BirthYearTag      string = `byr`
	IssueYearTag      string = `iyr`
	ExpirationYearTag string = `eyr`
	HeightTag         string = `hgt`
	HairColorTag      string = `hcl`
	EyeColorTag       string = `ecl`
	PassportIDTag     string = `pid`
	CountryIDTag      string = `cid`
)

var (
	ErrBirthYearInvalid      error = errors.New("invalid BirthYear")
	ErrIssueYearInvalid      error = errors.New("invalid IssueYear")
	ErrExpirationYearInvalid error = errors.New("invalid ExpirationYear")
	ErrHeightInvalid         error = errors.New("invalid Height")
	ErrHairColorInvalid      error = errors.New("invalid HairColor")
	ErrEyeColorInvalid       error = errors.New("invalid EyeColor")
	ErrPassportIDInvalid     error = errors.New("invalid PassportID")
	ErrCountryIDInvalid      error = errors.New("invalid CountryID")
)

type InvalidPassportErr struct {
	Errs []error
}

func (ipe *InvalidPassportErr) Add(e error) {
	ipe.Errs = append(ipe.Errs, e)
}

func (ipe InvalidPassportErr) Nil() bool {
	return len(ipe.Errs) == 0
}

func (ipe InvalidPassportErr) Has(e error) bool {
	for _, err := range ipe.Errs {
		if errors.Is(err, e) {
			return true
		}
	}
	return false
}

func (ipe InvalidPassportErr) HasOnly(e error) bool {
	return len(ipe.Errs) == 1 && errors.Is(ipe.Errs[0], e)
}

type Passport struct {
	BirthYear      string `passport:"byr"`
	IssueYear      string `passport:"iyr"`
	ExpirationYear string `passport:"eyr"`
	Height         string `passport:"hgt"`
	HairColor      string `passport:"hcl"`
	EyeColor       string `passport:"ecl"`
	PassportID     string `passport:"pid"`
	CountryID      string `passport:"cid"`
}

type Validator func(Passport) InvalidPassportErr

func (p Passport) ValidWithValidator(v Validator) InvalidPassportErr {
	return v(p)
}

func Unmarshal(data []byte, p *Passport) error {
	if len(data) == 0 {
		return nil
	}
	str := string(data)
	lines := strings.Split(str, "\n")
	lines = lines[:len(lines)-1]
	for _, l := range lines {
		fields := strings.Split(l, " ")
		for _, field := range fields {
			f := strings.Split(field, ":")
			k, v := f[0], f[1]
			switch k {
			case BirthYearTag:
				p.BirthYear = v
			case IssueYearTag:
				p.IssueYear = v
			case ExpirationYearTag:
				p.ExpirationYear = v
			case HeightTag:
				p.Height = v
			case HairColorTag:
				p.HairColor = v
			case EyeColorTag:
				p.EyeColor = v
			case PassportIDTag:
				p.PassportID = v
			case CountryIDTag:
				p.CountryID = v
			}
		}
	}

	return nil
}
