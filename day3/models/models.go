package models

import "strings"

const (
	OpenSquare CellType = `.`
	Tree       CellType = `#`
)

type CellType string

func (ct CellType) String() string {
	return string(ct)
}

type Cell struct {
	Type CellType
}

func (c Cell) IsOpenSquare() bool {
	return c.Type == OpenSquare
}

func (c Cell) IsTree() bool {
	return c.Type == Tree
}

type Step struct {
	Right int
	Down  int
}

type Map struct {
	grid [][]Cell
	rows int
	cols int
}

func NewMap(g [][]Cell, r, c int) Map {
	return Map{
		grid: g,
		rows: r,
		cols: c,
	}
}
func (m Map) At(i, j int) Cell {
	return m.grid[i%m.rows][j%m.cols]
}
func (m Map) Row(i int) []Cell {
	return m.grid[i%m.rows]
}
func (m Map) Col(j int) []Cell {
	columnI := make([]Cell, 0, m.rows)
	for i := range m.grid {
		columnI = append(columnI, m.grid[i][j])
	}
	return columnI
}
func (m Map) Size() (rows, cols int) {
	return m.rows, m.cols
}

func (m Map) String() string {
	var sb strings.Builder
	for i := range m.grid {
		for j := range m.grid[i] {
			sb.WriteString(m.grid[i][j].Type.String())
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
