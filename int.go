package typedenv

import (
	"os"
	"strconv"
)

// Int returns the int value of env by the specified key
func Int(key string, defaults ...int) int {
	env, ok := os.LookupEnv(key)

	if !ok {
		if len(defaults) == 1 {
			return defaults[0]
		}

		if value, ok := defaultInt[key]; ok {
			return value
		}
	}

	val, err := strconv.Atoi(env)
	if err != nil {
		return 0
	}

	return val
}
