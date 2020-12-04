package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/indiependente/aoc2020/day2/models"
	"github.com/indiependente/aoc2020/day2/part1"
	"github.com/indiependente/aoc2020/day2/part2"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	f, err := os.Open("./input")
	if err != nil {
		return err
	}
	defer f.Close()
	input, err := parseInput(f)
	if err != nil {
		return err
	}
	validPassNo := part1.IsValidBulk(input)
	fmt.Println(validPassNo, " valid password(s) found")
	validPassNo = part2.IsValidBulk(input)
	fmt.Println(validPassNo, " valid password(s) found")
	return nil
}

func parseInput(r io.Reader) ([]models.Password, error) {
	var passwords []models.Password
	br := bufio.NewReader(r)
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		p := models.Password{}
		fmt.Sscanf(line, "%d-%d %c: %s\n",
			&p.Policy.Low,
			&p.Policy.High,
			&p.Policy.Letter,
			&p.Password)
		passwords = append(passwords, p)
	}
	return passwords, nil
}
