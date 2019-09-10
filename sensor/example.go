package sensor

import (
	"github.com/sensornet/graph"
)

/*
示意图：
           4
           |
 0 -- 1 -- 2 -- 3
 |    |    |
 5 -- 6 -- 7
 |    |    |
 8 -- 9 -- 10

*/
func ExampleGraph1() {
	n0 := &graph.Node{Num: 0}
	n1 := &graph.Node{Num: 1}
	n2 := &graph.Node{Num: 2}
	n3 := &graph.Node{Num: 3}
	n4 := &graph.Node{Num: 4}
	n5 := &graph.Node{Num: 5}
	n6 := &graph.Node{Num: 6}
	n7 := &graph.Node{Num: 7}
	n8 := &graph.Node{Num: 8}
	n9 := &graph.Node{Num: 9}
	n10 := &graph.Node{Num: 10}
	nodes := []*graph.Node{
		n0, n1, n2, n3, n4, n5, n6, n7, n8, n9, n10,
	}
	routes := []*graph.Route{
		graph.NewRoute(n0, n1),
		graph.NewRoute(n1, n2),
		graph.NewRoute(n2, n4),
		graph.NewRoute(n2, n3),
		graph.NewRoute(n0, n5),
		graph.NewRoute(n1, n6),
		graph.NewRoute(n2, n7),
		graph.NewRoute(n5, n6),
		graph.NewRoute(n6, n7),
		graph.NewRoute(n5, n8),
		graph.NewRoute(n6, n9),
		graph.NewRoute(n7, n10),
		graph.NewRoute(n8, n9),
		graph.NewRoute(n9, n10),
	}
	egGraph, err := graph.GenerateGraphWithNodesAndRoutes(nodes, routes)
	if err != nil {
		panic(err)
	}
	InstallSensor1(egGraph)
}

/*
 0 - 1 - 2 - 3 - 4
     |       |   |
     5 -- -- 6 - 7
     |       |
     8 -- -- 9
*/
func ExampleGraph2() {
	n0 := &graph.Node{Num: 0}
	n1 := &graph.Node{Num: 1}
	n2 := &graph.Node{Num: 2}
	n3 := &graph.Node{Num: 3}
	n4 := &graph.Node{Num: 4}
	n5 := &graph.Node{Num: 5}
	n6 := &graph.Node{Num: 6}
	n7 := &graph.Node{Num: 7}
	n8 := &graph.Node{Num: 8}
	n9 := &graph.Node{Num: 9}
	nodes := []*graph.Node{
		n0, n1, n2, n3, n4, n5, n6, n7, n8, n9,
	}
	routes := []*graph.Route{
		graph.NewRoute(n0, n1),
		graph.NewRoute(n1, n2),
		graph.NewRoute(n2, n3),
		graph.NewRoute(n3, n4),
		graph.NewRoute(n1, n5),
		graph.NewRoute(n3, n6),
		graph.NewRoute(n4, n7),
		graph.NewRoute(n5, n6),
		graph.NewRoute(n6, n7),
		graph.NewRoute(n5, n8),
		graph.NewRoute(n6, n9),
		graph.NewRoute(n8, n9),
	}
	egGraph, err := graph.GenerateGraphWithNodesAndRoutes(nodes, routes)
	if err != nil {
		panic(err)
	}
	InstallSensor1(egGraph)
}
