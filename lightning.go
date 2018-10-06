package lightning

import (
	"errors"
	"fmt"
	"log"
)

var severityLevels = map[int]string{
	0: "FATAL",
	1: "ERROR",
	2: "WARNING",
	3: "INFO",
	4: "DEBUG",
}

func newLightning(minSeverity int) *Lightning {
	lightning := &Lightning{
		minSeverity: minSeverity,
	}

	return lightning
}

// GetMinSeverity returns the minimal severity level, which is used by the logger to decide if log function calls should be processed
func (lightning *Lightning) GetMinSeverity() int {
	return lightning.minSeverity
}

// Init lets you create and configure a custom logger
func Init(minSeverity int) (*Lightning, error) {
	if minSeverity > 4 || minSeverity < 0 {
		return nil, errors.New("minSeverity must be in the range of 0 to 4")
	}

	lightning := newLightning(minSeverity)
	return lightning, nil
}

// Log is the central logging function, which uses an error, a map with tags and a severity level as inputs and decides based on the logger configurtion where to send / display the log
func (lightning *Lightning) Log(err error, tags map[string]string, severity int) error {
	if lightning.minSeverity < severity {
		return errors.New("no logging, since minSeverity in the logger is lower than given severity")
	}

	go logToConsole(err, tags, severity)
	return nil
}

func logToConsole(err error, tags map[string]string, severity int) {
	if tags != nil {
		log.Println(fmt.Sprintf("%s - %s -- %s", severityLevels[severity], tags, err))
		return
	}
	log.Println(fmt.Sprintf("%s - %s", severityLevels[severity], err))
}
