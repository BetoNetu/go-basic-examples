package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	//What does MustCompile mean for regex
	// What is `\b`+name+`\b`
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")

	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
		// why ``?
		// what is %q, %v and %#q???????
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloErrorHandling(t *testing.T) {
	// (t *testing.T) -> t é o enderço de testing.T, de modo que
	// t.Errorf é a mesma coisa que testing.T.Errorf, apenas para
	// simplicidade e garantir carregamento rápido
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
