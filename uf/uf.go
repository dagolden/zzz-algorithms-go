// Copyright 2015 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

// Implement union-find algorithms
package uf

// Common methods for union-find implementations
type UFinder interface {
	Union(p, q int)
	Find(p int) int
	Connected(p, q int) bool
	Count() int
	Size() int
}

// QuickFind implements union-finding with quick-find algorithm
type QuickFind struct {
	nodes []int
	count int
}

// NewQuickFind returns a QuickFind initialized for the given number of
// nodes.  If 'n' is negative, nil is returned.
func NewQuickFind(n int) *QuickFind {
	if n < 1 {
		return nil
	}
	nl := make([]int, n)
	for i, _ := range nl {
		nl[i] = i
	}
	return &QuickFind{nl, n}
}

// Size returns the number of nodes
func (u *QuickFind) Size() int {
	return len(u.nodes)
}

// Count returns the number of components
func (u *QuickFind) Count() int {
	return u.count
}

// Find returns component for node n
func (u *QuickFind) Find(n int) int {
	return u.nodes[n]
}

// Connected returns true if p and q are connected and false otherwise
func (u *QuickFind) Connected(q, p int) bool {
	return u.Find(p) == u.Find(q)
}

// Union joins p and q into a single component
func (u *QuickFind) Union(p, q int) {
	pID, qID := u.Find(p), u.Find(q)

	// already connected
	if pID == qID {
		return
	}

	// change pID components to qID
	for i, v := range u.nodes {
		if v == pID {
			u.nodes[i] = qID
		}
	}

	// one less component now
	u.count--

	return
}
