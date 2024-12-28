package typedenv

import (
	"os"
	"strings"
)

func Bool(key string, defaults ...bool) bool {
	env, ok := os.LookupEnv(key)

	if !ok && len(defaults) == 1 {
		return defaults[0]
	}

	return strings.ToLower(env) == "true"
}
