package fx

import (
	"fmt"
	"testing"
)

func TestResultOk(t *testing.T) {
	result := Ok(1)

	if !result.IsOk() {
		t.Errorf("Expected true, got false")
	}

	if result.Get() != 1 {
		t.Errorf("Expected 1, got %d", result.Get())
	}

	if value, err := result.Values(); value != 1 || err != nil {
		t.Errorf("Expected 1, got %d", value)
	}

	if result.IsErr() {
		t.Errorf("Expected false, got true")
	}

	if err := result.Error(); err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if result.OrElse(2) != 1 {
		t.Errorf("Expected 1, got %d", result.OrElse(2))
	}
}

func TestResultErr(t *testing.T) {
	result := Err[int](fmt.Errorf("test error"))

	if result.IsOk() {
		t.Errorf("Expected false, got true")
	}

	if result.Get() != 0 {
		t.Errorf("Expected 0, got %d", result.Get())
	}

	if value, err := result.Values(); value != 0 || err == nil {
		t.Errorf("Expected 0, got %d", value)
	}

	if !result.IsErr() {
		t.Errorf("Expected true, got false")
	}

	if err := result.Error(); err == nil {
		t.Errorf("Expected error, got nil")
	}

	if result.OrElse(1) != 1 {
		t.Errorf("Expected 1, got %d", result.OrElse(1))
	}
}
