package main

import "unicode"
import "unicode/utf8"

func genVarName(s string) string {
	r, _ := utf8.DecodeRuneInString(s)

	// Note: we are choosing to ignore the situation where we decode utf8.RuneError
	// this situation would only happen if the source contained an invalid utf8 code point...
	// which is impossible because it won't compile

	return string(unicode.ToLower(r))
}
