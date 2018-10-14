package lightning

import "testing"

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

func TestLogNegative(t *testing.T) {
	severity := 3
	logger, _ := Init(severity)

	if logger.Log(nil, nil, 4) == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestLogPositive(t *testing.T) {
	severity := 3
	logger, _ := Init(severity)

	if logger.Log(nil, nil, 2) == nil {
		t.Errorf("Expected error, got nil")
	}
}
