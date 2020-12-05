package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/indiependente/aoc2020/day3/models"
	"github.com/indiependente/aoc2020/day3/part1"
	"github.com/indiependente/aoc2020/day3/part2"
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
	fmt.Println(part1.Trees(input))
	fmt.Println(part2.Trees(input,
		models.Step{
			Right: 1,
			Down:  1,
		},
		models.Step{
			Right: 3,
			Down:  1,
		},
		models.Step{
			Right: 5,
			Down:  1,
		},
		models.Step{
			Right: 7,
			Down:  1,
		},
		models.Step{
			Right: 1,
			Down:  2,
		}))
	return nil
}

func parseInput(r io.Reader) (models.Map, error) {
	grid := make([][]models.Cell, 0)
	br := bufio.NewReader(r)
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return models.Map{}, err
		}
		row := make([]models.Cell, 0)
		for _, cell := range line {
			if cell == '\n' {
				continue
			}
			row = append(row, models.Cell{Type: models.CellType(cell)})
		}
		grid = append(grid, row)
	}
	return models.NewMap(grid, len(grid), len(grid[0])), nil
}
