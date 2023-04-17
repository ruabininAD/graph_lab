package graph

import (
	"fmt"
)

func (G *Graph) BellmanFord(a, b int) (len int, path []int, err error) {
	fmt.Printf("func (G *Graph) BellmanFord(%d, %d int) (%v int, %v []int, %s error))\n", a, b, len, path, err)
	return 0, nil, nil
}
