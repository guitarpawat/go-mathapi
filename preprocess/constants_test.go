package preprocess

import (
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
