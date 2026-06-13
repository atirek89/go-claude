# go-claude

A starter Go project for learning [Claude Code](https://claude.ai/code) with Go.

## Requirements

- Go 1.24 or newer

## Project layout

```
.
├── main.go                       # CLI entry point
├── internal/
│   └── greeting/
│       ├── greeting.go           # greeting logic
│       └── greeting_test.go      # table-driven tests
├── go.mod
├── Makefile
└── README.md
```

## Getting started

```sh
go run . Atirek
# Hello, Atirek!

go run .
# Hello, World!
```

## Common tasks

```sh
make build   # compile the binary
make test    # run tests
make fmt     # format the code
make vet     # run go vet
make run     # run the program
make clean   # remove build artifacts
```
