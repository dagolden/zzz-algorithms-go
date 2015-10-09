// Copyright 2015 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package rbt_test

import (
	"github.com/dagolden/zzz-algorithms-go/rbt"
	"testing"
)

func TestRBT(t *testing.T) {
	rb, err := rbt.New()
	if err != nil {
		t.Errorf("New() failed")
	}

	if size := rb.Size(); size != 0 {
		t.Errorf("Size() was %d, not 0", size)
	}

	data := []string{"S", "E", "A", "R", "C", "H"}
	for i, c := range data {
		err := rb.Put(c, i)
		if err != nil {
			t.Errorf("Put(%s, %d) failed", c, i)
		}
	}

	if size := rb.Size(); size != len(data) {
		t.Errorf("Size() was %d, not %d", size, len(data))
	}

	for i, c := range data {
		val, err := rb.Get(c)
		if err != nil {
			t.Errorf("Get(%s) failed: %s", c, err)
		} else if val != i {
			t.Errorf("Get(%s) was %d, not %d", c, val, i)
		}
	}

	it := rb.Iterator()
	str := ""
	for it.HasNext() {
		s, _ := it.Next()
		str += s
	}
	if str != "ACEHRS" {
		t.Errorf("Iterator returned '%s', but expected '%s'", str, "ACEHRS")
	}
}
