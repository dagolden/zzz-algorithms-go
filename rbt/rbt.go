// Copyright 2015 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

// Package rbt implements a red-black binary tree
package rbt

import (
	"fmt"
)

type tree struct {
	root *node
}

func New() (*tree, error) {
	return &tree{}, nil
}

func (t *tree) Put(k string, v int) error {
	if t.root == nil {
		t.root = &node{key: k, val: v, size: 1}
		return nil
	}
	return t.root.Put(k, v)
}

func (t tree) Get(k string) (int, error) {
	if t.root == nil {
		return 0, fmt.Errorf("Key %s was not found", k)
	}
	return t.root.Get(k)
}

func (t tree) Size() int {
	if t.root == nil {
		return 0
	}
	return sizeOf(t.root)
}

func (t tree) Iterator() *treeIterator {

	it := &treeIterator{make([]*node, 0)}
	it.descendLeft(t.root)

	return it
}

type node struct {
	key   string
	val   int
	size  int
	isRed bool
	left  *node
	right *node
}

func (n node) Get(k string) (int, error) {
	if k == n.key {
		return n.val, nil
	}
	if k < n.key {
		if n.left == nil {
			return 0, fmt.Errorf("Key %s was not found", k)
		}
		return n.left.Get(k)
	}
	if k > n.key {
		if n.right == nil {
			return 0, fmt.Errorf("Key %s was not found", k)
		}
		return n.right.Get(k)
	}
	return 0, fmt.Errorf("Key %s was not found", k)
}

func (n *node) Put(k string, v int) (err error) {
	switch {
	case k < n.key:
		if n.left == nil {
			n.left = &node{key: k, val: v, size: 1}
		} else {
			err = n.left.Put(k, v)
		}
	case k > n.key:
		if n.right == nil {
			n.right = &node{key: k, val: v, size: 1}
		} else {
			err = n.right.Put(k, v)
		}
	default:
		n.val = v
	}
	n.size = sizeOf(n.left) + sizeOf(n.right) + 1
	return nil
}

// sizeOf is conditional; it returns zero for a nil pointer
func sizeOf(n *node) int {
	if n == nil {
		return 0
	}
	return n.size
}

type treeIterator struct {
	stack []*node
}

func (it *treeIterator) descendLeft(current *node) {
	for current != nil {
		it.stack = append(it.stack, current)
		current = current.left
	}
	return
}

func (it treeIterator) HasNext() bool {
	return len(it.stack) != 0
}

func (it *treeIterator) Next() (string, int) {
	last := len(it.stack) - 1
	n := it.stack[last]
	it.stack = it.stack[:last]
	it.descendLeft(n.right)
	return n.key, n.val
}
