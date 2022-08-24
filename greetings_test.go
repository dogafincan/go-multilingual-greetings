package greetings

import (
	"math/rand"
	"testing"
)

func init() {
	rand.Seed(1)
}

func TestHello(t *testing.T) {
	// Change expected when the greetings array changes.
	expected := "こんにちは"
	actual := RandomGreeting()

	if actual != expected {
		t.Errorf("expected %v; got %v", expected, actual)
	}
}
