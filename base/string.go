// Copyright (c) 2015, The Tony Authors.
// All rights reserved.
//
// Author: Rentong Zhang <rentongzh@gmail.com>

package base

import (
	"strings"
)

func Clean(src string) string {
	ret := strings.TrimSpace(src)
	ret = strings.Replace(ret, "\n", " ", -1)
	ret = strings.Replace(ret, "\u0009", " ", -1)
	ret = strings.Replace(ret, "&nbsp;", " ", -1)
	runes := []rune(ret)
	j := 0
	startSpace := true
	for i := 0; i < len(runes); i++ {
		if runes[i] != rune(' ') {
			startSpace = false
			runes[j] = runes[i]
			j++
			continue
		}

		if startSpace {
			continue
		}
		startSpace = true
		runes[j] = runes[i]
		j++
	}
	return string(runes[0:j])
}

func Segment(str, start, end string) string {
	segs1 := strings.Split(str, start)
	if len(segs1) < 2 {
		return str
	}

	segs := strings.Split(segs1[1], end)
	if len(segs) < 1 {
		return str
	}
	return segs[0]
}
