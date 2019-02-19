package simplifyPath

import "strings"

func simplifyPath(path string) string {
	stack := []string{}

	for _, dir := range strings.Split(path, "/") {
		switch dir {
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // pop
			}
		case ".", "":
			continue
		default:
			stack = append(stack, dir)
		}

	}

	return "/" + strings.Join(stack, "/")
}
