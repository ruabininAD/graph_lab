package menu

import (
	"fmt"
	"graph_labs/src/pkg/generator"
	"graph_labs/src/pkg/graph"
	"log"
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
	cls.Run()
}

func (myMenu *Menu) MainMenu() {
	fmt.Printf("" +
		"1) Сгенерировать новый граф\n" +
		"2) Показать граф\n" +
		"3) Число дорог из A в B\n" +
		"4) Показать матрицу Шимбала\n" +
		"5) Применить алгоритм Дейкстры\n" +
		"6) Применить алгориитм Беллмана-Форда\n")
	choiseMainMenu := 0
	fmt.Scan(&choiseMainMenu)
	Cls()
	switch choiseMainMenu {
	case 1:
		myMenu.Generated()
	case 2:
		myMenu.Print()
	case 3:
		a, b := 0, 0
		fmt.Println("Введите начальную и конечную вершины")
		fmt.Scan(&a, &b)
		if a == b {
			fmt.Println("a == b")
			return
		}
		if a > myMenu.graph.GetVCount() || b > myMenu.graph.GetVCount() {
			fmt.Println("Bершины с таким индексом нет")
			break
		}
		//fmt.Println(myMenu.graph.HowManuRoads(a, b))

	case 4:
		fmt.Printf("" +
			"1) результирующая матрица шимбала\n" +
			"2) шаг матрицы шимбала\n")
		choiseShimbal := 0
		fmt.Scan(&choiseShimbal)
		switch choiseShimbal {
		case 1:
			fmt.Println("введите функцию min или max\n")
			fun := ""
			fmt.Scan(&fun)
			myMenu.graph.ShimbelDistanceMatrix(fun).PrintLabel("Результурующая матрица шимбала для функции " + fun)

		case 2:
			fmt.Println("введите шаг для матрицы Шимбала\n")
			choiseStep := 0
			fmt.Scan(&choiseStep)
			fmt.Println("введите функцию min или max\n")
			fun := ""
			fmt.Scan(&fun)
			myMenu.graph.Shimbel_step(choiseStep, fun).PrintLabel("Матрица Шимбала для шага " + strconv.Itoa(choiseStep) + "и функции " + fun)
		}
	case 5:
		fmt.Println("Введите стартовую вершину для алгоритма Дейкстры\n")
		startV := 0
		fmt.Scan(&startV)
		if startV > myMenu.graph.GetVCount() {
			fmt.Println("нет такой вершины")
			break
		}
		fmt.Println("тут должен быть алгоритм дейкстры")
	case 6:
		fmt.Println("Введите стартовую вершину для алгоритма Беллмана Форда\n")
		startV := 0
		fmt.Scan(&startV)
		if startV > myMenu.graph.GetVCount() {
			fmt.Println("нет такой вершины")
			break
		}
		fmt.Println("тут должен быть алгоритм Беллмана-Форда")
	}

	fmt.Scanln()
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
	ECount := 0

	fmt.Scan(&graphVeriant)
	Cls()
	fmt.Printf("количество вершин:\n")
	fmt.Scan(&VCount)

	fmt.Printf("количество ребер:\n")
	fmt.Scan(&ECount)

	var myGraph *graph.Graph
	var err error

	switch graphVeriant {
	case 1:

		myGraph, err = generator.NewErlingAcyclicOrientedGraph(VCount, ECount)
		if err != nil {

			fmt.Printf("Для %d вершин нельзя постровить ациклический граф с %d ребер\n", VCount, ECount)

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
			"2) рендер")

	ChoisePrint := 0

	fmt.Scan(&ChoisePrint)

	switch ChoisePrint {
	case 1:
		myMenu.graph.PrintLabel("граф:")

	case 2:
		myMenu.graph.Render()
	}
}

func ConsoleMenu() {
	Cls()
	var menu Menu
	for {
		menu.MainMenu()
		fmt.Scanln()
	}

}
