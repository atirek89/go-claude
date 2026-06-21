// Package greeting builds friendly greeting messages.
package greeting

import (
	"fmt"
	"strings"
)

// Hello returns a greeting for the given name. Falls back to "World" if empty.
func Hello(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

// Goodbye returns a farewell for the given name. Falls back to "World" if empty.
func Goodbye(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Goodbye, %s!", name)
}
