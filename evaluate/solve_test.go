package evaluate

import (
	"testing"
)

func Test_Solve_1_2_Add_ShouldBe_3(t *testing.T) {
	input := []string{"1", "2", "+"}
	expected := "3.0000"
	actual, err := Solve(input)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	if expected != actual {
		t.Errorf("expected: %s but get: %s", expected, actual)
	}
}

func Test_Solve_1_2_Sub_ShouldBe_Minus1(t *testing.T) {
	input := []string{"1", "2", "-"}
	expected := "-1.0000"
	actual, err := Solve(input)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	if expected != actual {
		t.Errorf("expected: %s but get: %s", expected, actual)
	}
}

func Test_Solve_3_2_Mul_ShouldBe_6(t *testing.T) {
	input := []string{"3", "2", "*"}
	expected := "6.0000"
	actual, err := Solve(input)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	if expected != actual {
		t.Errorf("expected: %s but get: %s", expected, actual)
	}
}

func Test_Solve_7_2_Mul_ShouldBe_3Point5(t *testing.T) {
	input := []string{"7", "2", "/"}
	expected := "3.5000"
	actual, err := Solve(input)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	if expected != actual {
		t.Errorf("expected: %s but get: %s", expected, actual)
	}
}

func Test_Solve_8_3_Mod_ShouldBe_2(t *testing.T) {
	input := []string{"8", "3", "%"}
	expected := "2.0000"
	actual, err := Solve(input)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	if expected != actual {
		t.Errorf("expected: %s but get: %s", expected, actual)
	}
}

func Test_Solve_3_4_Mod_ShouldBe_81(t *testing.T) {
	input := []string{"3", "4", "^"}
	expected := "81.0000"
	actual, err := Solve(input)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	if expected != actual {
		t.Errorf("expected: %s but get: %s", expected, actual)
	}
}

func Test_Solve_Empty_ShouldBe_Error(t *testing.T) {
	input := []string{}
	_, err := Solve(input)

	if err == nil {
		t.Errorf("expected error")
	}
}

func Test_Solve_Blank_ShouldBe_Error(t *testing.T) {
	input := []string{""}
	_, err := Solve(input)

	if err == nil {
		t.Errorf("expected error")
	}
}

func Test_Solve_1_Add_ShouldBe_Error(t *testing.T) {
	input := []string{"1", "+"}
	_, err := Solve(input)

	if err == nil {
		t.Errorf("expected error")
	}
}

func Test_Solve_Add_ShouldBe_Error(t *testing.T) {
	input := []string{"+"}
	_, err := Solve(input)

	if err == nil {
		t.Errorf("expected error")
	}
}

func Test_Solve_NotNumber_ShouldBe_Error(t *testing.T) {
	input := []string{"?", "@", "+"}
	_, err := Solve(input)

	if err == nil {
		t.Errorf("expected error")
	}
}

func Test_Solve_NotNumberMixed_ShouldBe_Error(t *testing.T) {
	input := []string{"?", "0", "+"}
	_, err := Solve(input)

	if err == nil {
		t.Errorf("expected error")
	}
}
