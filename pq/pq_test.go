// Copyright 2015 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package pq_test

import (
	"github.com/dagolden/zzz-algorithms-go/pq"
	"testing"
)

func TestPQInts(t *testing.T) {
	_, err := pq.New(&pq.IntHeap{0, 1, 2, 3})
	if err == nil {
		t.Errorf("New() with data didn't error")
	}
	q, err := pq.New(new(pq.IntHeap))
	if q == nil || err != nil {
		t.Errorf("New() failed")
	}
	if q.Len() != 0 {
		t.Errorf("Len() on empty wasn't 0")
	}
	if q.Peek() != nil {
		t.Errorf("Peek() on empty wasn't nil")
	}
	for i := 10; i >= 0; i-- {
		q.Push(i)
		// t.Log("Heap is", q)
	}
	for i := 0; i <= 10; i++ {
		// t.Log("Heap is", q)
		if q.Peek().(int) != i {
			t.Error("Peek() wasn't", i)
		}
		if q.Pop().(int) != i {
			t.Error("Pop() wasn't", i)
		}
	}
}
