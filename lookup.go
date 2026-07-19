package typedenv

import "os"

func Lookup(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		return true
	}

	if _, ok := defaultBool[key]; ok {
		return true
	}

	if _, ok := defaultString[key]; ok {
		return true
	}

	if _, ok := defaultInt[key]; ok {
		return true
	}

	if _, ok := defaultInt64[key]; ok {
		return true
	}

	if _, ok := defaultDuration[key]; ok {
		return true
	}

	return false
}
