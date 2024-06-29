package health

import (
	"testing"
)

func TestNewHealth(t *testing.T) {
	h := NewHealth()
	if h.Status != StatusUnknown {
		t.Errorf("Expected status %s, got %s", StatusUnknown, h.Status)
	}
	if h.RTT != 0 {
		t.Errorf("Expected RTT 0, got %d", h.RTT)
	}
	if h.Result != nil {
		t.Errorf("Expected result nil, got %v", h.Result)
	}
}

func TestHealth_Up(t *testing.T) {
	h := NewHealth().Up()
	if h.Status != StatusUp {
		t.Errorf("Expected status %s, got %s", StatusUp, h.Status)
	}
}

func TestHealth_Down(t *testing.T) {
	h := NewHealth().Down()
	if h.Status != StatusDown {
		t.Errorf("Expected status %s, got %s", StatusDown, h.Status)
	}
}

func TestHealth_Unknown(t *testing.T) {
	h := NewHealth().Unknown()
	if h.Status != StatusUnknown {
		t.Errorf("Expected status %s, got %s", StatusUnknown, h.Status)
	}
}

func TestHealth_SetResult(t *testing.T) {
	h := NewHealth().SetResult("test result")
	if h.Result != "test result" {
		t.Errorf("Expected result 'test result', got %v", h.Result)
	}
}

func TestHealth_SetRTT(t *testing.T) {
	h := NewHealth().SetRTT(123)
	if h.RTT != 123 {
		t.Errorf("Expected RTT 123.45, got %d", h.RTT)
	}

	h.SetRTT(-5)
	if h.RTT != 0 {
		t.Errorf("Expected RTT 0 for negative input, got %d", h.RTT)
	}
}

func TestHealth_IsUp(t *testing.T) {
	h := NewHealth().Up()
	if !h.IsUp() {
		t.Errorf("Expected IsUp to be true, got false")
	}
}

func TestHealth_IsDown(t *testing.T) {
	h := NewHealth().Down()
	if !h.IsDown() {
		t.Errorf("Expected IsDown to be true, got false")
	}
}

func TestHealth_IsUnknown(t *testing.T) {
	h := NewHealth().Unknown()
	if !h.IsUnknown() {
		t.Errorf("Expected IsUnknown to be true, got false")
	}
}

func TestHealth_HasResult(t *testing.T) {
	h := NewHealth()
	if h.HasResult() {
		t.Errorf("Expected HasResult to be false, got true")
	}
	h.SetResult("test result")
	if !h.HasResult() {
		t.Errorf("Expected HasResult to be true, got false")
	}
}

func TestHealth_HasRTT(t *testing.T) {
	h := NewHealth()
	if h.HasRTT() {
		t.Errorf("Expected HasRTT to be false, got true")
	}
	h.SetRTT(123)
	if !h.HasRTT() {
		t.Errorf("Expected HasRTT to be true, got false")
	}
}

func TestHealth_GetResult(t *testing.T) {
	h := NewHealth().SetResult("test result")
	if h.GetResult() != "test result" {
		t.Errorf("Expected result 'test result', got %v", h.GetResult())
	}
}

func TestHealth_GetRTT(t *testing.T) {
	h := NewHealth().SetRTT(123)
	if h.GetRTT() != 123 {
		t.Errorf("Expected RTT 123.45, got %d", h.GetRTT())
	}
}
