package jsx

import (
	"fmt"
	"strings"

	"myitcv.io/react"
)

func parseCSS(s string) *react.CSS {
	res := new(react.CSS)

	parts := strings.Split(s, ";")

	for _, p := range parts {
		kv := strings.Split(p, ":")
		if len(kv) != 2 {
			panic(fmt.Errorf("invalid key-val %q in %q", p, s))
		}

		k, v := kv[0], kv[1]

		k = strings.TrimSpace(k)
		v = strings.TrimSpace(v)
		v = strings.Trim(v, "\"")

		switch k {

		case "float":
			res.Float = v

		case "font-size":
			res.FontSize = v

		case "font-style":
			res.FontStyle = v

		case "height":
			res.Height = v

		case "left":
			res.Left = v

		case "margin-top":
			res.MarginTop = v

		case "max-height":
			res.MaxHeight = v

		case "min-height":
			res.MinHeight = v

		case "overflow":
			res.Overflow = v

		case "overflow-y":
			res.OverflowY = v

		case "position":
			res.Position = v

		case "resize":
			res.Resize = v

		case "top":
			res.Top = v

		case "width":
			res.Width = v

		case "z-index":
			res.ZIndex = v

		default:
			panic(fmt.Errorf("unknown CSS key %q in %q", k, s))
		}
	}

	return res
}
