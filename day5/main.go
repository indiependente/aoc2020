package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/indiependente/aoc2020/day5/models"
	"github.com/indiependente/aoc2020/day5/part1"
	"github.com/indiependente/aoc2020/day5/part2"
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
	fmt.Println(part1.HighestSeatID(input))
	fmt.Println(part2.MySeatID(input))
	return nil
}

func parseInput(r io.Reader) ([]*models.BoardingPass, error) {
	passports := make([]*models.BoardingPass, 0)
	br := bufio.NewReader(r)
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		passports = append(passports,
			models.NewBoardingPass(strings.Replace(line, "\n", "", -1)),
		)

	}
	return passports, nil
}
