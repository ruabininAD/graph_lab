package graph

import "fmt"

type Graph struct {
	VCount  int
	ECount  int
	Amatrix [][]int
	flags   map[string]bool
}

func NewGraph(VCount int) (*Graph, error) {

	if VCount < 1 {
		return nil, fmt.Errorf("Попытка создать граф с нулем веришн с помощью функции NewGreph\n")
	}

	G := Graph{VCount: VCount, Amatrix: make([][]int, VCount), flags: make(map[string]bool)}

	for i := 0; i < VCount; i++ {
		G.Amatrix[i] = make([]int, VCount)
	}

	G.flags["oriented"] = false
	G.flags["unoriented"] = false
	G.flags["tree"] = false
	G.flags[""] = false
	G.flags["acyclic"] = false

	return &G, nil
}

func (G *Graph) Set(row, col, value int) {
	if G.flags["oriented"] {

		G.Amatrix[row][col] = value
		G.ECount++

	} else {
		fmt.Errorf("граф не ориентированный")
	}
}

func (G *Graph) SetUnOrientedE(row, col, value int) {
	if G.flags["unoriented"] {

		G.Amatrix[row][col] = value
		G.Amatrix[col][row] = value

	} else {
		fmt.Errorf("граф ориентированный, нет флага unoriented")
	}

}
