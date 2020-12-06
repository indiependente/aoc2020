package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/indiependente/aoc2020/day4/models"
	"github.com/indiependente/aoc2020/day4/part1"
	"github.com/indiependente/aoc2020/day4/part2"
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
	fmt.Println(part1.CountValid(input))
	fmt.Println(part2.CountValid(input))
	return nil
}
func parseInput(r io.Reader) ([]models.Passport, error) {
	passports := make([]models.Passport, 0)
	br := bufio.NewReader(r)
	data := []byte{}
	for {
		line, err := br.ReadBytes('\n')
		if err == io.EOF {
			p, err := unmarshalData(data)
			if err != nil {
				return nil, err
			}
			passports = append(passports, p)
			break
		}
		if err != nil {
			return nil, err
		}
		if string(line) == "\n" {
			p, err := unmarshalData(data)
			if err != nil {
				return nil, err
			}
			passports = append(passports, p)
			data = []byte{}
			continue
		}
		data = append(data, line...)
	}
	return passports, nil
}

func unmarshalData(data []byte) (models.Passport, error) {
	p := models.Passport{}
	err := models.Unmarshal(data, &p)
	if err != nil {
		return models.Passport{}, err
	}
	return p, nil
}
