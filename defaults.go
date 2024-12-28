package typedenv

var (
	defaultBool   = make(map[string]bool)
	defaultString = make(map[string]string)
	defaultInt    = make(map[string]int)
)

func DefaultBool(key string, value bool) {
	defaultBool[key] = value
}

func DefaultString(key string, value string) {
	defaultString[key] = value
}

func DefaultInt(key string, value int) {
	defaultInt[key] = value
}
