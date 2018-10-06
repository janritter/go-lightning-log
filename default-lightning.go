package lightning

var defaultLightning = newLightning(4)

// GetMinSeverity returns the minimal severity level, which is used by the default logger to decide if log function calls should be processed
func GetMinSeverity() int {
	return defaultLightning.minSeverity
}

// Log is the central logging function, which uses an error, a map with tags and a severity level as inputs and decides based on the logger configurtion where to send / display the log
func Log(err error, tags map[string]string, severity int) {
	defaultLightning.Log(err, tags, severity)
}
