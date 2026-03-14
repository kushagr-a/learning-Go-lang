package greet

import (
	"strings"
)

// its is re usable package not executable like main
// yo write bussiness logic or anything in this package
// and use it in main package

// this is exported funstion start with Capital letter
// other package can call it -> p1, p2
func Hello(name string) string {
	cleanName := normalizeName(name)
	return "Hello " + cleanName
}

// hellper function
func normalizeName(name string) string {
	n := strings.TrimSpace(name)

	if n == "" {
		return "World"
	}
	return strings.ToUpper(n)
}
