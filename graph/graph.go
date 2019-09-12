package graph

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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

// 根据连接关系，生成一张图，简化图的初始化输入
// 每一条route的写法为用逗号分隔的字符串，比如 "0,1" 代表连接 节点 0，1 的一条路径
func GenerateGraphWithRoutes(rs []string) (g *Graph, err error) {
	var (
		nodeMap   = make(map[int]*Node)
		nodeList  = []*Node{}
		routeList = []*Route{}
	)
	for _, r := range rs {
		var (
			na, nb       int64
			kna, knb     int
			ns           = strings.Split(r, ",")
			nodea, nodeb *Node
			ok           bool
		)
		if len(ns) != 2 {
			err = errors.New(fmt.Sprintf("Invalid input route with %s", r))
			return
		}
		if na, err = strconv.ParseInt(ns[0], 10, 64); err != nil {
			return
		}
		kna = int(na)
		if nb, err = strconv.ParseInt(ns[1], 10, 64); err != nil {
			return
		}
		knb = int(nb)
		if nodea, ok = nodeMap[kna]; !ok {
			nodea = &Node{Num: kna}
			nodeMap[kna] = nodea
		}
		if nodeb, ok = nodeMap[knb]; !ok {
			nodeb = &Node{Num: knb}
			nodeMap[knb] = nodeb
		}
		routeList = append(routeList, NewRoute(nodea, nodeb))
	}
	// 去重后获取节点列表,需要保证顺序，节点编号与下标一致
	for i := 0; i < len(nodeMap); i++ {
		node, ok := nodeMap[i]
		if !ok {
			panic(fmt.Sprintf("err node num with %d", i))
		}
		nodeList = append(nodeList, node)
	}
	// 检查生成的所有点
	checkGenNodeListByNodePairStr(nodeList)
	// 检查生成的所有管道
	// checkGenRouteListByNodePairStr(routeList)
	return GenerateGraphWithNodesAndRoutes(nodeList, routeList)
}

// 检查输入的node结对分析后的nodeList结果
func checkGenNodeListByNodePairStr(nl []*Node) (err error) {
	fmt.Println("检查所有输入的点:")
	count := len(nl)
	nums := []int{}
	for _, n := range nl {
		nums = append(nums, n.Num)
	}
CheckIdx:
	for i := 0; i < count; i++ {
		for _, n := range nums {
			if n == i {
				continue CheckIdx
			}
		}
		err = errors.New(fmt.Sprintf("err node with %d", i))
		return
	}
	fmt.Printf("检查所有输入点succ，共有 %d 个点 \n", count)
	return
}

// 根据结点列表和连接关系列表，生成一张图
func GenerateGraphWithNodesAndRoutes(nodesOri []*Node, routes []*Route) (g *Graph, err error) {
	if len(nodesOri) == 0 {
		err = errors.New("nodes can't be empty.")
		return
	}
	var (
		matrix    [][]int
		nodeCount = len(nodesOri)
		nodes     []*Node
	)
	// 保证节点编号和下标一致有序
NodeCountLoop:
	for i := 0; i < nodeCount; i++ {
		for _, n := range nodesOri {
			if n.Num == i {
				nodes = append(nodes, n)
				continue NodeCountLoop
			}
		}
		panic(fmt.Sprintf("not found node num with %d", i))
	}

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
	// g.ShowMatrix()
	return
}

// 打印该图的邻接矩阵
func (g *Graph) ShowMatrix() {
	firstLine := []int{0}
	for i := 0; i < g.NodeCount; i++ {
		firstLine = append(firstLine, i)
	}
	firstStr := ""
	for _, i := range firstLine {
		firstStr += fmt.Sprintf("%.2d ", i)
	}
	fmt.Println(firstStr)
	for i, line := range g.matrix {
		tmpStr := ""
		for _, num := range append([]int{i}, line...) {
			tmpStr += fmt.Sprintf("%.2d ", num)
		}
		fmt.Println(tmpStr)
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
