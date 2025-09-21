package typedenv

import (
	"os"
	"strconv"
)

func Int64(key string, defaults ...int64) int64 {
	env, ok := os.LookupEnv(key)

	if !ok {
		if len(defaults) == 1 {
			return defaults[0]
		}

		if value, ok := defaultInt64[key]; ok {
			return value
		}
	}

	val, err := strconv.ParseInt(env, 10, 64)
	if err != nil {
		return 0
	}

	return val
}
