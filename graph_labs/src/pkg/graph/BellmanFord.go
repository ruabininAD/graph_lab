package graph

import "log"

func (G *Graph) BellmanFord(a, b int) (path []int, err error) {
	log.Printf("func (G *Graph) BellmanFord(%d, %d int) (%v []int, %s error))\n", a, b, path, err)
	return nil, nil
}
