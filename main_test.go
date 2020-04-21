package main

import "testing"

func TestGreet(t *testing.T) {
	expected := "Hello, Blog!"
	response := Greet("Blog")
	if expected != response {
		t.Errorf("Greet was incorrect, got: '%s', want: '%s'", expected, response)
	}
}
