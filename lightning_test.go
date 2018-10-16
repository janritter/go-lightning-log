package lightning

import (
	"bytes"
	"errors"
	"log"
	"os"
	"strings"
	"testing"
)

func TestInit(t *testing.T) {

	severity := 1
	logger, _ := Init(severity)

	if logger.minSeverity != severity {
		t.Error("Expected 1, got", logger.minSeverity)
	}
}

func TestInitInvalidMinSeverity(t *testing.T) {
	testData := []int{-1, 5}

	for _, data := range testData {
		_, err := Init(data)

		if err == nil {
			t.Errorf("Expected error for %d, got nil", data)
		}
	}
}

func TestGetMinSeverity(t *testing.T) {
	severity := 1
	logger, _ := Init(severity)

	if logger.GetMinSeverity() != severity {
		t.Error("Expected 1, got", logger.GetMinSeverity())
	}
}

func TestLogNegativeSeverityToHigh(t *testing.T) {
	severity := 3
	logger, _ := Init(severity)

	if logger.Log(nil, nil, 4) == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestLogPositive(t *testing.T) {
	severity := 3
	logger, _ := Init(severity)

	if logger.Log(errors.New("test error"), nil, 2) != nil {
		t.Errorf("Expected nil, got error")
	}
}

func TestLogToConsole(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	logToConsole(errors.New("test error"), map[string]string{"module": "controller/bookings"}, 2)

	if strings.Contains(buf.String(), "WARNING - map[module:controller/bookings] -- test error") == false {
		t.Log(buf.String())
		t.Errorf("Expected 'WARNING - map[module:controller/bookings] -- test error', got " + buf.String())
	}

	logToConsole(errors.New("test error"), nil, 2)

	if strings.Contains(buf.String(), "WARNING - test error") == false {
		t.Log(buf.String())
		t.Errorf("Expected 'WARNING - test error', got " + buf.String())
	}
}
