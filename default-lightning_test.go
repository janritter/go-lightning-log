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

func TestDefaultLogPositive(t *testing.T) {
	if Log(errors.New("test error"), nil, 4) != nil {
		t.Errorf("Expected nil, got error")
	}
}
