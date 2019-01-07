package lightning

import (
	"bytes"
	"errors"
	"log"
	"os"
	"strings"
	"testing"
	"time"
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

func TestLogWithResultNegativeSeverityToHigh(t *testing.T) {
	severity := 3
	logger, _ := Init(severity)

	if logger.LogWithResult(nil, nil, 4) == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestLogWithResultPositive(t *testing.T) {
	severity := 3
	logger, _ := Init(severity)

	if logger.LogWithResult(errors.New("test error"), nil, 2) != nil {
		t.Errorf("Expected nil, got error")
	}
}

func TestLogWithResultNegativeMissingError(t *testing.T) {
	severity := 3
	logger, _ := Init(severity)

	if logger.LogWithResult(nil, nil, 3) == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestLogToConsole(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

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

func TestLogNegativeSeverityToHigh(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	severity := 2
	logger, _ := Init(severity)

	logger.Log(errors.New("test error"), map[string]string{"module": "controller/bookings"}, 3)

	// Wait for go routine to finish
	time.Sleep(10000)

	if strings.Contains(buf.String(), "") == false {
		t.Log(buf.String())
		t.Errorf("Expected '', got " + buf.String())
		log.SetOutput(os.Stderr)
	}
}

func TestLogPositive(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	severity := 3
	logger, _ := Init(severity)

	logger.Log(errors.New("test error"), map[string]string{"module": "controller/bookings"}, 3)

	// Wait for go routine to finish
	time.Sleep(10000)

	if strings.Contains(buf.String(), "INFO - map[module:controller/bookings] -- test error") == false {
		t.Log(buf.String())
		t.Errorf("Expected 'INFO - map[module:controller/bookings] -- test error', got " + buf.String())
		log.SetOutput(os.Stderr)
	}
}

func TestLogNegativeMissingError(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	severity := 2
	logger, _ := Init(severity)

	logger.Log(nil, map[string]string{"module": "controller/bookings"}, 1)

	// Wait for go routine to finish
	time.Sleep(10000)

	if strings.Contains(buf.String(), "") == false {
		t.Log(buf.String())
		t.Errorf("Expected '', got " + buf.String())
		log.SetOutput(os.Stderr)
	}
}
