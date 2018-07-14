package evaluate

import (
	"fmt"
	"strconv"

	"github.com/guitarpawat/go-mathapi/stack"
)

// InfixToPostfix converts infix maths expression to postfix, that is easier to calculate
func InfixToPostfix(infix []string) ([]string, error) {
	postfix := []string{}
	temp := stack.NewStack()
	for _, v := range infix {
		priority := operator[v]
		priorityTOS := 0
		if !temp.IsEmpty() {
			priorityTOS = operator[temp.MustPeek()]
		}
		if isNumber(v) {
			postfix = append(postfix, v)
		} else if v == "par" {
			temp.Push(v)
		} else if v == "end" {
			for !temp.IsEmpty() {
				op := temp.MustPop()
				if op == "par" {
					break
				} else {
					postfix = append(postfix, op)
				}
			}
		} else if priority == 0 {
			return nil, fmt.Errorf("unknown operator: %s", v)
		} else if priority > priorityTOS {
			temp.Push(v)
		} else {
			for priority <= priorityTOS {
				postfix = append(postfix, temp.MustPop())
				if temp.IsEmpty() {
					break
				} else {
					priorityTOS = operator[temp.MustPeek()]
				}
			}
			temp.Push(v)
		}
		//fmt.Println(temp, postfix)
	}
	for !temp.IsEmpty() {
		postfix = append(postfix, temp.MustPop())
		//fmt.Println(temp, postfix)
	}
	return postfix, nil
}

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
