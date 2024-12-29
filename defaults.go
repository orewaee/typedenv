package typedenv

var (
	defaultBool   = make(map[string]bool)
	defaultString = make(map[string]string)
	defaultInt    = make(map[string]int)
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
