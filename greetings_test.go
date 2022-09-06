package greetings

import (
	"math/rand"
	"testing"
)

func init() {
	rand.Seed(1)
}

func TestEmptyStringGreeting(t *testing.T) {
	for _, greeting := range greetings {
		if greeting == "" {
			t.Fatal("expected a non-empty string")
		}
	}
}

func TestRandomGreeting(t *testing.T) {
	// Change expected when the greetings array changes.
	expected := "こんにちは"
	actual := RandomGreeting()

	if actual != expected {
		t.Errorf("expected %v; got %v", expected, actual)
	}
}
