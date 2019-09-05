package sensor

import (
	"fmt"

	"github.com/sensornet/graph"
)

/*
示意图：
           5
           |
 1 -- 2 -- 3 -- 4
 |    |    |
 6 -- 7 -- 8
 |    |    |
 9 -- 10 - 11

*/
func ExampleGraph() {
	nodes := []*graph.Node{
		&graph.Node{Num: 1},
		&graph.Node{Num: 2},
		&graph.Node{Num: 3},
		&graph.Node{Num: 4},
		&graph.Node{Num: 5},
		&graph.Node{Num: 6},
		&graph.Node{Num: 7},
		&graph.Node{Num: 8},
		&graph.Node{Num: 9},
		&graph.Node{Num: 10},
		&graph.Node{Num: 11},
	}
	routes := []*graph.Route{
		&graph.Route{1, 2},
		&graph.Route{2, 3},
		&graph.Route{3, 5},
		&graph.Route{3, 4},
		&graph.Route{1, 6},
		&graph.Route{2, 7},
		&graph.Route{3, 8},
		&graph.Route{6, 7},
		&graph.Route{7, 8},
		&graph.Route{6, 9},
		&graph.Route{7, 10},
		&graph.Route{8, 11},
		&graph.Route{9, 10},
		&graph.Route{10, 11},
	}
	egGraph, err := graph.GenerateGraphWithNodesAndRoutes(nodes, routes)
	if err != nil {
		panic(err)
	}
	fmt.Println(egGraph)
}
