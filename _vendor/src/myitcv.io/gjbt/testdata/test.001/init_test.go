// +build js

package main_test

import (
	"fmt"
	"testing"

	"github.com/gopherjs/gopherjs/js"
)

func TestMain(m *testing.M) {
	i := m.Run()

	js.Global.Call("eval", fmt.Sprintf("window.$GopherJSTestResult = %v", i))
}
