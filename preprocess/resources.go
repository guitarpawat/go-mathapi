package preprocess

// InvalidKey contains all reserved key.
var InvalidKey = map[string]struct{}{
	"add": struct{}{},
	"sub": struct{}{},
	"mul": struct{}{},
	"div": struct{}{},
	"mod": struct{}{},
	"pow": struct{}{},
	"par": struct{}{},
	"end": struct{}{},
	"mem": struct{}{},
	"del": struct{}{},
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
