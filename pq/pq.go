// Copyright 2015 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

// Package pq implements a priority queue
package pq

import (
	"errors"
	"fmt"
	"sort"
)

// The HeapSlice defines an interface that a slice type must implement
// to be used as the storage element of a priority queue
type HeapSlice interface {
	// sort.Interface defines methods to test and manipulate elements
	// Note: the "lesser" element will be the first returned from the
	// priority queue.
	sort.Interface
	// Push appends an element to the end of the slice, doing whatever
	// cast is necessary.
	Push(v interface{})
	// Pop removes an element from the end of the slice, leaving its
	// length one less than before.
	Pop() interface{}
	// Peek returns the first element of the slice.
	Peek() interface{}
}

// Queue is a priority queue with an unlimited number of elements
type Queue struct {
	heap HeapSlice
}

// New returns a new, unbounded priority queue.  The heap argument
// must be an empty slice of a type that implements HeapSlice
func New(heap HeapSlice) (*Queue, error) {
	// XXX eventually, heapify the slice if it has elements
	if heap.Len() > 0 {
		return nil, errors.New("slice argument to New() was not empty")
	}
	return &Queue{heap}, nil
}

func (q Queue) Len() int {
	return q.heap.Len()
}

func (q Queue) String() string {
	return fmt.Sprint(q.heap)
}

func (q Queue) Peek() interface{} {
	if q.Len() == 0 {
		return nil
	}
	return q.heap.Peek()
}

func (q *Queue) Push(v interface{}) {
	q.heap.Push(v)
	q.swim(q.Len() - 1)
	return
}

func (q *Queue) Pop() interface{} {
	n := q.Len()
	q.heap.Swap(0, n-1)
	v := q.heap.Pop()
	q.sink(0)
	return v
}

func (q *Queue) sink(k int) {
	n := q.Len()
	for 2*k+2 <= n-1 {
		j := 2*k + 1
		if j < n-1 && !q.heap.Less(j, j+1) {
			j++
		}
		if q.heap.Less(k, j) {
			break
		}
		q.heap.Swap(j, k)
		k = j
	}
	return
}

func (q *Queue) swim(k int) {
	up := (k - 1) / 2
	for k > 0 && !q.heap.Less(up, k) {
		q.heap.Swap(up, k)
		k = up
		up = (up - 1) / 2
	}
	return
}

// TopQueue is a priority queue with an unlimited number of elements
type TopQueue struct {
	Queue
	cap int
}

// NewTop returns a new priority queue with a maximum capacity.  The heap
// argument must be an empty slice of a type that implements HeapSlice
func NewTop(heap HeapSlice, capacity int) (*TopQueue, error) {
	// XXX eventually, heapify the slice if it has elements
	if heap.Len() > 0 {
		return nil, errors.New("slice argument to New() was not empty")
	}
	return &TopQueue{heap, cap}, nil
}

func (q *TopQueue) Push(v interface{}) {
	q.heap.Push(v)
	n := q.heap.Len()
	q.swim(n - 1)
	// XXX wrong: lowest priority is not guaranteed to be last
	if n > q.cap {
		q.heap.Pop()
	}
	return
}

// IntHeap is a slice of integers that implements HeapSlice
type IntHeap []int

func (s IntHeap) Len() int           { return len(s) }
func (s IntHeap) Less(i, j int) bool { return s[i] < s[j] }
func (s IntHeap) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s *IntHeap) Push(v interface{}) {
	*s = append(*s, v.(int))
}

func (s *IntHeap) Pop() interface{} {
	n := len(*s) - 1
	v := (*s)[n]
	*s = (*s)[0:n]
	return v
}

func (s IntHeap) Peek() interface{} {
	return s[0]
}
