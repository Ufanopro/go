package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Pedro"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Pedro")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Pedro") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q. %v, quiere "", error`, msg, err)
	}
}
