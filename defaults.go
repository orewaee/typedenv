package typedenv

import "time"

var (
	defaultBool     = make(map[string]bool)
	defaultString   = make(map[string]string)
	defaultInt      = make(map[string]int)
	defaultInt64    = make(map[string]int64)
	defaultDuration = make(map[string]time.Duration)
)

// DefaultBool sets the global default value to the bool env
func DefaultBool(key string, value bool) {
	defaultBool[key] = value
}

// DefaultString sets the global default value to the string env
func DefaultString(key string, value string) {
	defaultString[key] = value
}

// DefaultInt sets the global default value to the int env
func DefaultInt(key string, value int) {
	defaultInt[key] = value
}

// DefaultInt64 sets the global default value to the int env
func DefaultInt64(key string, value int64) {
	defaultInt64[key] = value
}

// DefaultDuration sets the global default value to the duration env
func DefaultDuration(key string, value time.Duration) {
	defaultDuration[key] = value
}
