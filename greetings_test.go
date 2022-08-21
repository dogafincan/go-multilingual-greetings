package greetings

import (
	"testing"
)

func TestHello(t *testing.T) {
	for index, greeting := range greetings {
		if greeting == "" {
			t.Errorf("Greetings contains an empty string at index %v", index)
		}
	}
}
