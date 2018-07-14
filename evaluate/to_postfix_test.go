package evaluate

import (
	"reflect"
	"testing"
)

type testTable struct {
	input    []string
	expected []string
}

func Test_InfixToPostfix_Normal(t *testing.T) {
	table := []testTable{
		{[]string{"2", "add", "2"}, []string{"2", "2", "add"}},
		{[]string{"3", "add", "4", "mul", "5", "div", "6"}, []string{"3", "4", "5", "mul", "6", "div", "add"}},
		{[]string{"par", "8", "add", "9", "mul", "par", "4", "add", "5", "mul", "7", "add", "6", "end", "end", "add", "3", "mul", "5", "add", "4"},
			[]string{"8", "9", "4", "5", "7", "mul", "add", "6", "add", "mul", "add", "3", "5", "mul", "add", "4", "add"}},
		{[]string{"par", "300", "add", "23", "end", "mul", "par", "43", "sub", "21", "end", "div", "par", "84", "add", "7", "end"},
			[]string{"300", "23", "add", "43", "21", "sub", "mul", "84", "7", "add", "div"}},
		{[]string{"par", "4", "add", "8", "end", "mul", "par", "6", "sub", "5", "end", "div", "par", "par", "3", "sub", "2", "end", "mul", "par", "2", "add", "2", "end", "end"},
			[]string{"4", "8", "add", "6", "5", "sub", "mul", "3", "2", "sub", "2", "2", "add", "mul", "div"}},
	}
	for _, v := range table {
		actual, err := InfixToPostfix(v.input)
		if err != nil {
			t.Errorf("error: %s", err)
		} else if !reflect.DeepEqual(v.expected, actual) {
			t.Errorf("input: %s expected: %s but get: %s", v.input, v.expected, actual)
		}
	}
}

func Test_InfixToPostfix_Error(t *testing.T) {
	input := []string{"1", "?", "9"}
	_, err := InfixToPostfix(input)
	if err == nil {
		t.Errorf("expected error")
	}
}
