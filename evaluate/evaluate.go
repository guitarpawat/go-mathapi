package evaluate

// Evaluate receives infix expression and solve the math problem
func Evaluate(infix []string) (string, error) {
	posfix, err := InfixToPostfix(infix)
	if err != nil {
		return "", err
	}

	result, err := Solve(posfix)
	if err != nil {
		return "", err
	}

	return result, nil
}
