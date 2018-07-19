package preprocess

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_ConstantToNumber(t *testing.T) {
	input := []string{"pi", "fun", "e", "2", "phi"}
	expected := []string{"3.141593", "fun", "2.718282", "2", "1.618034"}
	actual := ConstantToNumber(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v but get: %v", expected, input)
	}
}

func Test_GetConstant_Found(t *testing.T) {
	con, ok := GetConstant("e")
	if !ok {
		t.Error("not found constant: e")
		return
	}

	actual := fmt.Sprintf("%.10f", con)
	expected := fmt.Sprintf("%.10f", consts["e"])
	if actual != expected {
		t.Errorf("expected: %s but get: %s", expected, actual)
	}
}

func Test_GetConstant_NotFound(t *testing.T) {
	_, ok := GetConstant("not a constant")
	if ok {
		t.Error("found constant")
	}
}
