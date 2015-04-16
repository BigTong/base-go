// Copyright (c) 2015, The Tony Authors.
// All rights reserved.
//
// Author: Rentong Zhang <rentongzh@gmail.com>

package base

import (
	"testing"
)

func TestStringClean(t *testing.T) {
	s1 := Clean("  aa")
	if s1 != "aa" {
		t.Error("not pass s1:", s1)
	}

	s2 := Clean("aa    bb cc   ")
	if s2 != "aa bb cc" {
		t.Error("not pass s2:", s2)
	}

	s3 := Clean("aa    bb cc  dd  eee    ")
	if s3 != "aa bb cc dd eee" {
		t.Error("not pass s3:", s3)
	}

	s4 := Clean("         aa    bb cc  dd         eee    ")
	if s4 != "aa bb cc dd eee" {
		t.Error("not pass s4:", s4)
	}

	s5 := Clean("         aa    bb cc  dd         eee    fff")
	if s5 != "aa bb cc dd eee fff" {
		t.Error("not pass s5:", s5)
	}
}
