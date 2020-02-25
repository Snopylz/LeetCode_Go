package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("{[()]}"))
}

func isValid(s string) bool {
	if s == "" {
		return true
	} else {
		if len(s)%2 != 0 {
			return false
		}
	}
	stack := make([]int32, len(s))
	stackLength := 0
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack[stackLength] = v
			stackLength++
		} else {
			if stackLength == 0 {
				return false
			}
			if (v == ')' && stack[stackLength-1] == '(') || (v == ']' && stack[stackLength-1] == '[') || (v == '}' && stack[stackLength-1] == '{') {
				stackLength--
			} else {
				return false
			}
		}

	}
	return stackLength == 0
}
