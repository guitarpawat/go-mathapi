package preprocess

import (
	"fmt"
	"math"
)

var consts = map[string]float64{}

func init() {
	consts["e"] = math.E
	consts["pi"] = math.Pi
	consts["phi"] = math.Phi
}

// ConstantToNumber replaces constants specify in consts map.
func ConstantToNumber(in []string) []string {
	for i := 0; i < len(in); i++ {
		con, ok := GetConstant(in[i])
		if ok {
			in[i] = fmt.Sprintf("%.6f", con)
		}
	}
	return in
}

// GetConstant by key and return value and boolean indicate validity.
func GetConstant(key string) (float64, bool) {
	con, ok := consts[key]
	return con, ok
}
