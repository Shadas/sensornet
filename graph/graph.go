package graph

import "fmt"

// 一个被处理的网络被视为一个可连结的图
type Graph struct {
	// 这个图所包含的所有结点
	nodes []*Node
	// 从某个node开始引出的结点
	leadNode *Node
	// validate用的map
	validMap map[*Node]struct{}
}

func GenerateGraphWithNodePairs() *Graph {
	return &Graph{}
}

func GenerateGraph(root *Node) *Graph {
	g := &Graph{
		leadNode: root,
	}
	return g
}

// 验证一个node为起始结点的graph是否为合法的网络
func (g *Graph) validateGraph() (b bool) {
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
