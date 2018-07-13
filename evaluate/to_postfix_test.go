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
		{[]string{"2", "+", "2"}, []string{"2", "2", "+"}},
		{[]string{"3", "+", "4", "*", "5", "/", "6"}, []string{"3", "4", "5", "*", "6", "/", "+"}},
		{[]string{"(", "8", "+", "9", "*", "(", "4", "+", "5", "*", "7", "+", "6", ")", ")", "+", "3", "*", "5", "+", "4"},
			[]string{"8", "9", "4", "5", "7", "*", "+", "6", "+", "*", "+", "3", "5", "*", "+", "4", "+"}},
		{[]string{"(", "300", "+", "23", ")", "*", "(", "43", "-", "21", ")", "/", "(", "84", "+", "7", ")"},
			[]string{"300", "23", "+", "43", "21", "-", "*", "84", "7", "+", "/"}},
		{[]string{"(", "4", "+", "8", ")", "*", "(", "6", "-", "5", ")", "/", "(", "(", "3", "-", "2", ")", "*", "(", "2", "+", "2", ")", ")"},
			[]string{"4", "8", "+", "6", "5", "-", "*", "3", "2", "-", "2", "2", "+", "*", "/"}},
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
