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
示意图：可见 ./examplepic/example2.jpeg
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

/*
示意图：可见 ./examplepic/example3.jpeg
*/
func ExampleGraph3() {
	input := []string{
		"0,1",
		"1,2",
		"2,3",
		"3,4",
		"4,5",
		"0,15",
		"3,20",
		"4,18",
		"5,6",
		"15,24",
		"24,23",
		"23,20",
		"20,18",
		"18,6",
		"15,14",
		"24,25",
		"23,22",
		"20,21",
		"6,7",
		"14,25",
		"25,22",
		"22,21",
		"19,7",
		"14,13",
		"22,11",
		"21,10",
		"19,9",
		"7,8",
		"13,12",
		"12,11",
		"11,10",
		"10,9",
		"9,8",
		"12,16",
		"10,17",
	}
	egGraph, err := graph.GenerateGraphWithRoutes(input)
	if err != nil {
		panic(err)
	}
	InstallSensor1(egGraph)
}
