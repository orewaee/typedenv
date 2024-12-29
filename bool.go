package typedenv

import (
	"os"
	"strings"
)

// Bool returns the bool value of env by the specified key
func Bool(key string, defaults ...bool) bool {
	env, ok := os.LookupEnv(key)

	if !ok {
		if len(defaults) == 1 {
			return defaults[0]
		}

		if value, ok := defaultBool[key]; ok {
			return value
		}
	}

	return strings.ToLower(env) == "true"
}
