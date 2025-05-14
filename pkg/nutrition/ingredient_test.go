package nutrition

import (
	"testing"
)

func TestNewIngredient(t *testing.T) {
	if _, err := NewIngredient(0, 0, 0, 0); err == nil {
		t.Error("Expected an error, but got nil")
	}

	if _, err := NewIngredient(0, 0, -1, 1); err == nil {
		t.Error("Expected an error, but got nil")
	}
}
