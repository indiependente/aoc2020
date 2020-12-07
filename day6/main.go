package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/indiependente/aoc2020/day6/models"
	"github.com/indiependente/aoc2020/day6/part1"
	"github.com/indiependente/aoc2020/day6/part2"
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
	fmt.Println(part1.SumQuestionCount(input))
	fmt.Println(part2.SumOnlyCommonQuestionCount(input))
	return nil
}

func parseInput(r io.Reader) ([]models.Group, error) {
	groups := make([]models.Group, 0)
	br := bufio.NewReader(r)
	data := []string{}
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			groups = append(groups, models.NewGroup(data))
			break
		}
		if err != nil {
			return nil, err
		}
		if string(line) == "\n" {
			groups = append(groups, models.NewGroup(data))
			data = []string{}
			continue
		}
		data = append(data, line)
	}
	return groups, nil
}
