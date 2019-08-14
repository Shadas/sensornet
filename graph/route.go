package graph

// 路径，用于连结各个结点
type Route struct {
	endpointA *Node
	endpointB *Node
}

func (r *Route) validate(node *Node) bool {
	// route 的端点不应该为空
	if r.endpointA == nil || r.endpointB == nil {
		return false
	}
	// route不应该是一个圈
	if r.endpointA == r.endpointB {
		return false
	}
	// 如果检查带着非空node，则检查此node是否为该路径的端点之一
	if node != nil {
		if node != r.endpointA && node != r.endpointB {
			return false
		}
	}
	return true
}

func (r *Route) theOtherNode(node *Node) *Node {
	if r.endpointA == node {
		return r.endpointB
	}
	if r.endpointB == node {
		return r.endpointA
	}
	return nil
}
