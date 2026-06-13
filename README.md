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

**CLI mode**

```sh
go run . Atirek
# Hello, Atirek!

go run .
# Hello, World!
```

**Server mode**

```sh
go run . --serve           # listens on :8080
go run . --serve --port 9090

# Endpoints
curl http://localhost:8080/hello
# Hello, World!
curl http://localhost:8080/hello?name=Atirek
# Hello, Atirek!
```

## Common tasks

```sh
make build   # compile the binary
make run     # run CLI
make serve   # start HTTP server on :8080
make test    # run tests
make fmt     # format the code
make vet     # run go vet
make clean   # remove build artifacts
```
