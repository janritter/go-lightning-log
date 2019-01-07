package lightning

import (
	"errors"
	"testing"
)

func TestDefaultGetMinSeverity(t *testing.T) {
	severity := 4

	if GetMinSeverity() != severity {
		t.Error("Expected 4, got", GetMinSeverity())
	}
}

func TestDefaultLogWithResponseNegativeMissingError(t *testing.T) {
	if LogWithResult(nil, nil, 4) == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestDefaultLogWithResponsePositive(t *testing.T) {
	if LogWithResult(errors.New("test error"), nil, 4) != nil {
		t.Errorf("Expected nil, got error")
	}
}

func TestDefaultLogPositive(t *testing.T) {
	Log(errors.New("test error"), nil, 4)
}
