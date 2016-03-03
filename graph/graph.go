// Copyright 2015 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

// Package graph implements graph data structures and associated
// algorithms.
package graph

import (
	"bufio"
	"fmt"
	"io"
)

// UniGraph represents a uni-directional graph of vertices and edges
type UniGraph struct {
	vertCount int           // number of vertices
	edgeCount int           // number of edges
	adj       map[int][]int // adjacency lists
}

func NewUniGraph() *UniGraph {
	adj := make(map[int][]int)
	return &UniGraph{0, 0, adj}
}

func ReadUniGraph(r io.Reader) (*UniGraph, error) {
	g := NewUniGraph()
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var v, w int
		fmt.Sscanf(scanner.Text(), "%d %d", &v, &w)
		g.AddEdge(v, w)
	}
	return g, scanner.Err()
}

func (g *UniGraph) V() int {
	return g.vertCount
}

func (g *UniGraph) E() int {
	return g.edgeCount
}

func (g *UniGraph) AddEdge(v, w int) {
	g.AddVertex(v)

	// no self-loops
	if v == w {
		return
	}

	g.AddVertex(w)
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.edgeCount++
	return
}

func (g *UniGraph) AddVertex(v int) {
	if _, ok := g.adj[v]; !ok {
		g.adj[v] = []int{}
		g.vertCount++
	}
	return
}

func (g *UniGraph) Adjacent(v int) []int {
	// copy so original stays immutable
	adj := append([]int{}, g.adj[v]...)
	return adj
}

func (g *UniGraph) DFS(s int) Paths {
	return g.path(s, true)
}

func (g *UniGraph) BFS(s int) Paths {
	return g.path(s, false)
}

func (g *UniGraph) path(s int, DFS bool) Paths {
	marked := make([]bool, g.vertCount)
	edgeTo := make([]int, g.vertCount)
	queue := []int{s}

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		marked[v] = true
		next := []int{}
		for _, w := range g.Adjacent(v) {
			if !marked[w] {
				edgeTo[w] = v
				if DFS {
					// recurse: go to w, then restart with v
					next = append(next, w, v)
					break
				} else {
					next = append(next, w)
					marked[w] = true
				}
			}
		}
		if DFS {
			queue = append(next, queue...)
		} else {
			queue = append(queue, next...)
		}
	}

	return NewPaths(s, marked, edgeTo)
}

// Paths represents a set of paths from a source vertex
type Paths struct {
	s      int
	marked []bool
	edgeTo []int
}

func NewPaths(s int, m []bool, e []int) Paths {
	return Paths{s, m, e}
}

func (p Paths) HasPathTo(v int) bool {
	return p.marked[v]
}

func (p Paths) PathTo(v int) ([]int, bool) {
	path := []int{}

	if !p.HasPathTo(v) {
		return path, false
	}

	for x := v; x != p.s; x = p.edgeTo[x] {
		path = append(path, x)
	}
	path = append(path, p.s)

	// path built up from end, so reverse it
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path, true
}
