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
	input := []string{
		"0,1",
		"1,2",
		"2,3",
		"2,4",
		"0,5",
		"1,6",
		"2,7",
		"5,6",
		"6,7",
		"5,8",
		"6,9",
		"7,10",
		"8,9",
		"9,10",
	}
	egGraph, err := graph.GenerateGraphWithRoutes(input)
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
	input := []string{
		"0,1",
		"1,2",
		"2,3",
		"3,4",
		"1,5",
		"3,6",
		"4,7",
		"5,6",
		"6,7",
		"5,8",
		"6,9",
		"8,9",
	}
	egGraph, err := graph.GenerateGraphWithRoutes(input)
	if err != nil {
		panic(err)
	}
	InstallSensor1(egGraph)
}
