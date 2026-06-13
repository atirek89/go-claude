BINARY := go-claude

.PHONY: build run serve test fmt vet clean

build:
	go build -o $(BINARY) .

run:
	go run .

serve:
	go run . --serve

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

clean:
	rm -f $(BINARY) coverage.out coverage.html
