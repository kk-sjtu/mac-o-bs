第四章 使用内置包

4.2 Dijkstra 路径查找算法

Dijkstra算法是一种贪心算法，用于计算图中的最短路径。它的基本思想是从起点开始，逐步扩展到其他节点，直到到达终点。在扩展的过程中，每次选择距离起点最近的节点，然后更新与该节点相邻的节点的距离。这样，直到所有节点都被访问过，就得到了起点到终点的最短路径。
注意是，单源最短路径。单源。
![4.2.png](../img/4.2.png)

![4.3 init.png](../img/4.3%20init.png)

对每个节点（结构体）的后驱节点进行遍历，然后计算开销。

先定义结构体

```go
type Node struct{
    Name string
    links []Edge
}

type Edge struct{
    from *Node
    to *Node
    cost int
}

type Graph struct{
    nodes map[string] *Node
}

```

还定义了一个辅助函数，创建一个新的图结构体实例。

```go
func NewGraph() *Graph{
    return &Graph{nodes:map[string]*Node{}}
}
```

需要先编写一些基础的结构，方便用户能够与实际的路径查找算法进行交互。这意味着必须能够获取图。

```go
func (g *Graph) AddNodes(names ...string){
    for _,name := range names{ //遍历所有的节点
        if _,ok := g.nodes[name];!ok{//如果没有这个节点
            g.nodes[name] = &Node{Name:name,links:[]Edghe{}}
            }
        }
}
```

```go
func (g *Graph)AddLink(a,b string,cost int){
	aNode := g.nodes[a]
	bNode := g.nodes[b]
	aNode.links = append(aNode.links,
		Edge{from:aNode,to:bNode,cost:uint(cost)})
}
```

前者用于添加没有连接的新节点，另一个用于向现有节点添加连接。

下面代码用来计算开销和父节点
![@蓝不过还啊 Dijkstra算法.png](../img/%40%E8%93%9D%E4%B8%8D%E8%BF%87%E8%BF%98%E5%95%8A%20Dijkstra%E7%AE%97%E6%B3%95.png)
```go
func (g *Graph) Dijkstra(source string)(map[string]uint,
	map[string]string){
	dist,prev := map[string]uint{},map[string]string{}
	
	for _,node := range g.nodes{
		dist[node.Name] = INFINITY
		prev[node.Name] = ""
    }
	visited := map[string]bool{}
	dist[source] = 0
	// 上述代码为初始化
	
	for u:= source;u!="";u=getClosestNonVisitedNode(
		dist,visited){
		uDist := dist[u]
		for _,link := range g.nodes[u].links{
			if _,ok := visited[link.to.Name];ok{
				continue
			}
			alt := uDist + link.cost
			v := link.to.Name
			if alt<dist[v]{
				dist[v] = alt
				prev[v] = u
            }   
		}
		visited[u] = true
    }
	return dist,prev
	
}

```
这被实现为Graph结构体的接收器。在其中
```go
    dist,prev := map[string]uint{},map[string]string{}
	
	for _,node := range g.nodes{
		dist[node.Name] = INFINITY
		prev[node.Name] = ""
    }
	visited := map[string]bool{}
	
	dist[source]=0
```
第一行创建开销和记录父节点的字典。对开销和父节点进行初始化，并把源头的开销设置为0

下一部分代码
```go
for 
```

下面是取最近的没有visit的节点
```go
func getClosestNonVisitedNode(dist map[string]uint,
	visited map[string]bool) string{
	lowestCost := INFINITY
	lowestNode := ""
	for key,dis := range dist{
        if _,ok:= visited[key];dis == INFINITY || ok{
		    continue
	        }
	    if dis < lowestCost{
			lowestCost = dis
			lowestNode = key
        }	
    }
	return lowestNode
}

```
第二个函数可以获取告诉我们节点开销的字典。以及告诉我们是否访问过节点的映射，并找出开销最低且尚未访问过的节点。

```go
for u:=source;u!="";u = getClosestNonVisitedNode(dist,visited){
	uDist := dist[u]
}
```

逻辑是，当该函数返回空字符串时，表示已经访问了所有节点。这个函数就是为了寻找，u的最近点？

```go
for _,link := range g.nodes[u].links{
	if _,ok := visited[link.to.Name];ok{
    continue    
	}
	alt := uDist + link.cost
	v := link.to.Name
	if alt<dist[v]{
        dist[v] = alt
		prev[v] = u
	}
}
visited[u] = true
}
```
负责遍历该节点的输出连接，并运行迪杰斯特拉迭代。

最后是制表函数
```go
func DijkstraString(dist map[string]uint, prev map[string] string)string {
    buf := &bytes.Buffer{}
    writer := tabwriter.NewWriter(buf,1,5,2,'',0)writer.Write([]byte("Node\tDistance\tPrevious Node\t\n"))
	for key, value := range dist{
        writer.Write([]byte(key + "\t"))
		writer.Write([]byte(strconv.FormatUint(uint64(value),10)+"\t"))
        writer.Write([]byte(prev[key]+"\t\n"))
	}
	writer.Flush()
    return buf.string()
	
}
```



4.3 康威生命游戏

游戏规则：
1. 如果该细胞存活，
   1. 当有少于2个存活邻居，则死亡
   2. 有2-3个存活邻居，则存活
   3. 有超过3个存活邻居，则死亡
2. 如果该细胞死亡，当有3个存活邻居时，存活。


![细胞自动机.png](../img/%E7%BB%86%E8%83%9E%E8%87%AA%E5%8A%A8%E6%9C%BA.png)

具体代码，见ch4_conway_game

4.4 工作量证明

在区块链中，目标是验证交易块的真实性，以确定任何特定钱包地址中有多少货币。可以通过为了使得区块合法化而投入了多少工作量来验证区块真实性。

工作量证明的目标是找到一个特定的哈希值，该哈希值满足一定的条件。这个条件是哈希值的前几位必须是0。这个条件是可调的，可以根据需要调整。这个条件越高，找到满足条件的哈希值的难度就越大。

用例子举例，比如要找到一个哈希值，使得前两位是0。这个哈希值是一个随机数，所以我们可以通过不断尝试不同的随机数来找到满足条件的哈希值。这个过程是一个耗时的过程，因为我们不知道哪个随机数会满足条件。这就是工作量证明的基本原理。
以下面字符串为例
"Hello, World!"
SHA哈希可表示为xxx，而我们想，找到一个特定字节序列加在hello world后面，使得SHA哈希值的前几位是0。GO语言1秒能处理400万个哈希随机。

```go

