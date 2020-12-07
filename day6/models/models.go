package models

import (
	"fmt"
	"strings"
)

type Group struct {
	People []Person
}

func NewGroup(lines []string) Group {
	g := Group{
		People: []Person{},
	}
	for _, l := range lines {
		p := Person{
			Questions: []rune{},
		}
		l = strings.ReplaceAll(l, "\n", "")
		for _, c := range l {
			p.AddQuestion(c)
		}
		g.AddPerson(p)
	}
	return g
}

func (g *Group) AddPerson(p Person) {
	g.People = append(g.People, p)
}

func (g Group) DistinctQuestions() questions {
	set := make(map[rune]struct{})
	dq := make([]rune, 0)
	for _, p := range g.People {
		for _, q := range p.Questions {
			_, ok := set[q]
			if ok {
				continue
			}
			set[q] = struct{}{}
			dq = append(dq, q)
		}
	}
	return dq
}

func (g Group) OnlyCommon() questions {
	set := make(map[rune]int)
	cq := make([]rune, 0)
	for _, p := range g.People {
		for _, q := range p.Questions {
			set[q] += 1
		}
	}
	num_people := len(g.People)
	for k, v := range set {
		if v == num_people {
			cq = append(cq, k)
		}
	}
	return cq
}

func (g Group) String() string {
	pronoun := "people"
	if len(g.People) == 1 {
		pronoun = "person"
	}
	questions := g.DistinctQuestions()
	return fmt.Sprintf("%d %s answered \"yes\" to %d questions: %s",
		len(g.People),
		pronoun,
		len(questions),
		questions,
	)
}

type Person struct {
	Questions questions
}
type questions []rune

func (qq questions) String() string {
	if len(qq) == 0 {
		return ""
	}
	var sb strings.Builder
	sb.WriteRune(qq[0])
	for i := 1; i < len(qq); i++ {
		sb.WriteString(", ")
		sb.WriteString(string(qq[i]))
	}
	return sb.String()
}

func (p *Person) AddQuestion(q rune) {
	p.Questions = append(p.Questions, q)
}
