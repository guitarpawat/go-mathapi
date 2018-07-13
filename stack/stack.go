// Package stack provides convenient using in this project, without casting the interface{}
package stack

import "errors"

// Stack is the stack of string for using in this project
type Stack struct {
	data []string
}

// NewStack return the pointer to new stack
func NewStack() *Stack {
	return new(Stack)
}

// IsEmpty returns true only when length of data in the stack is zero
func (o *Stack) IsEmpty() bool {
	return o.Length() == 0
}

// Length of this stack
func (o *Stack) Length() int {
	return len(o.data)
}

// Push string to the stack
func (o *Stack) Push(s string) {
	o.data = append(o.data, s)
}

// Peek (get) the lastest data from the stack
func (o *Stack) Peek() (string, error) {
	if o.IsEmpty() {
		return "", errors.New("error: stack is empty")
	}
	return o.MustPeek(), nil
}

// MustPeek peeks data from the stack or panic
func (o *Stack) MustPeek() string {
	return o.data[o.Length()-1]
}

// Pop (get and remove) the lastest data from the stack
func (o *Stack) Pop() (string, error) {
	if o.IsEmpty() {
		return "", errors.New("error: stack is empty")
	}
	return o.MustPop(), nil
}

// MustPop pops data from the stack or panic
func (o *Stack) MustPop() string {
	temp := o.MustPeek()
	o.data = o.data[:o.Length()-1]
	return temp
}
