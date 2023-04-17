package graph

import (
	"fmt"
	"log"
)

func (G *Graph) Shimbel_step(step int, fun string) *Graph {
	//step соответствует степени
	res := G
	for i := 0; i < step-1; i++ {
		res = ShimbelMultiply(res, G, fun)
	}
	return res
}

func (G *Graph) ShimbelDistanceMatrix(fun string) *Graph {
	log.Printf("start ShimbelDistanceMatrix  fun =%s", fun)
	// min  минимальный маршрут от точки до точки
	// max  максимальный маршрут от точки до точки
	res, err := NewGraph(G.GetVCount())

	if err != nil {
		log.Print(err)
	}

	ShimbelSteps := make([]*Graph, 0)
	for i := 1; i < G.GetVCount(); i++ {
		tmp := G.Shimbel_step(i, fun)
		ShimbelSteps = append(ShimbelSteps, tmp)
	}

	for i := 0; i < G.vCount; i++ {
		for j := 0; j < G.vCount; j++ {
			arrIJ := make([]int, G.vCount)
			log.Printf("start generated tab i = %v, j =%v", i, j)
			for step := 0; step < G.vCount-1; step++ {
				//log.Printf("append(arrIJ, ShimbelSteps[step].Get(i, j)) i = %v, j =%v", i, j)
				arrIJ = append(arrIJ, ShimbelSteps[step].Get(i, j))
			}

			value := 0
			if fun == "max" {
				value = max(arrIJ)
			} else {
				value = min(arrIJ)
			}
			res.Set(i, j, value)
		}
	}
	return res
}

func ShimbelMultiply(b, a *Graph, fun string) *Graph {
	if a.vCount != b.vCount {
		log.Printf("The matrices cannot be multiplied: n1 =%v, n2 =  %v", a.vCount, b.vCount)
		panic("The matrices cannot be multiplied")
	}

	result, err := NewGraph(a.vCount)

	if err != nil {
		log.Println(err)
	}

	for i := 0; i < a.vCount; i++ {
		for j := 0; j < a.vCount; j++ {
			arr := make([]int, a.vCount) //тут массив
			for k := 0; k < a.vCount; k++ {
				if a.Amatrix[i][k] == 0 || b.Amatrix[k][j] == 0 {
					arr = append(arr, 0)
					continue
				}
				arr = append(arr, a.Amatrix[i][k]+b.Amatrix[k][j])
				// добавление в массив
			}
			res := 0
			if fun == "max" {
				res = max(arr)
			} else {
				res = min(arr)
			}

			result.Amatrix[i][j] = res // выбор наибольшего из массива.
		}
	}

	return result
}

func (G *Graph) HowManuRoads(start, stop int) string {
	log.Printf("start HowManuRoads\n")

	if start == stop {
		return fmt.Sprintf("%v is %v:   Minimum road = %v \n", start, stop, 0)
	}

	// Инициализация массива для хранения количества путей до каждой вершины
	counts := make([]int, G.vCount)
	counts[start] = 1 // Количество путей до начальной вершины равно 1

	// Проход по матрице смежности с использованием алгоритма динамического программирования
	for i := 0; i < G.vCount; i++ {
		for j := 0; j < G.vCount; j++ {
			if G.Amatrix[i][j] == 1 {
				counts[j] += counts[i] // Обновление количества путей до вершины j
			}
		}
	}
	//return counts[end] // Количество путей до конечной вершины
	return fmt.Sprintf("From %v to %v: There are %v roads.  Minimum road = %v \n", start, stop, counts[stop])

}

func countNZero(arrRoad []int) (countNZero int) {
	count := 0
	log.Printf("проверка на нули массива: %q\n", fmt.Sprint(arrRoad))
	for lenRoad, countRoad := range arrRoad {
		if lenRoad != 0 {
			log.Printf("имеется %v дорог длинной в %v\n", countRoad, lenRoad)
			count += 1
		}
	}
	return count
}
