package graph

import (
	"testing"
)

var (
	InvalidNodes []*Node
	ValidNodes   []*Node
)

func initData() {
	genInvalidNodes()
	genValidNodes()
}

func genValidNodes() {
	// 单点node
	n1 := &Node{
		routes: [4]*Route{},
	}
	ValidNodes = append(ValidNodes, n1)
	// 双点node
	n2 := &Node{}
	n2a := &Node{}
	r2 := &Route{n2, n2a}
	n2.routes = [4]*Route{
		r2,
	}
	n2a.routes = [4]*Route{
		r2,
	}
	ValidNodes = append(ValidNodes, n2)
}

func genInvalidNodes() {
	// 有route的一端为空的，不应通过测试
	n1 := &Node{}
	n1.routes = [4]*Route{
		&Route{n1, nil},
	}
	// 有route为回环的
	n2 := &Node{}
	n2.routes = [4]*Route{
		&Route{n2, n2},
	}
	InvalidNodes = append(InvalidNodes, n1)
}

func TestGraphValidate(t *testing.T) {
	initData()
	testGraphValidateWithInvalidCases(t)
	testGraphValidateWithValidCases(t)
}

func testGraphValidateWithValidCases(t *testing.T) {
	for idx, vn := range ValidNodes {
		g := GenerateGraph(vn)
		if !g.validateGraphWithLeadNode() {
			t.Errorf("%d node validate should be true", idx)
			break
		}
	}
}

func testGraphValidateWithInvalidCases(t *testing.T) {
	for idx, in := range InvalidNodes {
		g := GenerateGraph(in)
		if g.validateGraphWithLeadNode() {
			t.Errorf("%d node validate should be wrong", idx)
		}
	}
}
