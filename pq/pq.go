// Copyright 2015 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

// Package pq implements a priority queue
package pq

import (
	"fmt"
)

type Comparator func(interface{}, interface{}) bool

type Queue struct {
	inOrder Comparator
	heap    []interface{}
}

func New(c Comparator) *Queue {
	return &Queue{c, make([]interface{}, 1)}
}

func (q Queue) String() string {
	return fmt.Sprint(q.heap)
}

func (q Queue) Size() int {
	return len(q.heap) - 1
}

func (q Queue) Peek() interface{} {
	if len(q.heap) == 1 {
		return nil
	}
	return q.heap[1]
}

func (q *Queue) Push(v interface{}) {
	q.heap = append(q.heap, v)
	q.swim(q.Size())
	return
}

func (q *Queue) Pop() (v interface{}) {
	n := q.Size()
	v = q.heap[1]
	q.heap[1], q.heap[n] = q.heap[n], q.heap[1]
	q.heap = q.heap[0:n]
	q.sink(1)
	return
}

func (q *Queue) sink(k int) {
	n := q.Size()
	for 2*k <= n {
		j := 2 * k
		if j < n && !q.inOrder(q.heap[j], q.heap[j+1]) {
			j++
		}
		if q.inOrder(q.heap[k], q.heap[j]) {
			break
		}
		q.heap[j], q.heap[k] = q.heap[k], q.heap[j]
		k = j
	}
	return
}

func (q *Queue) swim(k int) {
	for k > 1 && !q.inOrder(q.heap[k/2], q.heap[k]) {
		q.heap[k/2], q.heap[k] = q.heap[k], q.heap[k/2]
		k /= 2
	}
	return
}
