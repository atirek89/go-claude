// Command go-claude is a starter CLI / HTTP server used to learn Claude Code with Go.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/atirek89/go-claude/internal/greeting"
	"github.com/atirek89/go-claude/internal/server"
)

func main() {
	serve := flag.Bool("serve", false, "start the HTTP server instead of printing a greeting")
	port := flag.String("port", "8080", "port for the HTTP server")
	flag.Parse()

	if *serve {
		mux := server.New()
		if err := server.Run(":"+*port, mux); err != nil {
			log.Fatal(err)
		}
		return
	}

	name := "World"
	if flag.NArg() > 0 {
		name = flag.Arg(0)
	}
	fmt.Println(greeting.Hello(name))
}
