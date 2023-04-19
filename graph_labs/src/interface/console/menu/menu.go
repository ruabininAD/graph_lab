package menu

import (
	"fmt"
	"graph_labs/src/pkg/generator"
	"graph_labs/src/pkg/graph"
	"log"
	"math"
	"os"
	"os/exec"
	"strconv"
)

type Menu struct {
	logo  string
	graph *graph.Graph
}

func Cls() {
	cls := exec.Command("cmd", "/c", "cls")
	cls.Stdout = os.Stdout
	err := cls.Run()

	if err != nil {
		log.Printf("Ошибка очистки консоли\n")
	}
}

func (myMenu *Menu) MainMenu() {
	fmt.Printf("" +
		"1) Сгенерировать новый граф\n" +
		"2) Показать граф\n" +
		"3) Число дорог из A в B\n" +
		"4) Показать матрицу Шимбала\n" +
		"5) Применить алгоритм Дейкстры\n" +
		"6) Применить алгоритм Беллмана-Форда\n" +
		"7) Применить алгоритм Флойда\n" +
		"8) Добавить веса\n")
	choiceMainMenu := 0
	_, _ = fmt.Scan(&choiceMainMenu)
	Cls()
	switch choiceMainMenu {
	case 1:
		myMenu.Generated()
	case 2:
		myMenu.Print()
	case 3:
		a, b := 0, 0
		fmt.Println("Введите начальную и конечную вершины")
		_, _ = fmt.Scan(&a, &b)
		if a == b {
			fmt.Println("a == b")
			return
		}
		if a > myMenu.graph.GetVCount() || b > myMenu.graph.GetVCount() {
			fmt.Println("Bершины с таким индексом нет")
			break
		}
		count, shortestPath := myMenu.graph.CountPaths(a, b)
		if count == -1 {

			fmt.Printf("из %d в %d нет пути.\n", a, b)

		} else {

			fmt.Printf("из %d в %d есть %d путей. Самый короткий %d\n", a, b, count, shortestPath)

		}
	case 4:
		fmt.Printf("" +
			"1) результирующая матрица шимбала\n" +
			"2) шаг матрицы шимбала\n")
		choiseShimbal := 0
		_, _ = fmt.Scan(&choiseShimbal)
		switch choiseShimbal {
		case 1:
			fmt.Println("введите функцию min или max\n")
			fun := ""
			_, _ = fmt.Scan(&fun)
			myMenu.graph.ShimbelDistanceMatrix(fun).PrintLabel("Результурующая матрица шимбала для функции " + fun)

		case 2:
			fmt.Println("введите шаг для матрицы Шимбала\n")
			choiseStep := 0
			_, _ = fmt.Scan(&choiseStep)
			fmt.Println("введите функцию min или max\n")
			fun := ""
			_, _ = fmt.Scan(&fun)
			myMenu.graph.ShimbelStep(choiseStep, fun).PrintLabel("Матрица Шимбала для шага " + strconv.Itoa(choiseStep) + "и функции " + fun)
		}
	case 5:

		fmt.Println("Введите стартовую вершину для алгоритма Дейкстры\n")
		startV := 0
		_, _ = fmt.Scan(&startV)

		fmt.Println("Введите конечную вершину для алгоритма Дейкстры\n")
		finishV := 0
		_, _ = fmt.Scan(&finishV)

		if startV > myMenu.graph.GetVCount() || finishV > myMenu.graph.GetVCount() {

			fmt.Println("нет такой вершины")
			break

		}

		distance, path, err := myMenu.graph.Dijkstra(startV, finishV)
		if err != nil {
			log.Print(err)
			fmt.Println("ошибка")
		}

		if len(path) == 0 {
			fmt.Printf("между вершинами %d и  %d пути нет\n", startV, finishV)
			break
		}

		fmt.Printf("между вершинами %d и  %d путь длинной %d: %v ", startV, finishV, distance, path)
	case 6:

		fmt.Println("Введите стартовую вершину для алгоритма Беллмана Форда\n")
		startV := 0
		_, _ = fmt.Scan(&startV)

		fmt.Println("Введите конечную вершину для алгоритма  Беллмана Форда\n")
		finishV := 0
		_, _ = fmt.Scan(&finishV)

		if startV > myMenu.graph.GetVCount() {
			fmt.Println("нет такой вершины")
			break
		}
		path, distance, err := myMenu.graph.BellmanFord(startV, finishV)
		if err != nil {
			log.Print(err)
			fmt.Println("ошибка")
		}

		fmt.Printf("между вершинами %d и  %d путь длинной %d: %v ", startV, finishV, distance, path)
	case 7:

		dist, next, paths := myMenu.graph.Floid()
		printResFloid(dist, next, paths)

	case 8:
		fmt.Printf("" +
			"1) только положительные значения\n" +
			"2) любые значения\n")
		chioceWeight := 0
		_, _ = fmt.Scan(&chioceWeight)
		switch chioceWeight {
		case 1:
			myMenu.graph.SetRandomWeight("+")
		case 2:
			myMenu.graph.SetRandomWeight("-")
		default:
			fmt.Println("не та кнопочка")
		}

	default:
		fmt.Println("не та кнопочка")
	}

	_, _ = fmt.Scanln()
}

