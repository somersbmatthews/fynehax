package graph

type graphRenderer struct {
	graph *GraphWidget
}

type GraphWidget struct {
	Nodes map[string]GraphNode
	// Edges map[string]GraphEdge
}
