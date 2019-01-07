package lightning

var defaultLightning = newLightning(4)

// GetMinSeverity returns the minimal severity level, which is used by the default logger to decide if log function calls should be processed
func GetMinSeverity() int {
	return defaultLightning.minSeverity
}

// LogWithResult is similar to the standard "Log" command with the main difference, that all validation errors are getting returned. This can be
// usefull for debugging
func LogWithResult(err error, tags map[string]string, severity int) error {
	return defaultLightning.LogWithResult(err, tags, severity)
}

// Log is the central logging function, which uses an error, a map with tags and a severity level as inputs and decides based on the logger
// configurtion where to send / display the log entry
// In comparison to the "LogWithResult" command, all validation errors (for example: empty error parameter) get handled internally and not returned
// to the user
func Log(err error, tags map[string]string, severity int) {
	defaultLightning.Log(err, tags, severity)
}
