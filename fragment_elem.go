// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

// FragmentElem is the special React Fragment element definition. Fragments let
// you group a list of children without adding extra nodes to the DOM. See
// https://reactjs.org/docs/fragments.html for more details.
type FragmentElem struct {
	Element
}

// Fragment creates a new instance of a <React.Fragment> element with the
// provided children
func Fragment(children ...Element) *FragmentElem {
	return &FragmentElem{
		Element: createElement(symbolFragment, nil, children...),
	}
}
