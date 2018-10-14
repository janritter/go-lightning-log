package lightning

import "testing"

func TestDefaultGetMinSeverity(t *testing.T) {
	severity := 4

	if GetMinSeverity() != severity {
		t.Error("Expected 4, got", GetMinSeverity())
	}
}

// func TestDefaultLog(t *testing.T) {
// 	if Log(nil, nil, 4) == nil {
// 		t.Errorf("Expected error, got nil")
// 	}
// }
