package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

type Edge struct {
    source int
    destination int
    weight int
}

func main() {
    var numNodes int
    var edges []Edge
    var choice int

    fmt.Print("Enter 1 to enter the graph manually or otherkey to use the default graph: ")
    fmt.Scan(&choice)

    switch choice {
    case 1:
        var numNodes int
        fmt.Print("Enter the number of nodes: ")
        fmt.Scan(&numNodes)

        var numEdges int
        fmt.Print("Enter the number of edges: ")
        fmt.Scan(&numEdges)

        edges := make([]Edge, numEdges)
        for i := 0; i < numEdges; i++ {
            fmt.Printf("Enter edge %d (from, to, weight): ", i+1)
            var from, to, weight int
            _, err := fmt.Scan(&from, &to, &weight)
            if err != nil {
                fmt.Println(err)
                return
            }    
            edges[i] = Edge{from, to, weight}
        }
    default:
        numNodes = 5
        edges = []Edge{
		{0, 1, 4},
		{0, 2, 1},
		{1, 3, 1},
		{2, 1, 2},
		{2, 3, 5},
		{3, 4, 3},
	}
    }

    startNode := 0
    shortestPaths, predecessors := bellmanFord(edges, numNodes, startNode)

    if shortestPaths != nil {
		fmt.Println("Shortest paths from node", startNode, ":")
		for node, distance := range shortestPaths {
			fmt.Printf("Node %d -> Distance: %d\n", node, distance)
		}
	}
    
    visualize(shortestPaths,edges,predecessors)
}

func visualize(shortestPaths map[int]int,edges []Edge,predecessors map[int]int) {
    g := graph.New(graph.IntHash, graph.Directed())

    for _, edge := range edges {
        color := "black"
        if predecessors[edge.destination] == edge.source {
            color = "red"
        }
        _ = g.AddVertex(edge.source)
        _ = g.AddVertex(edge.destination)
        _ = g.AddEdge(edge.source, edge.destination, graph.EdgeAttribute("label", strconv.Itoa(edge.weight)), graph.EdgeAttribute("color", color), graph.EdgeWeight(edge.weight))
    }


    file, _ := os.Create("sp-graph.gv")
    _ = draw.DOT(g, file)
}

func bellmanFord(edges []Edge, numNodes, startNode int) (map[int]int, map[int]int) {
    distances := make(map[int]int)
    predecessors := make(map[int]int)

    for i := 0; i < numNodes; i++ {
		distances[i] = math.MaxInt64
	}

    distances[0] = 0

    for i := 0; i < numNodes-1; i++ {
        for _, edge := range edges {
            if  distances[edge.source] != math.MaxInt64 && distances[edge.source] + edge.weight < distances[edge.destination] {
                distances[edge.destination] = distances[edge.source] + edge.weight
                predecessors[edge.destination] = edge.source
            }
        }
    }

    for _, edge := range edges {
		if distances[edge.source] != math.MaxInt64 && distances[edge.source]+edge.weight < distances[edge.destination] {
			fmt.Println("The graph contains a cycle with negative weight.")
			return nil, nil
		}
	}


    return distances, predecessors
}