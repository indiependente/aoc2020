package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/indiependente/aoc2020/day1/part1"
	"github.com/indiependente/aoc2020/day1/part2"
)

func main() {
	if err := run(); err != nil {
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
	mul, err := part1.Sum2020(input)
	if err != nil {
		return err
	}
	fmt.Printf("Part 1 - The result is: %d\n", mul)
	mul, err = part2.Sum2020(input)
	if err != nil {
		return err
	}
	fmt.Printf("Part 2 - The result is: %d\n", mul)
	return nil
}

func parseInput(r io.Reader) ([]int, error) {
	input := make([]int, 0, 200)
	br := bufio.NewReader(r)
	i := 0
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		var n int
		fmt.Sscanf(line, "%d", &n)
		input = append(input, n)
		i++
	}
	return input, nil
}
