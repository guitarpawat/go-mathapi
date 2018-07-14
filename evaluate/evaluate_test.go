package evaluate

import (
	"testing"
)

func Test_1_Add_2_ShouldBe_3(t *testing.T) {
	input := []string{"1", "add", "2"}
	expected := "3.0000"
	actual, err := Evaluate(input)
	if err != nil {
		t.Errorf("error: %s", err)
		return
	}
	if expected != actual {
		t.Errorf("expected: %s but get: %s", expected, actual)
	}
}

func Test_Parentheses_1_Add_1_ShouldBe_2(t *testing.T) {
	input := []string{"(", "1", "add", "1", ")"}
	expected := "2.0000"
	actual, err := Evaluate(input)
	if err != nil {
		t.Errorf("error: %s", err)
		return
	}
	if expected != actual {
		t.Errorf("expected: %s but get: %s", expected, actual)
	}
}

func Test_2_Minus_1_ShouldBe_1(t *testing.T) {
	input := []string{"2", "sub", "1"}
	expected := "1.0000"
	actual, err := Evaluate(input)
	if err != nil {
		t.Errorf("error: %s", err)
		return
	}
	if expected != actual {
		t.Errorf("expected: %s but get: %s", expected, actual)
	}
}

func Test_2_Multiply_2_ShouldBe_4(t *testing.T) {
	input := []string{"2", "mul", "2"}
	expected := "4.0000"
	actual, err := Evaluate(input)
	if err != nil {
		t.Errorf("error: %s", err)
		return
	}
	if expected != actual {
		t.Errorf("expected: %s but get: %s", expected, actual)
	}
}

func Test_4_Divide_2_ShouldBe_2(t *testing.T) {
	input := []string{"4", "div", "2"}
	expected := "2.0000"
	actual, err := Evaluate(input)
	if err != nil {
		t.Errorf("error: %s", err)
		return
	}
	if expected != actual {
		t.Errorf("expected: %s but get: %s", expected, actual)
	}
}

func Test_5_Mod_3_ShouldBe_2(t *testing.T) {
	input := []string{"5", "mod", "3"}
	expected := "2.0000"
	actual, err := Evaluate(input)
	if err != nil {
		t.Errorf("error: %s", err)
		return
	}
	if expected != actual {
		t.Errorf("expected: %s but get: %s", expected, actual)
	}
}

func Test_Complex_Operation(t *testing.T) {
	input := []string{"60", "add", "60", "mul", "0", "add", "1"}
	expected := "61.0000"
	actual, err := Evaluate(input)
	if err != nil {
		t.Errorf("error: %s", err)
		return
	}
	if expected != actual {
		t.Errorf("expected: %s but get: %s", expected, actual)
	}
}

func Test_Invalid_Infix_ShouldBe_Error(t *testing.T) {
	input := []string{"60", "add", "60", "equ", "120"}
	_, err := Evaluate(input)
	if err == nil {
		t.Errorf("expected error")
		return
	}
}

func Test_Invalid_Postfix_ShouldBe_Error(t *testing.T) {
	input := []string{"60", "add", "add"}
	_, err := Evaluate(input)
	if err == nil {
		t.Errorf("expected error")
		return
	}
}
