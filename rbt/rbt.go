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
		t.root = &node{key: k, val: v}
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

type node struct {
	key   string
	val   int
	nodes int
	isRed bool
	left  *node
	right *node
}

func (n *node) Get(k string) (int, error) {
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

func (n *node) Put(k string, v int) error {
	if k < n.key {
		if n.left == nil {
			n.left = &node{key: k, val: v}
			return nil
		}
		return n.left.Put(k, v)
	}
	if k > n.key {
		if n.right == nil {
			n.right = &node{key: k, val: v}
			return nil
		}
		return n.right.Put(k, v)
	}
	n.val = v
	return nil
}
