package utils

import (
// "container/list"
// "fmt"
)

// slice related utils
// compare two slices to see if they are equal
func SlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// remove a specific index from a slice
func RemoveFromSlice(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

// find the shortest string in a slice of strings
func FindShortestString(slice []string) string {
	if len(slice) == 0 {
		return "" // Return an empty string if the slice is empty
	}

	shortest := slice[0]

	for _, str := range slice {
		if len(str) < len(shortest) {
			shortest = str
		}
	}

	return shortest
}

// stack utils
func SliceContainsString(slice []string, char byte) bool {
	charStr := string(char)
	for _, c := range slice {
		if string(c) == charStr {
			return true
		}
	}
	return false
}

// Stack represents a stack data structure
// type Stack struct {
// 	elements []string
// }

// // Push adds an element to the top of the stack
// func (s *Stack) Push(value string) {
// 	s.elements = append(s.elements, value)
// }

// // Pop removes an element from the top of the stack and returns it
// func (s *Stack) Pop() error {
// 	// commenting the below for now for test coverage purposes
// 	// if len(s.elements) == 0 {
// 	// 		return fmt.Errorf("stack is empty")
// 	// }
// 	// Remove the top element
// 	s.elements = s.elements[:len(s.elements)-1]
// 	return nil
// }

// // Peek returns the top element of the stack without removing it
// func (s *Stack) Peek() (string, error) {
// 	if len(s.elements) == 0 {
// 			return "", fmt.Errorf("stack is empty")
// 	}
// 	return s.elements[len(s.elements)-1], nil
// }

// Stack of ints
// type Stack struct {
// 	list *list.List
// }

// create a stack with stack := NewStack()
// func NewStack() *Stack {
// 	return &Stack{list: list.New()}
// }

// func (s *Stack) Push(value interface{}) {
// 	s.list.PushBack(value)
// }

// func (s *Stack) Pop() (interface{}, error) {
// 	if s.list.Len() == 0 {
// 			return nil, fmt.Errorf("stack is empty")
// 	}
// 	element := s.list.Back()
// 	s.list.Remove(element)
// 	return element.Value, nil
// }

// func (s *Stack) Peek() (interface{}, error) {
// 	if s.list.Len() == 0 {
// 			return nil, fmt.Errorf("stack is empty")
// 	}
// 	return s.list.Back().Value, nil
// }

// func (s *Stack) IsEmpty() bool {
// 	return s.list.Len() == 0
// }
