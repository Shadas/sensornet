package graph

import (
	"errors"
	"fmt"
)

// 一个被处理的网络被视为一个可连结的图
type Graph struct {
	NodeCount int
	// 这个图所包含的所有结点
	nodes []*Node
	// 这个图的邻接矩阵
	matrix [][]int
	// 原始管道
	oriRoutes []*Route
}

// 根据结点列表和连接关系列表，生成一张图
func GenerateGraphWithNodesAndRoutes(nodes []*Node, routes []*Route) (g *Graph, err error) {
	if len(nodes) == 0 {
		err = errors.New("nodes can't be empty.")
		return
	}
	var (
		matrix [][]int
	)
	// 初始化邻接矩阵
	for i := 0; i < len(nodes); i++ {
		line := []int{}
		for j := 0; j < len(nodes); j++ {
			line = append(line, 0)
		}
		matrix = append(matrix, line)
	}
	// 填充邻接矩阵
	for idx, r := range routes {
		if !r.validate(nil) {
			err = errors.New(fmt.Sprintf("invalid route with %d route.", idx))
			return
		}
		x, y := r.endpointA.Num, r.endpointB.Num
		matrix[x][y], matrix[y][x] = 1, 1
	}
	g = &Graph{
		nodes:     nodes,
		matrix:    matrix,
		NodeCount: len(nodes),
		oriRoutes: routes,
	}
	return
}

// 打印该图的邻接矩阵
func (g *Graph) ShowMatrix() {
	for _, line := range g.matrix {
		fmt.Println(line)
	}
}

// 获取该图的邻接矩阵
func (g *Graph) Matrix() [][]int {
	return g.matrix
}

// 获取该图的节点列表
func (g *Graph) Nodes() []*Node {
	return g.nodes
}

// 获取该图的原版管道信息
func (g *Graph) OriRoute() []*Route {
	return g.oriRoutes
}
