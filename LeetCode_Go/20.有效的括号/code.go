package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid2("{[()]}"))
}

func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 != 0 {
		return false
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

func isValid2(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	var stack []byte
	stackLength := len(stack)
	mymap := map[byte]byte{')': '(', ']': '[', '}': '{'}
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 {
			if _, dp := mymap[byte(s[i])]; dp {
				if stack[stackLength-1] == mymap[byte(s[i])] {
					stack = stack[:stackLength-1]
					stackLength--
					continue
				}
			}
		}
		stackLength++
		stack = append(stack, byte(s[i]))
		fmt.Println("length: ", stackLength)
	}
	return stackLength == 0
}
