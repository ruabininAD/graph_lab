package test_graph

import (
	"graph_labs/src/pkg/graph"
	"testing"
)

func TestNewGraph(t *testing.T) { // fixme

	// arrange
	input := 1

	expected, err := graph.NewGraph(input)
	if err != nil {
		t.Error(err)
	}

	// act
	res, err := graph.NewGraph(input)
	if err != nil {
		t.Error(err)
	}

	// asserts

	if res.Amatrix[0][0] != expected.Amatrix[0][0] {
		t.Errorf("не корректное сравнение графов")
	}

}
