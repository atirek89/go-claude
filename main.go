// Command go-claude is a small starter CLI used to learn Claude Code with Go.
package main

import (
	"fmt"
	"os"

	"github.com/atirek89/go-claude/internal/greeting"
)

func main() {
	name := "World"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	fmt.Println(greeting.Hello(name))
}