func (myMenu *Menu) Generated() {
	fmt.Printf("Создать:\n" +
		"1)  ориентированный граф с помощью распределения Элдинга\n" +
		//"2) ориентированный граф\n" +
		//"3) не ориентированный граф\n" +
		//"4) ориентированный ациклическй граф\n"
		"")

	graphVeriant := 0
	VCount := 0

	_, _ = fmt.Scan(&graphVeriant)
	Cls()
	fmt.Printf("количество вершин:\n")
	_, _ = fmt.Scan(&VCount)

	var myGraph *graph.Graph
	var err error

	switch graphVeriant {
	case 1:

		myGraph, err = generator.NewErlingAcyclicOrientedGraph(VCount)
		if err != nil {
			fmt.Printf("Для %d вершин нельзя постровить ациклический граф с %d ребер\n", VCount)
			return
		}
	default:
		fmt.Println("не та кнопка")
	}

	log.Printf("граф сгенерирован в консоли\n")

	myMenu.graph = myGraph

	myMenu.graph.PrintLabel("Граф сгенерирован:")

}

func (myMenu *Menu) Print() {
	fmt.Println(
		"1) вывести в консоль\n" +
			"2) рендер\n" +
			"3) показать свойства\n")

	ChoisePrint := 0

	_, _ = fmt.Scan(&ChoisePrint)

	switch ChoisePrint {
	case 1:
		myMenu.graph.PrintLabel("граф:")

	case 2:
		myMenu.graph.Render()

	case 3:
		for key, v := range myMenu.graph.Flags {

			if v == true {
				fmt.Printf("%s ", key)
			}

		}
		fmt.Println()

	default:
		fmt.Println("не та кнопка")
	}
}

func ConsoleMenu() {
	Cls()
	var menu Menu
	for {
		Cls()
		menu.MainMenu()
		_, _ = fmt.Scanln()
	}

}

func printResFloid(dist, next [][]int, paths map[string][]int) {

	// вывод матрицы расстояний
	fmt.Println("Матрица расстояний:")

	for _, row := range dist {
		for _, v := range row {
			if v == math.MaxInt32 {
				fmt.Printf("inf\t")
			} else {
				fmt.Printf("%d\t", v)
			}
		}
		fmt.Println()
	}

	// вывод матрицы следующих вершин на пути
	//fmt.Println("Матрица следующих вершин:")
	//for _, row := range next {
	//	fmt.Println(row)
	//}

	// вывод всех путей
	fmt.Println("Все пути:")
	for key, value := range paths {
		fmt.Printf("%s: %v\n", key, value)
	}
}
