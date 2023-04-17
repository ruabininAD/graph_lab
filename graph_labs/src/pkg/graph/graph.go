package graph

import (
	"fmt"
	"log"
	"math/rand"
	"os"
)

type Graph struct {
	vCount  int
	eCount  int
	Amatrix [][]int
	Flags   map[string]bool
}

func NewGraph(VCount int) (*Graph, error) {

	if VCount < 1 {
		return nil, fmt.Errorf("Попытка создать граф с нулем веришн с помощью функции NewGreph\n")
	}

	G := Graph{vCount: VCount, Amatrix: make([][]int, VCount), Flags: make(map[string]bool)}

	for i := 0; i < VCount; i++ {
		G.Amatrix[i] = make([]int, VCount)
	}

	G.Flags["oriented"] = false
	G.Flags["unoriented"] = false
	G.Flags["tree"] = false
	G.Flags[""] = false
	G.Flags["acyclic"] = false

	return &G, nil
}

func (G *Graph) Set(row, col, value int) {
	if G.Flags["oriented"] {

		G.Amatrix[row][col] = value
		G.eCount++

	} else {
		fmt.Errorf("граф не ориентированный")
		log.Printf("попытка внести ориентированное ребро в неориентированный граф")
	}
}

func (G *Graph) SetUnOrientedE(row, col, value int) {
	if G.Flags["unoriented"] {

		G.Amatrix[row][col] = value
		G.Amatrix[col][row] = value

	} else {
		fmt.Errorf("граф ориентированный, нет флага unoriented")
		log.Printf("попытка внести не ориентированное ребро в ориентированный граф")
	}

}

func (G *Graph) Get(row, col int) int {
	return G.Amatrix[row][col]
}

func (G *Graph) Print() {

	for i := 0; i < G.vCount; i++ {
		for j := 0; j < G.vCount; j++ {

			fmt.Printf("\t%d", G.Amatrix[i][j])

		}

		fmt.Println() // перевод строки
	}

	fmt.Println() //переход на новую строку
}

func (G *Graph) PrintLabel(text string) {
	fmt.Println(text)
	G.Print()
}

func (G *Graph) Render() {
	file, err := os.OpenFile("src/graph.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	text := ""
	for i := 0; i < G.vCount; i++ {
		text = ""
		for j := 0; j < G.vCount; j++ {

			text += fmt.Sprintf("%d, ", G.Amatrix[i][j])

		}
		file.WriteString(text + "\n")
	}

}

func (G *Graph) GetECount() int {
	return G.eCount
}

func (G *Graph) GetVCount() int {
	return G.vCount
}

func (G *Graph) SetRandomWeigh() {
	for i := 0; i < G.vCount; i++ {
		for j := 0; j < G.vCount; j++ {
			if G.Get(i, j) != 0 {
				G.Set(i, j, rand.Intn(100))
			}
		}
	}
}

func max(arr []int) int {
	max := arr[0]
	for _, element := range arr {
		if element > max {
			max = element
		}
	}
	return max
}

func min(arr []int) int {
	min := max(arr)
	for _, element := range arr {
		if element < min && element != 0 {
			min = element
		}
	}
	return min
}
