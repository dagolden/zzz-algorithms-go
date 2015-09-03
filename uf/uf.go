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
	return &QuickFind{make([]int, n), 0}
}

// Size returns the number of nodes
func (u QuickFind) Size() int {
	return len(u.nodes)
}
