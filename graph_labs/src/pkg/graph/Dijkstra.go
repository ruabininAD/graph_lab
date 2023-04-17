package graph

import "log"

func (G *Graph) Dijkstra(a, b int) (path []int, err error) {
	log.Printf("func (G *Graph) Dijkstra(%d, %d int) (%v []int, %s error))\n", a, b, path, err)
	return nil, nil
}
