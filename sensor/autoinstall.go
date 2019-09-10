package sensor

import (
	"fmt"

	"github.com/sensornet/graph"
)

// 安装传感器的算法1
func InstallSensor1(g *graph.Graph) {
	var (
		// 末端节点
		endPointsList = []*graph.Node{}
		// 连通节点
		connectedPointsList = []*graph.Node{}
		// 普通3路节点
		normal3PointsList = []*graph.Node{}
		// 普通4路节点
		normal4PointsList = []*graph.Node{}
		// 末端管道
		endRoutes = []*graph.Route{}
		// 连通管道
		connectedRoutes = []*graph.Route{}
		// 普通管道
		normalRoutes = []*graph.Route{}
	)
	// 获取所有末端节点，末端节点在邻接矩阵的行和列之和为1
	// 获取所有连通节点，末端节点在邻接矩阵的行和列之和为2
	for i := 0; i < g.NodeCount; i++ {
		var lineCount, columnCount int
		// 计算行的和
		for _, pos := range g.Matrix()[i] {
			lineCount += pos
		}
		// 计算列的和
		for _, line := range g.Matrix() {
			columnCount += line[i]
		}
		if lineCount == 1 && columnCount == 1 {
			endPointsList = append(endPointsList, g.Nodes()[i])
		}
		if lineCount == 2 && columnCount == 2 {
			connectedPointsList = append(connectedPointsList, g.Nodes()[i])
		}
		if lineCount == 3 && columnCount == 3 {
			normal3PointsList = append(normal3PointsList, g.Nodes()[i])
		}
		if lineCount == 4 && columnCount == 4 {
			normal4PointsList = append(normal4PointsList, g.Nodes()[i])
		}
	}
routeloop:
	for _, route := range g.OriRoute() {
		// 如果route的一端是末端节点，则该route为末端管道
		for _, ep := range endPointsList {
			if route.EndpointPos(ep) != "" {
				route.InstallSensor("A", graph.SensorTypeEntity)
				route.InstallSensor("B", graph.SensorTypeEntity)
				endRoutes = append(endRoutes, route)
				continue routeloop
			}
		}
		// 如果route的一端是连通节点，且该route不为末端管道，则该route为连通管道
		for _, ep := range connectedPointsList {
			if route.EndpointPos(ep) != "" {
				connectedRoutes = append(connectedRoutes, route)
				continue routeloop
			}
		}
		// 剩下的是普通管道
		normalRoutes = append(normalRoutes, route)
	}
	// 处理连通管道
	for _, r := range connectedRoutes {
		isAConnected, isBConnected := false, false
		for _, ep := range connectedPointsList {
			pos := r.EndpointPos(ep)
			if pos == "A" {
				isAConnected = true
			}
			if pos == "B" {
				isBConnected = true
			}
		}
		if isAConnected && isBConnected {
			// 如果两端都是连通节点，则不放置传感器
			continue
		}
		if isAConnected {
			r.InstallSensor("B", graph.SensorTypeEntity)
		}
		if isBConnected {
			r.InstallSensor("A", graph.SensorTypeEntity)
		}
	}
	// 遍历所有3路普通节点
	for _, p := range normal3PointsList {
		var (
			already int
			routes  []*graph.Route
		)
		for _, r := range g.OriRoute() {
			switch r.SensorWithPos(r.EndpointPos(p)) {
			case graph.SensorTypeEntity:
				already++
			case graph.SensorTypeNone:
				routes = append(routes, r)
			case graph.SensorTypeInvalid:
				continue
			}
		}
		for already < 2 && len(routes) > 0 {
			r := routes[0]
			pos := r.EndpointPos(p)
			if pos == "A" {
				r.InstallSensor("A", graph.SensorTypeEntity)
			}
			if pos == "B" {
				r.InstallSensor("B", graph.SensorTypeEntity)
			}
			routes = routes[1:]
			already++
		}
	}
	// 遍历所有4路普通节点
	for _, p := range normal4PointsList {
		var (
			already int
			routes  []*graph.Route
		)
		for _, r := range g.OriRoute() {
			switch r.SensorWithPos(r.EndpointPos(p)) {
			case graph.SensorTypeEntity:
				already++
			case graph.SensorTypeNone:
				routes = append(routes, r)
			case graph.SensorTypeInvalid:
				continue
			}
		}
		for already < 3 && len(routes) > 0 {
			r := routes[0]
			pos := r.EndpointPos(p)
			if pos == "A" {
				r.InstallSensor("A", graph.SensorTypeEntity)
			}
			if pos == "B" {
				r.InstallSensor("B", graph.SensorTypeEntity)
			}
			routes = routes[1:]
			already++
		}
	}
	fmt.Println("末端节点：")
	for _, p := range endPointsList {
		fmt.Println(p.Num)
	}
	fmt.Println("连通节点：")
	for _, p := range connectedPointsList {
		fmt.Println(p.Num)
	}
	fmt.Println("普通3路节点：")
	for _, p := range normal3PointsList {
		fmt.Println(p.Num)
	}
	fmt.Println("普通4路节点：")
	for _, p := range normal4PointsList {
		fmt.Println(p.Num)
	}
	fmt.Println("末端管道：")
	for _, r := range endRoutes {
		r.ShowRoute()
	}
	fmt.Println("连通管道：")
	for _, r := range connectedRoutes {
		r.ShowRoute()
	}
	fmt.Println("普通管道：")
	for _, r := range normalRoutes {
		r.ShowRoute()
	}
}
