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

func TestConstruction(t *testing.T) {
	// test that constructed size is correct
	qf := uf.NewQuickFind(10)
	if qf.Size() != 10 {
		t.Error("QuickFind nodes had incorrect size ", qf.Size())
	}

	// test that arguments must be positive integers
	for i := 0; i > -2; i-- {
		if uf.NewQuickFind(i) != nil {
			t.Errorf("NewQuickFind(%d) was non-nil", i)
		}
	}
}
