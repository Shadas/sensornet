package graph

import (
	"fmt"
)

type SensorType int

const (
	SensorTypeInvalid SensorType = -1
	SensorTypeNone    SensorType = 0
	SensorTypeEntity  SensorType = 1
)

// 路径，用于连结各个结点
type Route struct {
	endpointA *Node
	endpointB *Node
	posA      SensorType // 0为默认值，默认为空；1 为实体传感器
	posB      SensorType
}

func NewRoute(epa, epb *Node) *Route {
	r := &Route{
		endpointA: epa,
		endpointB: epb,
	}
	return r
}

// 安装传感器到对应位置, pos 为A或B
func (r *Route) InstallSensor(pos string, st SensorType) {
	if pos == "A" {
		r.posA = st
	}
	if pos == "B" {
		r.posB = st
	}
}

//返回对应位置的传感器
func (r *Route) SensorWithPos(pos string) SensorType {
	switch pos {
	case "A":
		return r.posA
	case "B":
		return r.posB
	default:
		return SensorTypeInvalid
	}
}

// 展示route的端点结构
func (r *Route) ShowRoute() {
	sa, sb := "", ""
	if r.posA == SensorTypeEntity {
		sa = "*"
	}
	if r.posB == SensorTypeEntity {
		sb = "*"
	}
	fmt.Printf("%d%s -- %s%d\n", r.endpointA.Num, sa, sb, r.endpointB.Num)
}

// 判断此节点是否为该route的端点之一, 如果返回值为 ""，则不是，否则返回"A"或"B"
func (r *Route) EndpointPos(node *Node) string {
	if node == r.endpointA {
		return "A"
	} else if node == r.endpointB {
		return "B"
	} else {
		return ""
	}
}

// 返回路径的A节点
func (r *Route) NodeA() *Node {
	return r.endpointA
}

// 返回路径的B节点
func (r *Route) NodeB() *Node {
	return r.endpointB
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
