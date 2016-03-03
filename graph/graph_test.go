// Copyright 2015 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package graph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/xdg/testy"
)

var tinyG = `
0 5
4 3
0 1
9 12
6 4
5 4
0 2
11 12
9 10
0 6
7 8
9 11
5 3
`

func TestUniGraph(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	// constructor
	g := NewUniGraph()

	is.Equal(g.V(), 0)
	is.Equal(g.E(), 0)

	// add edge
	g.AddEdge(1, 2)

	is.Equal(g.V(), 2)
	is.Equal(g.E(), 1)

	// add vertex
	g.AddVertex(0)
	is.Equal(g.V(), 3)
	is.Equal(g.E(), 1)

	// self loops shouldn't be added
	g.AddEdge(3, 3)
	is.Equal(g.V(), 4)
	is.Equal(g.E(), 1)

	// adjacency list
	g.AddEdge(3, 4)
	g.AddEdge(3, 5)
	g.AddEdge(3, 2)
	is.Equal(g.Adjacent(3), []int{4, 5, 2})

	// adjacency is read-only
	adj := g.Adjacent(3)
	adj[0] = 0
	is.Equal(g.Adjacent(3), []int{4, 5, 2})
}

func TestUniPaths(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	// build from Reader
	b := bufio.NewReader(strings.NewReader(tinyG))
	g, err := ReadUniGraph(b)
	is.Equal(err, nil)
	is.Equal(g.V(), 13)
	is.Equal(g.E(), 13)
	is.Equal(g.Adjacent(0), []int{5, 1, 2, 6})

	// test DFS
	p := g.DFS(0)
	is.True(p.HasPathTo(4))
	is.False(p.HasPathTo(8))
	if path, ok := p.PathTo(3); ok {
		is.Equal(path, []int{0, 5, 4, 3})
	} else {
		is.Error("PathTo(3) failed")
	}

	// test BFS
	p = g.BFS(0)
	is.True(p.HasPathTo(4))
	is.False(p.HasPathTo(8))
	if path, ok := p.PathTo(3); ok {
		is.Equal(path, []int{0, 5, 3})
	} else {
		is.Error("PathTo(3) failed")
	}
}
