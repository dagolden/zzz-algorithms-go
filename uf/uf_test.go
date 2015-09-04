// Copyright 2015 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package uf_test

import (
	"github.com/dagolden/zzz-algorithms-go/uf"
	"testing"
)

func TestQuickFind(t *testing.T) {

	// test that constructor arguments must be positive integers
	for i := 0; i > -2; i-- {
		if uf.NewQuickFind(i) != nil {
			t.Errorf("NewQuickFind(%d) was non-nil", i)
		}
	}

	// test that constructed size is correct
	qf := uf.NewQuickFind(10)
	if qf.Size() != 10 {
		t.Error("QuickFind had incorrect size ", qf.Size())
	}
	if qf.Count() != 10 {
		t.Errorf("QuickFind had incorrect count. Got %d, expected %d", qf.Count(), 10)
	}

	// test find returns identity before union
	for i := 0; i < 10; i++ {
		if qf.Find(i) != i {
			t.Errorf("node %d not linked to self", i)
		}
	}

	// test union
	if qf.Count() != 10 {
		t.Errorf("pre-union component count: want %d, got %d", 10, qf.Count())
	}
	qf.Union(0, 1)
	if !qf.Connected(0, 1) {
		t.Errorf("nodes 0 and 1 not connected")
	}
	if qf.Count() != 9 {
		t.Errorf("post-union component count: want %d, got %d", 9, qf.Count())
	}
}
