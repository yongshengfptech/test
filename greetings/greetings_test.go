package greetings

import (
	"regexp"
	"testing"

	"rsc.io/quote"
)

// Test the function Greeting(name) with a "name"
func TestGreeting_Name(t *testing.T) {
	name := "John"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Greeting(name)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("John") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// Test the function Greeting(name) with empty name
func TestGreeting_EmptyName(t *testing.T) {
	msg, err := Greeting("")

	if msg != quote.Hello() || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want %q, error message`, msg, err, quote.Hello())
	}
}
