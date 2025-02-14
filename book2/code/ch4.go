package main

import (
	"bytes"
	"fmt"
	"strconv"
	"text/tabwriter"
)

const INFINITY = ^uint(0)

type Node struct {
	Name  string
	links []Edge
}

type Edge struct {
	from *Node
	to   *Node
	cost int
}

type Graph struct {
	nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{nodes: map[string]*Node{}}
}

func (g *Graph) AddNodes(names ...string) {
	for _, name := range names { //遍历所有的节点
		if _, ok := g.nodes[name]; !ok { //如果没有这个节点
			g.nodes[name] = &Node{Name: name, links: []Edge{}}
		}
	}
}

func (g *Graph) AddLink(a, b string, cost int) {
	aNode := g.nodes[a]
	bNode := g.nodes[b]
	aNode.links = append(aNode.links,
		Edge{from: aNode, to: bNode, cost: uint(cost)})
}

func (g *Graph) Dijkstra(source string) (map[string]uint,
	map[string]string) {
	dist, prev := map[string]uint{}, map[string]string{}

	for _, node := range g.nodes {
		dist[node.Name] = INFINITY
		prev[node.Name] = ""
	}
	visited := map[string]bool{}
	dist[source] = 0
	// 上述代码为初始化

	for u := source; u != ""; u = getClosestNonVisitedNode(
		dist, visited) {
		uDist := dist[u]
		for _, link := range g.nodes[u].links {
			if _, ok := visited[link.to.Name]; ok {
				continue
			}
			alt := uDist + link.cost
			v := link.to.Name
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
		visited[u] = true
	}
	return dist, prev

}
func getClosestNonVisitedNode(dist map[string]uint,
	visited map[string]bool) string {
	lowestCost := INFINITY
	lowestNode := ""
	for key, dis := range dist {
		if _, ok := visited[key]; dis == INFINITY || ok {
			continue
		}
		if dis < lowestCost {
			lowestCost = dis
			lowestNode = key
		}
	}
	return lowestNode
}
func DijkstraString(dist map[string]uint, prev map[string]string) string {
	buf := &bytes.Buffer{}
	writer := tabwriter.NewWriter(buf, 1, 5, 2, ' ', 0)
	writer.Write([]byte("Node\tDistance\tPrevious Node\t\n"))
	for key, value := range dist {
		writer.Write([]byte(key + "\t"))
		writer.Write([]byte(strconv.FormatUint(uint64(value), 10) + "\t"))
		writer.Write([]byte(prev[key] + "\t\n"))
	}
	writer.Flush()
	return buf.String()

}

func main() {
	g := NewGraph()
	g.AddNodes("a", "b", "c", "d", "e")
	g.AddLink("a", "b", 6)
	g.AddLink("d", "a", 1)
	g.AddLink("b", "e", 2)
	g.AddLink("b", "d", 1)
	g.AddLink("c", "e", 5)
	g.AddLink("c", "b", 5)
	g.AddLink("e", "d", 1)
	g.AddLink("e", "c", 4)
	dist, prev := g.Dijkstra("a")
	fmt.Println(DijkstraString(dist, prev))
}
