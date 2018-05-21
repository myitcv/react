// +build js

package main_test

import (
	"testing"

	"github.com/gopherjs/gopherjs/js"
)

func Test(t *testing.T) {
	w := js.Global.Get("window")

	if w == nil {
		t.Fatalf("expected to be able to access window")
	}
}
