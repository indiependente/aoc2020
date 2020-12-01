package main

import (
	"fmt"
	"log"
	"os"

	"github.com/indiependente/aoc2020/day1/part1"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	f, err := os.Open("./part1/input")
	if err != nil {
		return err
	}
	mul, err := part1.Sum2020(f)
	if err != nil {
		return err
	}
	fmt.Printf("The result is: %d\n", mul)
	return nil
}
