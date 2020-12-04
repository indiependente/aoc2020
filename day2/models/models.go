package models

type Password struct {
	Policy   Policy
	Password string
}

type Policy struct {
	Low    int
	High   int
	Letter rune
}
