package greetings

import (
	"math/rand"
	"testing"
)

func init() {
	rand.Seed(1)
}

func TestHello(t *testing.T) {
	for index, greeting := range greetings {
		if greeting == "" {
			t.Errorf("greetings contains an empty string at index %v", index)
		}
	}

	// Change expected when the greetings array changes.
	expected := "こんにちは"
	actual := Hello()

	if actual != expected {
		t.Errorf("expected %v; got %v", expected, actual)
	}
}
