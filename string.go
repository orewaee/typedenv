package typedenv

import "os"

func String(key string, defaults ...string) string {
	env, ok := os.LookupEnv(key)

	if !ok && len(defaults) == 1 {
		return defaults[0]
	}

	return env
}
