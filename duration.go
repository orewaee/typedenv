package typedenv

import (
	"os"
	"time"
)

// Duration returns the duration value of env by the specified key
func Duration(key string, defaults ...time.Duration) time.Duration {
	env, ok := os.LookupEnv(key)

	if !ok {
		if len(defaults) == 1 {
			return defaults[0]
		}

		if value, ok := defaultDuration[key]; ok {
			return value
		}
	}

	val, err := time.ParseDuration(env)
	if err != nil {
		return 0
	}

	return val
}
