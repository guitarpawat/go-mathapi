package preprocess

// InvalidKey contains all reserved key.
var InvalidKey = map[string]struct{}{
	"add": {},
	"sub": {},
	"mul": {},
	"div": {},
	"mod": {},
	"pow": {},
	"par": {},
	"end": {},
	"mem": {},
	"del": {},
}

// ProcessAll change all the infix to the list on map.
func ProcessAll(in []string) []string {
	res := ConstantToNumber(in)
	res = MemToNumber(res)
	return res
}

// IsKeyInAnyResource returns true if any map has the key.
func IsKeyInAnyResource(key string) bool {
	_, ok1 := GetConstant(key)
	if ok1 {
		return ok1
	}
	_, ok2 := GetFromMem(key)
	if ok2 {
		return ok2
	}

	return false
}

// IsValidKey returns false if that key had reserved.
func IsValidKey(key string) bool {
	_, ok := InvalidKey[key]
	return !ok
}
