// +build js

package react_test

import (
	"testing"

	"honnef.co/go/js/dom"

	"myitcv.io/react"
	"myitcv.io/react/testutils"
)

func TestDivElem(t *testing.T) {
	class := "test"

	div := react.Div(
		&react.DivProps{
			ClassName: class,
			DataSet: react.DataSet{
				"toggle": "dropdown",
			},
		},
	)

	x := testutils.Wrapper(div)
	cont := testutils.RenderIntoDocument(x)

	el := testutils.FindRenderedDOMComponentWithClass(cont, class)

	if _, ok := el.(*dom.HTMLDivElement); !ok {
		t.Fatal("Failed to find <Div> element")
	}
}
