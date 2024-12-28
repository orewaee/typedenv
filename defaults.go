package typedenv

var (
	defaultBool   = make(map[string]bool)
	defaultString = make(map[string]string)
	defaultInt    = make(map[string]int)
)

func SetDefaultBool(key string, value bool) {
	defaultBool[key] = value
}

func SetDefaultString(key string, value string) {
	defaultString[key] = value
}

func SetDefaultInt(key string, value int) {
	defaultInt[key] = value
}
