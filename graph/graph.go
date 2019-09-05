package graph

import (
	"errors"
	"fmt"
)

// 一个被处理的网络被视为一个可连结的图
type Graph struct {
	// 这个图所包含的所有结点
	nodes []*Node
	// 这个图的邻接矩阵
	matrix [][]int
	// 从某个node开始引出的结点
	leadNode *Node
	// validate用的map
	validMap map[*Node]struct{}
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
		nodes:  nodes,
		matrix: matrix,
	}
	return
}

func GenerateGraph(root *Node) *Graph {
	g := &Graph{
		leadNode: root,
	}
	return g
}

// 通过起始结点验证一个node为起始结点的graph是否为合法的网络
func (g *Graph) validateGraphWithLeadNode() (b bool) {
	g.validMap = make(map[*Node]struct{})
	if g.leadNode == nil {
		fmt.Println("leadNode is nil")
		return false
	}
	return g.validateGraphCursion(g.leadNode)
}

func (g *Graph) validateGraphCursion(node *Node) (b bool) {
	if node == nil {
		fmt.Println("cursion node is nil")
		return false
	}
	g.validMap[node] = struct{}{}
	if !validateNode(node) {
		fmt.Println("not valid node")
		return false
	}
	for _, route := range node.routes {
		if route == nil {
			continue
		}
		if !route.validate(node) {
			fmt.Println("not valid route")
			return false
		}
		// 因为这里之前校验过route的端点，所以这里不用对另一个端点判断是否为nil
		on := route.theOtherNode(node)
		if _, ok := g.validMap[on]; ok {
			continue
		}
		b = g.validateGraphCursion(on)
		if !b {
			return b
		}
	}
	return true
}
