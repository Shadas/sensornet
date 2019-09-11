package graph

// 定义节点结构
// 每一个节点包括：
//
type Node struct {
	Num int `json:"num"`
}

func NewNode(num int) *Node {
	n := &Node{
		Num: num,
	}
	return n
}
