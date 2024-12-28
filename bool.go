package typedenv

import (
	"os"
	"strings"
)

func Bool(key string, defaults ...bool) bool {
	env, ok := os.LookupEnv(key)

	if !ok {
		if len(defaults) == 1 {
			return defaults[0]
		}

		if value, ok := defaultBool[key]; !ok {
			return value
		}
	}

	return strings.ToLower(env) == "true"
}
