package greeting

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "simple name", input: "Atirek", want: "Hello, Atirek!"},
		{name: "empty falls back to World", input: "", want: "Hello, World!"},
		{name: "whitespace falls back to World", input: "   ", want: "Hello, World!"},
		{name: "trims surrounding whitespace", input: "  Go  ", want: "Hello, Go!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.input); got != tt.want {
				t.Errorf("Hello(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
