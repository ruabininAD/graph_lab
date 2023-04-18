package graph

import (
	"log"
	"math"
)

// func (G *Graph) BellmanFord(a, b int) (len int, path []int, err error) {
//
//		fmt.Printf("func (G *Graph) BellmanFord(%d, %d int) (%v int, %v []int, %s error))\n", a, b, len, path, err)
//
//		dist := make([]int, G.vCount)     // слайс для хранения расстояний от начальной вершины до остальных вершин
//		visited := make([]bool, G.vCount) // слайс для отслеживания посещенных вершин
//		prev := make([]int, G.vCount)     // слайс для хранения предыдущих вершин на кратчайшем пути
//
//		for i := 0; i < G.vCount; i++ {
//			dist[i] = int(^uint(0) >> 1) // устанавливаем бесконечное расстояние для всех вершин, кроме начальной
//			visited[i] = false
//			prev[i] = -1
//		}
//
//		dist[start] = 0 // расстояние от начальной вершины до самой себя равно 0
//
//
//
//
//
//		return 0, nil, nil
//	}
func (G *Graph) BellmanFord(start, end int) ([]int, int, error) {
	finish := end
	// initialize distances to infinity
	distance := make([]int, G.vCount)
	for i := range distance {
		distance[i] = math.MaxInt32
	}

	// set distance to start vertex to 0
	distance[start] = 0

	// relax edges |V|-1 times
	for i := 0; i < G.vCount-1; i++ {
		for u := 0; u < G.vCount; u++ {
			for v := 0; v < G.vCount; v++ {
				if G.Get(u, v) != 0 && distance[u]+G.Get(u, v) < distance[v] {
					distance[v] = distance[u] + G.Get(u, v)
				}
			}
		}
	}

	// check for negative weight cycles
	for u := 0; u < G.vCount; u++ {
		for v := 0; v < G.vCount; v++ {
			if G.Get(u, v) != 0 && distance[u]+G.Get(u, v) < distance[v] {
				return nil, -1, nil
			}
		}
	}

	// find shortest path
	path := []int{end}
	for end != start {
		for u := 0; u < G.vCount; u++ {
			if G.Get(u, end) != 0 && distance[end]-G.Get(u, end) == distance[u] {
				path = append([]int{u}, path...)
				end = u
			}
		}
	}

	log.Printf("вектор расстояний %v мое : %d\n", distance, distance[finish])
	return path, distance[finish], nil
}
