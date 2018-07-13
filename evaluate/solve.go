package evaluate

import (
	"fmt"
	"math"
	"strconv"

	"github.com/guitarpawat/go-mathapi/stack"
)

// Solve the math problem, requires postfix input.
func Solve(postfix []string) (string, error) {
	temp := stack.NewStack()
	for _, v := range postfix {
		priority := operator[v]
		if priority > 0 {
			n1, err := temp.Pop()
			if err != nil {
				return "", fmt.Errorf("solve error: %s", err)
			}

			n2, err := temp.Pop()
			if err != nil {
				return "", fmt.Errorf("solve error: %s", err)
			}

			res, err := cal(n1, n2, v)
			if err != nil {
				return "", fmt.Errorf("solve error: %s", err)
			}

			temp.Push(res)
		} else {
			temp.Push(v)
		}
	}

	if temp.Length() != 1 {
		return "", fmt.Errorf("solve error: invalid input format")
	}

	res := temp.MustPop()
	if !isNumber(res) {
		return "", fmt.Errorf("%s is not a number", res)
	}

	return res, nil
}

func cal(n1, n2, op string) (string, error) {
	f1, err := strconv.ParseFloat(n1, 64)
	if err != nil {
		return "", fmt.Errorf("%s is not a number", n1)
	}

	f2, err := strconv.ParseFloat(n2, 64)
	if err != nil {
		return "", fmt.Errorf("%s is not a number", n2)
	}

	switch op {
	case "+":
		return fmt.Sprintf("%.4f", f2+f1), nil
	case "-":
		return fmt.Sprintf("%.4f", f2-f1), nil
	case "*":
		return fmt.Sprintf("%.4f", f2*f1), nil
	case "/":
		return fmt.Sprintf("%.4f", f2/f1), nil
	case "%":
		return fmt.Sprintf("%.4f", float64(int(f2)%int(f1))), nil
	case "^":
		return fmt.Sprintf("%.4f", math.Pow(f2, f1)), nil
	default:
		// Should be impossible to reach here!
		panic(fmt.Sprintf("unknown operator: %s", op))
	}
}
