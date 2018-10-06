package test

import (
	"testing"

	lightning "github.com/janritter/go-lightning-log"
)

func TestInit(t *testing.T) {

	severity := 1
	logger, _ := lightning.Init(severity)

	if logger.GetMinSeverity() != severity {
		t.Error("Expected 1, got", logger.GetMinSeverity())
	}
}

func TestInitInvalidMinSeverity(t *testing.T) {
	testData := []int{-1, 5}

	for _, data := range testData {
		_, err := lightning.Init(data)

		if err == nil {
			t.Errorf("Expected error for %d, got nil", data)
		}
	}
}
