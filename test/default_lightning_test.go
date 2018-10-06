package test

import (
	"testing"

	lightning "github.com/janritter/go-lightning-log"
)

func TestDefaultInit(t *testing.T) {
	severity := 4

	if lightning.GetMinSeverity() != severity {
		t.Error("Expected 4, got", lightning.GetMinSeverity())
	}
}
