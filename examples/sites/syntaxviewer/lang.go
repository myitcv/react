package main

import "strconv"

const _lang_name = "GoShell"

var _lang_index = [...]uint8{0, 2, 7}

func (i lang) String() string {
	if i < 0 || i >= lang(len(_lang_index)-1) {
		return "lang(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _lang_name[_lang_index[i]:_lang_index[i+1]]
}
