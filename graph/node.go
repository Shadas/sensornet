package graph

// 定义节点结构
// 每一个节点包括：
//
type Node struct {
	Num    int `json:"num"`
	routes [4]*Route
}

func NewNode(num int) *Node {
	n := &Node{
		Num:    num,
		routes: [4]*Route{},
	}
	return n
}

func validateNode(node *Node) bool {
	if len(node.routes) != 4 {
		return false
	}
	return true
}
