package simplifypath

import "strings"

func simplifyPath(path string) string {
	var curr []byte = make([]byte, 0, len(path))
	var out []string = make([]string, 0, len(path))

	for _, sym := range []byte(path[1:] + "/") {
		if sym != '/' {
			curr = append(curr, sym)
		} else if len(curr) > 0 {
			if curr[0] == '.' {
				if len(curr) > 1 {
					if curr[1] == '.' && len(curr) == 2 {
						out = out[:max(0, len(out)-1)]
					} else {
						out = append(out, string(curr))
					}
				}
			} else {
				out = append(out, string(curr))
			}
			curr = curr[:0]
		}
	}
	return "/" + strings.Join(out, "/")
}
