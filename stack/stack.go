package stack

import "errors"

type Stack struct {
	data []string
}

func NewStack() *Stack {
	return new(Stack)
}

func (o *Stack) IsEmpty() bool {
	return o.Length() == 0
}

func (o *Stack) Length() int {
	return len(o.data)
}

func (o *Stack) Push(s string) {
	o.data = append(o.data, s)
}

func (o *Stack) Peek() (string, error) {
	if o.IsEmpty() {
		return "", errors.New("error: stack is empty")
	}
	return o.data[o.Length()-1], nil
}

func (o *Stack) Pop() (string, error) {
	temp, err := o.Peek()
	if err != nil {
		return "", err
	}
	o.data = o.data[:o.Length()-1]
	return temp, nil
}
