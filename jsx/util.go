package jsx

import "strconv"

func parseBool(s string) bool {
	v, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}

	return v
}
