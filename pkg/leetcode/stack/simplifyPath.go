package stack

import (
	"learn-golang/pkg/util"
	"strings"
)

func SimplifyPath(path string) string {

	pathStrings := strings.Split(path, "/")
	stack := &util.Stack[string]{}

	for _, value := range pathStrings {

		if value == "" || value == "." {
			continue
		} else if value == ".." {

			if !stack.IsEmpty() {
				stack.Pop()
			}

		} else {
			stack.Push(value)
		}

	}

	if stack.IsEmpty() {
		return "/"
	}

	result := ""

	for !stack.IsEmpty() {
		top, _ := stack.Pop()
		result = "/" + top + result
	}

	return result

}
