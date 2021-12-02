package apple

import (
	"fmt"
)

type Stack struct {
	stack []rune
}

func (s *Stack) pop() {
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *Stack) push(c rune) {
	s.stack = append(s.stack, c)
}

func (s Stack) isEmpty() bool {
	if len(s.stack) == 0 {
		return true
	} else {
		return false
	}
}

func IsValid(s string) bool {
	if len(s) == 1 {
		return false
	}
	dict := map[rune]rune{')': '(', '}': '{', ']': '['}
	var stack Stack

	for _, char := range s {
		fmt.Println(stack.stack)
		if char == '{' || char == '(' || char == '[' {
			stack.push(char)
		} else if stack.isEmpty() {
			return false
		} else if ((char == ')') || (char == '}') || (char == ']')) && dict[char] == stack.stack[len(stack.stack)-1] {
			stack.pop()
		} else {
			return false
		}
	}
	if len(stack.stack) == 0 {
		return true
	} else {
		return false
	}

}
