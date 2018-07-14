package evaluate

import (
	"testing"
)

func Test_1_Add_2_ShouldBe_3(t *testing.T) {
	input := []string{"1", "+", "2"}
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
	input := []string{"(", "1", "+", "1", ")"}
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
	input := []string{"2", "-", "1"}
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
	input := []string{"2", "*", "2"}
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
	input := []string{"4", "/", "2"}
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
	input := []string{"5", "%", "3"}
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
	input := []string{"60", "+", "60", "*", "0", "+", "1"}
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
	input := []string{"60", "+", "60", "=", "120"}
	_, err := Evaluate(input)
	if err == nil {
		t.Errorf("expected error")
		return
	}
}

func Test_Invalid_Postfix_ShouldBe_Error(t *testing.T) {
	input := []string{"60", "+", "+"}
	_, err := Evaluate(input)
	if err == nil {
		t.Errorf("expected error")
		return
	}
}
