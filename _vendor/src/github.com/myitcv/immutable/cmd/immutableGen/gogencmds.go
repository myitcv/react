// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package main

import "fmt"

type gogenCmds []string

func (g *gogenCmds) Set(value string) error {
	*g = append(*g, value)
	return nil
}

func (g *gogenCmds) String() string {
	return fmt.Sprint(*g)
}
