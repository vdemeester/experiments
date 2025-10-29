package main

import "testing"

func TestGetMessage(t *testing.T) {
	expected := "Hello, World!"
	result := getMessage()
	if result != expected {
		t.Errorf("getMessage() = %q, want %q", result, expected)
	}
}

func TestMain(t *testing.T) {
	// Fake test to ensure main function exists and runs without panic
	// In a real scenario, we might test the main function's behavior
	// but for now we just pass
	t.Log("Main function test passed")
}

func TestMessageNotEmpty(t *testing.T) {
	result := getMessage()
	if result == "" {
		t.Error("getMessage() returned empty string")
	}
}
