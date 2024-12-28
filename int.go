package typedenv

import (
	"os"
	"strconv"
)

func Int(key string, defaults ...int) int {
	env, ok := os.LookupEnv(key)

	if !ok && len(defaults) == 1 {
		return defaults[0]
	}

	val, err := strconv.Atoi(env)
	if err != nil {
		return 0
	}

	return val
}
