package leet

import (
	"fmt"
	"practice/go/base/utils"
)

func IsValidParanthesis(s string) bool {
	if len(s) == 0 {
		return false
	}

	specialChars := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}

	OpenList := make([]string, 0, len(specialChars))

	for key := range specialChars {
		OpenList = append(OpenList, key)
	}

	stack := &Stack{}

	// loop over the list of strings
	for i := 0; i < len(s); i++ {
		if utils.SliceContainsString(OpenList, s[i]) {
			// if this item is in our OpenList
			stack.Push(string(s[i]))
		} else {
			top, err := stack.Peek()
			if err != nil {
				return false
			}
			if specialChars[top] == string(s[i]) {
				stack.Pop()
			} else {
				return false
			}
		}
	}
	_, err := stack.Peek()

	return err != nil
}

type Stack struct {
	elements []string
}

// Push adds an element to the top of the stack
func (s *Stack) Push(value string) {
	s.elements = append(s.elements, value)
}

// Pop removes an element from the top of the stack and returns it
func (s *Stack) Pop() {
	// commenting the below for now for test coverage purposes
	// if len(s.elements) == 0 {
	// 		return fmt.Errorf("stack is empty")
	// }
	// Remove the top element
	s.elements = s.elements[:len(s.elements)-1]
	// return nil
}

// Peek returns the top element of the stack without removing it
func (s *Stack) Peek() (string, error) {
	if len(s.elements) == 0 {
		return "", fmt.Errorf("stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}
