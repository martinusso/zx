package cmd

import "testing"

func TestGenerateLorem(t *testing.T) {
	got, err := generateLorem([]string{})
	if err != nil {
		t.Errorf("Expected not error, got '%s'", err.Error())
	}
	expected := lorem
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}

func TestGenerateLoremInvalidArgumentSize(t *testing.T) {
	got, err := generateLorem([]string{"-1"})
	if err != nil {
		t.Errorf("Expected not error, got '%s'", err.Error())
	}
	expected := lorem
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}

func TestGenerateLoremAtferLimitSize(t *testing.T) {
	got, err := generateLorem([]string{"524"})
	if err != nil {
		t.Errorf("Expected not error, got '%s'", err.Error())
	}
	expected := lorem
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}

func TestGenerateLoremWithValidArgumentSize(t *testing.T) {
	got, err := generateLorem([]string{"26"})
	if err != nil {
		t.Errorf("Expected not error, got '%s'", err.Error())
	}
	expected := "Lorem ipsum dolor sit amet"
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}
