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

func compareInts(i, j interface{}) bool {
	return i.(int) < j.(int)
}

func TestPQInts(t *testing.T) {
	q := pq.New(compareInts)
	if q == nil {
		t.Errorf("New() wasn't't be nil")
	}
	if q.Size() != 0 {
		t.Errorf("Size() on empty wasn't 0")
	}
	if q.Peek() != nil {
		t.Errorf("Peek() on empty wasn't nil")
	}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(0)
	for i := 0; i < 4; i++ {
		t.Log("Heap is", q)
		if q.Peek().(int) != i {
			t.Error("Peek() wasn't", i)
		}
		if q.Pop().(int) != i {
			t.Error("Pop() wasn't", i)
		}
	}
}
