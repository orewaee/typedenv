package typedenv

import "os"

func String(key string, defaults ...string) string {
	env, ok := os.LookupEnv(key)

	if !ok {
		if len(defaults) == 1 {
			return defaults[0]
		}

		if value, ok := defaultString[key]; ok {
			return value
		}
	}

	return env
}
